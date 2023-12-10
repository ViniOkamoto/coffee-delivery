package main

import (
	"coffee-delivery/main/application"
	"context"
	"fmt"
)

func main() {
	app := application.NewApp()

	err := app.Run(context.TODO())
	if err != nil {
		fmt.Println(err)
	}
}
