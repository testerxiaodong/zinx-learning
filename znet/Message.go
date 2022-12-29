package znet

type Message struct {
	Id      uint32 // 消息Id
	DateLen uint32 // 消息长度
	Data    []byte // 消息内容
}

func NewMessage(id uint32, data []byte) *Message {
	return &Message{
		id,
		uint32(len(data)),
		data,
	}
}

func (m *Message) GetMessageId() uint32 {
	return m.Id
}

func (m *Message) GetMessageLen() uint32 {
	return m.DateLen
}

func (m *Message) GetMessageData() []byte {
	return m.Data
}

func (m *Message) SetMessageId(u uint32) {
	m.Id = u
}

func (m *Message) SetMessageLen(u uint32) {
	m.DateLen = u
}

func (m *Message) SetMessageData(data []byte) {
	m.Data = data
}
