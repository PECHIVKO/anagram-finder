package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strings"
	"unicode"
)

type AnagramSlice []string

type Dictionary map[string]AnagramSlice

const (
	defaultPort = ":8080"
)

var Dict Dictionary = make(Dictionary)

func stringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, unicode.ToLower(runeValue))
	}
	return r
}

func sortStringByCharacter(s string) string {
	r := stringToRuneSlice(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func isDuplicate(as AnagramSlice, word string) bool {
	for _, a := range as {
		if a == word {
			return true
		}
	}
	return false
}

func (d Dictionary) loadDictionary(input []string) {
	for _, word := range input {
		sortedWord := sortStringByCharacter(word)
		if _, found := d[sortedWord]; found {
			if !isDuplicate(d[sortedWord], word) {
				d[sortedWord] = append(d[sortedWord], word)
			}
		} else {
			d[sortedWord] = AnagramSlice{word}
		}
	}
}

func finder(word string, d Dictionary) (result AnagramSlice, err error) {
	sortedWord := strings.ToLower(sortStringByCharacter(word))
	if _, found := d[sortedWord]; found {
		result = d[sortedWord]
		return result, nil

	}
	err = fmt.Errorf("There is no anagrams for word \"%s\" in dictionary", word)
	return nil, err
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/get", returnAnagrams)
	http.HandleFunc("/load", loadJSON)
	err := http.ListenAndServe(defaultPort, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	fmt.Fprint(w, "welcome home")
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "NOT FOUND (404)")
	}
}

func returnAnagrams(w http.ResponseWriter, r *http.Request) {
	word, ok := r.URL.Query()["word"]
	if !ok || len(word[0]) < 1 {
		log.Println("Url Param 'word' is missing")
		return
	}
	response, err := finder(word[0], Dict)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(response)
}

func loadJSON(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data []string
	err := decoder.Decode(&data)
	if err != nil {
		log.Println("Unexpected JSON structure:", err)
	} else {
		Dict.loadDictionary(data)
		log.Println("JSON loaded")
	}
	defer r.Body.Close()
}

func main() {
	handleRequests()
}
