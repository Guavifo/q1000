package swearjar

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/nlopes/slack"

	"github.com/quakkels/q1000/bot"
	"github.com/quakkels/q1000/data"
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
			"fuuk",
			"asshat",
			"shit",
			"damn",
			"wtf",
			"crappy",
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
			var message = ""
			swearCount, err := b.incrementSwearCount(ev.User)
			username := bot.GetUsername(ev.User)
			if time.Now().Month() == time.March {
				if strings.ToLower(username) == "hammrock" {
					username = "Shamrock"
				}
			}

			if err != nil {
				fmt.Fprintln(os.Stderr, "Could not write to swearjar: ", err)
				message = fmt.Sprintf("%v, you owe swearbucks to the swear jar. Pay up!", username)
			} else {
				message = fmt.Sprintf("%v, you owe %v swearbucks to the swear jar. Pay up!", username, swearCount)
			}

			bot.MessageChannel(ev.Channel, message)
			return nil
		}
	}

	if strings.HasPrefix(strings.ToLower(text), "pay swearbuck") {
		count, _ := b.decrementSwearCount(ev.User)
		if count == 0 {
			bot.MessageChannel(ev.Channel, "All debts are payed.")
		} else {
			bot.MessageChannel(ev.Channel, fmt.Sprintf("Payment accepted. You now owe %d swearbucks.", count))
		}
	}
	return nil
}

func (b *Behavior) incrementSwearCount(userID string) (int, error) {
	result, err := b.store.Query("select swearcount from swearjar where userid=?", userID)
	if err != nil {
		return 0, err
	}

	count := 0
	if result.Next() {
		err = result.Scan(&count)
		if err != nil {
			return 0, err
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
		return 0, err
	}

	_, err = stmt.Exec(userID, count)
	return count, err
}

func (b *Behavior) decrementSwearCount(userID string) (int, error) {
	rows, err := b.store.Query(
		"select swearcount from swearjar where userid=?",
		userID)
	if err != nil {
		return 0, err
	}

	count := 0
	if rows.Next() {
		rows.Scan(&count)
		if count > 0 {
			count--
			stmt, err := b.store.Prepare("update swearjar set swearcount=? where userid=?")
			if err != nil {
				return count, err
			}
			_, err = stmt.Exec(count, userID)
			if err != nil {
				return count, err
			}
		}
	}

	return count, nil
}
