package main

import (
	"fmt"
	"log"
	"myapp/api/router"
	"myapp/config"
	"net/http"
)

//  @title          MYAPP API
//  @version        1.0
//  @description    This is a sample RESTful API with a CRUD

//  @contact.name   Mahdi Younesi

// @host       localhost:8080
// @basePath   /v1
func main() {
	c := config.New()

	r := router.New()
	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", c.Server.Port),
		Handler:      r,
		ReadTimeout:  c.Server.TimeoutRead,
		WriteTimeout: c.Server.TimeoutWrite,
		IdleTimeout:  c.Server.TimeoutIdle,
	}

	log.Println("Starting server " + s.Addr)
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed")
	}
}

func hiThere(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi there!"))
}
