package controllers

import (
	"context"
	"net/http"

	goravelhttp "github.com/goravel/framework/http"

	proto "github.com/go-dyn/app/grpc/protos/user"
	"github.com/go-dyn/app/services"
)

type UserController struct {
	proto.UnimplementedUserServiceServer
	BaseController
	service services.IUser
}

func NewUserController() *UserController {
	return &UserController{
		service: services.NewUserImpl(),
	}
}

func (c *UserController) List(context.Context, *proto.ListRequest) (*proto.ListResponse, error) {
	list, err := c.service.List()

	if err != nil {
		return &proto.ListResponse{
			Code:    http.StatusUnauthorized,
			Message: err.Error(),
		}, nil
	}
	items := []*proto.Item{}
	for _, item := range list {
		items = append(items, &proto.Item{
			Id:         uint64(item.ID),
			Identifier: item.Identifier,
		})
	}
	return &proto.ListResponse{
		Code:  http.StatusOK,
		Items: items,
	}, nil
}

func (c *UserController) Paginate(context.Context, *proto.PaginateRequest) (*proto.PaginateResponse, error) {
	list, total, err := c.service.Paginate(1, 10)

	if err != nil {
		return &proto.PaginateResponse{
			Code:    http.StatusUnauthorized,
			Message: err.Error(),
		}, nil
	}
	items := []*proto.Item{}
	for _, item := range list {
		items = append(items, &proto.Item{
			Id:         uint64(item.ID),
			Identifier: item.Identifier,
		})
	}
	return &proto.PaginateResponse{
		Code:  http.StatusOK,
		Total: total,
		Items: items,
	}, nil
}

func (c *UserController) Show(ctx context.Context, req *proto.ShowRequest) (*proto.ItemResponse, error) {
	httpCtx := goravelhttp.Background()

	item, err := c.service.GetAuthUser(httpCtx)
	if err != nil {
		return &proto.ItemResponse{
			Code:    http.StatusUnauthorized,
			Message: err.Error(),
		}, nil
	}

	return &proto.ItemResponse{
		Code: http.StatusOK,
		Item: &proto.Item{
			Id: uint64(item.ID),
		},
	}, nil
}

// Create implements user.UserServiceServer.
func (*UserController) Create(context.Context, *proto.CreateRequest) (*proto.ItemResponse, error) {
	panic("unimplemented")
}

// Update implements user.UserServiceServer.
func (*UserController) Update(context.Context, *proto.UpdateRequest) (*proto.ItemResponse, error) {
	panic("unimplemented")
}
