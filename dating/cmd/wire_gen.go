// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/sterligov/otus_highload/dating/internal/config"
	"github.com/sterligov/otus_highload/dating/internal/gateway/sql"
	"github.com/sterligov/otus_highload/dating/internal/server/http"
	"github.com/sterligov/otus_highload/dating/internal/server/http/handler/v1"
	"github.com/sterligov/otus_highload/dating/internal/server/http/middleware"
	"github.com/sterligov/otus_highload/dating/internal/usecase/auth"
	"github.com/sterligov/otus_highload/dating/internal/usecase/city"
	"github.com/sterligov/otus_highload/dating/internal/usecase/user"
)

// Injectors from wire.go:

func setup(configConfig *config.Config) (*internalhttp.Server, func(), error) {
	db, err := sql.NewDatabase(configConfig)
	if err != nil {
		return nil, nil, err
	}
	userGateway := sql.NewUserGateway(db)
	useCase := user.NewUserUseCase(userGateway)
	userHandler := v1.NewUserHandler(useCase)
	cityGateway := sql.NewCityGateway(db)
	cityUseCase := city.NewCityUseCase(cityGateway)
	cityHandler := v1.NewCityHandler(cityUseCase)
	authUseCase := auth.NewAuthUseCase(userGateway)
	ginJWTMiddleware := middleware.Auth(configConfig, authUseCase)
	handler := v1.NewHandler(userHandler, cityHandler, ginJWTMiddleware)
	server, err := internalhttp.NewServer(configConfig, handler)
	if err != nil {
		return nil, nil, err
	}
	return server, func() {
	}, nil
}