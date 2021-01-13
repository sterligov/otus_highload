// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/sterligov/otus_highload/dating/internal/config"
	"github.com/sterligov/otus_highload/dating/internal/domain"
	"github.com/sterligov/otus_highload/dating/internal/gateway/sql"
	internalhttp "github.com/sterligov/otus_highload/dating/internal/server/http"
	v1 "github.com/sterligov/otus_highload/dating/internal/server/http/handler/v1"
	"github.com/sterligov/otus_highload/dating/internal/server/http/middleware"
	"github.com/sterligov/otus_highload/dating/internal/usecase/auth"
	"github.com/sterligov/otus_highload/dating/internal/usecase/user"
)

func setup(*config.Config) (*internalhttp.Server, func(), error) {
	panic(wire.Build(
		wire.Bind(new(domain.UserGateway), new(*sql.UserGateway)),
		sql.NewDatabase,
		sql.NewUserGateway,
		user.NewUserUseCase,
		auth.NewAuthUseCase,
		v1.NewUserHandler,
		middleware.Auth,
		v1.NewHandler,
		internalhttp.NewServer,
	))
}
