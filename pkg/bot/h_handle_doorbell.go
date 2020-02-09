package bot

import (
	"context"

	joehttp "github.com/go-joe/http-server"
)

// HandleDoorbell listens for GET requests at the doorbell path and reports that someone is at the door
func (b *Bot) HandleDoorbell(ctx context.Context, evt joehttp.RequestEvent) error {
	if evt.URL.Path == "/doorbell" {
		b.Say("#general", "Someone is at the door!")
		b.Logger.Info("I can log using the Bot logger")
	}
	return nil
}
