package main

func MorrisZhongxu(head *TreeNode) []int {
	cur := head
	var mostRight *TreeNode
	res := []int{}
	for cur != nil {
		mostRight = cur.Left
		if mostRight == nil {
			res = append(res, cur.Val)
			cur = cur.Right
		} else {
			for mostRight.Right != nil && mostRight.Right != cur {
				mostRight = mostRight.Right
			}
			if mostRight.Right == cur {
				res = append(res, cur.Val)
				mostRight.Right = nil
				cur = cur.Right
			} else {
				mostRight.Right = cur
				cur = cur.Left
			}
		}
	}
	return res
}

//func MorrisZhongxu(root *TreeNode) []int {
//	result := make([]int, 0)
//	current := root
//	var predecessor *TreeNode
//
//	for current != nil {
//		if current.Left == nil {
//			result = append(result, current.Val)
//			current = current.Right
//		} else {
//			// Find the inorder predecessor
//			predecessor = current.Left
//			for predecessor.Right != nil && predecessor.Right != current {
//				predecessor = predecessor.Right
//			}
//
//			// Make current as the right child of its inorder predecessor
//			if predecessor.Right == nil {
//				predecessor.Right = current
//				current = current.Left
//			} else {
//				// Revert the changes made in the 'if' part to restore the original tree
//				predecessor.Right = nil
//				result = append(result, current.Val)
//				current = current.Right
//			}
//		}
//	}
//
//	return result
//}
