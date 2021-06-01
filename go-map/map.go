package main

import "fmt"

// AssocArray is the list that holds (key, value) pairs
// in the map
type AssocArray struct {
	key   string
	value interface{}
	next  *AssocArray
}

// a new associative array
func newAssocArray(key string, value interface{}) *AssocArray {
	return &AssocArray{key: key, value: value}
}

// getIf returns value if the current node has the key
func (a *AssocArray) getIf(key string) interface{} {
	if a != nil && key == a.key {
		return a.value
	}
	return nil
}

// Search represents the searching mechanisms
// used to look for values given the key
type Search struct {
	key          string
	value        interface{}
	cursor, memo *AssocArray
}

func newSearch(a *AssocArray, key string) *Search {
	return &Search{
		key:    key,
		cursor: a,
	}
}

// step checks if current cursor holds the value
func (s *Search) step() {
	s.value = s.cursor.getIf(s.key)
}

// indicate whether we have to continue or stop
// searching
func (s *Search) searching() bool {
	return s.value == nil && s.cursor != nil
}

// Given the head of an associative array, find traverses
// the list looking for the value for key
func find(a *AssocArray, key string) *Search {
	s := newSearch(a, key)
	for s.step(); s.searching(); s.step() {
		s.memo = s.cursor
		s.cursor = s.cursor.next
	}
	return s
}

// Map represents a new hash table for (key, value) pairs.
type Map struct {
	size  int
	chain []*AssocArray
}

func newMap(size int) *Map {
	return &Map{
		size:  size,
		chain: make([]*AssocArray, size),
	}
}

// raise b to the power p
func pow(b, p int) int {
	if p == 0 {
		return 1
	}
	for i := 0; i < p-1; i++ {
		b *= b
	}
	return b
}

const (
	UINT32_MAX = 0xffffffff
)

// hashFunc returns hash value of the string
func hashFunc(key string) int {
	var hash uint32
	n := len(key)
	for i := 0; hash < UINT32_MAX && i < n; i++ {
		hash += uint32(key[i]) * uint32(pow(0x1f, n-i+1))
	}
	return int(hash)
}

// getBucket takes the hash of the key and finds the right bucket
// for the key.
func (m *Map) getBucket(key string) int {
	return hashFunc(key) % m.size
}

// Get returns the value of the key if there is one.
func (m *Map) Get(key string) interface{} {
	s := find(m.chain[m.getBucket(key)], key)
	return s.value
}

// Set stores the (key, value) pair in the map
// for efficient lookup.
func (m *Map) Set(key string, value interface{}) {
	b := m.getBucket(key)
	a := m.chain[b]
	s := find(a, key)
	if s.value != nil { /* (key, value) pair already exists in map, change value to new value */
		s.cursor.value = value
	} else {
		n := newAssocArray(key, value)
		if s.cursor == a { /* if a is the first element or nil */
			n.next = s.cursor
			m.chain[b] = n
		} else if s.cursor == nil {
			s.memo.next = n
		} else { /* cursor not nil, memo not nil */
			n.next = s.cursor
			s.memo.next = n
		}
	}
}

// TODO(Joe-Degs):
// implement Delete method for the map

func main() {
	m := newMap(10)

	m.Set("Joe", "Jollof")
	m.Set("hello", "world")
	m.Set("string slice", []string{"joe"})
	m.Set("func", func(s string) {
		fmt.Println(s)
	})

	fmt.Println(m.Get("Joe"))
	fmt.Println(m.Get("hello"))
	fmt.Println(m.Get("string slice"))
	m.Get("func").(func(string))("works just fine")
}
