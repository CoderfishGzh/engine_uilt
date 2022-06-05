package uilt

import (
	"io/fs"
	_ "io/fs"
	"os"
	"strings"
	"syscall"
)

// index文件形式：
// key,[value]\n

// Set by file size
// Todo 内存有换页机制，映射时间长了，可能会被换走
const defaultMemMapSize = 128 * (1 << 20) // 假设映射的内存大小为 128M

// 索引文件信息
var Index_Info fs.FileInfo

func Index_file_init() (Fd *os.File, Index []byte) {
	var err error
	Index_Info, err = os.Stat("/home/oem/下载/dictionary1.txt")
	// en index file
	Fd, err = os.OpenFile(
		"/home/oem/下载/dictionary1.txt",
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
		//defaultMemMapSize,
		syscall.PROT_READ,
		syscall.MAP_SHARED,
	)
	if err != nil {
		panic(err)
	}
	return Fd, Index
}

// Index_by_enter
// 返回索引所有的kv对
func Index_by_enter(Index []byte) []string {
	str := strings.Split(string(Index), "\n")
	return str
}

// Separation_kv
// 从kv对中分离key和value
// 成功返回1，失败返回0
func Separation_kv(kv_pair string) (key []byte, value []byte, ret int) {
	str := strings.Split(kv_pair, ",")

	if len(str) < 2 {
		return key, value, 0
	}
	key = []byte(str[0])
	value = []byte(str[1])
	return key, value, 1
}

// Separation_value
// 划分value中的docid
func Separation_value(values_be string) (docids []string) {
	// 切分[docid1 docid2],取出docic数组
	values_l := strings.Split(values_be, "[")
	values := strings.Split(values_l[1], "]")
	// 切分docid1 docid2,取出docic数组
	docids = strings.Split(values[0], " ")
	return docids
}
