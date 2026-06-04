func search(nums []int, target int) int {
    l:=0
    r:=len(nums)-1

    for l<=r {
        t := r-l/2
        if nums[t] == target {
            return t
        } else if nums[t] > target {
            r = t-1
        } else {
            l = t+1
        }
    }
    return -1
}


//Here we will try the iterative solution