package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("API Started!")
	env_setup()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})
	var host_url = os.Getenv("HOST") + ":" + os.Getenv("HOST_PORT")
	http.ListenAndServe(host_url, nil)
}

func env_setup() {
	os.Setenv("HOST", "localhost")
	os.Setenv("HOST_PORT", "8080")
}
