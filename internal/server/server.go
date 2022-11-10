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
	"github.com/ankeesler/spirits/pkg/api"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
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

	spiritRepo := spiritmemory.New(r)
	spiritService := spiritservice.New(spiritRepo)

	battleRepo := battlememory.New(r)
	battleService := battleservice.New(battleRepo, actionQueue)

	s := grpc.NewServer(grpc.UnaryInterceptor(unaryLogFunc), grpc.StreamInterceptor(streamLogFunc))
	api.RegisterSpiritServiceServer(s, spiritService)
	api.RegisterActionServiceServer(s, actionService)
	api.RegisterBattleServiceServer(s, battleService)

	battleRunner, err := runner.Wire(battleRepo)
	if err != nil {
		return nil, fmt.Errorf("wire battle controller: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	if err := builtin.Load[*api.Spirit, *spiritpkg.Spirit](
		ctx,
		c.SpiritBuiltinDir,
		func() *api.Spirit { return &api.Spirit{} },
		convertspirit.FromAPI,
		spiritRepo,
	); err != nil {
		return nil, fmt.Errorf("load builtin spirits: %w", err)
	}

	if err := builtin.Load[*api.Action, *actionpkg.Action](
		ctx,
		c.ActionBuiltinDir,
		func() *api.Action { return &api.Action{} },
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
	log.Printf("Unary req: %s: %v", info.FullMethod, req)
	rsp, err := handler(ctx, req)
	log.Printf("Unary rsp: %s: %v %v", info.FullMethod, rsp, err)
	return rsp, err
}

func streamLogFunc(
	srv interface{}, ss grpc.ServerStream,
	info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Printf("Stream req: %s", info.FullMethod)
	err := handler(srv, ss)
	log.Printf("Stream rsp: %s: %v", info.FullMethod, err)
	return err
}
