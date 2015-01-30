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

type Game struct {
	inProgress  bool
	players     map[string]*Player
	playerOrder []string
	mu          sync.Mutex
}

func NewGame() *Game {
	return &Game{
		inProgress: false,
		players:    make(map[string]*Player),
		mu:         sync.Mutex{}}
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
	g.inProgress = true

	// Populate playerOrder and make sure it's shuffled
	g.playerOrder = make([]string, nPlayers)
	perm := rand.Perm(nPlayers)
	i := 0
	for name, _ := range g.players {
		g.playerOrder[perm[i]] = name
		i += 1
	}

	return fmt.Sprintf("Game has begun.  Player order: %s",
		strings.Join(g.playerOrder, ", ")), err
}
