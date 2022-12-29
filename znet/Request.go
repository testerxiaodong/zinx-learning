package znet

import "zinx-learning/ziface"

type Request struct {
	// 已经和客户端建立好的链接
	conn ziface.IConnection
	// 客户端请求的数据
	msg ziface.IMessage
}

func (r *Request) GetCurrConn() ziface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.msg.GetMessageData()
}

func (r *Request) GetMessageId() uint32 {
	return r.msg.GetMessageId()
}
