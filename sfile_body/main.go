package sfile

import (
	"fmt"
	"os"
)

const (
	VERSION = "2.0Lab"
	SITE    = "https://brotherhoodhk.org"
)

var ROORPATH = os.Getenv("SFILE_HOME")
var softwareinfo = "version: " + VERSION + "\nour site: " + SITE
var LOGPATH = ROORPATH + "/log/"
var errorlog = LogInit("error")
var filemap = ROORPATH + "/conf/filemap.cnf"
var siteconf = ROORPATH + "/conf/site.cnf"

func SfileStart(args []string) {
	if len(args) == 0 {
		fmt.Println(softwareinfo)
		return
	}
	switch args[0] {
	case "add":
		if len(args) < 2 {
			Error()
			return
		}
		AddFile(args[1])
	case "list":
		ShowList()
	case "get":
		if len(args) < 2 {
			Error()
			return
		}
		GetFile(args[1])
	case "rm":
		if len(args) < 2 {
			Error()
			return
		}
		RemoveFile(args[1])
	case "upload":
		if len(args) < 2 {
			Error()
			return
		}
		UploadFile(args[1])
	case "pull":
		if len(args) < 2 {
			Error()
			return
		}
		PullFile(args[1])
	}

}
func Error() {
	fmt.Println("not this command")
}

// put file into file system
func AddFile(filename string) {
	list := ParseList(filemap)
	nowpath, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	filepath := nowpath + "/" + filename
	if !Exist_File(filepath) {
		fmt.Println("file" + filename + " dont exist")
		os.Exit(1)
	}
	if _, ok := list[filename]; !ok {
		list[filename] = filepath
	} else {
		var buff = make([]byte, wrbuffsize)
		f, err := os.OpenFile(filepath, os.O_RDONLY, 0666)
		if err != nil {
			fmt.Println("open target file failed")
			errorlog.Println(err)
			return
		}
		lang, err := f.Read(buff)
		if err != nil {
			fmt.Println("read from cache failed")
			errorlog.Println(err)
			return
		}
		f, err = os.OpenFile(list[filename], os.O_TRUNC|os.O_WRONLY, 0666)
		if err != nil {
			fmt.Println("origin file is not exsit,this file will instead it")
			errorlog.Println("origin file " + list[filename] + " is not exsit,this file " + filepath + " will instead it")
			list[filename] = filepath
		} else {
			_, err := f.Write(buff[:lang])
			if err != nil {
				fmt.Println("write to origin file failed")
				errorlog.Println("write to " + list[filename] + " failed")
			}
		}
		f.Close()
	}
	FormatList(list, filemap)
}

// show all local file system
func ShowList() {
	list := ParseList(filemap)
	resout := ""
	for k, v := range list {
		resout += fmt.Sprintf("%-20s %s\n", k, v)
	}
	fmt.Println(resout)
}

// delete file info from filesystem
func RemoveFile(filename string) {
	list := ParseList(filemap)
	if _, ok := list[filename]; ok {
		delete(list, filename)
		FormatList(list, filemap)
	} else {
		fmt.Println("not this file")
	}
}

// get file from file system
func GetFile(filename string) {
	list := ParseList(filemap)
	if _, ok := list[filename]; !ok {
		fmt.Println(filename + " dont exsit in filesystem")
		return
	}
	var buff = make([]byte, wrbuffsize)
	f, err := os.OpenFile(list[filename], os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("cant open origin file,it might not exist")
		errorlog.Println("open " + list[filename] + " failed")
		return
	}
	lang, err := f.Read(buff)
	if err != nil {
		fmt.Println("cant read from cache")
		errorlog.Println(err)
		return
	}
	nowpath, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		errorlog.Println(err)
		return
	}
	filepath := nowpath + "/" + filename
	f, err = os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("cant open file in current dir")
		errorlog.Println(err)
	}
	_, err = f.Write(buff[:lang])
	if err != nil {
		fmt.Println("cant write to file")
		errorlog.Println("cant write to " + filepath)
	}
}
