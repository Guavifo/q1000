package dictionary

import (
	"fmt"
	"regexp"

	"github.com/nlopes/slack"

	"github.com/quakkels/q1000/bot"
)

// Behavior for dictionary command parsing
type Behavior struct {
	regex *regexp.Regexp
}

// NewBehavior creates a new dictionary Behavior
func NewBehavior() *Behavior {
	return &Behavior{
		regex: regexp.MustCompile("(define:|Define:) ([a-z|A-Z|-|']+)"),
	}
}

// Evaluate slack messages to parse definition requests
func (b *Behavior) Evaluate(ev *slack.MessageEvent, bot *bot.Bot) error {
	if ev.BotID != "" {
		return nil
	}

	words := b.parseText(ev.Text)
	message := b.buildMessage(words)
	if message != "" {
		bot.MessageChannel(ev.Channel, message)
	}

	return nil
}

func (b *Behavior) parseText(text string) []string {
	matches := b.regex.FindAllString(text, -1)
	var words []string
	for _, match := range matches {
		words = append(words, match[7:])
	}

	return words
}

func (b *Behavior) buildMessage(words []string) string {
	message := ""
	length := len(words)
	for index, word := range words {
		message += fmt.Sprintf("<http://www.dictionary.com/browse/%s|%s>", word, word)
		if index < length-1 {
			message += " • "
		}
	}

	return message
}
