package main

import (
	"fmt"

	. "github.com/Joe-Degs/software-machines/go-heap/memory"
)

func main() {
	m := make(Memory, 2)
	b := m.Bytes()
	s := m.Serialise()
	fmt.Println("m (cells) =", len(m), "of", cap(m), ":", m)
	fmt.Println("b (bytes) =", len(b), "of", cap(b), ":", b)
	fmt.Println("s (bytes) =", len(s), "of", cap(s), ":", s)
	fmt.Println("")

	// overwrite with memory type
	m.Overwrite(Memory{3, 4})
	fmt.Println("overwrite with Memory{3,4}")
	fmt.Println("m (cells) =", len(m), "of", cap(m), ":", m)
	fmt.Println("b (bytes) =", len(b), "of", cap(b), ":", b)
	fmt.Println("s (bytes) =", len(s), "of", cap(s), ":", s)
	fmt.Println("")

	// serialise and overwrite with a byte slice
	s = m.Serialise()
	m.Overwrite([]byte{8, 7, 6, 5, 4, 3, 2, 1})
	fmt.Println("m (cells) =", len(m), "of", cap(m), ":", m)
	fmt.Println("b (cells) =", len(b), "of", cap(b), ":", b)
	fmt.Println("s (cells) =", len(s), "of", cap(s), ":", s)
	fmt.Println("")

	// overwrite byte slice
	m.Overwrite([]byte{10})
	fmt.Println("m (cells) =", len(m), "of", cap(m), ":", m)
	fmt.Println("b (cells) =", len(b), "of", cap(b), ":", b)
	fmt.Println("s (cells) =", len(s), "of", cap(s), ":", s)
	fmt.Println("")

	// proofing byte addressability
	b[15] = 0xa
	fmt.Println("m (cells) =", len(m), "of", cap(m), ":", m)
	fmt.Println("b (cells) =", len(b), "of", cap(b), ":", b)
	fmt.Println("s (cells) =", len(s), "of", cap(s), ":", s)
	fmt.Println("")

	// proofing byte addressability
	b[8] = 0
	fmt.Println("m (cells) =", len(m), "of", cap(m), ":", m)
	fmt.Println("b (cells) =", len(b), "of", cap(b), ":", b)
	fmt.Println("s (cells) =", len(s), "of", cap(s), ":", s)
	fmt.Println("")

	// serialise and proof that s a duplicate copy chnages to
	// it does not affect the memory buffer itself.
	s = m.Serialise()
	s[9] = 10
	fmt.Println("m (cells) =", len(m), "of", cap(m), ":", m)
	fmt.Println("b (cells) =", len(b), "of", cap(b), ":", b)
	fmt.Println("s (cells) =", len(s), "of", cap(s), ":", s)
	fmt.Println("")
}
