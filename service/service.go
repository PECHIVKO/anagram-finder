package service

import (
	"fmt"
	"strings"

	"github.com/PECHIVKO/anagram-finder/models"
)

func (d models.Dictionary) LoadDictionary(input []string) {
	for _, word := range input {
		sortedWord := sortStringByCharacter(word)
		if _, found := d[sortedWord]; found {
			if !isDuplicate(d[sortedWord], word) {
				d[sortedWord] = append(d[sortedWord], word)
			}
		} else {
			d[sortedWord] = models.AnagramSlice{word}
		}
	}
}

func (d models.Dictionary) Finder(word string) (result models.AnagramSlice, err error) {
	sortedWord := strings.ToLower(sortStringByCharacter(word))
	if _, found := d[sortedWord]; found {
		result = d[sortedWord]
		return result, nil

	}
	err = fmt.Errorf("There is no anagrams for word \"%s\" in dictionary", word)
	return nil, err
}
