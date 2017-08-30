package main

import (
	"fmt"
	"os"

	"q1000/bibleverse"
	"q1000/bot"
	"q1000/data"
	"q1000/swearjar"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Fprintln(os.Stderr, "Argument is missing. Cannot start. Need: <token> <dataResourceName>")
		os.Exit(1)
	}

	store, err := data.OpenStore(args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error opening data store. ", err)
		os.Exit(1)
	}
	defer store.Close()

	behaviors := []bot.Behavior{
		swearjar.NewBehavior(store),
		bibleverse.NewBehavior(),
	}

	theBot, err := bot.NewBot(args[0], behaviors)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error when newing a bot. ", err)
		os.Exit(1)
	}

	theBot.Run()
}
