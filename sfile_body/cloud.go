package sfile

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
)

// 上传单文件至云端
func UploadFile(filename string) {
	list := ParseList(filemap)
	if _, ok := list[filename]; !ok {
		fmt.Println(list[filename], " is not in filesystem")
		return
	}
	//从本地读取文件至内存
	f, err := os.OpenFile(list[filename], os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("cant read target file in local")
		errorlog.Println(err)
		return
	}
	var buff = make([]byte, wrbuffsize)
	lang, err := f.Read(buff)
	if err != nil {
		fmt.Println("cant read target file in local")
		errorlog.Println(err)
		return
	}
	//配置连接讯息
	config := ParseList(siteconf)
	value, ok := config["cloud"]
	if !ok || !strings.ContainsRune(value, '@') {
		fmt.Println("host not set")
		return
	}
	sendmsg := &SendMsg{Content: buff[:lang], Action: 1, MessBox: filename}
	SendContentToHost(value, sendmsg, 1)
}

// 从云端拉取文件至本地
func PullFile(filename string) {
	//配置连接讯息
	config := ParseList(siteconf)
	value, ok := config["cloud"]
	if !ok || !strings.ContainsRune(value, '@') {
		fmt.Println("host not set")
		return
	}
	sendmsg := &SendMsg{Action: 2, MessBox: filename}
	SendContentToHost(value, sendmsg, 2)
}
func SendContentToHost(coninfo string, content any, actionid int) {
	coninfoarr := strings.Split(coninfo, "@")
	if len(coninfoarr) != 2 {
		fmt.Println("host info is bad")
		return
	}
	hostadd := "ws://" + coninfoarr[1] + "/singlefile?otoken=" + coninfoarr[0]
	dl := websocket.Dialer{}
	ws, _, err := dl.Dial(hostadd, nil)
	if err != nil {
		fmt.Println("cant not connect to host")
		errorlog.Println(err)
		ws.Close()
	} else {
		var statuschannel = make(chan bool)
		go func() {
		sendcontent:
			err := ws.WriteJSON(content)
			if err != nil {
				fmt.Println("cant write data to host")
				errorlog.Println(err)
				ws.Close()
				return
			}
			for {
				select {
				case c := <-statuschannel:
					if !c {
						fmt.Println("resend")
						goto sendcontent
					} else {
						return
					}
				}
			}
		}()
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			var resp Response
			for {
				err := ws.ReadJSON(&resp)
				if err == nil {
					switch resp.StatusCode {
					case 200:
						statuschannel <- true
						if actionid == 2 {
							SaveFile(resp.Footer, resp.Content)
						}
						wg.Done()
						break
					case 500:
						statuschannel <- false
					case 400:
						fmt.Println("cloud file system dont have this file")
						os.Exit(1)
					case 401:
						fmt.Println("args are not correct")
						os.Exit(1)
					}
				}
			}
		}()
		wg.Wait()
		fmt.Println("upload file success")
	}
}

// clean the server target file
func CleanFile(filename string) {
	config := ParseList(siteconf)
	value, ok := config["cloud"]
	if !ok || !strings.ContainsRune(value, '@') {
		fmt.Println("host not set")
		return
	}
	cmd := &CommonCommand{Header: filename, Actionid: 3}
	valarr := strings.Split(value, "@")
	hostadd := valarr[len(valarr)-1]
	url := fmt.Sprintf("ws://%v/cmdline", hostadd)
	ConnectWithWebsocket(url, cmd)
}
func uploadprivatefile(filehead string) {
	filearr := strings.Split(filehead, "/")
	if len(filearr) != 2 || len(filearr[0]) < 1 || len(filearr[1]) < 1 {
		//命名不符合 dirname/filename规范
		return
	}
	filelist := ParseList(filemap)
	if _, ok := filelist[filearr[1]]; !ok {
		fmt.Println(filearr[1], " dont exist in file system")
		return
	}
	f, err := os.OpenFile(filelist[filearr[1]], os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("can open file ", filearr[1])
		errorlog.Println(err)
		return
	}
	var buff = make([]byte, wrbuffsize)
	lang, err := f.Read(buff)
	if err != nil {
		fmt.Println("cant read file from memory")
		errorlog.Println(err)
		return
	}
	CommonFileUpload(filehead, buff[:lang], 41)

}

