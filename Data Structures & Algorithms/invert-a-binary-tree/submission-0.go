/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
    if root == nil || (root.Right == nil && root.Left == nil) {
        return root
    }
    //Don't know if the root == nil guard clause is necessary
    tempLeft := root.Left
    tempRight := root.Right

    root.Left = invertTree(tempRight)
    root.Right = invertTree(tempLeft)

    return root
}
