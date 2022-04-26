package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ankeesler/spirits/internal/log"
)

func Run(ctx context.Context, addr string, h http.Handler) error {
	server := http.Server{
		Addr:    addr,
		Handler: h,
	}

	shutdownCtx, cancelShutdownCtx := context.WithCancel(context.TODO())
	defer cancelShutdownCtx()

	go func() {
		select {
		case <-ctx.Done():
		case <-shutdownCtx.Done():
		}

		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Info(fmt.Sprintf("server shutdown error: %s", err.Error()))
		}
		cancelShutdownCtx()
	}()

	log.Info(fmt.Sprintf("server listening on %s", addr))
	err := server.ListenAndServe()
	log.Info(fmt.Sprintf("server listen and serve returned error: %s", err.Error()))
	cancelShutdownCtx()

	<-shutdownCtx.Done()

	return nil
}
