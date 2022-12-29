package ziface

type IMessage interface {
	// GetMessageId 获取消息ID
	GetMessageId() uint32
	// GetMessageLen 获取消息的长度
	GetMessageLen() uint32
	// GetMessageData 获取消息的内容
	GetMessageData() []byte
	// SetMessageId 设置消息的ID
	SetMessageId(uint32)
	// SetMessageLen 设置消息的长度
	SetMessageLen(uint32)
	// SetMessageData 设置消息的内容
	SetMessageData(data []byte)
}
