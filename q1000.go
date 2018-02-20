package main

import (
	"fmt"
	"os"

	"github.com/quakkels/q1000/bibleverse"
	"github.com/quakkels/q1000/bot"
	"github.com/quakkels/q1000/data"
	"github.com/quakkels/q1000/swearjar"
	"github.com/quakkels/q1000/troll"
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
		troll.NewBehavior(),
	}

	theBot, err := bot.NewBot(args[0], behaviors)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error when newing a bot. ", err)
		os.Exit(1)
	}

	theBot.Run()
}
