package main

import (
	"strconv"
	"werewolf/room"
)

var (
	croom *room.Room
)

func InitCmds() (cm map[string]func(int) string) {
	cm = make(map[string]func(int) string, 10)
	cm["c"] = create
	cm["d"] = getDeads
	//cm["j"] = Join
	cm["k"] = kill
	cm["r"] = rescue
	cm["s"] = seekWolf
	return
}

func create(n int) string {
	croom = room.NewRoom(n)
	return JoinMsg
}

func Join(n int, isW bool) string {
	var result string
	begin := croom.IsBegin()
	if !begin {
		croom.Join(n, isW)
		result = strconv.Itoa(n) + "号就位"
	} else {
		result = "游戏开始了,发送 c 重新创建"
	}
	return result
}

func getDeads(n int) string {
	if len(croom.Deads) < 1 {
		return "没人死"
	}
	var result string
	for _, n := range croom.Deads {
		result += strconv.Itoa(n) + ","
	}
	result += "死了"
	return result
}

func kill(n int) string {
	croom.Kill(n)
	return "啊!!"
}

func seekWolf(n int) string {
	if croom.CheckWolf(n) {
		return "狼"
	}
	return "好人"
}

func rescue(n int) string {
	if croom.Cure(n) {
		return strconv.Itoa(n) + "号活了"
	}
	return "救错了"
}
