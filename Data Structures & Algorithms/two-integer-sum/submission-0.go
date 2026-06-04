func twoSum(nums []int, target int) []int {
    //Again hash map - this time map of remainder -> index

    m := make(map[int]int)
    for i, num := range nums {
        j, exists := m[num]
        if exists {
            return []int{j,i}
        }
        m[target-num] = i
    }
    return []int{}
}
