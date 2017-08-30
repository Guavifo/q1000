package swearjar

import (
	"fmt"
	"os"
	"strings"

	"github.com/nlopes/slack"

	"q1000/bot"
	"q1000/data"
)

// Behavior handles tracking swears and asking for fake payments
type Behavior struct {
	swears []string
	store  *data.Store
}

// NewBehavior returns a new swearjar behavior
func NewBehavior(store *data.Store) *Behavior {
	return &Behavior{
		swears: []string{
			"fuck",
			"shit",
			"damn",
		},
		store: store,
	}
}

// Evaluate slack messages for swears
func (b *Behavior) Evaluate(ev *slack.MessageEvent, bot *bot.Bot) error {
	if ev.BotID != "" {
		return nil
	}

	text := strings.ToLower(ev.Text)
	for _, swear := range b.swears {
		if strings.Contains(text, swear) {
			err := b.incrementSwearCount(ev.User)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Could not write to swearjar: ", err)
			}
			message := bot.GetUsername(ev.User) + ", you need to make a deposit to the swear jar. Pay up!"
			bot.MessageChannel(ev.Channel, message)
			return nil
		}
	}

	if strings.Contains(text, "pay up") {
		bot.MessageChannel(ev.Channel, "All debts are payed.")
	}
	return nil
}

func (b *Behavior) incrementSwearCount(userID string) error {
	result, err := b.store.Query("select swearcount from swearjar where userid=?", userID)
	if err != nil {
		return err
	}

	count := 0
	if result.Next() {
		err = result.Scan(&count)
		if err != nil {
			return err
		}
	}
	count++

	stmt, err := b.store.Prepare(
		`insert into swearjar 
		(userid, swearcount) values (?, ?) 
		on duplicate key update 
			userid = values(userid), 
			swearcount = values(swearcount)`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(userID, count)
	return err
}
