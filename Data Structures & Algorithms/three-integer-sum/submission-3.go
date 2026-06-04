import (
    "slices"
)

func threeSum(nums []int) [][]int {
    out := [][]int{}
    slices.Sort(nums)
    for i, num := range nums {
        if i != 0 && nums[i] == nums[i-1] {
            continue
        }
        if num > 0 {
            break
        }
        left := i + 1
        right := len(nums) - 1
        for left < right {
            l := nums[left]
            r := nums[right]
            if l + r ==  -num {
                out = append(out, []int{num, l, r})
                left++
                right--
                for left < right && nums[left] == nums[left-1] {
                    left ++
                }
                for left < right && nums[right] == nums[right + 1]{
                    right --
                }
            } else if l + r > -num {
                right --
            } else {
                left ++
            }
        } 
    }
    return out
}


//This looks like something we want to set up a hashmap of our numbers
//Then we can loop over the list - we can do better

//HOW ABOUT:
//Two pointer approach without a hashmap would reduce the space complexity
//and the time complexity will still be n^2, we do the approach for each index i


//Checking duplicates: make sure our starting value < 0

