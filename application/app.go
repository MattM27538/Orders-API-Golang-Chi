// package application

// import(
// 	"fmt"
// 	"context"
// 	"net/http"
// 	"time"
	
// 	"github.com/redis/go-redis/v9"
// )

// type App struct{
// 	router http.Handler
// 	redisDB *redis.Client
// }

// func NewApp() *App{
// 	app:=&App{
// 		router:loadRoutes(),
// 		redisDB:redis.NewClient(&redis.Options{}),
// 	}
// 	return app
// }

// func (a *App) Start(ctx context.Context) error{
// 	server:=&http.Server{
// 		Addr:":3000",
// 		Handler:a.router,
// 	}

// 	err:=a.redisDB.Ping(ctx).Err()
// 	if err!= nil {
// 		return fmt.Errorf("failed to connect to redis: %w", err)
// 	}

// 	defer func(){
// 		if err:=a.redisDB.Close(); err!=nil{
// 			fmt.Println("failed to close redis", err)
// 		}
// 	}()


// 	fmt.Println("Starting server")

// 	channel:=make(chan error, 1)

// 	go func(){
// 		err=server.ListenAndServe()
// 		if err!=nil{
// 			channel<-fmt.Errorf("failed to start server:%w", err)
// 			}
// 		close(channel)
// 	}()
	
// 	select{
// 	case err:=<-channel:
// 		return err
// 	case <-ctx.Done():
// 		timeout, cancel:=context.WithTimeout(context.Background(),time.Second*10)
// 		defer cancel()

// 		return server.Shutdown(timeout)
// 	}

// 	return nil
// }

package application

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

type App struct {
	router http.Handler
	redisDB    *redis.Client
  config Config
}

func NewApp(config Config) *App {
	app := &App{
	  redisDB:    redis.NewClient(&redis.Options{
      Addr: config.RedisAddress,
    }),
    config: config,
	}

	app.loadRoutes()

	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", a.config.ServerPort),
		Handler: a.router,
	}

	err := a.redisDB.Ping(ctx).Err()
	if err != nil {
		return fmt.Errorf("failed to connect to redis: %w", err)
	}

	defer func() {
		if err := a.redisDB.Close(); err != nil {
			fmt.Println("failed to close redis", err)
		}
	}()

	fmt.Println("Starting server")

	ch := make(chan error, 1)

	go func() {
		err = server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("failed to start server: %w", err)
		}
		close(ch)
	}()

	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		return server.Shutdown(timeout)
	}

	return nil
}