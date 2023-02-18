package sfile

import "github.com/gorilla/websocket"

// 客户端发向服务端格式
type SendMsg struct {
	Content []byte `json:"content"`
	Action  int    `json:"action"`
	MessBox string `json:"messbox"`
}
type Response struct {
	StatusCode int
	Content    []byte
	Footer     string
}

// 通用指令协议
type CommonCommand struct {
	Header   string
	Cmd      map[string]string
	Actionid int
}

type RemoteMethod interface {
	todo(*websocket.Conn, RemoteResponse)
}
type RemoteResponse interface {
	GetStatus() int
}
