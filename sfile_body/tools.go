package sfile

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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
func SaveFile(filename string,content []byte) {
	list:=ParseList(filemap)
	if _,ok:=list[filename];!ok{
		nowpath,err:=os.Getwd()
		if err!=nil{
			fmt.Println("get path error")
			errorlog.Println(err)
			os.Exit(-1)
		}
		fullpath:=nowpath+"/"+filename
		list[filename]=fullpath
		FormatList(list,filemap)
	}
	val:=list[filename]
	f,err:=os.OpenFile(val,os.O_CREATE|os.O_TRUNC|os.O_WRONLY,0666)
	if err!=nil{
		fmt.Println("write error")
		errorlog.Println(err)
		f.Close()
		os.Exit(-1)
	}
	_,err=f.Write(content)
	if err!=nil{
		fmt.Println("write error")
		errorlog.Println(err)
	}
	f.Close()
}
