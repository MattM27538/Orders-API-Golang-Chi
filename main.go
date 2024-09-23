package main

import (
	"fmt"
	"context"
	"github.com/MattM27538/Microservice1/application"
)

func main(){
	app:=application.NewApp()
	err:=app.Start(context.TODO())
	if err!=nil{
		fmt.Println("failed to start app:", err)
	}
} 