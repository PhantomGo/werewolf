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
	return &Wall{Points: make([]*PointRecord, 0, 1000)}
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

func (w *Wall) Show() (pw []*PointWall) {
	pw = make([]*PointWall, 0, 40)

	type PM struct {
		Total int
		Count float32
	}
	pMap := make(map[string]*PM, 40)
	for _, p := range w.Points {
		if _, pm := pMap[p.OpenID]; !pm {
			pMap[p.OpenID] = &PM{p.Point, 1}
		} else {
			pMap[p.OpenID].Count++
			pMap[p.OpenID].Total += p.Point
		}
	}
	for n, pm := range pMap {
		pw = append(pw, &PointWall{Name: n, PointTotal: pm.Total, PointAverage: float32(pm.Total) / pm.Count})
	}
	return
}

type PointWall struct {
	Name         string  `json:"name"`
	PointTotal   int     `json:"total"`
	PointAverage float32 `json:"rank"`
}
