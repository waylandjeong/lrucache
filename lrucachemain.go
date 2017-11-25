package main

import (
	"fmt"
	"container/list"
	"math/rand"
)

type Node struct {
	value string
	uuid int
}

type CacheEntry struct {
	address uint32
	value uint32
}

// CacheMem - Cachemem object for implementing the LRU cache
// Cache modeled as a hashmap with address as the key and a cachentry as the value
// The cachentry has an address and value pair
// Methods:
//    addEntry
//    delEntry
//    getSize
//    getEntry
//
type CacheMem struct {
	cMap map[uint32] *CacheEntry
	lruHist *list.List
	cSize uint
}

// getSize method
func (c CacheMem) getSize() uint {
	return c.cSize
}

// setSize method
func (c *CacheMem) setSize(s uint) {
	c.cSize = s
}

// constructor method
func makeCache() *CacheMem {
	m := new(CacheMem)
	m.cMap = make(map[uint32] *CacheEntry)
	m.lruHist = list.New()
	m.cSize = 0
	return m
}

// isFull method
func (c CacheMem) isFull() bool {
	return len(c.cMap) >= CACHE_SIZE
}

// removeLRU method
func (c *CacheMem) removeLRU() {

	beforeSize := len(c.cMap)
	temp := c.lruHist.Front()
	_ = c.lruHist.Remove(temp)
	cNode := temp.Value.(*CacheEntry)
	delete(c.cMap, cNode.address)
	c.cSize = uint(len(c.cMap))
	fmt.Println("DEBUG: cnode.address = ", cNode.address, " b/a = ", beforeSize, " / ", c.cSize, " lrhHist Size = ", c.lruHist.Len())

}

// addEntry method
func (c *CacheMem) addEntry(address uint32, value uint32) *CacheEntry {

	// check first if cache is full
	if c.isFull() {
		fmt.Println("DEBUG: cache is FULL in addEntry, removing LRU")
		c.removeLRU()
	}

	// check if address is in range
	if address > ADDRESS_RANGE {
		return nil
	}

	// add entry into cache
	e := new(CacheEntry)
	e.value = value
	e.address = address
	c.cMap[address] = e

	// add entry into LRU history
	c.lruHist.PushBack(e)

	// update cache size
	c.cSize = uint(len(c.cMap))

	// return cache entry
	return e

}

const CACHE_SIZE =128
const ADDRESS_RANGE = 1024

func main() {

	// Define and initialize the cache
	var lruCache *CacheMem
	lruCache = makeCache()

	fmt.Println("Initialized cache with size ", lruCache.getSize())

	if lruCache.isFull() {
		fmt.Println("Cache is FULL")
	} else {
		fmt.Println("Cache is NOT FULL")
	}

	// load the cache
	for i := 0; i < 2 * CACHE_SIZE; i++ {
		k := rand.Uint32() % ADDRESS_RANGE
		v := uint32(i)
		// fmt.Println("DEBUG: adding k = ", k, "v = ", v)
		e := lruCache.addEntry(k, v)
		if e == nil {
			fmt.Println("ERROR! addEntry failed k = ", k, "v = ", v)
		}
	}

	fmt.Println("Filled cache with size ", lruCache.getSize())

	// define hash table
	var m map[int]int
	m = make(map[int]int)

	for i := 0; i < CACHE_SIZE; i++ {
		m[rand.Intn(ADDRESS_RANGE)] = i
	}

	fmt.Println("Welcome to my LRU Cache code")

	fmt.Println("Cache initialized to size", len(m))

	//for key, value := range m {
	//	fmt.Println("Key:", key, "Value:", value)
	//}

	linkedlist := list.New()

	fmt.Println("Size before ", linkedlist.Len())

	linkedlist.PushBack(Node{"aaaa", 1234})
	linkedlist.PushBack(Node{"bbbb", 5678})
	linkedlist.PushBack(Node{"cccc", 9101})

	fmt.Println("Size after ", linkedlist.Len())

	var anode Node
	for e := linkedlist.Front(); e != nil; e = e.Next() {
		anode = e.Value.(Node)
		fmt.Println("value = ", anode.value, " uuid = ", anode.uuid)
	}

	frontE := linkedlist.Front()
	tempE := linkedlist.Remove(frontE)
	tempN := tempE.(Node)
	fmt.Printf("Removed [%s] from list\n", tempN.value)

	fmt.Println("List after removal")
	for e := linkedlist.Front(); e != nil; e = e.Next() {
		anode = e.Value.(Node)
		fmt.Println("value = ", anode.value, " uuid = ", anode.uuid)
	}
}
