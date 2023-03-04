package sfile

import (
	"fmt"
	"log"
	"os"

	"github.com/gorilla/websocket"
)

const (
	KB = 1024
	MB = 1024 * KB
	GB = 1024 * MB
)

var wrbuffsize = 100 * MB

func LogInit(logname string) *log.Logger {
	f, err := os.OpenFile(LOGPATH+logname+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed,", err)
		return nil
	}
	newlog := log.New(f, "["+logname+"]", log.LUTC|log.Lshortfile|log.LstdFlags)
	return newlog
}
func Exist_File(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}
func (s *CommonCommand) todo(con *websocket.Conn, resp RemoteResponse) {
senddata:
	err := con.WriteJSON(s)
	if err != nil {
		fmt.Println("write json to websocket failed")
		errorlog.Println(err)
		os.Exit(-1)
	}
	err = con.ReadJSON(resp)
	if err != nil {
		//if data broken ,resend data
		goto senddata
	}
	switch resp.GetStatus() {
	case 200:
		fmt.Println("task done")
	case 400:
		fmt.Println("task failed")
	case 401:
		fmt.Println("args not correct")
	case 402:
		fmt.Println("permission denied")
	case 500:
		goto senddata
	}
}
func (s *SendMsgPlus) todo(con *websocket.Conn, resp RemoteResponse) {
senddata:
	err := con.WriteJSON(s)
	if err != nil {
		fmt.Println("write json to websocket failed")
		errorlog.Println(err)
		os.Exit(-1)
	}
	err = con.ReadJSON(resp)
	if err != nil {
		//if data broken ,resend data
		goto senddata
	}
	if status(resp.GetStatus()) {
		goto senddata
	}
}
func (s *Response) GetStatus() int {
	return s.StatusCode
}

// if return true ,it's resend signle
func status(code int) bool {
	switch code {
	case 200:
		fmt.Println("task done")
	case 400:
		fmt.Println("task failed")
	case 401:
		fmt.Println("args not correct")
	case 402:
		fmt.Println("permission denied")
	case 500:
		return true
	}
	return false
}
