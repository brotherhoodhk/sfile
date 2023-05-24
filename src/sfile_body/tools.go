package sfile

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/oswaldoooo/octools/toolsbox"
)

func ParseList(path string) map[string]string {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		errorlog.Println(err)
		return nil
	} else if len(f) < 3 {
		return make(map[string]string)
	}
	content := string(f)
	basicarr := strings.Split(content, "\n")
	var namelist = make(map[string]string)
	for _, v := range basicarr {
		if len(v) > 2 {
			resarr := strings.Split(v, "=")
			if len(resarr) == 2 {
				//name=path
				namelist[resarr[0]] = resarr[1]
			}
		}
	}
	return namelist
}
func FormatList(origin map[string]string, path string) bool {
	recordmsg := ""
	for k, v := range origin {
		recordmsg += k + "=" + v + "\n"
	}
	err := ioutil.WriteFile(path, []byte(recordmsg), 0666)
	if err != nil {
		fmt.Println("write list to file error")
		errorlog.Println(err)
		return false
	}
	return true
}

// 将文件储存至本地file system
func SaveFile(filename string, content []byte) {
	list := ParseList(filemap)
	if _, ok := list[filename]; !ok {
		nowpath, err := os.Getwd()
		if err != nil {
			fmt.Println("get path error")
			errorlog.Println(err)
			os.Exit(-1)
		}
		fullpath := nowpath + "/" + filename
		list[filename] = fullpath
		FormatList(list, filemap)
	}
	val := list[filename]
	f, err := os.OpenFile(val, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("write error")
		errorlog.Println(err)
		f.Close()
		os.Exit(-1)
	}
	_, err = f.Write(content)
	if err != nil {
		fmt.Println("write error")
		errorlog.Println(err)
	}
	f.Close()
}

func ConnectWithWebsocket(url string, rmd RemoteMethod) {
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
	rmd.todo(ws, resp)
}
func isprivatefilename(filename string) bool {
	if !strings.ContainsRune(filename, '/') {
		return false
	}
	namearr := strings.Split(filename, "/")
	if len(namearr) != 2 || len(namearr[0]) < 1 || len(namearr[1]) < 1 {
		return false
	}
	return true
}

// check the remote info
func CheckRemoteInfo(target string) bool {
	if strings.ContainsRune(target, '@') && strings.Count(target, ":") >= 2 && len(target) > 10 {
		addarr := strings.Split(target, "@")
		if strings.ContainsRune(addarr[0], ':') && len(addarr[0]) > 3 && strings.ContainsRune(addarr[1], ':') && len(addarr[1]) > 6 {
			return true
		}
	}
	return false
}

// 获取认证信息
func GetAuthInfo() (AuthMethod, bool) {
	siteinfo := ParseList(siteconf)
	if ve, ok := siteinfo["cloud"]; ok {
		infoarr := strings.Split(ve, "@")
		usrinfo := strings.Join(infoarr[:len(infoarr)-1], "@")
		usrinfoarr := strings.Split(usrinfo, ":")
		usrname := usrinfoarr[0]
		pwd := strings.Join(usrinfoarr[1:], ":")
		auth := AuthMethod{Usrname: []byte(usrname), Key: toolsbox.Sha256([]byte(pwd))}
		return auth, true
	}
	return AuthMethod{}, false
}

// 从路径中取得文件名
func getfilename(originname string) (finalname string) {
	if strings.ContainsRune(originname, '/') {
		namearr := strings.Split(originname, "/")
		finalname = namearr[len(namearr)-1]
	} else {
		finalname = originname
	}
	return
}

// 清除本地记录的所有链接
func ClearLocalFS() {
	filelist := ParseList(filemap)
	filelist = make(map[string]string)
	FormatList(filelist, filemap)
}
