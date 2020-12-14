package service

import (
	"fmt"
	"strings"
)

type AnagramSlice []string

type Dictionary map[string]AnagramSlice

// LoadDictionary loads words to Dictionary
func (d Dictionary) LoadDictionary(input []string) {
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

// FindAnagrams finds all anagrams in Dictionary for word
func (d Dictionary) FindAnagrams(word string) (result AnagramSlice, err error) {
	if len(d) == 0 {
		err = fmt.Errorf("Dictionary is empty. Please load words.")
		return result, err
	}
	sortedWord := strings.ToLower(sortStringByCharacter(word))
	if _, found := d[sortedWord]; found {
		result = d[sortedWord]
		return result, nil

	}
	err = fmt.Errorf("There is no anagrams for word \"%s\" in dictionary", word)
	return nil, err
}
