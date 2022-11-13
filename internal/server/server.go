package server

import (
	"context"
	"fmt"
	"io/fs"
	"log"
	"math/rand"
	"net"
	"time"

	actionpkg "github.com/ankeesler/spirits/internal/action"
	convertaction "github.com/ankeesler/spirits/internal/action/convert"
	actionqueue "github.com/ankeesler/spirits/internal/action/queue/memory"
	actionservice "github.com/ankeesler/spirits/internal/action/service"
	"github.com/ankeesler/spirits/internal/battle/runner"
	battleservice "github.com/ankeesler/spirits/internal/battle/service"
	battlememory "github.com/ankeesler/spirits/internal/battle/storage/memory"
	"github.com/ankeesler/spirits/internal/builtin"
	spiritpkg "github.com/ankeesler/spirits/internal/spirit"
	convertspirit "github.com/ankeesler/spirits/internal/spirit/convert"
	spiritservice "github.com/ankeesler/spirits/internal/spirit/service"
	spiritmemory "github.com/ankeesler/spirits/internal/spirit/storage/memory"
	"github.com/ankeesler/spirits/internal/storage/memory"
	"github.com/ankeesler/spirits/pkg/api/spirits/v1"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Config struct {
	Port int

	SpiritBuiltinDir fs.FS
	ActionBuiltinDir fs.FS
}

type Server struct {
	port int

	s            *grpc.Server
	battleRunner *runner.Runner
}

func Wire(c *Config) (*Server, error) {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	actionRepo := memory.New[*actionpkg.Action](r)
	actionQueue := actionqueue.New()
	actionService := actionservice.New(actionRepo)

	spiritRepo := spiritmemory.New(r, actionRepo)
	spiritService := spiritservice.New(spiritRepo)

	battleRepo := battlememory.New(r, spiritRepo)
	battleService := battleservice.New(battleRepo, actionQueue, actionQueue)

	s := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			streamLogFunc,
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_opentracing.StreamServerInterceptor(),
			grpc_prometheus.StreamServerInterceptor,
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			unaryLogFunc,
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_opentracing.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)

	spiritsv1.RegisterSpiritServiceServer(s, spiritService)
	spiritsv1.RegisterActionServiceServer(s, actionService)
	spiritsv1.RegisterBattleServiceServer(s, battleService)

	battleRunner := runner.New(battleRepo)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	if err := builtin.Load[*spiritsv1.Spirit, *spiritpkg.Spirit](
		ctx,
		c.SpiritBuiltinDir,
		func() *spiritsv1.Spirit { return &spiritsv1.Spirit{} },
		convertspirit.FromAPI,
		spiritRepo,
	); err != nil {
		return nil, fmt.Errorf("load builtin spirits: %w", err)
	}

	if err := builtin.Load[*spiritsv1.Action, *actionpkg.Action](
		ctx,
		c.ActionBuiltinDir,
		func() *spiritsv1.Action { return &spiritsv1.Action{} },
		convertaction.FromAPI,
		actionRepo,
	); err != nil {
		return nil, fmt.Errorf("load builtin actions: %w", err)
	}

	return &Server{
		port: c.Port,

		s:            s,
		battleRunner: battleRunner,
	}, nil
}

func (s *Server) Serve(ctx context.Context) error {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port)) // Closed by grpc.Serve()
	if err != nil {
		return fmt.Errorf("listen: %w", err)
	}

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		go func() {
			<-ctx.Done()
			s.s.GracefulStop()
		}()

		log.Printf("server listening on %s", l.Addr().String())
		return s.s.Serve(l)
	})

	g.Go(func() error {
		return s.battleRunner.Run(ctx)
	})

	return g.Wait()
}

func unaryLogFunc(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	log.Printf("unary req: %s: %v", info.FullMethod, textproto(req))
	rsp, err := handler(ctx, req)
	log.Printf("unary rsp: %s: %v %v", info.FullMethod, textproto(req), err)
	return rsp, err
}

func streamLogFunc(
	srv interface{}, ss grpc.ServerStream,
	info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Printf("stream req: %s", info.FullMethod)
	err := handler(srv, &loggingStream{info: info, ServerStream: ss})
	log.Printf("stream rsp: %s: %v", info.FullMethod, err)
	return err
}

type loggingStream struct {
	info *grpc.StreamServerInfo
	grpc.ServerStream
}

func (l *loggingStream) SendMsg(m interface{}) error {
	log.Printf("stream rsp: %s: %v...", l.info.FullMethod, textproto(m))
	return l.ServerStream.SendMsg(m)
}

func textproto(i any) string {
	return prototext.MarshalOptions{
		Multiline: true,
	}.Format(i.(protoreflect.ProtoMessage))
}
