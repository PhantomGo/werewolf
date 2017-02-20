package main

import (
	"strconv"
	"strings"

	"github.com/wizjin/weixin"
)

const (
	HelpMsg   = "创建游戏发送:c 人数 杀人:k 号码 验人:s 号码 救人:r 号码"
	JoinMsg   = " 加入游戏请发 'j 号码 g' 狼请发 'j 号码 w'"
	ForbidMsg = "你没权限"
)

var (
	cMap = InitCmds()
)

func Echo(w weixin.ResponseWriter, r *weixin.Request) {
	txt := r.Content // 获取用户发送的消息
	cmds := strings.Split(txt, " ")
	if len(cmds) < 2 {
		w.ReplyText(HelpMsg)
		return
	}
	c := cmds[0]
	c = strings.ToLower(c)
	nStr := cmds[1]
	n, err := strconv.Atoi(nStr)
	if err != nil && c != "d" && c != "name" {
		w.ReplyText(HelpMsg)
		return
	}

	if _, ok := cMap[c]; !ok {
		if c == "j" {
			var sk uint
			s := strings.ToLower(cmds[2])
			switch s {
			case "w":
				sk = 1
			case "s":
				sk = 2
			case "wt":
				sk = 3
			default:
				sk = 0
			}
			w.ReplyText(Join(n, r.FromUserName, sk))
			return
		}
		w.ReplyText(HelpMsg)
		return
	}
	w.ReplyText(cMap[c](r.FromUserName, n))
}

func Subscribe(w weixin.ResponseWriter, r *weixin.Request) {
	w.ReplyText(HelpMsg) // 有新人关注，返回欢迎消息
}
