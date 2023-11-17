package main

func MorrisHouxu(head *TreeNode) []int {
	cur := head
	var mostRight *TreeNode
	res := []int{}
	for cur != nil {
		mostRight = cur.Left
		if mostRight == nil {
			cur = cur.Right
		} else {
			for mostRight.Right != nil && mostRight.Right != cur {
				mostRight = mostRight.Right
			}
			if mostRight.Right == cur {
				mostRight.Right = nil
				_reverse(cur.Left, &res)
				cur = cur.Right
			} else {
				mostRight.Right = cur
				cur = cur.Left
			}
		}
	}
	_reverse(head, &res)
	return res
}

func _reverse(head *TreeNode, res *[]int) {
	temp := []int{}
	for head != nil {
		temp = append(temp, head.Val)
		head = head.Right
	}

	for i := 0; i < len(temp)/2; i++ {
		temp[i], temp[len(temp)-1-i] = temp[len(temp)-1-i], temp[i]
	}

	*res = append(*res, temp...)
}
