package room

type Player struct {
	OpenId     string
	Number     int
	IsWerewolf bool
	IsSeeker   bool
	IsWitch    bool
	IsDead     bool
}

func NewPlayer(i string, number int, skill uint) *Player {
	return &Player{IsDead: false, Number: number, IsWerewolf: skill == 1,
		IsSeeker: skill == 2, IsWitch: skill == 3, OpenId: i}
}

func (p *Player) CanKill() bool {
	return p.IsWerewolf || p.IsWitch
}
