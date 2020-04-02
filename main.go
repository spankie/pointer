package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/spankie/pointer/router"
)

func main() {
	handler := router.NewRouter()
	handler.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}
	time.AfterFunc(5*time.Second, func() { server.Shutdown(context.Background()) })
	fmt.Println(server.ListenAndServe())
}
