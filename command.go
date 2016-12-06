package main

import "werewolf/room"

var (
	croom *room.Room
)

func InitCommand() map[string]string {
	cmds := make(map[string]string)
	cmds["c"] = create(1)
	cmds["j"] = join(1)
	cmds["d"] = getDeads()
	cmds["k"] = kill(1)
	cmds["s"] = checkWolf(1)
	return cmds
}

func create(n int) string {
	croom = room.NewRoom(n)
	return "加入游戏请发 j"
}

func join(n int) string {
	var result string
	begin := croom.IsBegin()
	if !begin {
		croom.Join(n, true)
		result = "你是" + string(n) + "号"
	} else {
		result = "游戏开始了,发送 c 重新创建"
	}
	return result
}

func getDeads() string {
	var result string
	for _, n := range croom.Deads {
		result = result + string(n) + ","
	}
	result = "死了"
	return result
}

func kill(n int) string {
	croom.Kill(n)
	return "啊!!"
}

func checkWolf(n int) string {
	if croom.CheckWolf(n) {
		return "狼"
	}
	return "好人"
}
