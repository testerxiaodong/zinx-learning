package main

import (
	"fmt"
	"zinx-learning/ziface"
	"zinx-learning/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

//func (p *PingRouter) PreHandle(req ziface.IRequest) {
//	fmt.Println("Call Router PreHandle")
//	_, err := req.GetCurrConn().GetTCPConnection().Write([]byte("before ping...\n"))
//	if err != nil {
//		fmt.Println("Call Back PreHandle error: ", err)
//		return
//	}
//}

func (p *PingRouter) Handle(req ziface.IRequest) {
	fmt.Println("Call Router Handle")
	fmt.Println("recv from client: MsgId = ", req.GetMessageId(), "data = ", string(req.GetData()))
	err := req.GetCurrConn().Send(1, []byte("ping...ping...ping"))
	if err != nil {
		fmt.Println(err)
	}
}

//func (p *PingRouter) PostHandle(req ziface.IRequest) {
//	fmt.Println("Call Router PostHandle")
//	_, err := req.GetCurrConn().GetTCPConnection().Write([]byte("after ping...\n"))
//	if err != nil {
//		fmt.Println("Call Back PostHandle error: ", err)
//		return
//	}
//}

func main() {
	s := znet.NewServer("Zinx V0.2")
	// Serve函数会阻塞主程序，用go异步执行
	s.AddRouter(&PingRouter{})
	go s.Serve()
	// 创建客户端链接服务端，测试服务端程序的逻辑
	znet.CreateClient()
}
