package znet

import (
	"fmt"
	"io"
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
		//_, err := conn.Write([]byte("Hello World"))
		//if err != nil {
		//	fmt.Println("write conn error: ", err)
		//	return
		//}
		//buf := make([]byte, 512)
		//ctn, err := conn.Read(buf)
		//if err != nil {
		//	fmt.Println("read conn error:", err)
		//	return
		//}
		//fmt.Printf("read from server: %s, ctn: %d\n", buf, ctn)
		dp := NewDatePack()
		binaryMsg, err := dp.Pack(NewMessage(0, []byte("Zinx v0.5 client Test Message")))
		if err != nil {
			fmt.Println("Pack Message error: ", err)
			return
		}
		if _, err := conn.Write(binaryMsg); err != nil {
			fmt.Println("write msg error: ", err)
			return
		}

		DataHead := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(conn, DataHead); err != nil {
			fmt.Println("Read DataHead error: ", err)
			break
		}
		msgHead, err := dp.Unpack(DataHead)
		if err != nil {
			fmt.Println("client unpack server msg error; ", err)
			break
		}
		if msgHead.GetMessageLen() > 0 {
			msg := msgHead.(*Message)
			msg.Data = make([]byte, msg.GetMessageLen())
			if _, err := io.ReadFull(conn, msg.Data); err != nil {
				fmt.Println("read msg data error: ", err)
				return
			}
			fmt.Println("recv Server Msg: Id = ", msg.Id, "len = ", msg.DateLen, "data = ", string(msg.Data))
		}
		time.Sleep(time.Second)
	}
}
