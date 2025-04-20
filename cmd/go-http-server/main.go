package main

import (
	"fmt"
	"github.com/kant-github/go-http-server/pkg/config"
	"log"
	"net/http"
)

func main() {
	//load config
	cfg := config.MustLoad()

	//setup router
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello boys, go server is listening"))
	})

	server := http.Server{
		Addr:    cfg.Http_server.Addr,
		Handler: router,
	}
	fmt.Printf("Server started at %s\n", cfg.Http_server.Addr) // Add newline and fix message

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal("Server failed to start:", err) // Include actual error
	}

}
