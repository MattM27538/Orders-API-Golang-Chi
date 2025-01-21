package main

import (
	"fmt"
	"context"
	"os"
	"os/signal"
	"github.com/MattM27538/Orders-API-Golang-Chi/application"
)

func main(){
	app:=application.NewApp(application.LoadConfig())

	ctx, cancel:=signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err:=app.Start(ctx)
	if err!=nil{
		fmt.Println("failed to start app:", err)
	}
} 