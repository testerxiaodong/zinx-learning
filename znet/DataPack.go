package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"zinx-learning/utils"
	"zinx-learning/ziface"
)

type DataPack struct {
}

func (d DataPack) GetHeadLen() uint32 {
	// Id uint32(4字节) + DataLen uint32(4字节)
	return 8
}

func (d DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	// 创建一个存放bytes字节的缓冲
	dataBuff := bytes.NewBuffer([]byte{})
	// 将DataLen写进dataBuffer中
	err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMessageLen())
	if err != nil {
		fmt.Println("write buffer error: ", err)
		return nil, err
	}
	// 将MsgId写进dataBuffer中
	err = binary.Write(dataBuff, binary.LittleEndian, msg.GetMessageId())
	if err != nil {
		fmt.Println("write buffer error: ", err)
		return nil, err
	}
	// 将data数据写进dataBuffer中
	err = binary.Write(dataBuff, binary.LittleEndian, msg.GetMessageData())
	if err != nil {
		fmt.Println("write buffer error: ", err)
		return nil, err
	}
	return dataBuff.Bytes(), nil
}

func (d DataPack) Unpack(binaryData []byte) (ziface.IMessage, error) {
	// 创建一个从输入二进制数据的ioReader
	dataBuff := bytes.NewReader(binaryData)
	// 只解压head信息，得到dataLen和MsgId
	msg := &Message{}
	// 读dataLen
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.DateLen); err != nil {
		return nil, err
	}
	// 读MsgId
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}
	// 判断dataLen是否超过我们允许的最大包长度
	if utils.GlobalObject.MaxPackageSize > 0 && msg.DateLen > utils.GlobalObject.MaxPackageSize {
		return nil, errors.New("too large msg data recv")
	}
	return msg, nil
}

func NewDatePack() *DataPack {
	return &DataPack{}
}
