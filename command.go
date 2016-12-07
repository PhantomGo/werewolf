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
	cm["v"] = vote
	return
}

func create(n int) string {
	croom = room.NewRoom(n)
	return JoinMsg
}

func Join(n int, isW bool) string {
	var result string
	begin := croom.IsBegin()
	if !begin && n <= croom.Count && n > 0 {
		croom.Join(n, isW)
		result = strconv.Itoa(n) + "号就位"
	} else {
		result = "游戏开始了,发送 c 重新创建"
	}
	return result
}

func getDeads(n int) string {
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

func vote(n int) string {
	if croom.Vote(n) {
		return "游戏结束"
	}
	return "进入黑夜"
}
