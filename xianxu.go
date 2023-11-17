package main

func MorrisXianxu(head *TreeNode) []int {
	res := []int{}
	cur := head
	var mostRight *TreeNode
	for cur != nil {
		mostRight = cur.Left
		if mostRight == nil {
			res = append(res, cur.Val)
			cur = cur.Right
		} else {
			for mostRight.Right != nil && mostRight.Right != cur {
				mostRight = mostRight.Right
			}
			if mostRight.Right == nil {
				res = append(res, cur.Val)
				mostRight.Right = cur
				cur = cur.Left
			} else {
				mostRight.Right = cur
				cur = cur.Right
			}
		}
	}
	return res
}

//func morrisPreorderTraversal(root *TreeNode) []int {
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
//				result = append(result, current.Val)
//				predecessor.Right = current
//				current = current.Left
//			} else {
//				// Revert the changes made in the 'if' part to restore the original tree
//				predecessor.Right = nil
//				current = current.Right
//			}
//		}
//	}
//
//	return result
//}
