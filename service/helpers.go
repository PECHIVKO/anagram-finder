package service

import (
	"sort"
	"unicode"

	"github.com/PECHIVKO/anagram-finder/models"
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

func isDuplicate(as models.AnagramSlice, word string) bool {
	for _, a := range as {
		if a == word {
			return true
		}
	}
	return false
}
