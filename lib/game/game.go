package game

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"sync"
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
	inProgress  bool
	players     map[string]*Player
	playerOrder []string
	enemies     []Being
	turn        Turn
	mu          sync.Mutex
}

func NewGame() (g *Game, err error) {
	g = &Game{
		inProgress: false,
		players:    make(map[string]*Player),
		mu:         sync.Mutex{}}
	g.enemies, err = LoadEnemies()
	return g, err
}

func (g *Game) AddPlayer(name string) (err error) {
	g.mu.Lock()
	defer g.mu.Unlock()
	_, ok := g.players[name]
	if ok {
		return errors.New(fmt.Sprintf(
			"Player \"%s\" is already in the game.", name))
	}
	g.players[name] = NewPlayer(name)
	return err
}

func (g *Game) PlayerStats(name string) (stats string, err error) {
	g.mu.Lock()
	defer g.mu.Unlock()
	player, ok := g.players[name]
	if !ok {
		return "", errors.New(fmt.Sprintf(
			"Player \"%s\" is not in the game.", name))
	}
	return player.String(), err
}

func (g *Game) Start() (msg string, err error) {
	g.mu.Lock()
	defer g.mu.Unlock()
	nPlayers := len(g.players)
	if nPlayers < minPlayers {
		return "", errors.New(fmt.Sprintf(
			"You need at least %d players; %d is not enough!",
			minPlayers, nPlayers))
	}
	if g.inProgress {
		return "", errors.New("The game has already begun.")
	}
	g.inProgress = true

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
	g.mu.Lock()
	defer g.mu.Unlock()
	return g.playerOrder[g.turn.playerNumber]
}

func (g *Game) InProgress() bool {
	g.mu.Lock()
	defer g.mu.Unlock()
	return g.inProgress
}

func (g *Game) RandomEnemy() Being {
	g.mu.Lock()
	defer g.mu.Unlock()
	return g.enemies[rand.Intn(len(g.enemies))]
}
