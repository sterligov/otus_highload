package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/sterligov/otus_highload/dating/internal/logger"

	"go.uber.org/zap"

	"github.com/sterligov/otus_highload/dating/internal/config"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "/etc/configs/config.yml", "Path to configuration file")
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

	var rerr error
	defer func() {
		if rerr != nil {
			log.Fatalln(rerr)
		}
	}()

	cfg, err := config.New(configFile)
	if err != nil {
		rerr = err
		return
	}

	if err := logger.InitGlobal(cfg); err != nil {
		rerr = err
		return
	}

	server, cleanup, err := setup(cfg)
	if err != nil {
		rerr = err
		return
	}
	defer cleanup()

	go func() {
		if err := server.Start(); err != nil {
			zap.L().Error("grpc server start failed", zap.Error(err))
		}
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT)

	<-signals
	signal.Stop(signals)

	if err := server.Stop(context.Background()); err != nil && !errors.Is(err, http.ErrServerClosed) {
		zap.L().Warn("http server stop failed", zap.Error(err))
	}
}
