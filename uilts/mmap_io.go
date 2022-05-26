package uilt

import (
	"fmt"
	"os"
	"strings"
	"syscall"
)

// Set by file size
// if to big, will be painc()
const defaultMemMapSize = 128 * (1 << 20) // 假设映射的内存大小为 128M

func Index_file_init() (Fd *os.File, Index []byte) {
	var err error
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
		1<<8,
		syscall.PROT_READ,
		syscall.MAP_SHARED,
	)
	if err != nil {
		panic(err)
	}

	return Fd, Index
}

func Index_split(Index []byte) (str []string) {
	str = strings.Split(string(Index), "\n")
	return str
}

func Index_fmt(Index []byte) {
	fmt.Println(Index_split(Index))
}

// Todo
func Index_insert(Index []byte) {

}
