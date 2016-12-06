package room

type Room struct {
	Id      int
	Count   int
	PIndex  int
	Players []*Player
	Deads   []int
}

func NewRoom(count int) *Room {
	r := &Room{Count: count, PIndex: 0, Id: 0}
	r.Players = make([]*Player, count)
	r.Deads = make([]int, 6)
	return r
}

func (this *Room) Join(n int, wolf bool) *Player {
	p := NewPlayer(n, wolf)
	this.Players[this.PIndex] = p
	this.PIndex++
	return p
}

func (this *Room) IsBegin() bool {
	return this.PIndex+1 == this.Count
}

func (this *Room) Kill(number int) {
	this.Players[number-1].IsDead = true
	this.Deads = append(this.Deads, number)
}

func (this *Room) CheckWolf(n int) bool {
	for _, p := range this.Players {
		if p.Number == n {
			return p.IsWerewolf
		}
	}
	return false
}
