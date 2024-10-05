package dummy

import (
	"context"
	desc "github.com/yesmishgan/go-pokeball/internal/pb/api/dummy"
	"github.com/yesmishgan/go-pokeball/pkg/app"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Implementation struct {
	desc.UnimplementedDummyServiceServer
}

func NewDummyService() *Implementation {
	return &Implementation{}
}

func (i *Implementation) GetDescription() app.ServiceDesc {
	return desc.NewDummyServiceServiceDesc(i)
}

func (i *Implementation) Ping(_ context.Context, _ *desc.PingRequest) (*desc.PingResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Ping method unimplemented")
}
