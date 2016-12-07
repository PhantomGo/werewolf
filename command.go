package main

import (
	"strconv"
	"werewolf/room"
)

var (
	croom *room.Room
)

func Create(n int) string {
	croom = room.NewRoom(n)
	return "加入游戏请发 j号码g 狼请发 j号码w"
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

func GetDeads() string {
	var result string
	for _, n := range croom.Deads {
		result += strconv.Itoa(n) + ","
	}
	result += "死了"
	return result
}

func Kill(n int) string {
	croom.Kill(n)
	return "啊!!"
}

func SeekWolf(n int) string {
	if croom.CheckWolf(n) {
		return "狼"
	}
	return "好人"
}

func Rescue(n int) string {
	if croom.Cure(n) {
		return strconv.Itoa(n) + "号活了"
	}
	return "救错了"
}
