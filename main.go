package main

import (
	"fmt"
	"io"
	"net/http"
	"runtime"

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
	http.Handle("/", mux) // 注册接收微信服务器数据的接口URI
	http.HandleFunc("/ss", hello)
	err := http.ListenAndServe(":80", nil) // 启动接收微信数据服务器
	if err != nil {
		fmt.Println(err)
	}
	InitSignal()
}

func hello(rw http.ResponseWriter, req *http.Request) {
	io.WriteString(rw, "hello phantom")
}

func Echo(w weixin.ResponseWriter, r *weixin.Request) {
	txt := r.Content // 获取用户发送的消息
	dict := InitCommand()
	w.ReplyText(dict[txt]) // 回复一条文本消息
	//w.PostText("Post:" + txt) // 发送一条文本消息
}

func Subscribe(w weixin.ResponseWriter, r *weixin.Request) {
	w.ReplyText("欢迎关注") // 有新人关注，返回欢迎消息
}
