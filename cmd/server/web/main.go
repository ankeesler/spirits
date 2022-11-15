package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	spiritsv1 "github.com/ankeesler/spirits/pkg/api/gateway/spirits/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

var (
	webAssetsDir = flag.String(
		"web-assets-dir", "build", "Path to web assets")
	upstreamAPIServer = flag.String(
		"upstream-api-server", "http://127.0.0.1:12345", "URL for upstream API server")
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
	mux.Handle("/api", gatewayMux(ctx, *upstreamAPIServer))

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux); err != nil {
		log.Fatalf("listen and serve: %s", err.Error())
	}
}

func gatewayMux(ctx context.Context, upstreamAPIServer string) http.Handler {
	gatewayMux := runtime.NewServeMux()

	conn, err := grpc.DialContext(ctx, upstreamAPIServer)
	if err != nil {
		log.Fatalf("dial upstream API server: %s", err.Error())
	}

	if err := spiritsv1.RegisterActionServiceHandler(ctx, gatewayMux, conn); err != nil {
		log.Fatalf("register action service upstream: %s", err.Error())
	}

	if err := spiritsv1.RegisterSpiritServiceHandler(ctx, gatewayMux, conn); err != nil {
		log.Fatalf("register spirit service upstream: %s", err.Error())
	}

	return gatewayMux
}
