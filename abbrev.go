package abbrev

import (
	"regexp"
)

// Abbrev calculates abbreviations of words list.
// It returns unique abbreviations of words passed in words list.
// Example: ruby -> map[r: ruby, ru:ruby, rub:ruby, ruby:ruby].
// Example: ruby, rule ->  map[rub:ruby, ruby:ruby, rul:rule, rule:rule] and the non unique ones are removed. (r, ru)
func Abbrev(words []string) map[string]string {
	uniqAbbrevs := make(map[string]string)
	for _, word := range words {
		i := 0
		for i <= len(word) {
			maybeAbbrev := word[0:i]
			if _, ok := uniqAbbrevs[maybeAbbrev]; ok {
				// key already exists. remove it.
				delete(uniqAbbrevs, maybeAbbrev)
			} else {
				uniqAbbrevs[maybeAbbrev] = word

			}
			i++
		}
	}
	return uniqAbbrevs
}

// AbbrevMatching calculates abbreviations of words list.
// It returns unique abbreviations of words passed in words list for every word that matches pattern pat.
// Example: ruby -> map[r: ruby, ru:ruby, rub:ruby, ruby:ruby].
// Example: ruby, rule ->  map[rub:ruby, ruby:ruby, rul:rule, rule:rule] and the non unique ones are removed. (r, ru)
// Example: ruby, python and pat is "p+" you will get map[p:python, py:python, pyt:python, pyth:python, pytho: python, python:python]
func AbbrevMatching(words []string, pat string) map[string]string {
	// First filter words to the ones that matches pat.
	filtered := make([]string, len(words))
	for _, word := range words {
		matched, _ := regexp.MatchString(pat, word)
		if matched {
			filtered = append(filtered, word)
		}
	}
	// then return Abbrev of the filtered words.
	return Abbrev(filtered)
}
