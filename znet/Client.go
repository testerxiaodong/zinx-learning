package znet

import (
	"fmt"
	"net"
	"time"
)

func CreateClient() {
	fmt.Println("client start")

	time.Sleep(time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("dail tcp error: ", err)
		return
	}
	for {
		_, err := conn.Write([]byte("Hello World"))
		if err != nil {
			fmt.Println("write conn error: ", err)
			return
		}
		buf := make([]byte, 512)
		ctn, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read conn error:", err)
			return
		}
		fmt.Printf("read from server: %s, ctn: %d\n", buf, ctn)
		time.Sleep(time.Second)
	}
}
