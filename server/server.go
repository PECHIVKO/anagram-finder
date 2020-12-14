package server

import (
	"log"
	"net/http"

	"github.com/PECHIVKO/anagram-finder/api/router"
	"github.com/PECHIVKO/anagram-finder/service"
)

func Run(port string, d *service.Dictionary) {
	// Init Session
	Session := &http.Server{
		Addr:    port,
		Handler: router.NewRouter(d),
	}
	log.Fatal(Session.ListenAndServe())
}
