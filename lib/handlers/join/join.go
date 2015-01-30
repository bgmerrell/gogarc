package join

import (
	"fmt"

	"github.com/bgmerrell/gogarc/lib/command"
	hr "github.com/bgmerrell/gogarc/lib/handlerregistry"
)

const commandName = "join"

func init() {
	hr.Registry.Add(commandName, &JoinHandler{})
}

type JoinHandler struct{}

func (h *JoinHandler) Handle(c *command.Command, outputCh chan string) {
	outputCh <- fmt.Sprintf("%s joined the game", c.Nick)
}
