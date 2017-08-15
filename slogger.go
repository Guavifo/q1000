package main

import (
	"fmt"
	"github.com/nlopes/slack"
	"os"
	"slogger/chatlog"
)

var (
	api      *slack.Client
	channels map[string]string
	users    map[string]string
)

func main() {

	args := os.Args[1:]
	if len(args) != 1 {
		panic("Token argument is missing. Cannot start.")
	}

	api = slack.New(args[0])

	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.ConnectedEvent:
			fmt.Println("Info: ", ev.Info)
			fmt.Println("Connection Counter: ", ev.ConnectionCount)
			rtm.SendMessage(rtm.NewOutgoingMessage("HI!!!.", "#general"))

		case *slack.MessageEvent:
			fmt.Printf("Message: %v\n", ev)
			err := chatlog.WriteLog(
				getChannel(ev.Channel),
				getUsername(ev.User),
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

func setChannelList() {
	chans, err := api.GetChannels(true)
	if err != nil {
		fmt.Println("Error when getting channels. ", err)
		return
	}

	channels = make(map[string]string)
	for _, ch := range chans {
		channels[ch.ID] = ch.Name
		fmt.Printf("id: %s, name: %s\n", ch.ID, ch.Name)
	}
}

func getChannel(id string) string {
	if channelName, ok := channels[id]; ok {
		return channelName
	}

	setChannelList()
	if channelName, ok := channels[id]; ok {
		return channelName
	}

	return id
}

func getUsername(user string) string {

	if username, ok := users[user]; ok {
		return username
	}

	userInfo, err := api.GetUserInfo(user)
	if err != nil {
		return user
	}

	if users == nil {
		users = make(map[string]string)
	}

	users[user] = userInfo.Name
	return userInfo.Name
}
