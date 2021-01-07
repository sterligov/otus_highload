// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/sterligov/otus_highload/dating/internal/config"
	"github.com/sterligov/otus_highload/dating/internal/server"
)

func setup(*config.Config) (*server.Server, func(), error) {
	panic(wire.Build())
}
