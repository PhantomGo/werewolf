package wall

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	w := NewWall()
	w.Add(1, 1, "s")
	t.Log("ok")
}

func TestShow(t *testing.T) {
	w := NewWall()
	w.Add(1, 1, "a")
	w.Add(1, 1, "b")
	w.Add(2, 2, "b")
	for _, p := range w.Show() {
		fmt.Println(fmt.Sprintf("%s %d %f", p.Name, p.PointTotal, p.PointAverage))
	}
}
