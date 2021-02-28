package main

import (
	"fmt"
	"main/controllers"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", controllers.GetURLHandler)
	mux.HandleFunc("/create", controllers.CreateURLHandler)

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", mux)
}
