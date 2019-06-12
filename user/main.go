package main

import (
	"fmt"
	"goapi/handlers"
	"net/http"
	"os"
)

func main() { //we want the handler to respond to the root user

	http.HandleFunc("/", handlers.RootHandler)
	err := http.ListenAndServe("localhost:11111", nil) //nil because we will not use any handler
	if err != nil {
		// panic(err) //terminate the program and display error
		fmt.Println(err)
		os.Exit(1)
	}
}
