package bot

import (
	"github.com/nlopes/slack"
)

// Behavior interface is how to add behaviors that react to events
type Behavior interface {
	Evaluate(*slack.MessageEvent, *Bot) error
}
