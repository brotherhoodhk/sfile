package main

import (
	"fmt"
	"strings"
)

var StatusCode = 909

func RespParser(footer string, content []byte) {
	if len(footer) > 0 && footer != " " {
		footer = strings.ReplaceAll(footer, " ", "\n")
	}
	fmt.Println(footer)
}
