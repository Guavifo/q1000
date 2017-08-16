package bot

import (
	"fmt"

	"github.com/nlopes/slack"
)

// Bot represents an instance of a _bot_
type Bot struct {
	api *slack.Client
	rtm *slack.RTM
}

// NewBot will make an instance of... a _new... bot_
func NewBot(token string) (*Bot, error) {
	if token == "" {
		return nil, fmt.Errorf("token was empty")
	}

	api := slack.New(token)
	return &Bot{
			api: api,
			rtm: api.NewRTM(),
		},
		nil
}
