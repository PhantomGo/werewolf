package main

import (
	"fmt"
	"net/http"
	"runtime"
	"strconv"

	"github.com/wizjin/weixin"
)

func main() {
	fmt.Println("This is webserver base!")
	runtime.GOMAXPROCS(runtime.NumCPU())
	mux := weixin.New("werewolf", "wx94f59dce884960bc", "e743cfb5ab3b914740f2f7d3a9de6f7b")
	// 注册文本消息的处理函数
	mux.HandleFunc(weixin.MsgTypeText, Echo)
	// 注册关注事件的处理函数
	mux.HandleFunc(weixin.MsgTypeEventSubscribe, Subscribe)
	http.Handle("/", mux)                  // 注册接收微信服务器数据的接口URI
	err := http.ListenAndServe(":80", nil) // 启动接收微信数据服务器
	if err != nil {
		fmt.Println(err)
	}
	InitSignal()
}

func Echo(w weixin.ResponseWriter, r *weixin.Request) {
	txt := r.Content // 获取用户发送的消息
	if len(txt) < 2 {
		w.ReplyText("输入错误")
	} else {
		c := Substr(txt, 0, 1)
		n, err := strconv.Atoi(Substr(txt, 1, 1))
		if err != nil {
			w.ReplyText("输入错误")
		} else {
			switch c {
			case "c":
				w.ReplyText(Create(n)) // 回复一条文本消息
			case "k":
				w.ReplyText(Kill(n)) // 回复一条文本消息
			case "s":
				w.ReplyText(CheckWolf(n)) // 回复一条文本消息
			case "dd":
				w.ReplyText(GetDeads()) // 回复一条文本消息
			case "j":
				w.ReplyText(Join(n)) // 回复一条文本消息
			}
		}
		//w.PostText("Post:" + txt) // 发送一条文本消息
	}
}

func Subscribe(w weixin.ResponseWriter, r *weixin.Request) {
	w.ReplyText("创建游戏发送 c人数 杀人发 k号码 验人发 s号码") // 有新人关注，返回欢迎消息
}

func Substr(str string, start int, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}
