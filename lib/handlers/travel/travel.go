package travel

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bgmerrell/gogarc/lib/command"
	"github.com/bgmerrell/gogarc/lib/game"
	hr "github.com/bgmerrell/gogarc/lib/handlerregistry"
)

const (
	commandName    = "travel"
	attackDieSides = 6
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	hr.Registry.Add(commandName, &TravelHandler{})
}

type TravelHandler struct{}

func (h *TravelHandler) Handle(g *game.Game, c *command.Command, outputCh chan string) {
	if !g.InProgress() || c.Nick != g.CurrentPlayer() {
		outputCh <- fmt.Sprintf("%s: It's not your turn.", c.Nick)
		return
	}
	enemy := g.RandomEnemy()

	// The enemy attacks with whatever they are strongest with.  If their
	// wits and vigor are the same then one "attack type" is chosen
	// randomly.
	var attackType string
	var attackMod int
	attackTypes := []string{"wits", "vigor"}
	witsIdx := 0
	vigorIdx := 1
	if enemy.Stats.Wits > enemy.Stats.Vigor {
		attackType = attackTypes[witsIdx]
	} else if enemy.Stats.Wits < enemy.Stats.Vigor {
		attackType = attackTypes[vigorIdx]
	} else {
		attackType = attackTypes[rand.Intn(len(attackTypes))]
	}
	if attackTypes[witsIdx] == attackType {
		attackMod = enemy.Stats.Wits
	} else {
		attackMod = enemy.Stats.Vigor
	}
	outputCh <- fmt.Sprintf(
		"%s: You encounter %s that challenges your %s.",
		c.Nick, enemy.String(), attackType)
	roll := rand.Intn(attackDieSides)
	outputCh <- fmt.Sprintf(
		"%s: Your enemy's attack score is %d (%d +%d)",
		c.Nick, roll+attackMod, roll, attackMod)
}
