package room

type Room struct {
	Id      int
	Count   int
	Players []*Player
	Deads   []int
}

func NewRoom(count int) *Room {
	r := &Room{Count: count, Id: 0}
	r.Players = make([]*Player, count)
	r.Deads = make([]int, 0, 6)
	return r
}

func (this *Room) Join(n int, wolf bool) *Player {
	p := NewPlayer(n, wolf)
	this.Players[n-1] = p
	return p
}

func (this *Room) IsBegin() bool {
	for _, p := range this.Players {
		if p == nil {
			return false
		}
	}
	return true
}

func (this *Room) Kill(number int) {
	p := this.Players[number-1]
	if !p.IsDead {
		p.IsDead = true
		this.Deads = append(this.Deads, number)
	}
}

func (this *Room) CheckWolf(n int) bool {
	return this.Players[n-1].IsWerewolf
}

func (this *Room) Cure(n int) bool {
	for i, p := range this.Deads {
		if p == n {
			this.Deads = append(this.Deads[:i], this.Deads[i+1:]...)
			return true
		}
	}
	return false
}

func (this *Room) Vote(n int) bool {
	this.Players[n-1].IsDead = true
	this.Deads = make([]int, 0, 6)
	good := make([]int, 0, 4)
	wolf := make([]int, 0, 4)
	for _, p := range this.Players {
		if p.IsDead {
			continue
		}
		if p.IsWerewolf {
			good = append(good, p.Number)
		} else {
			wolf = append(wolf, p.Number)
		}
	}
	return len(good) == len(wolf)
}
