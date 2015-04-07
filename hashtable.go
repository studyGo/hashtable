package main

import (
	"container/list"
	"fmt"
	"hash/crc32"
)

const max int = 100000

type HashTable struct {
	size int
	item []*Items
}

type Items struct {
	length int
	l      *list.List
}

type Item struct {
	key   string
	value interface{}
}

func getIndex(key string) int {
	return int(crc32.ChecksumIEEE([]byte(key))) % max

}

func New(size int) *HashTable {
	return &HashTable{size: size, item: make([]*Items, max)}
}

func (table *HashTable) Add(key string, value interface{}) {
	index := getIndex(key)
	if table.item[index] == nil {
		table.item[index] = &Items{0, list.New()}
		table.item[index].l.PushFront(&Item{key, value})
		table.item[index].length = 1
	} else {
		table.item[index].l.PushBack(&Item{key, value})
		table.item[index].length += 1

	}
}

func (table *HashTable) Get(key string) interface{} {
	index := getIndex(key)
	l := table.item[index].l
	for s := l.Front(); s != nil; s = s.Next() {
		sInfo := s.Value.(*Item)
		if sInfo.key == key {
			return sInfo.value
		}
	}
	return nil
}

func main() {
	s := New(100)
	s.Add("demo", 100)
	fmt.Println(s.Get("demo"))

}
