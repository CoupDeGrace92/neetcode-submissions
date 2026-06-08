func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 {
        return false
    }
    max := len(matrix) * len(matrix[0])
    //this is just modular arithmatic over mod len(matrix[0])
    mod := len(matrix[0])

    left := 0
    right := max - 1

    for left <= right {
        t := (left + right)/2
        row := t / mod
        idx := t % mod

        if matrix[row][idx] == target {
            return true
        } else if matrix[row][idx] < target {
            left = t + 1
        } else {
            right = t - 1
        }
    }

    return false
}


func mod (i, mod int) int {
    return i%mod
}
//That matrix is the same as a sorted array - 
//We have to have an algo that gives us a x,y coord from an absolute index

//We can do this with floor division