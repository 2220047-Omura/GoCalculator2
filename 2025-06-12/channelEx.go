package main

import (
	"fmt"
	"math/rand"
)

// A Tree is a binary tree with integer values.
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// New returns a new, random binary tree holding the values k, 2k, ..., 10k.
func New(k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(10) {
		t = insert(t, (1+v)*k)
	}
	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
	} else {
		t.Right = insert(t.Right, v)
	}
	return t
}

/*
func (t *Tree) String() string {
	if t == nil {
		return "()"
	}
	s := ""
	if t.Left != nil {
		s += t.Left.String() + " "
	}
	s += fmt.Sprint(t.Value)
	if t.Right != nil {
		s += " " + t.Right.String()
	}
	return "(" + s + ")"
}
*/

// Walk 関数では、Treeを探索しながら、Treeに含まれる全ての
// 値を、値が小さい物から順にチャネル chに送信する
func Walk(t *Tree, ch chan int) {
	WalkSub(t, ch)
	close(ch)

}

// 再起呼び出し用の関数
func WalkSub(t *Tree, ch chan int) {
	if t.Left != nil {
		WalkSub(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		WalkSub(t.Right, ch)
	}

}

// Same 関数は、2つのTree t1 とt2が全く同じ値を含んでいるものかどうか
// 確認する
func Same(t1, t2 *Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)
	// goroutine によりWalk関数を並行実行
	go Walk(t1, c1)
	go Walk(t2, c2)

	// それぞれのチャネルが受信した値を取り出し
	// 一致するか評価
	for i := range c1 {
		j := <-c2
		if i != j {
			return false
		}
	}
	return true

}

func main() {
	fmt.Println(Same(New(1), New(1)))
	fmt.Println(Same(New(1), New(2)))
}
