package main

import (
	"strconv"
	"strings"

	"github.com/wizjin/weixin"
)

var (
	HelpMsg = "创建游戏发送 c人数 杀人 k号码 验人 s号码 救人 r号码"
	JoinMsg = "加入游戏请发 j号码g 狼请发 j号码w"
	cMap    = InitCmds()
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
	if err != nil && c != "d" {
		w.ReplyText(HelpMsg)
		return
	}

	if _, ok := cMap[c]; !ok {
		if c == "j" {
			w.ReplyText(Join(n, cmds[2] == "w"))
			return
		}
		w.ReplyText(HelpMsg)
		return
	}

	w.ReplyText(cMap[c](n))
	//w.PostText("Post:" + txt) // 发送一条文本消息
}

func Subscribe(w weixin.ResponseWriter, r *weixin.Request) {
	w.ReplyText(HelpMsg) // 有新人关注，返回欢迎消息
}
