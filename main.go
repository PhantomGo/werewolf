package main

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/wizjin/weixin"
)

func main() {
	fmt.Println("This is webserver base!")
	runtime.GOMAXPROCS(runtime.NumCPU())
	mux := weixin.New("werewolf", "wx94f59dce884960bc", "e743cfb5ab3b914740f2f7d3a9de6f7b")
	CreateMenu(mux)
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
