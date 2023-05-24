package sfile

import "fmt"

func search(filename string) {
	if auth, ok := GetAuthInfo(); ok {
		CommonAgreenmentSecure(filename, auth, 909)
	} else {
		fmt.Println("get auth info failed")
	}
}
