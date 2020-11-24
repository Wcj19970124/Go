package socket

import (
	"github.com/gorilla/websocket"
)

//升级请求 http->webSocket
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

//GetMessage 获取消息
func GetMessage() {

}

//SendMessage 发送消息
func SendMessage() {

}
