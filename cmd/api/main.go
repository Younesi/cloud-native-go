package main

import (
	"fmt"
	"log"
	"myapp/config"
	"net/http"
)

func main() {
	c := config.New()

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hiThere)

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", c.Server.Port),
		Handler:      mux,
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
