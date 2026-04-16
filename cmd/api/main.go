package main

import (
	"github.com/yasm3/prevently/internal/api"
)

func main() {
	router := api.NewRouter()
	server := api.NewAPIServer(router)
	server.Run()
}
