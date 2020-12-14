package service

import (
	"sort"
	"unicode"
)

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

// isDuplicate checks if word is already in Dictionary
func isDuplicate(as AnagramSlice, word string) bool {
	for _, a := range as {
		if a == word {
			return true
		}
	}
	return false
}
