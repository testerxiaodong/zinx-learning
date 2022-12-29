package ziface

import "net"

// IConnection 定义链接模块的抽象层
type IConnection interface {
	// Start 启动链接 让当前链接准备开始工作
	Start()
	// Stop 停止链接 结束当前链接的工作
	Stop()
	// GetTCPConnection 获取当前链接绑定的socket conn
	GetTCPConnection() *net.TCPConn
	// GetConnID 获取当前链接的链接ID
	GetConnID() uint32
	// GetRemoteAddr 获取远程客户端的TCP状态 IP Port
	GetRemoteAddr() net.Addr
	// Send 发送数据，将数据发送给远程的客户端
	Send(MsgId uint32, data []byte) error
}

type HandleFunc func(*net.TCPConn, []byte, int) error
