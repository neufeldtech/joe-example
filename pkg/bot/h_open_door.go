package bot

import (
	"net/http"

	"github.com/go-joe/joe"
)

// OpenDoor makes a call to the door solenoid server to open it
// Since this method is 'standalone' we are able to register it with joetest for testing
// However, one shortcoming is that we don't have access to the
// parent bot object here for logging, or other custom functionality like using the Slack client
func OpenDoor(msg joe.Message) error {
	resp, err := http.Get("http://httpbin.org/get?action=OpenDoor")
	if err != nil || resp.StatusCode != 200 {
		msg.Respond("I had an error opening the door %v HTTP %v", err, resp.StatusCode)
		return err
	}
	msg.Respond("I opened the door!")
	return nil
}
