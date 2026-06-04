func dailyTemperatures(temperatures []int) []int {
    out := make([]int, len(temperatures))
    stack := []tempIdx{}

    for i, temp := range temperatures {
        for len(stack) > 0 && temp > stack[len(stack)-1].temp {
            top := stack[len(stack)-1]
            out[top.idx] = i-top.idx
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, tempIdx{
            temp: temp,
            idx: i,
        })
    }
    return out
}

type tempIdx struct {
    temp int
    idx int
}

//Keep a stack of strictly decreasing temp values, each of thsoe values should probably also have the index attached