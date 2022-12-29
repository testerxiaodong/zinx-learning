package ziface

type IRequest interface {
	// GetCurrConn 得到当前链接
	GetCurrConn() IConnection
	// GetData 得到请求的消息数据
	GetData() []byte
	// GetMessageId 得到消息的Id
	GetMessageId() uint32
}
