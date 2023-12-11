package main

import (
	"coffee-delivery/main/cmd/application"
	"context"
	"fmt"
)

func main() {
	fmt.Println("User Created")
	app := application.NewApp()

	err := app.Run(context.TODO())

	if err != nil {
		fmt.Println(err)
	}
}
