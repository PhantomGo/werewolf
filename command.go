package main

import (
	"strconv"
	"werewolf/room"
	"werewolf/wall"
)

var (
	croom *room.Room
	PW    *wall.Wall
)

func InitCmds() (cm map[string]func(string, int) string) {
	cm = make(map[string]func(string, int) string, 10)
	cm["c"] = create
	cm["d"] = getDeads
	//cm["j"] = Join
	cm["k"] = kill
	cm["r"] = rescue
	cm["s"] = seekWolf
	cm["v"] = vote
	cm["l"] = plist
	cm["pw"] = addPoint
	PW = wall.NewWall()
	return
}

func create(id string, n int) string {
	croom = room.NewRoom(n)
	return "房间号" + strconv.Itoa(croom.ID) + JoinMsg
}

func Join(n int, id string, skill uint) string {
	var result string
	begin := croom.IsBegin()
	if !begin && n <= croom.Count && n > 0 {
		croom.Join(id, n, skill)
		result = strconv.Itoa(n) + "号就位"
	} else {
		result = "游戏开始了,发送 c 重新创建"
	}
	return result
}

func getDeads(id string, n int) string {
	if len(croom.Deads) < 1 {
		return "平安夜"
	}
	var result string
	if len(croom.Deads) == 2 {
		if croom.Deads[0] > croom.Deads[1] {
			result = strconv.Itoa(croom.Deads[1]) + " " + strconv.Itoa(croom.Deads[0]) + "双死"
		} else {
			result = strconv.Itoa(croom.Deads[0]) + " " + strconv.Itoa(croom.Deads[1]) + "双死"
		}
	} else {
		result = strconv.Itoa(croom.Deads[0]) + "死了"
	}
	return result
}

func kill(id string, n int) string {
	if p := findP(id); p != nil && p.CanKill() {
		croom.Kill(n)
		return "啊!!"
	}
	return "你不能杀人"
}

func seekWolf(id string, n int) string {
	if p := findP(id); p != nil && p.IsSeeker {
		if croom.CheckWolf(n) {
			return "狼"
		}
		return "好人"
	}
	return ForbidMsg
}

func rescue(id string, n int) string {
	if p := findP(id); p != nil && p.IsWitch {
		if croom.Cure(n) {
			return strconv.Itoa(n) + "号活了"
		}
		return "救错了"
	}
	return ForbidMsg
}

func vote(id string, n int) string {
	if croom.Vote(n) {
		return "游戏结束"
	}
	return "进入黑夜"
}

func plist(id string, n int) string {
	var result string
	for i, p := range croom.Players {
		if p != nil {
			result += strconv.Itoa(i) + ","
		}
	}
	result += "已加入游戏"
	return result
}

func addPoint(id string, n int) string {
	PW.Add(croom.ID, n, id)
	return "success"
}

func findP(id string) *room.Player {
	for _, p := range croom.Players {
		if p.OpenId == id {
			return p
		}
	}
	return nil
}
