/*
 * @lc app=leetcode.cn id=117 lang=golang
 *
 * [117] 填充每个节点的下一个右侧节点指针 II
 */
package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	queue := []*Node{}
	queue = append(queue, root)
	for len(queue) > 0 {
		n := len(queue)
		var preNode *Node
		for i := 0; i < n; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if i == 0 {
				cur.Next = nil
			} else {
				cur.Next = preNode
			}
			preNode = cur
		}
	}
	return root
}

// @lc code=end
func main() {
	root := &Node{
		Val: 1,
		Left: &Node{
			Val: 2,
			Left: &Node{
				Val: 4,
			},
			Right: &Node{
				Val: 5,
			},
		},
		Right: &Node{
			Val: 3,
			Right: &Node{
				Val: 7,
			},
		},
	}
	fmt.Println(connect(root))
	fmt.Println("====")
}
