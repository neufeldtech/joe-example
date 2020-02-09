package bot

import (
	"testing"

	"github.com/go-joe/joe"
	"github.com/go-joe/joe/joetest"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestOpenDoor(t *testing.T) {
	defer gock.Off()
	gock.Observe(gock.DumpRequest)
	gock.New("http://httpbin.org").
		Get("/get").
		MatchParam("action", "OpenDoor").
		Reply(200)

	bot := joetest.NewBot(t)
	bot.Respond("open sesame", OpenDoor)
	bot.Start()
	bot.EmitSync(joe.ReceiveMessageEvent{
		Text: "open sesame",
	})
	bot.Stop()
	assert.Regexp(t, "^test > I opened the door!", bot.ReadOutput())
	assert.True(t, gock.IsDone())
}
