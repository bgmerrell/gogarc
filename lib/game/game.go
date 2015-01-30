package game

import (
	"errors"
	"fmt"
	"sync"
)

type Game struct {
	inProgress bool
	players    map[string]*Player
	mu         sync.Mutex
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
