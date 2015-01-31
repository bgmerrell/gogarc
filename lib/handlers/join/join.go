package join

import (
	"fmt"

	"github.com/bgmerrell/gogarc/lib/command"
	"github.com/bgmerrell/gogarc/lib/game"
	hr "github.com/bgmerrell/gogarc/lib/handlerregistry"
)

const commandName = "join"

func init() {
	hr.Registry.Add(commandName, &JoinHandler{})
}

type JoinHandler struct{}

func (h *JoinHandler) Handle(g *game.Game, c *command.Command, outputCh chan string) {
	if g.InProgress() {
		outputCh <- fmt.Sprintf(
			"%s: Sorry, the game has already started", c.Nick)
		return
	}
	err := g.AddPlayer(c.Nick)
	if err != nil {
		outputCh <- err.Error()
		return
	}
	outputCh <- fmt.Sprintf("%s joined the game", c.Nick)
}
