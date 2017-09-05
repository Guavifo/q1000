package troll

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/nlopes/slack"

	"q1000/bot"
)

// Behavior for troll joke messages
type Behavior struct {
	regex *regexp.Regexp
}

// NewBehavior returns a new troll behavior
func NewBehavior() *Behavior {
	return &Behavior{
		regex: regexp.MustCompile(`((is kind of )|(was kind of )|(it's kind of )|(that's kind of ))([a-z]?|[A-Z]?|( ))+(.?|!?)`),
	}
}

// Evaluate slack messages for trolly joke opportunities
func (b *Behavior) Evaluate(ev *slack.MessageEvent, bot *bot.Bot) error {
	if ev.BotID != "" {
		return nil
	}

	trollMessage := b.getTrollMessage(ev.Text)
	if trollMessage == "" {
		return nil
	}

	bot.MessageChannel(ev.Channel, trollMessage)
	return nil
}

func (b *Behavior) getTrollMessage(message string) string {
	results := b.regex.FindAllString(message, -1)

	if len(results) == 0 {
		return ""
	}
	match := results[len(results)-1]
	match = strings.TrimPrefix(match, "is kind of ")
	match = strings.TrimPrefix(match, "was kind of ")
	match = strings.TrimPrefix(match, "it's kind of ")
	match = strings.TrimPrefix(match, "that's kind of ")

	return fmt.Sprintf("Your face is kind of %s", match)
}
