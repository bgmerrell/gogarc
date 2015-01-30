package start

import (
	"github.com/bgmerrell/gogarc/lib/command"
	"github.com/bgmerrell/gogarc/lib/game"
	hr "github.com/bgmerrell/gogarc/lib/handlerregistry"
)

const commandName = "start"

func init() {
	hr.Registry.Add(commandName, &StatsHandler{})
}

type StatsHandler struct{}

func (h *StatsHandler) Handle(g *game.Game, c *command.Command, outputCh chan string) {
	msg, err := g.Start()
	if err != nil {
		outputCh <- err.Error()
		return
	}
	outputCh <- msg
}
