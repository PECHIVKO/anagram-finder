package router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Handle func(http.ResponseWriter, *http.Request)

func (rt *Router) homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		rt.errorHandler(w, r, http.StatusNotFound)
		return
	}
	fmt.Fprint(w, "welcome home")
}

func (rt *Router) errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "NOT FOUND (404)")
	}
}

func (rt *Router) returnAnagrams(w http.ResponseWriter, r *http.Request) {
	word, ok := r.URL.Query()["word"]
	if !ok || len(word[0]) < 1 {
		log.Println("Url Param 'word' is missing")
		return
	}
	response, err := rt.dict.Finder(word[0])
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(response)
}

func (rt *Router) loadJSON(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data []string
	err := decoder.Decode(&data)
	if err != nil {
		log.Println("Unexpected JSON structure:", err)
	} else {
		Dict.service.LoadDictionary(data)
		log.Println("JSON loaded")
	}
	defer r.Body.Close()
}
