package game

import (
	"fmt"
)

// Customize in the future based on race or something fun
const defaultPlayerStat = 4

type playerStats struct {
	beingStats
	Luck int
}

func (s *playerStats) String() string {
	return fmt.Sprintf("Health: %d, Luck: %d, Vigor: %d, Wits: %d",
		s.Health, s.Luck, s.Vigor, s.Wits)
}

func newPlayerStats() playerStats {
	return playerStats{
		beingStats: beingStats{
			Health: defaultPlayerStat,
			Vigor:  defaultPlayerStat,
			Wits:   defaultPlayerStat},
		Luck: defaultPlayerStat}
}

type Player struct {
	Name  string
	Stats playerStats
}

func NewPlayer(name string) *Player {
	return &Player{
		Name:  name,
		Stats: newPlayerStats()}
}

func (p *Player) String() string {
	return fmt.Sprintf("%s: (%s)", p.Name, p.Stats.String())
}
