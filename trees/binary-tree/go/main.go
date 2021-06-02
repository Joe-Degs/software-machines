package main

import "fmt"

// a binary tree with linked lists.
type Ltree struct {
	data        int
	left, right *Ltree
}

func newLtree(data int) *Ltree {
	return &Ltree{
		data: data,
	}
}

// if value is bigger it goes in the right tree.
// if its lesser than parent it goes in left tree.
func (l *Ltree) Lappend(data int) {
	if l == nil {
		l = newLtree(data)
		return
	}

	if data == l.data {
		return
	}

	if data > l.data {
		if l.right == nil {
			l.right = newLtree(data)
		}
		l.right.Lappend(data)
	} else {
		if l.left == nil {
			l.left = newLtree(data)
		}
		l.left.Lappend(data)
	}
}

func tab(level int) {
	for i := 0; i < level; i++ {
		fmt.Printf("-")
	}
}

func printTree(root *Ltree, level int) {
	if root == nil {
		tab(level)
		fmt.Println("<n>")
		return
	} else {
		tab(level)
		fmt.Println(root.data)
	}

	printTree(root.right, level+1)
	printTree(root.left, level+1)
}

func main() {
	t := newLtree(5)
	t.Lappend(4)
	t.Lappend(6)
	t.Lappend(7)
	t.Lappend(11)
	t.Lappend(2)
	t.Lappend(1)
	t.Lappend(8)
	t.Lappend(10)
	t.Lappend(20)
	t.Lappend(0)

	printTree(t, 0)
}
