package app

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

func (a *App) runGRPC() {
	go func() {
		if err := a.grpcServer.Serve(a.listeners.grpc); err != nil {
			log.Fatalf(fmt.Errorf("grpc: %w", err).Error())
		}
	}()
}

func (a *App) runPublicHTTP() {
	publicServer := &http.Server{Handler: a.publicServer}
	go func() {
		if err := publicServer.Serve(a.listeners.http); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf(fmt.Errorf("http.public: %w", err).Error())
		}
	}()
}

func (a *App) runAdminHTTP() {
	adminServer := &http.Server{Handler: a.adminServer}
	go func() {
		if err := adminServer.Serve(a.listeners.httpAdmin); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf(fmt.Errorf("http.admin: %w", err).Error())
		}
	}()
}
