package sfile

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
