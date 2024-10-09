package utils

import (
	"strings"
	"sync"
	"time"
	"unicode"
)

var (
	SubscriberRCVTimeout = time.Second * 5

	config MotadataMap

	WaitGroup = sync.WaitGroup{}
)

type Tokenizer struct {
	Tokens []string

	startIndex, Counts, index int
}

func NewTokenizer(tokens int) *Tokenizer {
	return &Tokenizer{
		Tokens: make([]string, tokens),
	}
}

// tokenizer methods

func (tokenizer *Tokenizer) Split(value, delimiter string) {

	tokenizer.startIndex, tokenizer.Counts, tokenizer.index = 0, 0, 0

	for {

		tokenizer.index = strings.Index(value[tokenizer.startIndex:], delimiter)

		if tokenizer.index == -1 {
			break
		}

		tokenizer.Tokens[tokenizer.Counts] = value[tokenizer.startIndex : tokenizer.startIndex+tokenizer.index]

		tokenizer.Counts++

		tokenizer.startIndex += tokenizer.index + len(delimiter)
	}

	tokenizer.Tokens[tokenizer.Counts] = value[tokenizer.startIndex:]

	tokenizer.Counts++
}

func Tokenize(word string, tokens []string, startIndices []int, endIndices []int) []string {

	return tokenize(word, tokens, startIndices, endIndices, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	})
}

func tokenize(word string, tokens []string, startIndices []int, endIndices []int, f func(rune) bool) []string {

	count := 0

	start := -1 // valid span start if >= 0

	for end, runes := range word {

		if f(runes) {

			if start >= 0 {

				startIndices[count] = start

				endIndices[count] = end

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

		startIndices[count] = start

		endIndices[count] = len(word)

		count++
	}

	// Create strings from recorded field indices.

	index := 0

	for i := 0; i < count; i++ {

		if endIndices[i]-startIndices[i] > 0 {

			// Check if the token contains at least one letter

			if strings.IndexFunc(word[startIndices[i]:endIndices[i]], unicode.IsLetter) >= 0 {

				tokens[index] = strings.ToLower(word[startIndices[i]:endIndices[i]])

				index++
			}
		}
	}

	return tokens[:index]
}

func InitConfig(configurations MotadataMap) {

	config = configurations

	SetLogLevel(config.GetIntValue(SystemLogLevel))
}

func GetMaxWorker() int {

	if config.Contains("max.worker") {

		return config.GetIntValue("max.worker")
	}

	return 10
}

func GetMaxChannelBuffer() int {

	if config.Contains("max.channel.buffer") {

		return config.GetIntValue("max.channel.buffer")
	}

	return 1000000
}

func GetHost() MotadataString {
	return "localhost"
}

func GetSubscriberPort() MotadataString {

	if config.Contains("event.subscriber.port") {

		return config.GetMotadataStringValue("event.subscriber.port")
	}
	return "8888"
}

func GetPprofPort() string {

	if config.Contains("pprof.port") {

		return config.GetStringValue("pprof.port")
	}
	return "6161"
}

func GetFlushTimer() int {

	if config.Contains("log.pattern.flush.timer.seconds") {

		return config.GetIntValue("log.pattern.flush.timer.seconds")
	}
	return 120
}

func GetPublisherPort() MotadataString {

	if config.Contains("event.publisher.port") {

		return config.GetMotadataStringValue("event.publisher.port")
	}
	return "8889"
}
