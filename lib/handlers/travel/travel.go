package travel

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bgmerrell/gogarc/lib/command"
	"github.com/bgmerrell/gogarc/lib/game"
	hr "github.com/bgmerrell/gogarc/lib/handlerregistry"
)

const commandName = "travel"

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	hr.Registry.Add(commandName, &TravelHandler{})
}

type TravelHandler struct{}

func (h *TravelHandler) Handle(g *game.Game, c *command.Command, outputCh chan string) {
	if !g.InProgress {
		outputCh <- fmt.Sprintf("The game has not started.")
		return
	}
	if c.Nick != g.CurrentPlayer().Name {
		outputCh <- fmt.Sprintf("%s: It's not your turn.", c.Nick)
		return
	}
	if g.Turn.HasTraveled {
		outputCh <- fmt.Sprintf("%s: You have already traveled this turn.", c.Nick)
		return
	}
	g.Turn.HasTraveled = true
	enemy := g.RandomEnemy()

	// The enemy attacks with whatever they are strongest with.  If their
	// wits and vigor are the same then one "attack type" is chosen
	// randomly.
	var attackMod int
	if enemy.Stats.Wits > enemy.Stats.Vigor {
		g.Turn.AttackType = game.AttackTypeWits
	} else if enemy.Stats.Wits < enemy.Stats.Vigor {
		g.Turn.AttackType = game.AttackTypeVigor
	} else {
		attackTypes := []string{game.AttackTypeWits, game.AttackTypeVigor}
		g.Turn.AttackType = attackTypes[rand.Intn(len(attackTypes))]
	}
	switch g.Turn.AttackType {
	case game.AttackTypeWits:
		attackMod = enemy.Stats.Wits
	case game.AttackTypeVigor:
		attackMod = enemy.Stats.Vigor
	}
	outputCh <- fmt.Sprintf(
		"%s: You encounter %s that challenges your %s.",
		c.Nick, enemy.String(), g.Turn.AttackType)
	roll := rand.Intn(game.AttackDieSides) + 1
	g.Turn.EnemyAttack = roll + attackMod
	outputCh <- fmt.Sprintf(
		"%s: Your enemy's attack score is %d (%d +%d)",
		c.Nick, g.Turn.EnemyAttack, roll, attackMod)
}
