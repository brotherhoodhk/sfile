package sfile

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/oswaldoooo/octools/toolsbox"
)

func Test() {
	siteinfo := ParseList(siteconf)
	if ve, ok := siteinfo["cloud"]; ok {
		infoarr := strings.Split(ve, "@")
		usrinfo := strings.Join(infoarr[:len(infoarr)-1], "@")
		usrinfoarr := strings.Split(usrinfo, ":")
		usrname := usrinfoarr[0]
		pwd := strings.Join(usrinfoarr[1:], ":")
		auth := AuthMethod{Usrname: []byte(usrname), Key: toolsbox.Sha256([]byte(pwd))}
		CommonAgreenmentSecure("", auth, 900)
	}
}
func TestMkdir() {
	siteinfo := ParseList(siteconf)
	if ve, ok := siteinfo["cloud"]; ok {
		infoarr := strings.Split(ve, "@")
		usrinfo := strings.Join(infoarr[:len(infoarr)-1], "@")
		usrinfoarr := strings.Split(usrinfo, ":")
		usrname := usrinfoarr[0]
		pwd := strings.Join(usrinfoarr[1:], ":")
		auth := AuthMethod{Usrname: []byte(usrname), Key: toolsbox.Sha256([]byte(pwd))}
		CommonAgreenmentSecure("testdir", auth, 940)
	}
}
func TestUploadPrivateFile() {
	f, err := os.OpenFile(ROORPATH+"/go.mod", os.O_RDONLY, 0700)
	if err != nil {
		fmt.Println(err)
		return
	}
	read := bufio.NewReader(f)
	buff := make([]byte, 10*MB)
	n, err := read.Read(buff)
	if err != nil {
		fmt.Println(err)
		return
	}
	if auth, ok := GetAuthInfo(); ok {
		NewFileUploadSecure("testdir/go.mod", buff[:n], auth, 941)
	}
}
func TestPullFile() {
	if auth, ok := GetAuthInfo(); ok {
		CommonExchangeFilePlus("testdir/go.mod", auth, 942)
	}
}
func TestDeleteDir() {
	if auth, ok := GetAuthInfo(); ok {
		CommonAgreenmentSecure("testdir", auth, 943)
	}
}
func TestDeleteFile() {
	if auth, ok := GetAuthInfo(); ok {
		CommonAgreenmentSecure("testdir/go.mod", auth, 9431)
	}
}
