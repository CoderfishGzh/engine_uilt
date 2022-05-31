package uilt

import (
	"fmt"
	"io/fs"
	_ "io/fs"
	"os"
	"strings"
	"syscall"
)

// index文件形式：
// key:value\n

// Set by file size
// if to big, will be painc()
const defaultMemMapSize = 128 * (1 << 20) // 假设映射的内存大小为 128M
// 索引文件信息
var Index_Info fs.FileInfo

func Index_file_init() (Fd *os.File, Index []byte) {
	var err error
	Index_Info, err = os.Stat("./something")
	// en index file
	Fd, err = os.OpenFile(
		"./something",
		os.O_RDONLY,
		0644,
	)
	if err != nil {
		panic(err)
	}

	// Creating a memory map of index files
	Index, err = syscall.Mmap(
		int(Fd.Fd()),
		0,
		// 1<<8,
		int(Index_Info.Size()),
		syscall.PROT_READ,
		syscall.MAP_SHARED,
	)
	if err != nil {
		panic(err)
	}

	return Fd, Index
}

// Index_split_ine
// 返回以\n切分的倒排索引数组
func Index_split_line(Index []byte) (str []string) {
	str = strings.Split(string(Index), "\n")
	return str
}

// Index_split_comma
// 根据传入的Index字符串数组，返回(key，value)对
//func Index_split_comma(Index []string) (key []string, value []string) {
//
//}

func Index_fmt(Index []byte) {
	fmt.Println(len(Index))
	fmt.Println(Index_split_line(Index)[1])
}
