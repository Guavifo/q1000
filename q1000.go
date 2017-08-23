package main

import (
	"fmt"
	"os"

	"q1000/bibleverse"
	"q1000/bot"
	"q1000/swearjar"
)

func main() {
	args := os.Args[1:]
	if len(args) > 1 {
		fmt.Fprintln(os.Stderr, "Too many arguments. Expect only token.")
		os.Exit(1)
	}
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "Token argument is missing. Cannot start.")
		os.Exit(1)
	}

	behaviors := []bot.Behavior{
		swearjar.NewBehavior(),
		bibleverse.NewBehavior(),
	}

	theBot, err := bot.NewBot(args[0], behaviors)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error when newing a bot. ", err)
		os.Exit(1)
	}

	theBot.Run()
}
