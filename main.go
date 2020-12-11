package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

var wordSlice []string = []string{"foobar", "aaBb", "Baba", "boofar", "test", "жито", "тижо", "спів", "півс"}

type dictionary map[string][]string

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

func (d dictionary) loadDictionary(input []string) {
	for _, word := range input {
		sortedWord := sortStringByCharacter(word)
		if _, found := d[sortedWord]; found {
			d[sortedWord] = append(d[sortedWord], word)
		} else {
			d[sortedWord] = []string{word}
		}
	}
}

func finder(word string, d dictionary) (result []string, err error) {
	sortedWord := sortStringByCharacter(word)
	if _, found := d[sortedWord]; found {
		result = d[sortedWord]
		return result, nil

	}
	err = fmt.Errorf("not found")
	return nil, err
}

func main() {
	dict := make(dictionary)
	var userInput string
	dict.loadDictionary(wordSlice)
	for k, v := range dict {
		fmt.Println(k, v)
	}
	fmt.Scanf("%s", &userInput)
	output, findError := finder(strings.ToLower(userInput), dict)
	if findError != nil {
		fmt.Printf("%#v\n", nil)
	} else {
		fmt.Printf("%v", output)
	}
}
