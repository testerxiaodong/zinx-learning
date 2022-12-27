package znet

import (
	"fmt"
	"net"
	"zinx-learning/ziface"
)

type Connection struct {
	// 当前链接的socket TCP套接字
	Conn *net.TCPConn
	// 当前链接ID
	ConnID uint32
	// 当前链接的状态
	isClosed bool
	// 当前链接所绑定的处理业务方法API
	handleAPI ziface.HandleFunc
	// 告知当前链接已经退出/停止的 channel
	ExitChan chan bool
}

// StartReader 链接的读业务方法
func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running")
	defer fmt.Println("ConnID = ", c.ConnID, "Reader is exit, remote addr is ", c.Conn.RemoteAddr().String())
	defer c.Stop()
	for {
		// 读取客户端数据到buffer中
		buf := make([]byte, 512)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("read buffer error: ", err)
			continue
		}
		if err := c.handleAPI(c.Conn, buf, cnt); err != nil {
			fmt.Println("Conn handle error: ", err)
			continue
		}
	}
}

func (c *Connection) Start() {
	fmt.Println("Conn Start, ConnID = ", c.ConnID)
	// 启动当前链接的读数据业务
	go c.StartReader()
	// Todo 启动当前链接的写数据业务
}

func (c *Connection) Stop() {
	fmt.Println("Conn Stop, ConnID = ", c.ConnID)
	// 如果当前链接已经关闭
	if c.isClosed == true {
		return
	}
	c.isClosed = true
	c.Conn.Close()
	close(c.ExitChan)
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) GetRemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) Send(data []byte) error {
	//TODO implement me
	panic("implement me")
}

// NewConnection 链接对象的构造函数
func NewConnection(Conn *net.TCPConn, ConnID uint32, handleAPI ziface.HandleFunc) *Connection {
	return &Connection{
		Conn,
		ConnID,
		false,
		handleAPI,
		make(chan bool),
	}
}
