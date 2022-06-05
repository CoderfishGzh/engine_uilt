package main

import (
	"engine_utils/DB"
	"engine_utils/codec"
	"engine_utils/model"
	uilt "engine_utils/uilts"
	"fmt"
	"os"
	"time"
)

var Fd *os.File
var Index []byte
var Kv_pair []string

func main() {

	DB.MysqlInit()

	Fd, Index = uilt.Index_file_init()

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

	start := time.Now() // 获取当前时间
	// get the value entry from list, type is []byte
	search_entry := list.Search([]byte("仙入"))
	search_value := search_entry.Value

	docids := uilt.Separation_value(string(search_value))
	var Docs []model.Docs
	Docs = DB.Get_docid_value(docids)

	//for _, docid := range docids {
	//	Docs = append(Docs, DB.Get_docid_value(docid))
	//}

	for _, doc := range Docs {
		fmt.Println(doc)
	}
	cost := time.Since(start) // 计算此时与start的时间差
	fmt.Println(cost)
}
