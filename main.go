package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"MServer/config"
	"MServer/globalVar"
	"MServer/initRouter"
	//"MServer/models"
)

var err error

func main() {

	//load Env & setup Logger
	config.LoadEnvAndSetupLogger()

	//setup DB
	config.DB = config.NewDatabase()
	defer config.CloseDatabase()
	//config.DB.AutoMigrate(&models.User{})

	config.RedisClient = config.NewRedisClient()
	defer config.RedisClient.Close()

	router := initRouter.SetupRouter()
	router.Static("/static", "./static")
	//router.StaticFile("/favicon.ico","./static/favicon.ico")
	router.LoadHTMLGlob("tpl/**/*")
	Srv := &http.Server{
		Addr:    ":" + globalVar.USE_PORT,
		Handler: router,
	}

	go func() {
		// service connections
		if err := Srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			config.Logger.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	//close() GRPC Websocket etc.
	config.Logger.Println("system closing step 1")

	ctxEnd, _ := context.WithTimeout(context.Background(), 1*time.Second)
	//defer cancelEnd()
	if err := Srv.Shutdown(ctxEnd); err != nil {
		config.Logger.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 1 seconds.
	config.Logger.Println("timeout of 1 seconds.")
	<-ctxEnd.Done()
	config.Logger.Println("Server exiting")
}
