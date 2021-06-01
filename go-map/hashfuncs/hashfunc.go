package main

import "fmt"

/*
* Hash functions map keys to smaller intergers interget(buckets);
* It does this in a random like manner to ensure that keys appear to be
* evenly distributed even if there are irregularities in the input data.
*
* Mapping intergers is cool breeze because int can almost always be represented as
* as binary digits, but strings of characters or other complex data types
* contain separate strings binary digits making up the whole. And that's where
* the problem is.
*
* I will try few hash funcs in this package to see how the whole thing
* works.
* */

const (
	UINT32_MAX uint32 = 0xffffffff
)

// this function i saw in Eleanor's McHughs talk explaining how to implement
// maps and golang.
func eleanorHashFunc(key string) int {
	var h uint32
	for i := len(key) - 1; h < UINT32_MAX && i > 0; i-- {
		h = h << 8          // h * (pow(2, 8)); shift h left by 8 bits
		h += uint32(key[i]) // add ki to h
	}
	return int(h)
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

// Java's implementation of hashfunction for strings.
// s[0]*pow(31, n-2) + s[1]*pow(31, n-2) + ... + s[n-1]
func javaHashFunc(key string) int {
	var h uint32
	n := len(key)
	for i := 0; h < UINT32_MAX && i < n; i++ {
		h += uint32(key[i]) * uint32(pow(0x1f, n-i+1))
	}
	return int(h)
}

// 5-bit left circular shift of h, then xor in key[i].
func crcVariant(key string) int {
	var h uint32
	n := len(key)
	for i := 0; h < UINT32_MAX && i < n; i++ {
		var hb uint32
		hb = h & 0xf8000000 // extract first 5 high-order bits from h
		h = h << 5          // shift left by 5 bits
		h ^= (hb >> 27)     // move highorder 5 bits to end and xor into h
		h ^= uint32(key[i]) // xor h and key[i]
	}
	return int(h)
}

func getBucket(key string, size int, hash func(string) int) int {
	return hash(key) % size
}

func main() {
	hashfuncs := map[string]func(string) int{
		"eleanor's hash implementation":     eleanorHashFunc,
		"java's string maps implementation": javaHashFunc,
		"CRC hash implementation":           crcVariant,
	}
	keys := []string{"rose", "kelvin", "joe", "kpakpa", "jude", "mesopotamia", "lil dicky",
		"the quick brown fox jumps over the lazy foxoxooxoxoxs",
		"i am eating in the cut like a champ",
		"hashes for days mi amigo",
	}
	size := 100
	for name, hash := range hashfuncs {
		fmt.Printf("%s: Key -> Hash -> Bucket\n", name)
		for _, key := range keys {
			fmt.Printf("\t %s -> %x -> %d\n", key, hash(key), getBucket(key, size, hash))
		}
	}
}

/*
* Lessons learnt so far
* 1. the smaller the bucket size, the higher the chance of getting collisions.
* 2. there is an awful lot of mathematics involved in this thing. learn some
*    maths to save you life.
* 3.
* */
