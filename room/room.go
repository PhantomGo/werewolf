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
	return len(this.Players) == this.Count
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
