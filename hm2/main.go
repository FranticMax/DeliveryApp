package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/health", getHealth)

	fmt.Println("start listening on 8000")

	err := http.ListenAndServe(":8000", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func getHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got health request")
	io.WriteString(w, "{\"status\": \"OK\"}")
}
