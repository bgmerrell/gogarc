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
	name  string
	stats playerStats
}

func NewPlayer(name string) *Player {
	return &Player{
		name:  name,
		stats: newPlayerStats()}
}

func (p *Player) String() string {
	return fmt.Sprintf("%s: (%s)", p.name, p.stats.String())
}
