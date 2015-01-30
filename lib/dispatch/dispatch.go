package dispatch

import (
	"fmt"

	"github.com/bgmerrell/gogarc/lib/command"
	hr "github.com/bgmerrell/gogarc/lib/handlerregistry"
)

type EventDispatcher struct {
	Output chan string
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		Output: make(chan string)}
}

func (d *EventDispatcher) Send(cmd *command.Command) {
	h, err := hr.Registry.Find(cmd.Command)
	if err != nil {
		d.Output <- fmt.Sprintf("%s: Unknown command: %s",
			cmd.Nick, cmd.Command)
		return
	}
	h.Handle(cmd, d.Output)
	close(d.Output)
}
