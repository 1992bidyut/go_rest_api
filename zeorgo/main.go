package main

import (
	"os"
	"zeorgo/api"
)

func main() {
	server := api.NewAPIServer(":8080")
	server.Run()
}

func env_setup() {
	os.Setenv("HOST", "localhost")
	os.Setenv("HOST_PORT", "8080")
}
