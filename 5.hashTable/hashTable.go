package main

import "fmt"

//ArraySize is the size of the hash table array
const ArraySize = 7

//hastable will hold an array
type HashTable struct {
	myarray [ArraySize]*bucket
}

//bucket is a linkedlist in each of the slot of the hashTable array
type bucket struct {
	head *bucketNode
}

//bucketnode structure
type bucketNode struct {
	key  string
	next *bucketNode
}

//insert will take in a key and add it to the hash table array
func (h *HashTable) Insert(key string) {
	index := Hash(key)
	h.myarray[index].insert(key)
}

//search will take in a key and return true if that key is stored in the hash
func (h *HashTable) Search(key string) bool {
	index := Hash(key)
	return h.myarray[index].search(key)
}

//delete will take in a key and delete it from the hash table
func (h *HashTable) Delete(key string) {
	index := Hash(key)
	h.myarray[index].delete(key)
}

//insert will take in a key, create a node with the key
func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println(k, "already exist")
	}

}

//search will take in a key and return true if the bucket has that key
func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

//delete will take in a key and delete the node from the bucket
func (b *bucket) delete(k string) {

	if b.head.key == k {
		b.head = b.head.next
		return
	}
	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == k {
			//delete
			previousNode.next = previousNode.next.next
		}
		previousNode = previousNode.next
	}
}

func Hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

//init will create a bucket in each slot of the hash table
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.myarray {
		result.myarray[i] = &bucket{}
	}
	return result
}

func main() {
	hashTable := Init()

	list := []string{
		"ERIC",
		"KENNY",
		"KYLE",
		"STAN",
		"RANDY",
		"BUTTERS",
		"TOKEN",
	}

	for _, v := range list {
		hashTable.Insert(v)
	}

	hashTable.Delete("RANDY")

	fmt.Println(hashTable.Search("ERIC"))
	fmt.Println(hashTable.Search("RANDY"))

}
