package stats

import (
	"fmt"
	"sort"
	"strings"

	"github.com/bgmerrell/gogarc/lib/command"
	"github.com/bgmerrell/gogarc/lib/game"
	hr "github.com/bgmerrell/gogarc/lib/handlerregistry"
)

const commandName = "stats"

func init() {
	hr.Registry.Add(commandName, &StatsHandler{})
}

type StatsHandler struct{}

func getSortedNicksFromArgs(rawArgs string) []string {
	nicks := strings.Fields(rawArgs)
	sort.Strings(nicks)
	return nicks
}

func (h *StatsHandler) Handle(g *game.Game, c *command.Command, outputCh chan string) {
	nicks := getSortedNicksFromArgs(c.Args)
	// If there are no arguments, return the stats for the user that made
	// issued the command.
	if len(nicks) == 0 {
		stats, err := g.PlayerStats(c.Nick)
		if err != nil {
			outputCh <- err.Error()
			return
		}
		outputCh <- fmt.Sprintf("%s: %s", c.Nick, stats)
		return
	}

	for _, nick := range nicks {
		stats, err := g.PlayerStats(nick)
		if err != nil {
			outputCh <- err.Error()
			continue
		}
		outputCh <- fmt.Sprintf("%s: %s's stats: %s", c.Nick, nick, stats)
	}
}
