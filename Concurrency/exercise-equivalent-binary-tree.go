package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	var _walk func(t *tree.Tree, ch chan int)
	_walk = func (t *tree.Tree, ch chan int) {
		if t == nil {
			return
		}
		_walk(t.Left, ch)
		ch <- t.Value
		_walk(t.Right, ch)
	}
	_walk(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	a1 := collect(ch1)
	a2 := collect(ch2)
	if len(a1) != len(a2) {
		return false
	} else {
		for i := 0; i < len(a1); i ++ {
			if a1[i] != a2[i] {
				return false
			}
		}
		return true
	}
}

func collect(ch chan int) []int  {
	a := make([]int, 10)
	for i := range ch {
		a = append(a, i)
	}
	return a
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := range ch {
		fmt.Println(i)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}