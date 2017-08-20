package bot

import (
	"fmt"
	"os"

	"github.com/nlopes/slack"
)

// Bot represents an instance of a bot
type Bot struct {
	api       *slack.Client
	rtm       *slack.RTM
	behaviors []Behavior

	channels map[string]string
	users    map[string]string
}

// NewBot will make an instance of a bot
func NewBot(token string, behaviors []Behavior) (*Bot, error) {
	if token == "" {
		return nil, fmt.Errorf("token was empty")
	}

	api := slack.New(token)
	rtm := api.NewRTM()

	return &Bot{
			api:       api,
			rtm:       rtm,
			behaviors: behaviors,
		},
		nil
}

// Run the bot
func (b *Bot) Run() {
	go b.rtm.ManageConnection()

	for msg := range b.rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.ConnectedEvent:
		case *slack.MessageEvent:
			for _, beh := range b.behaviors {
				err := beh.Evaluate(ev, b)
				if err != nil {
					fmt.Fprintln(os.Stderr, "Error when evaluating a behavior. ", err)
				}
			}

		case *slack.RTMError:
			fmt.Printf("Error: %s\n", ev.Error())

		case *slack.InvalidAuthEvent:
			fmt.Printf("Invalid credentials")
			return
		}
	}
}

func (b *Bot) setChannelList() {
	chans, err := b.api.GetChannels(true)
	if err != nil {
		fmt.Println("Error when getting channels. ", err)
		return
	}

	b.channels = make(map[string]string)
	for _, ch := range chans {
		b.channels[ch.ID] = ch.Name
		fmt.Printf("id: %s, name: %s\n", ch.ID, ch.Name)
	}
}

// GetChannel gets the channel name that matches the ID
func (b *Bot) GetChannel(id string) string {
	if channelName, ok := b.channels[id]; ok {
		return channelName
	}

	b.setChannelList()
	if channelName, ok := b.channels[id]; ok {
		return channelName
	}

	return id
}

// GetUsername gets the username for the user
func (b *Bot) GetUsername(user string) string {
	if username, ok := b.users[user]; ok {
		return username
	}

	userInfo, err := b.api.GetUserInfo(user)
	if err != nil {
		return user
	}

	if b.users == nil {
		b.users = make(map[string]string)
	}

	b.users[user] = userInfo.Name
	return userInfo.Name
}
