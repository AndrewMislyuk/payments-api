package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/AndrewMislyuk/payments-api/internal/handler"
	"github.com/AndrewMislyuk/payments-api/internal/service"
	"github.com/AndrewMislyuk/payments-api/pkg/server"
	"github.com/AndrewMislyuk/payments-api/pkg/stripe"
)

// @title API Product Subscription
// @version 1.0
// @description API for frontend cliend

// @host localhost:3000
// @BasePath /
func main() {
	stripeMethod := stripe.NewStripe()

	apiService := service.NewService(stripeMethod)
	apiHandler := handler.NewHandler(apiService)

	srv := new(server.Server)

	go func() {
		if err := srv.Run("3000", apiHandler.InitRouter()); err != nil {
			log.Fatal(err)
		}
	}()

	log.Println("server has been running...")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("server was stopped")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("error occurred on server shutting down: %s", err.Error())
	}
}
