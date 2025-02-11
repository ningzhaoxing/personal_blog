package server

import (
	"context"
	"fmt"
	"gin-blog/initialize"
	"gin-blog/pkg/globals"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() {
	initialize.AppInit()

	InitRouter(globals.E)

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", globals.C.App.Host, globals.C.App.Port),
		Handler: globals.E,
	}

	// 平滑关机
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println("server-->", err)
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)
	<-sigs
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("stutdown err %v \n", err)
	}
	log.Println("shutdown-->ok")
}
