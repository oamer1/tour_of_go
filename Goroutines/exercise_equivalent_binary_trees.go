package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

/* // A Tree is a binary tree with integer values.
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
} */

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)

	var walkerRecurse func(t *tree.Tree)
	walkerRecurse = func(t *tree.Tree) {
		if t == nil {
			return
		}
		walkerRecurse(t.Left)
		ch <- t.Value
		walkerRecurse(t.Right)

	}
	walkerRecurse(t)

}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		n1, ok1 := <-ch1
		n2, ok2 := <-ch2
		if n1 != n2 || ok1 != ok2 {
			return false
		}
		if !ok1 {
			break
		}
	}
	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
