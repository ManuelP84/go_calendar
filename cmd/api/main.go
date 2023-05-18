package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/ManuelP84/calendar/infra/http/handlers"
)

func main() {
	webApp := CreateApp()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	handlers.SetRoutes(webApp)

	if err := webApp.Run(ctx); err != nil {
		log.Fatalf("Fatal error %s", err.Error())
	}

}
