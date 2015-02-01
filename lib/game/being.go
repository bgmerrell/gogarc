package game

import (
	"fmt"
)

type beingStats struct {
	Health int
	Vigor  int
	Wits   int
}

func (s *beingStats) String() string {
	return fmt.Sprintf("Health: %d, Vigor: %d, Wits: %d",
		s.Health, s.Vigor, s.Wits)
}

type Being struct {
	Name  string
	Stats beingStats `json:"stats"`
}

func (b *Being) String() string {
	return fmt.Sprintf("%s (%s)", b.Name, b.Stats.String())
}
