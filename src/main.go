package main

import "os"
import (
	"sfile/sfile_body"
)

func main() {
	args := os.Args[1:]
	sfile.SfileStart(args)
}
