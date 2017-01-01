package wall

import (
	"fmt"
	"strconv"
	"time"
)

type Wall struct {
	Points []*PointRecord
}

type PointRecord struct {
	RoomID int
	OpenID string
	Point  int
	CTime  string
}

func NewWall() *Wall {
	return &Wall{Points: make([]*PointRecord, 1000)}
}

func (w *Wall) Add(rid, p int, oid string) (err error) {
	w.Points = append(w.Points, &PointRecord{
		RoomID: rid,
		Point:  p,
		OpenID: oid,
		CTime:  time.Now().String(),
	})
	for _, p := range w.Points {
		fmt.Println(strconv.Itoa(p.RoomID) + " " + p.OpenID + " " + p.CTime)
	}
	return err
}
