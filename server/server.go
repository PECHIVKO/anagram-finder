package server

import (
	"log"
	"net/http"

	"github.com/PECHIVKO/anagram-finder/api/router"
	"github.com/PECHIVKO/anagram-finder/models"
)

// var Server *http.Server

func Run(port string, d *models.Dictionary) {
	//
	Session := &http.Server{
		Addr:    port,
		Handler: router.NewRouter(d),
	}
	log.Fatal(Session.ListenAndServe())
}
