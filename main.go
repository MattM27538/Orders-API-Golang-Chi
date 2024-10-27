package main

import (
	"fmt"
	"context"
	"os"
	"os/signal"
	"github.com/MattM27538/Microservice1/application"
)

func main(){
	app:=application.NewApp()

	ctx, cancel:=signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err:=app.Start(ctx)
	if err!=nil{
		fmt.Println("failed to start app:", err)
	}
} 