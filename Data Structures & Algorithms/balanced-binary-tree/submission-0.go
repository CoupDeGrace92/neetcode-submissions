/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    return dfs(root).balance
}

type Result struct {
    balance bool
    height int
}


func dfs(root *TreeNode) Result {
    if root == nil {
        return Result{true, 0}
    }

    left := dfs(root.Left)
    right := dfs(root.Right)

    balanced := left.balance && right.balance && abs(left.height-right.height) <=1
    r := Result {
        balance: balanced,
        height: 1+max(left.height, right.height),
    }
    return r
}

func abs (val int) int {
    if val < 0 {
        val = -val
    }
    return val
}

func max(a, b int) int {
    if a>b {
        return a
    } 
    return b
}
