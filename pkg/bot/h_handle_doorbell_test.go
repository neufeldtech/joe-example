package bot

// This test aims to show the 'desired' behaviour: I'd like to register the handler HandleDoorbell but it only
// exists on my custom Bot struct and not on the joetest.Bot struct. I'm not sure how to accomplish this
// func TestHandleDoorbell(t *testing.T) {
// 	b := joetest.NewBot(t)
// 	b.Brain.RegisterHandler(b.HandleDoorbell)
// 	b.Start()
// 	u, _ := url.Parse("http://localhost:8080/doorbell")
// 	b.EmitSync(joehttp.RequestEvent{
// 		URL:    u,
// 		Method: "GET",
// 	})
// 	b.Stop()
// 	assert.Regexp(t, "^test > Someone is at the door!", b.ReadOutput())
// 	assert.True(t, gock.IsDone())
// }
