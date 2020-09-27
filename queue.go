package main

import "fmt"

type Node struct {
	Value      int
	ID         int
	ReportName string
}

type Queue []*Node

func (q *Queue) String() string {
	var id []int
	for _, v := range *q {
		id = append(id, v.ID)
	}
	return fmt.Sprint(id)

}

func (q *Queue) Push(n *Node) {
	*q = append(*q, n)
}

func (q *Queue) Pop() (n *Node) {
	n = (*q)[0]
	*q = (*q)[1:]
	return
}

func (q *Queue) Pop2() (n *Node) {
	if len(*q) > 0 {
		n = (*q)[0]
		if len(*q) >= 1 {
			*q = (*q)[1:]
		}
	} else {
		// make slice empty
		*q = (*q)[0:0]
		return
	}
	return
}

func (q *Queue) Len() int {
	return len(*q)
}

func main() {
	node := Node{
		Value:      1,
		ID:         1,
		ReportName: "a",
	}
	node2 := Node{
		Value:      2,
		ID:         2,
		ReportName: "b",
	}
	fmt.Println(node)
	var q Queue
	q.Push(&node)
	q.Push(&node2)
	q.Pop2()
	q.Pop2()
	q.Pop2()
	q.Pop2()
	q.Pop2()
	q.Pop2()
	fmt.Println(q)
	fmt.Println(q.String())
}
