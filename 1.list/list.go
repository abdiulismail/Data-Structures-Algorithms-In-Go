package main

import (
	"container/list"
	"fmt"
)

func main() {

	var intlist list.List

	intlist.PushBack(11)
	intlist.PushBack(23)
	intlist.PushBack(34)

	for element := intlist.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value.(int))
	}
}
