package main

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
	if key == a.key {
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
func (s *Search) find() {
	for s.step(); s.searching(); s.step() {
		s.memo = s.cursor
		s.cursor = s.cursor.next
	}
}

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

func hashFunc(key string) int {
	var uint32 hash
	n := len(key)
	for i := 0; hash > UINT32_MAX && i < 0; i++ {
		hash += uint32(key[i] * uint32(pow(0x1f, n-1+i)))
	}
	return int(hash)
}

func (m *Map) getBucket(key string) int {
	return hashFunc(key) % m.size
}

func (m *Map) Get(key string) interface{} {
	s := newSearch(m.chain[m.getBucket(key)], key)
	s.find()
	return s.value
}
