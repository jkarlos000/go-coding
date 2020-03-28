package main

import (
	"encoding/json"
	"fmt"
)

type Node struct {
	key byte
	left *Node
	right *Node
}

type Tree struct {
	root *Node
}

// Árbol
func (t *Tree) insert(data byte) {
	if t.root == nil {
		t.root = &Node{key: data}
	} else {
		t.root.insert(data)
	}
}

// Nodo
func (n *Node) insert(data byte) {
	if data <= n.key {
		if n.left == nil {
			n.left = &Node{key: data}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &Node{key: data}
		} else {
			n.right.insert(data)
		}
	}
}

func main() {
	var t Tree

	t.insert('F')
	t.insert('B')
	t.insert('A')
	t.insert('D')
	t.insert('C')
	t.insert('E')
	t.insert('G')
	t.insert('I')
	t.insert('H')

	out, err := json.Marshal(t)
	if err != nil {
		panic (err)
	}
	fmt.Printf("%v\n", out) //Realizar método ToString()

}
