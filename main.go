package main

import (
	"github.com/PECHIVKO/anagram-finder/server"
	"github.com/PECHIVKO/anagram-finder/service"
)

const (
	defaultPort = ":8080"
)

func main() {
	// Init Dictionary in memory
	var Dict service.Dictionary = make(service.Dictionary)

	server.Run(defaultPort, &Dict) // Pass adress to Dictionary in router to use in handlers
}
