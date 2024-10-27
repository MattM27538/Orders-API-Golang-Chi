package application

import(
	"fmt"
	"context"
	"net/http"
	"time"
	"github.com/redis/go-redis/v9"
)

type App struct{
	router http.Handler
	redisDB *redis.Client
}

func NewApp() *App{
	app:=&App{
		router:loadRoutes(),
		redisDB:redis.NewClient(&redis.Options{}),
	}
	return app
}

func (a *App) Start(ctx context.Context) error{
	server:=&http.Server{
		Addr:":3000",
		Handler:a.router,
	}

	err:=a.redisDB.Ping(ctx).Err()
	if err!= nil {
		return fmt.Errorf("failed to connect to redis: %w", err)
	}

	defer func(){
		if err:=a.redisDB.Close(); err!=nil{
			fmt.Println("failed to close redis", err)
		}
	}()


	fmt.Println("Starting server")

	channel:=make(chan error, 1)

	go func() {
		err=server.ListenAndServe()
		if err!=nil{
			channel<-fmt.Errorf("failed to start server:%w", err)
			}
		close(channel)
	}()
	
	select{
	case err:=<-channel:
		return err
	case <-ctx.Done():
		timeout, cancel:=context.WithTimeout(context.Background(),time.Second*10)
		defer cancel()

		return server.Shutdown(timeout)
	}

	return nil
}