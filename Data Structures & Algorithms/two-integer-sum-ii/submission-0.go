import(
    "slices"
)

func twoSum(numbers []int, target int) []int {
    slices.Sort(numbers)
    left := 0
    right := len(numbers) - 1
    for right > left {
        sum := numbers[right] + numbers[left]
        if sum == target {
            return []int{left + 1, right + 1}
        } else if sum > target {
            right --
        } else {
            left ++
        }
    }
    fmt.Println("We could not find a solution you dunce")
    return []int{}
}

//this particular solution requires us to use O(1) additional space
//IN this case, we can not just throw a hash map at it
//instead we will use a two pointers approach

//DUMB FUCKING ASSIGNMENT IS ASSUMING A 1-indexed array