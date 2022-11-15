package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"

	spiritsv1 "github.com/ankeesler/spirits/pkg/api/spirits/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

var (
	webAssetsDir = flag.String(
		"web-assets-dir", "build", "Path to web assets")
	upstreamAPIServer = flag.String(
		"upstream-api-server", "127.0.0.1:12345", "URL for upstream API server")
)

func main() {
	log.SetFlags(log.Lmicroseconds | log.Lshortfile)

	var port int
	if portEnv, ok := os.LookupEnv("PORT"); ok {
		var err error
		port, err = strconv.Atoi(portEnv)
		if err != nil {
			log.Fatalf("invalid port: %v", err)
		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(*webAssetsDir)))
	mux.Handle("/api/", http.StripPrefix("/api", gatewayMux(ctx, *upstreamAPIServer)))

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("listen: %s", err.Error())
	}
	defer l.Close()

	log.Printf("server listening on %s", l.Addr().String())
	if err := http.Serve(l, mux); err != nil {
		log.Fatalf("http server: %s", err.Error())
	}
}

func gatewayMux(ctx context.Context, upstreamAPIServer string) http.Handler {
	mux := runtime.NewServeMux(
		runtime.WithErrorHandler(
			runtime.ErrorHandlerFunc(
				func(
					ctx context.Context,
					mux *runtime.ServeMux,
					marshaler runtime.Marshaler,
					w http.ResponseWriter,
					r *http.Request,
					err error,
				) {
					log.Printf("gateway error: %s %s: %s", r.Method, r.URL, err.Error())
					runtime.DefaultHTTPErrorHandler(ctx, mux, marshaler, w, r, err)
				},
			),
		),
		runtime.WithRoutingErrorHandler(
			runtime.RoutingErrorHandlerFunc(
				func(
					ctx context.Context,
					mux *runtime.ServeMux,
					marshaler runtime.Marshaler,
					w http.ResponseWriter,
					r *http.Request,
					i int,
				) {
					log.Printf("gateway routing error: %s %s: %d", r.Method, r.URL, i)
					runtime.DefaultRoutingErrorHandler(ctx, mux, marshaler, w, r, i)
				},
			),
		),
		runtime.WithStreamErrorHandler(
			runtime.StreamErrorHandlerFunc(
				func(ctx context.Context, err error) *status.Status {
					log.Printf("gateway stream error: %s", err.Error())
					return runtime.DefaultStreamErrorHandler(ctx, err)
				},
			),
		),
	)

	conn, err := grpc.DialContext(
		ctx, upstreamAPIServer, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("dial upstream API server: %s", err.Error())
	}

	if err := spiritsv1.RegisterActionServiceHandler(ctx, mux, conn); err != nil {
		log.Fatalf("register action service upstream: %s", err.Error())
	}

	if err := spiritsv1.RegisterSpiritServiceHandler(ctx, mux, conn); err != nil {
		log.Fatalf("register spirit service upstream: %s", err.Error())
	}

	return mux
}
