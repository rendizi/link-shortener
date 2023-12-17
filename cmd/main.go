package main

import (
	"fmt"
	"module_name/http/server/handler"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handler.isValid)
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		if err == http.ErrServerClosed {
			fmt.Println("server closed")
		} else {
			fmt.Printf("error starting server: %s\n", err)
			os.Exit(1)
		}
	}
}
