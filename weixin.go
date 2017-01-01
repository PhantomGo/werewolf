package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/wizjin/weixin"
)

const (
	HelpMsg   = "创建游戏发送 c人数 杀人 k号码 验人 s号码 救人 r号码"
	JoinMsg   = " 加入游戏请发 j号码g 狼请发 j号码w"
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
	if err != nil && c != "d" {
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

func CreateMenu(wx *weixin.Weixin) {
	menu := &weixin.Menu{make([]weixin.MenuButton, 2)}
	menu.Buttons[0].Name = "死亡名单"
	menu.Buttons[0].Type = weixin.MenuButtonTypeKey
	menu.Buttons[0].Key = "D 1"
	menu.Buttons[1].Name = "创建游戏"
	menu.Buttons[1].SubButtons = make([]weixin.MenuButton, 5)
	menu.Buttons[1].SubButtons[0].Name = "8人局"
	menu.Buttons[1].SubButtons[0].Type = weixin.MenuButtonTypeKey
	menu.Buttons[1].SubButtons[0].Key = "c 8"
	menu.Buttons[1].SubButtons[1].Name = "9人局"
	menu.Buttons[1].SubButtons[1].Type = weixin.MenuButtonTypeKey
	menu.Buttons[1].SubButtons[1].Key = "c 9"
	menu.Buttons[1].SubButtons[2].Name = "9人局"
	menu.Buttons[1].SubButtons[2].Type = weixin.MenuButtonTypeKey
	menu.Buttons[1].SubButtons[2].Key = "c 10"
	menu.Buttons[1].SubButtons[3].Name = "10人局"
	menu.Buttons[1].SubButtons[3].Type = weixin.MenuButtonTypeKey
	menu.Buttons[1].SubButtons[3].Key = "c 10"
	menu.Buttons[1].SubButtons[4].Name = "12人局"
	menu.Buttons[1].SubButtons[4].Type = weixin.MenuButtonTypeKey
	menu.Buttons[1].SubButtons[4].Key = "c 12"

	err := wx.CreateMenu(menu)
	if err != nil {
		fmt.Println(err)
	}
}
