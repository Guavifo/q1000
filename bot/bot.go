package bot

import (
	"fmt"

	"github.com/nlopes/slack"

	"slogger/chatlog"
)

// Bot represents an instance of a bot
type Bot struct {
	api *slack.Client
	rtm *slack.RTM

	channels map[string]string
	users    map[string]string
}

// NewBot will make an instance of a bot
func NewBot(token string) (*Bot, error) {
	if token == "" {
		return nil, fmt.Errorf("token was empty")
	}

	api := slack.New(token)
	rtm := api.NewRTM()

	return &Bot{
			api: api,
			rtm: rtm,
		},
		nil
}

// Run the bot
func (b *Bot) Run(log *chatlog.Log) {
	go b.rtm.ManageConnection()

	for msg := range b.rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.ConnectedEvent:
			fmt.Println("Info: ", ev.Info)
			fmt.Println("Connection Counter: ", ev.ConnectionCount)
			b.rtm.SendMessage(b.rtm.NewOutgoingMessage("HI!!!.", "#general"))

		case *slack.MessageEvent:
			fmt.Printf("Message: %v\n", ev)
			err := log.WriteLog(
				b.getChannel(ev.Channel),
				b.getUsername(ev.User),
				ev.Text,
				ev.Timestamp)
			if err != nil {
				fmt.Println("Error in main writeing message log: ", err)
				recover()
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

func (b *Bot) getChannel(id string) string {
	if channelName, ok := b.channels[id]; ok {
		return channelName
	}

	b.setChannelList()
	if channelName, ok := b.channels[id]; ok {
		return channelName
	}

	return id
}

func (b *Bot) getUsername(user string) string {
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
