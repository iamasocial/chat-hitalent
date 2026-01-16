package main

import (
	"chat/internal/config"
	"chat/internal/db"
	"chat/internal/repository"
	"chat/internal/server"
	"chat/internal/service"
	handler "chat/internal/transport/http"
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	cfg, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg)

	dbConn, err := db.NewPostgres(&cfg.DB)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	repo := repository.NewRepository(dbConn)
	svc := service.NewService(repo)
	h := handler.NewHandler(svc)
	router := handler.NewRouter(h)
	srv := server.NewServer(&cfg.HTTP, router)

	go func() {
		if err := srv.Run(); err != nil && err != http.ErrServerClosed {
			log.Print("server error: %w", err)
			cancel()
		}
	}()

	log.Print("server started")

	<-ctx.Done()
}
