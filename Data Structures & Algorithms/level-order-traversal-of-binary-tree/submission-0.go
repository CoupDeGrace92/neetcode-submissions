/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }

    var queue []*TreeNode
    var result [][]int
    queue = append(queue, root)
    for len(queue) != 0 {
        var level []int
        l := len(queue)
        for i:=0; i<l; i++ {
            level = append(level, queue[i].Val)
            if queue[i].Left != nil {
                queue = append(queue, queue[i].Left)
            }
            if queue[i].Right != nil {
                queue = append(queue, queue[i].Right)
            }
        }
        result = append(result, level)
        queue = queue[l:]
    }

    return result
}

