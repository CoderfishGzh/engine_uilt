package main

import (
	"engine_utils/uilts"
	"os"
)

var Fd *os.File
var Index []byte

func main() {
	Fd, Index = uilt.Index_file_init()
	uilt.Index_fmt(Index)
}
