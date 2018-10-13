package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type App struct {
	Controller Controller
}

func (a *App) Run() {
	a.Controller = Controller{}

	routers := GetRoutes()
	for _, route := range routers {
		a.Controller.AddRoute(route.Path, route.Method, route.Fn)
	}

	srv := http.Server{
		Addr:         "127.0.0.1:8081",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      a.Controller.Router,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatalln("cannot start server, err: ", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	srv.Shutdown(ctx)
	defer cancel()

	log.Println("shutting down")
	os.Exit(0)
}
