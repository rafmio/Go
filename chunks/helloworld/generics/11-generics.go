// https://habr.com/ru/companies/otus/articles/782414/
package main

import "fmt"

// TreeNode представляет узел в бинарном дереве поиска
type TreeNode[T any] struct {
	Value T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

// Insert добавляет значение в бинарное дерево поиска
func (n *TreeNode[T]) Insert(value T, compare func(a, b T) int) {
	if compare(value, n.Value) < 0 {
		if n.Left == nil {
			n.Left = &TreeNode[T]{Value: value}
		} else {
			n.Left.Insert(value, compare)
		}
	} else {
		if n.Right == nil {
			n.Right = &TreeNode[T]{Value: value}
		} else {
			n.Right.Insert(value, compare)
		}
	}
}

// InOrder обходит дерево в порядке возрастания
func (n *TreeNode[T]) InOrder(visit func(T)) {
	if n == nil {
		return
	}
	n.Left.InOrder(visit)
	visit(n.Value)
	n.Right.InOrder(visit)
}

func main() {
	root := &TreeNode[int]{Value: 5}
	root.Insert(3, func(a, b int) int { return a - b })
	root.Insert(7, func(a, b int) int { return a - b })
	root.Insert(1, func(a, b int) int { return a - b })
	root.Insert(9, func(a, b int) int { return a - b })

	root.InOrder(func(value int) {
		fmt.Printf("%d\n", value)
	})
}
