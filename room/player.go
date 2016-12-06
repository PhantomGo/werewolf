package room

type Player struct {
	Number     int
	IsWerewolf bool
	IsDead     bool
}

func NewPlayer(number int, wolf bool) *Player {
	return &Player{IsDead: false, Number: number, IsWerewolf: wolf}
}
