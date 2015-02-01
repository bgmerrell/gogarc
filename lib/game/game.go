package game

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

const (
	// TODO: Put in config file
	minPlayers = 2

	AttackDieSides = 6

	AttackTypeWits  = "wits"
	AttackTypeVigor = "vigor"
)

type Turn struct {
	playerNumber int
	HasTraveled  bool
	AttackType   string
	EnemyAttack  int
}

type Game struct {
	InProgress  bool
	players     map[string]*Player
	playerOrder []string
	enemies     []Being
	Turn        Turn
}

func NewGame() (g *Game, err error) {
	g = &Game{players: make(map[string]*Player)}
	g.enemies, err = LoadEnemies()
	return g, err
}

func (g *Game) AddPlayer(name string) (err error) {
	_, ok := g.players[name]
	if ok {
		return errors.New(fmt.Sprintf(
			"Player \"%s\" is already in the game.", name))
	}
	g.players[name] = NewPlayer(name)
	return err
}

func (g *Game) PlayerStats(name string) (stats string, err error) {
	player, ok := g.players[name]
	if !ok {
		return "", errors.New(fmt.Sprintf(
			"Player \"%s\" is not in the game.", name))
	}
	return player.Stats.String(), err
}

func (g *Game) Start() (msg string, err error) {
	nPlayers := len(g.players)
	if nPlayers < minPlayers {
		return "", errors.New(fmt.Sprintf(
			"You need at least %d players; %d is not enough!",
			minPlayers, nPlayers))
	}
	if g.InProgress {
		return "", errors.New("The game has already begun.")
	}
	g.InProgress = true

	// Populate playerOrder and make sure it's shuffled
	g.playerOrder = make([]string, nPlayers)
	perm := rand.Perm(nPlayers)
	i := 0
	for name, _ := range g.players {
		g.playerOrder[perm[i]] = name
		i += 1
	}

	g.Turn = Turn{
		playerNumber: 0,
		HasTraveled:  false,
	}
	return fmt.Sprintf("Game has begun.  Player order: %s.",
		strings.Join(g.playerOrder, ", ")), err
}

func (g *Game) CurrentPlayer() *Player {
	return g.players[g.playerOrder[g.Turn.playerNumber]]
}

func (g *Game) RandomEnemy() Being {
	return g.enemies[rand.Intn(len(g.enemies))]
}
