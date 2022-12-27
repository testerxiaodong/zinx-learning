package main

import (
	"zinx-learning/znet"
)

func main() {
	s := znet.NewServer("Zinx V0.1")
	// Serve函数会阻塞主程序，用go异步执行
	go s.Serve()
	// 创建客户端链接服务端，测试服务端程序的逻辑
	znet.CreateClient()
}
