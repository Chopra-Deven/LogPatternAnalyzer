package utils

import (
	"strings"
	"unicode"
)

type Tokenizer struct {
	Tokens                   []string
	startIndices, endIndices []int
	Counts                   int
}

func NewTokenizer(tokens int) *Tokenizer {
	return &Tokenizer{
		Tokens:       make([]string, tokens),
		startIndices: make([]int, tokens),
		endIndices:   make([]int, tokens),
	}
}

func (tokenizer *Tokenizer) Tokenize(word string) []string {

	return tokenizer.tokenize(word, func(r rune) bool {
		return unicode.IsSpace(r)
	})
}

func (tokenizer *Tokenizer) tokenize(word string, f func(rune) bool) []string {

	count := 0

	start := -1 // valid span start if >= 0

	for end, runes := range word {

		if f(runes) {

			if start >= 0 {

				tokenizer.startIndices[count] = start

				tokenizer.endIndices[count] = end

				count++

				start = ^start
			}
		} else {

			if start < 0 {

				start = end
			}
		}
	}

	// Last field might end at EOF.
	if start >= 0 {

		tokenizer.startIndices[count] = start

		tokenizer.endIndices[count] = len(word)

		count++
	}

	// Create strings from recorded field indices.

	tokenizer.Counts = 0

	for i := 0; i < count; i++ {

		if tokenizer.endIndices[i]-tokenizer.startIndices[i] > 0 {

			// Check if the token contains at least one letter

			tokenizer.Tokens[tokenizer.Counts] = strings.ToLower(word[tokenizer.startIndices[i]:tokenizer.endIndices[i]])

			tokenizer.Counts++
		}
	}

	return tokenizer.Tokens[:tokenizer.Counts]
}
