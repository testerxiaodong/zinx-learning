package znet

import (
	"fmt"
	"net"
	"zinx-learning/ziface"
)

// Server 是IServer的实现类
type Server struct {
	// 服务器名称
	Name string
	// 服务器绑定的ip版本
	IPVersion string
	// 服务器监听的IP
	IP string
	// 服务器的端口
	Port int
}

// Start 启动服务器
func (s *Server) Start() {
	fmt.Printf("[Start] Server Listenner at IP :%s, Port: %d\n", s.IP, s.Port)
	// 获取一个TCP的Addr
	go func() {
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr error: ", err)
			return
		}
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen tcp error: ", err)
			return
		}
		fmt.Println("start Zinx server success", s.Name)
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("accept tcp error: ", err)
				continue
			}
			// 客户端建立连接成功，进行业务处理
			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("recv buf error: ", err)
						continue
					}
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("write buffer error: ", err)
						continue
					}
				}
			}()
		}
	}()
}

func (s *Server) Stop() {
	// Todo 将一些服务器的资源、状态或者一些已经开辟的链接信息进行停止或者回收
}

func (s *Server) Serve() {
	// 启动server的服务功能
	s.Start()
	// Todo 做一些启动服务器之后的额外业务
	// 阻塞状态
	select {}
}

func NewServer(name string) ziface.IServer {
	return &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
}
