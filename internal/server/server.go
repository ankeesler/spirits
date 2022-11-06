package server

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"math/rand"
	"net"
	"os"
	"time"

	actionpkg "github.com/ankeesler/spirits0/internal/action"
	memoryqueue "github.com/ankeesler/spirits0/internal/action/queue/memory"
	actionservice "github.com/ankeesler/spirits0/internal/action/service"
	"github.com/ankeesler/spirits0/internal/api"
	"github.com/ankeesler/spirits0/internal/battle/controller"
	battleservice "github.com/ankeesler/spirits0/internal/battle/service"
	battlememory "github.com/ankeesler/spirits0/internal/battle/storage/memory"
	"github.com/ankeesler/spirits0/internal/builtin"
	spiritservice "github.com/ankeesler/spirits0/internal/spirit/service"
	spiritmemory "github.com/ankeesler/spirits0/internal/spirit/storage/memory"
	"github.com/ankeesler/spirits0/internal/storage/memory"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

type Config struct {
	SpiritBuiltinDir fs.FS
	ActionBuiltinDir fs.FS
}

type Server struct {
	addr string

	s                *grpc.Server
	battleController *controller.Controller
}

func Wire(c *Config) (*Server, error) {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	actionRepo := memory.New[*api.Action](r)
	actionQueue := memoryqueue.New()
	actionService := actionservice.New(actionRepo, actionQueue)

	spiritRepo := spiritmemory.New(r)
	spiritService := spiritservice.New(spiritRepo, actionRepo)

	battleRepo := battlememory.New(r)
	battleService := battleservice.New(battleRepo, spiritRepo, actionQueue)

	s := grpc.NewServer()
	api.RegisterSpiritServiceServer(s, spiritService)
	api.RegisterActionServiceServer(s, actionService)
	api.RegisterBattleServiceServer(s, battleService)

	battleController, err := controller.Wire(battleRepo, actionRepo, actionQueue)
	if err != nil {
		return nil, fmt.Errorf("wire battle controller: %w", err)
	}

	if err := builtin.Load[*api.Spirit](
		c.SpiritBuiltinDir,
		spiritRepo,
		func() *api.Spirit { return &api.Spirit{} },
		func(spirit *api.Spirit) error {
			for _, action := range spirit.GetActions() {
				switch action.Definition.(type) {
				case *api.SpiritAction_ActionId:
					return errors.New("builtin spirit cannot have action from ID")
				}
			}
			return nil
		},
	); err != nil {
		return nil, fmt.Errorf("load builtin spirits: %w", err)
	}

	if err := builtin.Load[*api.Action](
		c.ActionBuiltinDir,
		actionRepo,
		func() *api.Action { return &api.Action{} },
		func(action *api.Action) error {
			_, err := actionpkg.FromAPI(action)
			return err
		},
	); err != nil {
		return nil, fmt.Errorf("load builtin actions: %w", err)
	}

	port, ok := os.LookupEnv("PORT")
	addr := fmt.Sprintf(":%s", port)
	if !ok {
		addr = ":0"
	}

	return &Server{
		addr: addr,

		s:                s,
		battleController: battleController,
	}, nil
}

func (s *Server) Serve(ctx context.Context) error {
	l, err := net.Listen("tcp", s.addr) // Closed by grpc.Serve()
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
		return s.battleController.Run(ctx)
	})

	return g.Wait()
}
