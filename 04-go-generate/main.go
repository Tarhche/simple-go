package main

import (
	"go-generate/api"
	"log"
)

//go:generate go tool swag init --generalInfo api/api.go --output ./docs
func main() {
	if err := api.StartServer(); err != nil {
		log.Fatalln(err)
	}
}
