package bibleverse

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/nlopes/slack"

	"github.com/quakkels/q1000/bot"
)

// Behavior for bible verse parsing and posting
type Behavior struct {
	books []string
	regex *regexp.Regexp
}

// NewBehavior returns a new Bible verse behavior
func NewBehavior() *Behavior {
	return &Behavior{
		books: []string{"Genesis", "Exodus", "Leviticus", "Numbers", "Deuteronomy", "Joshua", "Judges", "Ruth", "1 Samuel", "2 Samuel", "1 Kings", "2 Kings", "1 Chronicles", "2 Chronicles", "Ezra", "Nehemiah", "Esther", "Job", "Psalms", "Psalm", "Proverbs", "Ecclesiastes", "Song of Solomon", "Isaiah", "Jeremiah", "Lamentations", "Ezekiel", "Daniel", "Hosea", "Joel", "Amos", "Obadiah", "Jonah", "Micah", "Nahum", "Habakkuk", "Zephaniah", "Haggai", "Zechariah", "Malachi", "Matthew", "Mark", "Luke", "John", "Acts", "Romans", "1 Corinthians", "2 Corinthians", "Galatians", "Ephesians", "Philippians", "Colossians", "1 Thessalonians", "2 Thessalonians", "1 Timothy", "2 Timothy", "Titus", "Philemon", "Hebrews", "James", "1 Peter", "2 Peter", "1 John", "2 John", "3 John", "Jude", "Revelation"},
		regex: regexp.MustCompile("(\\d\\s)?([a-zA-Z]+)\\s(\\d+):(\\d+)(-\\d+)?"),
	}
}

// Evaluate slack messages for Bible verses
func (b *Behavior) Evaluate(ev *slack.MessageEvent, bot *bot.Bot) error {
	if ev.BotID != "" {
		return nil
	}

	verses := b.parseText(ev.Text)
	if len(verses) == 0 {
		return nil
	}

	bot.MessageChannel(ev.Channel, b.buildMessage(verses))
	return nil
}

func (b *Behavior) parseText(text string) []string {
	matches := b.regex.FindAllString(text, -1)
	var verses []string
	for _, match := range matches {
		for _, book := range b.books {
			if strings.Contains(match, book) {
				isAdded := false
				for _, verse := range verses {
					if match == verse {
						isAdded = true
						break
					}
				}
				if isAdded == false {
					verses = append(verses, match)
					continue
				}
			}
		}
	}

	return verses
}

func (b *Behavior) buildMessage(verses []string) string {
	message := ""
	length := len(verses)
	for index, verse := range verses {
		message += fmt.Sprintf("<https://www.biblegateway.com/passage/?search=%s|%s>", verse, verse)
		if index < length-1 {
			message += " • "
		}
	}

	return message
}
