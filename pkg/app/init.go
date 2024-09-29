package app

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/yesmishgan/go-project-template/pkg/swaggway"
)

func (a *App) initPublicHTTPHandlers() {
	a.publicServer = chi.NewMux()

	if a.desc != nil {
		mux, err := swaggway.NewMux(a.desc, a.publicServer)
		if err != nil {
			log.Fatalf("error while init gateway: %v", err)
		}

		if err := a.desc.RegisterGateway(context.Background(), mux); err != nil {
			log.Fatalf("error while register gateway: %v", err)
		}
	}
	// default handler preventing panic in case of no handlers: https://github.com/go-chi/chi/issues/362
	if len(a.publicServer.Routes()) == 0 {
		a.publicServer.HandleFunc("/", http.NotFound)
	}
}

func (a *App) initAdminHTTP() {
	a.adminServer = chi.NewMux()
}

func (a *App) initGRPCServer() {
	a.grpcServer = grpc.NewServer()

	a.desc.RegisterGRPC(a.grpcServer)
	reflection.Register(a.grpcServer)
}
