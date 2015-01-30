package game

type Player struct {
	health int
	wits   int
	vigor  int
	luck   int
}

// Customize in the future based on race or something fun
const defaultStat = 4

func NewPlayer() *Player {
	return &Player{
		health: defaultStat,
		wits:   defaultStat,
		vigor:  defaultStat,
		luck:   defaultStat}
}
