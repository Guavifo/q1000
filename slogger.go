package main

import (
	"fmt"
	"os"

	"slogger/bot"
	"slogger/chatlog"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "Token argument is missing. Cannot start.")
		os.Exit(1)
	}

	log, err := chatlog.OpenDefault()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot open log.", err)
		os.Exit(1)
	}
	defer log.Close()

	loggerBehavior := bot.NewLoggerBehavior(log)
	behaviors := []bot.Behavior{loggerBehavior}

	theBot, err := bot.NewBot(args[0], behaviors)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error when newing a bot. ", err)
		os.Exit(1)
	}

	theBot.Run(log)
}
