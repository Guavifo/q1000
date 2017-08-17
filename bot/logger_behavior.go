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
func NewLoggerBehavior(log *chatlog.Log) LoggerBehavior {
	return LoggerBehavior{logger: log}
}

// Evaluate will evalutate a slack message event and log appropriately
func (b LoggerBehavior) Evaluate(ev *slack.MessageEvent, rtm *slack.RTM) error {
	fmt.Printf("Message: %v\n", ev)

	if b.logger == nil {
		return errors.New("missing logger instance in LoggerBehavior")
	}

	err := b.logger.WriteLog(
		ev.Channel,
		ev.User,
		ev.Text,
		ev.Timestamp)
	if err != nil {
		return err
	}
	return nil
}
