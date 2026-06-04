func productExceptSelf(nums []int) []int {
    //This solution will be without using the division operator
    //The goal is to minimize repeated work - we can store the prefix/suffix products

    //For the suffix products - we could divide by i but that gets back to our more efficient solution 1
    //Instead, we are going to just look to build it up, maybe its already in our map.

    if len(nums) == 0 {
        return []int{}
    }

    prefix := make([]int, len(nums))
    suffix := make([]int, len(nums))
    for i, _ := range nums {
        if i == 0 {
            prefix[i] = 1
        } else {
            prefix[i] = nums[i-1] * prefix[i-1]
        }
    }
    
    for j := len(nums) - 1; j>=0; j-- {
        if j == len(nums) - 1 {
            suffix[j] = 1
        } else {
            suffix[j] = nums[j+1] * suffix[j+1]
        }
    }
    out := make([]int, len(nums))
    for i:=0; i<len(nums); i++ {
        out[i] = prefix[i] * suffix[i]
    }
    return out
}
