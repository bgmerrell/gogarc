package start

import (
	"fmt"

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
	if g.InProgress() {
		outputCh <- fmt.Sprintf(
			"%s: There is already a game in progress", c.Nick)
		return
	}
	msg, err := g.Start()
	if err != nil {
		outputCh <- err.Error()
		return
	}
	outputCh <- msg
}
