func search(nums []int, target int) int {
    if len(nums) == 0 {
        return -1
    }
    
    i := len(nums)/2
    if nums[i] == target {
        return i
    } else if nums[i] > target {
        return search(nums[:i], target)
    } else {
        result := search(nums[i+1:], target)
        if result == -1 {
            return -1
        }
        return result + i + 1
    }
}

