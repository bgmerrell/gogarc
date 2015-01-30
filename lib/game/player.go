package game

import (
	"fmt"
)

// Customize in the future based on race or something fun
const defaultStat = 4

type playerStats struct {
	health int
	luck   int
	vigor  int
	wits   int
}

func (s *playerStats) String() string {
	return fmt.Sprintf("health: %d, luck: %d, vigor: %d, wits: %d",
		s.health, s.luck, s.vigor, s.wits)
}

func newPlayerStats() *playerStats {
	return &playerStats{
		health: defaultStat,
		wits:   defaultStat,
		vigor:  defaultStat,
		luck:   defaultStat}
}

type Player struct {
	name  string
	stats *playerStats
}

func NewPlayer(name string) *Player {
	return &Player{
		name:  name,
		stats: newPlayerStats()}
}

func (p *Player) String() string {
	return fmt.Sprintf("%s: %s", p.name, p.stats.String())
}
