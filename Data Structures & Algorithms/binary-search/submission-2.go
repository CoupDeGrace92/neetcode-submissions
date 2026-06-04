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

//This is my version of the recursive approach - this is O(log(n)) in space complexity 
//because each recursion is creating a slice of the nums array to iterate over - with upper bound log(n)

//Because of that, this is not the best we can do for a binary search.  I chose this path not because of it being
//optimal space complexity but more so to illustrate that I can do recursion.

//In this case, there are two returning base cases, 0 len nums and hitting the target.