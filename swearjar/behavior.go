package swearjar

import (
	"strings"

	"github.com/nlopes/slack"

	"q1000/bot"
)

// Behavior handles tracking swears and asking for fake payments
type Behavior struct {
	swears []string
}

// NewBehavior returns a new swearjar behavior
func NewBehavior() *Behavior {
	return &Behavior{
		swears: []string{
			"fuck",
			"shit",
			"damn",
		},
	}
}

// Evaluate slack messages for swears
func (b *Behavior) Evaluate(ev *slack.MessageEvent, bot *bot.Bot) error {
	if ev.BotID != "" {
		return nil
	}

	text := strings.ToLower(ev.Text)
	for _, swear := range b.swears {
		if strings.Contains(text, swear) {
			message := bot.GetUsername(ev.User) + ", you need to make a deposit to the swear jar. Pay up!"
			bot.MessageChannel(ev.Channel, message)
			return nil
		}
	}

	if strings.Contains(text, "pay up") {
		bot.MessageChannel(ev.Channel, "All debts are payed.")
	}
	return nil
}