package main

import (
	"fmt"
	"github.com/nlopes/slack"
	"os"
)

func main() {

	args := os.Args[1:]
	if len(args) != 1 {
		panic("Token argument is missing. Cannot start.")
	}

	api := slack.New("token")

	groups, err := api.GetGroups(false)
	if err != nil {
		fmt.Println("Error getting groups. ", err)
		return
	}

	if len(groups) == 0 {
		fmt.Println("No groups found")
	}
	for _, group := range groups {
		fmt.Printf("id: %s, name: %s", group.ID, group.Name)
	}

	channels, err := api.GetChannels(true)
	if err != nil {
		fmt.Println("Error getting channels. ", err)
	}

	if len(channels) == 0 {
		fmt.Println("No channels found")
	}
	for _, channel := range channels {
		fmt.Printf("id:%s, name:%s\n", channel.ID, channel.Name)
	}
}
