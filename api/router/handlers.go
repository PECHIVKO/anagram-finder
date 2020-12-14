package router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (rt *Router) homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		rt.errorHandler(w, r, http.StatusNotFound)
		return
	}
	fmt.Fprint(w, "welcome home")
}

// custom errorHandler
func (rt *Router) errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "NOT FOUND (404)")
	}
}

// returnAnagrams handle get request
func (rt *Router) returnAnagrams(w http.ResponseWriter, r *http.Request) {
	word, ok := r.URL.Query()["word"]
	if !ok || len(word[0]) < 1 {
		log.Println("Url Param 'word' is missing")
		return
	}
	response, err := rt.dict.FindAnagrams(word[0])
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(response)
}

// loadJSON handle load request
func (rt *Router) loadJSON(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data []string
	err := decoder.Decode(&data)
	if err != nil {
		log.Println("Unexpected JSON structure:", err)
	} else {
		rt.dict.LoadDictionary(data)
		log.Println("JSON loaded")
	}
	defer r.Body.Close()
}
