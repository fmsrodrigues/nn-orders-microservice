package main

import (
	"context"
	"fmt"
	"orders/application"
)

func main() {
	app := application.NewApp()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("Failed to start server: ", err)
	}
}
