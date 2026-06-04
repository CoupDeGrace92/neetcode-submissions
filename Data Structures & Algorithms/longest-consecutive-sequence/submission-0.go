import (
    "slices"
)

func longestConsecutive(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    slices.Sort(nums)
    longest := 1
    current := 1
    for i, num := range nums {
        if i == 0 {
            continue
        }
        if num - 1 == nums[i-1] {
            current ++
            if current > longest {
                longest = current
            }
        } else if num == nums[i-1] {
            continue
        } else {
            current = 1
        }
    }
    return longest
}
