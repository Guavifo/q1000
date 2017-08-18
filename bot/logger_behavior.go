package bot

import (
	"errors"
	"fmt"

	"github.com/nlopes/slack"

	"slogger/chatlog"
)

// LoggerBehavior handles what get logged from chat
type LoggerBehavior struct {
	logger *chatlog.Log
}

// NewLoggerBehavior will make a new behavior for logging
func NewLoggerBehavior(log *chatlog.Log) (*LoggerBehavior, error) {
	if log == nil {
		return nil, errors.New("Cannot make LoggerBehavior with a nil chatlog.log")
	}
	return &LoggerBehavior{logger: log}, nil
}

// Evaluate will evalutate a slack message event and log appropriately
func (b LoggerBehavior) Evaluate(ev *slack.MessageEvent, bot *Bot) error {
	fmt.Printf("Message: %v\n", ev)

	if b.logger == nil {
		return errors.New("missing logger instance in LoggerBehavior")
	}

	err := b.logger.WriteLog(
		bot.getChannel(ev.Channel),
		bot.getUsername(ev.User),
		ev.Text,
		ev.Timestamp)
	if err != nil {
		return err
	}
	return nil
}
