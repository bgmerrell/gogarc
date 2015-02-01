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

// TODO: Put in config file
const minPlayers = 2

type Turn struct {
	playerNumber int
	hasTraveled  bool
}

type Game struct {
	InProgress  bool
	players     map[string]*Player
	playerOrder []string
	enemies     []Being
	turn        Turn
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
	return player.String(), err
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

	g.turn = Turn{
		playerNumber: 0,
		hasTraveled:  false,
	}
	return fmt.Sprintf("Game has begun.  Player order: %s.",
		strings.Join(g.playerOrder, ", ")), err
}

func (g *Game) CurrentPlayer() string {
	return g.playerOrder[g.turn.playerNumber]
}

func (g *Game) RandomEnemy() Being {
	return g.enemies[rand.Intn(len(g.enemies))]
}
