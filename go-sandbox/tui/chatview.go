package tui

import (
	"fmt"
	"github.com/marcusolsson/tui-go"
)

type SubmitMessageHandler func(string)

type ChatView struct {
	tui.Box
	frame    *tui.Box
	history  *tui.Box
	onSubmit SubmitMessageHandler
}

func NewChatView() *ChatView {
	view := &ChatView{}

	// ref: https://github.com/marcusolsson/tui-go/blob/master/example/chat/main.go
	view.history = tui.NewVBox()
	
	historyScroll := tui.NewScrollArea(view.history)
	historyScroll.SetAutoscrollToBottom(true)

	historyBox := tui.NewVBox(historyScroll)
	historyBox.SetBorder(true)

	input := tui.NewEntry()
	input.SetFocused(true)
	input.SetSizePolicy(tui.Expanding, tui.Maximum)

	input.OnSubmit(func(e *tui.Entry) {
		if e.Text() != "" {
			if view.onSubmit != nil {
				view.onSubmit(e.Text()) // send tui text as a message.
			}
			e.SetText("")
		}
	})

	inputBox := tui.NewHBox(input)
	inputBox.SetBorder(true)
	inputBox.SetSizePolicy(tui.Expanding, tui.Maximum)

	view.frame = tui.NewVBox(
		historyBox,
		inputBox,
	)

	view.frame.SetBorder(false)
	view.Append(view.frame)

	return view
}

// @param handler: function(string) which executed with new tui text
func (c *ChatView) OnSubmit(handler SubmitMessageHandler) {
	c.onSubmit = handler
}

func (c *ChatView) AddMessage(user string, nsg string) {
	c.history.Append(
		tui.NewHBox(
			tui.NewLabel(fmt.Sprintf("%v: %v", user, msg))
		),
	)
}
