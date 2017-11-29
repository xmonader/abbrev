package abbrev

import (
	"regexp"
	"testing"
)

func TestReturnsDataWithOne(t *testing.T) {
	testStrings := []string{"ruby"}
	abbreved := Abbrev(testStrings)
	if len(abbreved) == 0 {
		t.Errorf("abbreved data is empty.")
	}
}

func TestReturnsDataWithMultiple(t *testing.T) {
	testStrings := []string{"ruby", "python", "rules"}
	abbreved := Abbrev(testStrings)
	if len(abbreved) == 0 {
		t.Errorf("abbreved data is empty.")
	}
}
func TestReturnsDataUnique(t *testing.T) {
	testStrings := []string{"ruby", "python", "rules"}
	abbreved := Abbrev(testStrings)
	if _, ok := abbreved["ru"]; ok {
		t.Errorf("ru isn't unique between ruby and rul.")
	}
}

func TestAbbrevMatching(t *testing.T) {
	testStrings := []string{"ruby", "python", "rules"}
	pat := "p+"
	abbreved := AbbrevMatching(testStrings, pat)
	for k := range abbreved {
		if matched, _ := regexp.MatchString(pat, k); !matched {
			t.Errorf("key %s is added even if it doesn't match the pattern %s", k, pat)
		}
	}
}
