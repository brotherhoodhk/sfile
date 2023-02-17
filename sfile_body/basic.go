package sfile

import (
	"fmt"
	"log"
	"os"
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
