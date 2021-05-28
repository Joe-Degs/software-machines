package memory

import (
	"reflect"
	"unsafe"
)

// Memory is a heap data structure and its word aligned, contigous and byte
// addressable
type Memory []uintptr

var (
	// holds the type []byte
	BYTE_SLICE = reflect.TypeOf([]byte(nil))
	// holds the type Memory
	MEMORY = reflect.TypeOf(Memory{})
	// size of an element in the heap, might be
	// 4 or 8 bytes depending on the cpu architecture.
	MEMORY_BYTES = int(MEMORY.Elem().Size())
)

// newHeader returns a new slice header with the appropriate
// sizes depending on the cpu architecture
func (m Memory) newHeader() (h reflect.SliceHeader) {
	h = *(*reflect.SliceHeader)(unsafe.Pointer(&m))
	h.Len = len(m) * MEMORY_BYTES
	h.Cap = cap(m) * MEMORY_BYTES
	return
}

// Bytes returns a byte slice of the memory, basically pointing
// to the same memory location. This is what makes the memory byte
// addressable i think.
func (m *Memory) Bytes() []byte {
	h := m.newHeader()
	return *(*[]byte)(unsafe.Pointer(&h))
}

// Serialise returns a duplicate of the memory
func (m *Memory) Serialise() []byte {
	h := m.newHeader()
	b := make([]byte, h.Len)
	copy(b, *(*[]byte)(unsafe.Pointer(&h)))
	return b
}

// Overwrites the memory with the given interface.
// could be a byte slice or memory type
func (m *Memory) Overwrite(i interface{}) {
	switch i := i.(type) {
	case Memory:
		copy(*m, i)
	case []byte:
		copy(m.Bytes(), i)
	}
}
