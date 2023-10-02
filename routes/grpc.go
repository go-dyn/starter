package routes

import (
	"github.com/go-dyn/app/grpc/controllers"
	"github.com/go-dyn/app/grpc/protos/example"
	"github.com/go-dyn/app/grpc/protos/user"
	"github.com/goravel/framework/facades"
)

func Grpc() {
	server := facades.Grpc().Server()
	user.RegisterUserServiceServer(server, controllers.NewUserController())
	example.RegisterExampleServiceServer(server, controllers.NewExampleController())
}
