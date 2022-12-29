package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"zinx-learning/ziface"
)

type GlobalObj struct {
	/*
		Server
	*/
	TcpServer ziface.IServer // 当前Zinx全局的Server对象
	Host      string         // 当前服务器主机监听的IP
	TcpPort   int            // 当前服务器主机监听的Port
	Name      string         // 当前服务器的名称
	/*
		Client
	*/
	Version        string // 当前Zinx的版本号
	MaxConn        int    // 当前服务器主机允许的最大链接数
	MaxPackageSize uint32 // 当前Zinx框架数据包的最大值
}

var GlobalObject *GlobalObj

func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("conf/zinx.json")
	if err != nil {
		fmt.Println("read zinx config file error: ", err)
		return
	}
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		fmt.Println("config json file Unmarshal error: ", err)
	}
	fmt.Println("reload config success")
}

func init() {
	fmt.Println("init default server config")
	GlobalObject = &GlobalObj{
		Name:           "ZinxServerApp",
		Version:        "V0.4",
		TcpPort:        8999,
		Host:           "0.0.0.0",
		MaxConn:        1000,
		MaxPackageSize: 4096,
	}
	GlobalObject.Reload()
}
