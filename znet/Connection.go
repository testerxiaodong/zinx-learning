package znet

import (
	"errors"
	"fmt"
	"io"
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
	// 告知当前链接已经退出/停止的 channel
	ExitChan chan bool
	// 该链接处理的方法IRouter
	Router ziface.IRouter
}

// StartReader 链接的读业务方法
func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running")
	defer fmt.Println("ConnID = ", c.ConnID, "Reader is exit, remote addr is ", c.Conn.RemoteAddr().String())
	defer c.Stop()
	for {
		// 读取客户端数据到buffer中
		//buf := make([]byte, utils.GlobalObject.MaxPackageSize)
		//_, err := c.Conn.Read(buf)
		//if err != nil {
		//	fmt.Println("read buffer error: ", err)
		//	continue
		//}
		//if err := c.handleAPI(c.Conn, buf, cnt); err != nil {
		//	fmt.Println("Conn handle error: ", err)
		//	continue
		//}
		// 创建一个拆包对象
		dp := NewDatePack()
		// 读取客户端的Msg Head 二进制流 8个字节
		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(c.GetTCPConnection(), headData); err != nil {
			fmt.Println("read msg head error: ", err)
			break
		}
		// 拆包，得到msgId 和 DateLen 放在msg 中
		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("headData unpack error: ", err)
			break
		}
		if msg.GetMessageLen() > 0 {
			data := make([]byte, msg.GetMessageLen())
			_, err := io.ReadFull(c.GetTCPConnection(), data)
			if err != nil {
				fmt.Println("read data error: ", err)
				break
			}
			msg.SetMessageData(data)
		}
		// 封装Request
		req := &Request{
			conn: c,
			msg:  msg,
		}

		go func(req ziface.IRequest) {
			c.Router.PreHandle(req)
			c.Router.Handle(req)
			c.Router.PostHandle(req)
		}(req)
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

func (c *Connection) Send(MsgId uint32, data []byte) error {
	if c.isClosed {
		return errors.New("connection closed when send msg")
	}
	dp := NewDatePack()
	binaryMsg, err := dp.Pack(NewMessage(MsgId, data))
	if err != nil {
		return err
	}
	if _, err := c.Conn.Write(binaryMsg); err != nil {
		fmt.Println("write msg error: ", err)
	}
	return nil
}

// NewConnection 链接对象的构造函数
func NewConnection(Conn *net.TCPConn, ConnID uint32, router ziface.IRouter) *Connection {
	return &Connection{
		Conn,
		ConnID,
		false,
		make(chan bool),
		router,
	}
}
