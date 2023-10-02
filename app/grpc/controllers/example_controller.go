package controllers

import (
	"context"
	"net/http"

	proto "github.com/go-dyn/app/grpc/protos/example"
	"github.com/go-dyn/app/services"
)

type ExampleController struct {
	BaseController
	proto.UnimplementedExampleServiceServer
	service services.IExample
}

func NewExampleController() *ExampleController {
	return &ExampleController{
		service: services.NewExampleImpl(),
	}
}

// List implements user.UserServiceServer.
func (*ExampleController) List(context.Context, *proto.ListRequest) (*proto.ListResponse, error) {
	panic("unimplemented")
}

func (c *ExampleController) Show(ctx context.Context, req *proto.ShowRequest) (*proto.ItemResponse, error) {
	item, err := c.service.Get(req.Id)
	if err != nil {
		return &proto.ItemResponse{
			Code:    http.StatusUnauthorized,
			Message: "Error",
			Item: &proto.Item{
				Id: uint64(0),
			},
		}, err
	}

	return &proto.ItemResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Item: &proto.Item{
			Id: uint64(item.ID),
		},
	}, err
}

// Create implements user.UserServiceServer.
func (*ExampleController) Create(context.Context, *proto.CreateRequest) (*proto.ItemResponse, error) {
	panic("unimplemented")
}

// Update implements user.UserServiceServer.
func (*ExampleController) Update(context.Context, *proto.UpdateRequest) (*proto.ItemResponse, error) {
	panic("unimplemented")
}
