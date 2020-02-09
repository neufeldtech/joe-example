package bot

import (
	"github.com/go-joe/joe"
	"github.com/nlopes/slack"
)

type Config struct {
	SlackToken    string //Falls back to CLI if this is blank :)
	HTTPListen    string
	SlackAPIDebug bool
}
type Bot struct {
	*joe.Bot        // Anonymously embed the joe.Bot type so we can use its functions easily.
	conf     Config // You can keep other fields here as well.
	slack    *slack.Client
}
