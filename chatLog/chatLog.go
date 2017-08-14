package chatLog

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const logFileName string = "slackmessages.log"

// WriteLog will persist slack messages
func WriteLog(channel string, username string, message string, timeStamp string) error {
	entry := fmt.Sprintf("%s|%21s|%21s|\t%s\n", getTime(timeStamp), channel, username, message)

	file, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		return err
	}

	defer file.Close()

	byteSlice := []byte(entry)
	_, err = file.Write(byteSlice)

	return err
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
