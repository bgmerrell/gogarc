package dispatch

import (
	"fmt"
	"sync"

	"github.com/bgmerrell/gogarc/lib/command"
	"github.com/bgmerrell/gogarc/lib/game"
	hr "github.com/bgmerrell/gogarc/lib/handlerregistry"
)

// An EventDispatcher takes game commands and calls the appropriate handler
// for that command.  Because handlers may alter the state of the game,
// EventDispatcher only allows one handler to run at a time.
type EventDispatcher struct {
	mu sync.Mutex
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		mu: sync.Mutex{}}
}

func (d *EventDispatcher) Send(g *game.Game, cmd *command.Command, outputCh chan string) {
	h, err := hr.Registry.Find(cmd.Command)
	if err != nil {
		outputCh <- fmt.Sprintf("%s: Unknown command: %s",
			cmd.Nick, cmd.Command)
		return
	}
	d.mu.Lock()
	h.Handle(g, cmd, outputCh)
	d.mu.Unlock()
	close(outputCh)
}
