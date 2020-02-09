package bot

import (
	"errors"
	"fmt"

	joehttp "github.com/go-joe/http-server"
	"github.com/go-joe/joe"
	slackAPI "github.com/nlopes/slack"
)

// Modules creates a list of joe.Modules that can be used with this configuration.
func (conf Config) Modules() []joe.Module {
	var modules []joe.Module
	modules = append(modules, joehttp.Server(conf.HTTPListen))
	return modules
}

func (conf Config) Validate() error {
	if conf.HTTPListen == "" {
		return errors.New("missing HTTP listen address")
	}
	return nil
}

func NewBot(conf Config) (*Bot, error) {
	if err := conf.Validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %v", err)
	}

	bot := &Bot{
		Bot:   joe.New("Joe Example", conf.Modules()...),
		slack: slackAPI.New(conf.SlackToken, slackAPI.OptionDebug(conf.SlackAPIDebug)),
	}
	return bot, nil
}
