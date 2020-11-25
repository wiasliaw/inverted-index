package script

import (
	"strings"
	"unicode"

	snowballeng "github.com/kljensen/snowball/english"
)

func Analyze(text string) []string {
	tokens := tokenize(text)
	tokens = lowercaseFilter(tokens)
	tokens = commenWordFilter(tokens)
	tokens = stemmingFilter(tokens)
	return tokens
}

func tokenize(text string) []string {
	return strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
}

func lowercaseFilter(tokens []string) []string {
	r := make([]string, len(tokens))
	for i, token := range tokens {
		r[i] = strings.ToLower(token)
	}
	return r
}

func commenWordFilter(tokens []string) []string {
	r := make([]string, len(tokens))
	for _, token := range tokens {
		if _, isCommon := commonDict[token]; !isCommon {
			r = append(r, token)
		}
	}
	return r
}

func stemmingFilter(tokens []string) []string {
	r := make([]string, len(tokens))
	for i, token := range tokens {
		r[i] = snowballeng.Stem(token, false)
	}
	return r
}
