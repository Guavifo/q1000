package troll

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/nlopes/slack"

	"github.com/quakkels/q1000/bot"
)

// Behavior for troll joke messages
type Behavior struct {
	regex      *regexp.Regexp
	irishRegex *regexp.Regexp
}

// NewBehavior returns a new troll behavior
func NewBehavior() *Behavior {
	return &Behavior{
		regex:      regexp.MustCompile(`((is kind of )|(was kind of )|(it's kind of )|(that's kind of ))([a-z]?|[A-Z]?|( ))+(.?|!?)`),
		irishRegex: regexp.MustCompile(`^hamrock|[^shamrock]hamrock|hammock`),
	}
}

// Evaluate slack messages for trolly joke opportunities
func (b *Behavior) Evaluate(ev *slack.MessageEvent, bot *bot.Bot) error {
	if ev.BotID != "" {
		return nil
	}

	trollMessages := b.getTrollMessages(ev.Text)

	for _, msg := range trollMessages {
		bot.MessageChannel(ev.Channel, msg)
	}
	return nil
}

func (b *Behavior) getTrollMessages(message string) []string {
	messages := []string{}

	trollMessage := b.getTrollMessage(message)
	if trollMessage != "" {
		messages = append(messages, trollMessage)
	}

	now := time.Now()
	month := now.Month()
	if month == time.March {
		irishMessage := b.getIrishMessage(message)
		if irishMessage != "" {
			messages = append(messages, irishMessage)
		}
	}

	return messages
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

func (b *Behavior) getIrishMessage(message string) string {
	lowerMessage := strings.ToLower(message)
	results := b.irishRegex.FindAllString(lowerMessage, -1)
	if len(results) == 0 {
		return ""
	}
	return "Please refer to him as 'Shamrock' for the month."
}
