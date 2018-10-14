package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"
)

type App struct {
	Controller Controller
	srv        http.Server
}

func (a *App) Run(addr, port string) {
	a.Controller = Controller{}

	routers := GetRoutes()
	for _, route := range routers {
		a.Controller.AddRoute(route.Path, route.Method, route.Fn)
	}

	if strings.HasPrefix(addr, "http://") {
		addr = strings.TrimLeft(addr, "http://")
	}

	a.srv = http.Server{
		Addr:         addr + ":" + port,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      a.Controller.Router,
	}

	go func() {
		err := a.srv.ListenAndServe()
		if err != nil {
			log.Fatalln("cannot start server, err: ", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	a.srv.Shutdown(ctx)
	defer cancel()

	log.Println("shutting down")
	os.Exit(0)
}

func (a *App) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	a.srv.Shutdown(ctx)
	defer cancel()

	log.Println("shutting down")
	os.Exit(0)
}
