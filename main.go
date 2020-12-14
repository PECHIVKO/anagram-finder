package main

import (
	"github.com/PECHIVKO/anagram-finder/models"
	"github.com/PECHIVKO/anagram-finder/server"
)

const (
	defaultPort = ":8080"
)

// func handleRequests() {
// 	http.HandleFunc("/", homePage)
// 	http.HandleFunc("/get", returnAnagrams)
// 	http.HandleFunc("/load", loadJSON)
// 	err := http.ListenAndServe(defaultPort, nil)
// 	if err != nil {
// 		log.Fatal("ListenAndServe: ", err)
// 	}
// }

func main() {
	var Dict models.Dictionary = make(models.Dictionary)
	server.Run(defaultPort, &Dict)
}
