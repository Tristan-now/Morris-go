package main

import "fmt"

func main() {
	nl := []*TreeNode{}
	for i := 1; i < 8; i++ {
		nl = append(nl, &TreeNode{Val: i})
	}

	nl[0].Left = nl[1]
	nl[0].Right = nl[2]
	nl[1].Left = nl[3]
	nl[1].Right = nl[4]
	nl[2].Left = nl[5]
	nl[2].Right = nl[6]

	//res := MorrisXianxu(nl[0])
	//fmt.Printf("%+v", res)

	//res := MorrisZhongxu(nl[0])
	//fmt.Printf("%+v", res)

	res := MorrisHouxu(nl[0])
	fmt.Printf("%+v", res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func morris(head *TreeNode) {
	cur := head
	var mostRight *TreeNode
	for cur != nil {
		println(cur.Val)
		mostRight = cur.Left
		if mostRight == nil {
			cur = cur.Right
		} else {
			for mostRight.Right != nil && mostRight.Right != cur {
				mostRight = mostRight.Right
			}
			if mostRight.Right == nil {
				mostRight.Right = cur
				cur = cur.Left
			} else {
				mostRight.Right = nil
				cur = cur.Right
			}
		}
	}
}
