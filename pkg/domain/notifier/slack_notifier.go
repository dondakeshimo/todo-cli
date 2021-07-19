package notifier

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// SlackNotifier is a struct that notify with slack.
type SlackNotifier struct {
	webhookURL string
	mentionTo  string
}

type payload struct {
	Text      string `json:"text"`
	Username  string `json:"username"`
	IconEmoji string `json:"icon_emoji"`
}

const (
	username  = "todo-cli"
	iconEmoji = "memo"
)

// NewSlackNotifier is a constructor for SlackNotifier.
func NewSlackNotifier(webhookURL string, mentionTo string) (*SlackNotifier, error) {
	if webhookURL == "" {
		return nil, fmt.Errorf("webhookURL is empty")
	}
	return &SlackNotifier{webhookURL: webhookURL, mentionTo: mentionTo}, nil
}

// Push is a function that push notification.
func (sn *SlackNotifier) Push(r Request) (string, error) {
	text := buildText(r.Contents, sn.mentionTo)
	pl, err := json.Marshal(payload{Text: text, Username: username, IconEmoji: iconEmoji})
	if err != nil {
		return "", err
	}

	res, err := http.PostForm(sn.webhookURL, url.Values{"payload": {string(pl)}})
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	return "", nil
}

func buildText(contents string, mentionTo string) string {
	text := ""

	if mentionTo != "" {
		text = fmt.Sprintf("@%s\n", mentionTo)
	}

	text += contents

	return text
}
