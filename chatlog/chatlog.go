package chatlog

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const logFileName string = "slackmessages.log"

// Log is here and this doggone comment needs to be here or the linter will complain
type Log struct {
	logPath string
	file    *os.File
}

// Open the log
func Open(logPath string) (*Log, error) {
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		return nil, err
	}

	return &Log{
			logPath: logPath,
			file:    file,
		},
		nil
}

// OpenDefault will open the default log
func OpenDefault() (*Log, error) {
	return Open(logFileName)
}

// WriteLog will persist slack messages
func (l *Log) WriteLog(channel string, username string, message string, timeStamp string) error {
	if l.file == nil {
		err := errors.New("tried accessing log file that is nil")
		return err
	}

	entry := fmt.Sprintf("%s|%21s|%21s|\t%s\n", getTime(timeStamp), channel, username, message)
	_, err := l.file.Write([]byte(entry))

	return err
}

// Close the Log file
func (l *Log) Close() error {
	if l.file != nil {
		err := l.file.Close()
		l.file = nil
		return err
	}
	return nil
}

func getTime(ts string) string {
	parsingTime := strings.Split(ts, ".")

	seconds, err := strconv.ParseInt(parsingTime[0], 10, 64)
	if err != nil {
		panic(err)
	}

	nanoSeconds, err := strconv.ParseInt(parsingTime[1], 10, 64)
	if err != nil {
		panic(err)
	}

	timeStamp := time.Unix(seconds, nanoSeconds)
	return timeStamp.Format(time.RFC822)
}
