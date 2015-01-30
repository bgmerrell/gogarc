package stats

import (
	"github.com/bgmerrell/gogarc/lib/command"
	"github.com/bgmerrell/gogarc/lib/game"
	hr "github.com/bgmerrell/gogarc/lib/handlerregistry"
)

const commandName = "stats"

func init() {
	hr.Registry.Add(commandName, &StatsHandler{})
}

type StatsHandler struct{}

func (h *StatsHandler) Handle(g *game.Game, c *command.Command, outputCh chan string) {
	stats, err := g.PlayerStats(c.Nick)
	if err != nil {
		outputCh <- err.Error()
		return
	}
	outputCh <- stats
}
