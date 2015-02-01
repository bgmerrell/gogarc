package attack

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bgmerrell/gogarc/lib/command"
	"github.com/bgmerrell/gogarc/lib/game"
	hr "github.com/bgmerrell/gogarc/lib/handlerregistry"
)

const commandName = "attack"

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	hr.Registry.Add(commandName, &AttackHandler{})
}

type AttackHandler struct{}

func (h *AttackHandler) Handle(g *game.Game, c *command.Command, outputCh chan string) {
	player := g.CurrentPlayer()
	if !g.InProgress || c.Nick != player.Name {
		outputCh <- fmt.Sprintf("%s: It's not your turn.", c.Nick)
		return
	}
	// There hasn't been an encounter if the attack type isn't set
	if g.Turn.AttackType == "" {
		outputCh <- fmt.Sprintf("%s: There's nothing to attack.", c.Nick)
		return
	}
	var attackMod int
	if g.Turn.AttackType == game.AttackTypeWits {
		attackMod = player.Stats.Wits
	} else {
		attackMod = player.Stats.Vigor
	}
	roll := rand.Intn(game.AttackDieSides) + 1
	total := roll + attackMod
	var resultMsg string
	if total == g.Turn.EnemyAttack {
		resultMsg = "It's a draw."
	} else if total > g.Turn.EnemyAttack {
		resultMsg = "You win!"
	} else {
		resultMsg = "You lose!"
		player.Stats.Health -= 1
		if player.Stats.Health <= 0 {
			resultMsg += "You are dead!"
			// TODO: update game state
		}
	}
	outputCh <- fmt.Sprintf(
		"%s: Your attack score is %d (%d +%d). %s",
		c.Nick, total, roll, attackMod, resultMsg)
}
