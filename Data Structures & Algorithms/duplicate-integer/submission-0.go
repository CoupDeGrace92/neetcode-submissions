func hasDuplicate(nums []int) bool {
    m := make(map[int]struct{})
    for _, n := range nums {
        _, exists := m[n]
        if exists {
            return true
        }
        m[n] = struct{}{}
    }
    return false
}
