package main

import (
	"engine_utils/codec"
	"engine_utils/uilts"
	"fmt"
	"os"
	"time"
)

var Fd *os.File
var Index []byte
var Kv_pair []string

func main() {
	Fd, Index = uilt.Index_file_init()
	start := time.Now() // 获取当前时间
	Kv_pair = uilt.Index_by_enter(Index)
	fmt.Println(len(Kv_pair))
	// crate skiplist
	list := uilt.NewSkipList()
	// insert all index
	for i := 0; i < 12193028; i++ {
		key, value, err := uilt.Separation_kv(Kv_pair[i])
		if err == 0 {
			continue
		}
		list.Add(codec.NewEntry(key, value))
	}
	cost := time.Since(start) // 计算此时与start的时间差
	fmt.Println(cost)

	// time.Now().Sub(start)也可
	search_value := list.Search([]byte("回电瓶"))
	
	fmt.Println(string(search_value.Value))
}
