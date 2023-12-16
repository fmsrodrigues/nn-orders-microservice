package main

import (
	"context"
	"fmt"
	"orders/application"
	"os"
	"os/signal"
)

func main() {
	app := application.NewApp(application.LoadConfig())

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		fmt.Println("Failed to start server: ", err)
	}
}
