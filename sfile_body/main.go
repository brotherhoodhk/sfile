package sfile

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"strings"
)

const (
	VERSION      = "2.0Lab"
	SITE         = "https://brotherhoodhk.org/codelabcn/sfile"
	softwareinfo = "version: " + VERSION + "\nour site: " + SITE
)
const (
	AUTHGETWARN = "failed to get authentication"
)

var ROORPATH = os.Getenv("SFILE_HOME")
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
		// AddFile(args[1])
		AddfileInterface("", args[1:]...)
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
		switch len(args) {
		case 2:
			UploadFile(args[1])
		case 3:
			if args[1] == "--private" {
				if strings.ContainsRune(args[2], '/') {
					uploadprivatefile(args[2])
				}
			}
		default:
			Error()
			return
		}
	case "pull":
		switch len(args) {
		case 2:
			PullFile(args[1])
		case 3:
			if args[1] == "--private" && isprivatefilename(args[2]) {
				// CommonExchangeFile(args[2], 42)
				if auth, ok := GetAuthInfo(); ok {
					CommonExchangeFilePlus(args[2], auth, 842)
				} else {
					fmt.Println(AUTHGETWARN)
				}
			} else {
				Error()
			}
		default:
			Error()
		}
	case "clean":
		switch len(args) {
		case 2:
			CleanFile(args[1])
		case 3:
			if args[1] == "--private" {
				if strings.ContainsRune(args[2], '/') {
					// CommonAgreenment(args[2], 431)
					if auth, ok := GetAuthInfo(); ok {
						CommonAgreenmentSecure(args[2], auth, 8431)
					} else {
						fmt.Println(AUTHGETWARN)
					}
				} else {
					fmt.Println("filename is not correct")
				}
			} else if args[1] == "-r" {
				if strings.ContainsRune(args[2], '/') {
					fmt.Println("your dirname is not correct")
				} else {
					// CommonAgreenment(args[2], 43)
					if auth, ok := GetAuthInfo(); ok {
						CommonAgreenmentSecure(args[2], auth, 843)
					} else {
						fmt.Println(AUTHGETWARN)
					}
				}
			}
		}
	case "mkdir":
		if len(args) < 2 {
			Error()
			return
		}
		// CommonAgreenment(args[1], 40)
		if auth, ok := GetAuthInfo(); ok {
			CommonAgreenmentSecure(args[1], auth, 840)
		} else {
			fmt.Println(AUTHGETWARN)
		}
	case "config":
		if len(args) < 2 {
			Error()
			fmt.Println("learn about sfile command from here=>https://brotherhoodhk.org/codelabcn/sfile/tutorial")
			return
		}
		ConfigureSfile(args[1:])
	case "clear":
		if len(args) == 1 {
			ClearLocalFS()
		} else {
			Error()
		}
	case "test":
		// Test()
		TestMkdir()
		TestUploadPrivateFile()
		TestPullFile()
		TestDeleteFile()
		TestDeleteDir()
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
	filename = getfilename(filename)
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

// the outside addfile
func AddfileInterface(basename string, dirname ...string) {
	var fe fs.FileInfo
	var fearr []fs.FileInfo
	var err error
	for _, v := range dirname {
		basename += v + "/"
		fe, err = os.Stat(v)
		if err != nil {
			fmt.Println(v, "is not exist")
			errorlog.Println(err)
		} else {
			if fe.IsDir() { //if the name is directory name
				fearr, err = ioutil.ReadDir(v)
				if err != nil {
					fmt.Println("read directory", v, "failed")
					errorlog.Println(err)
				} else {
					for _, v := range fearr {
						if v.IsDir() {
							AddfileInterface(basename, v.Name())
						} else {
							AddFile(basename + v.Name())
						}
					}
				}
			} else {
				AddFile(v)
			}
		}
	}
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
