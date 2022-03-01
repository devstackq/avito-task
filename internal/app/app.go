package app

import (
	"avito/internal/config"
	"avito/internal/server"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

const confFile = ".env"

func init() {
	if err := godotenv.Load(confFile); err != nil {
		log.Printf(".env file does not exist ...\n app starting with default configuration")
	}
}

func Start() {
	cfg := config.GetConfig()

	app := server.NewApp(cfg)

	app.Init()
	//gracefull shutdown
	ctx, shutdown := context.WithCancel(context.Background())
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		shutdown()
	}()
	app.Run(ctx)
}