// 通用指令协议
func CommonAgreenment(bcmd string, act int) {
	config := ParseList(siteconf)
	value, ok := config["cloud"]
	if !ok || !strings.ContainsRune(value, '@') {
		fmt.Println("host not set")
		return
	}
	cmd := &CommonCommand{Header: bcmd, Actionid: act}
	valarr := strings.Split(value, "@")
	hostadd := valarr[len(valarr)-1]
	url := fmt.Sprintf("ws://%v/cmdline", hostadd)
	ConnectWithWebsocket(url, cmd)
}

// 通用指令加强版
func CommonAgreenmentPlus(heads string, cmd map[string]string, act int) {
	config := ParseList(siteconf)
	value, ok := config["cloud"]
	if !ok || !strings.ContainsRune(value, '@') {
		fmt.Println("host not set")
		return
	}
	cmde := &CommonCommand{Header: heads, Cmd: cmd, Actionid: act}
	valarr := strings.Split(value, "@")
	hostadd := valarr[len(valarr)-1]
	url := fmt.Sprintf("ws://%v/cmdline", hostadd)
	ConnectWithWebsocket(url, cmde)
}

// 通用指令协议(认证版)
func CommonAgreenmentSecure(bcmd string, auth AuthMethod, act int) {
	config := ParseList(siteconf)
	value, ok := config["cloud"]
	if !ok || !strings.ContainsRune(value, '@') {
		fmt.Println("host not set")
		return
	}
	cmd := &CommonCommand{Header: bcmd, Actionid: act, Auth: auth}
	valarr := strings.Split(value, "@")
	hostadd := valarr[len(valarr)-1]
	url := fmt.Sprintf("ws://%v/cmdline", hostadd)
	ConnectWithWebsocket(url, cmd)
}

// 通用文件传输协议
func CommonFileUpload(heads string, content []byte, act int) {
	sendmsg := &SendMsg{MessBox: heads, Content: content, Action: act}
	//配置连接讯息
	config := ParseList(siteconf)
	value, ok := config["cloud"]
	if !ok || !strings.ContainsRune(value, '@') {
		fmt.Println("host not set")
		return
	}
	SendContentToHost(value, sendmsg, act)
}
func CommonExchangeFile(heads string, act int) {
	cmd := &CommonCommand{Header: heads, Actionid: act}
	//配置连接讯息
	config := ParseList(siteconf)
	value, ok := config["cloud"]
	if !ok || !strings.ContainsRune(value, '@') {
		fmt.Println("host not set")
		return
	}
	valarr := strings.Split(value, "@")
	hostadd := valarr[len(valarr)-1]
	url := fmt.Sprintf("ws://%v/cmdline", hostadd)
	AcceptFile(url, cmd)
}

// 文件传输协议底层
func AcceptFile(url string, val any) {
	dl := websocket.Dialer{}
	ws, _, err := dl.Dial(url, nil)
	if err != nil {
		fmt.Println("connect to ", url, " failed")
		errorlog.Println(err)
		ws.Close()
		return
	}
	defer ws.Close()
	resp := new(Response)
	retry := 0
senddata:
	err = ws.WriteJSON(val)
	if err != nil {
		fmt.Println("cant write data to socket")
		errorlog.Println(err)
		os.Exit(-1)
	}
	err = ws.ReadJSON(resp)
	if err != nil {
		fmt.Println("accept data lost,resend data")
		if retry < 5 {
			fmt.Println("lost data start resend data")
			retry++
			goto senddata
		} else {
			fmt.Println("accept data lost")
			os.Exit(1)
		}
	}
	switch resp.StatusCode {
	case 200:
		if len(resp.Content) < 1 {
			fmt.Println("content is null")
			return
		}
		SaveFile(resp.Footer, resp.Content)
		fmt.Println("task done")
	case 400:
		fmt.Println("pull file failed")
		return
	case 401:
		fmt.Println("args not correct")
		return
	case 500:
		fmt.Println("data lost resend data")
		goto senddata
	}
}

// 文件传输协议(认证版)二代
func NewFileUploadSecure(heads string, content []byte, auth AuthMethod, act int) {
	config := ParseList(siteconf)
	value, ok := config["cloud"]
	if !ok || !strings.ContainsRune(value, '@') {
		fmt.Println("host not set")
		return
	}
	msg := &SendMsgPlus{Action: act, Auth: auth, Content: content, MessBox: heads}
	valarr := strings.Split(value, "@")
	hostadd := valarr[len(valarr)-1]
	url := fmt.Sprintf("ws://%v/fls", hostadd)
	ConnectWithWebsocket(url, msg)
}
