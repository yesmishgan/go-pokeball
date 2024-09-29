package app

import (
	"fmt"
	"net"
	"sync"

	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"
)

type listeners struct {
	http      net.Listener
	httpAdmin net.Listener
	grpc      net.Listener
}

type App struct {
	publicServer chi.Router
	adminServer  chi.Router
	grpcServer   *grpc.Server
	listeners    *listeners
	desc         ServiceDesc
}

func New() (*App, error) {
	lis, err := newListeners()
	if err != nil {
		return nil, err
	}

	return &App{
		listeners: lis,
	}, nil
}

func (a *App) Run(impl ...Service) {
	descs := make([]ServiceDesc, 0, len(impl))
	for _, i := range impl {
		descs = append(descs, i.GetDescription())
	}

	a.desc = NewCompoundServiceDesc(descs...)

	wg := sync.WaitGroup{}
	wg.Add(1)

	a.initGRPCServer()
	a.initPublicHTTPHandlers()
	a.initAdminHTTP()
	a.runPublicHTTP()
	a.runAdminHTTP()
	a.runGRPC()
	wg.Wait()
}

func (a *App) PublicServer() chi.Router {
	return a.publicServer
}

func newListeners() (*listeners, error) {
	grpcLis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8002))
	if err != nil {
		return nil, fmt.Errorf("net.Listen port=%d: %w", 8002, err)
	}
	httpLis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8000))
	if err != nil {
		return nil, fmt.Errorf("net.Listen port=%d: %w", 8000, err)
	}
	adminHttpLis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8001))
	if err != nil {
		return nil, fmt.Errorf("net.Listen port=%d: %w", 8001, err)
	}

	return &listeners{
		grpc:      grpcLis,
		http:      httpLis,
		httpAdmin: adminHttpLis,
	}, nil
}
