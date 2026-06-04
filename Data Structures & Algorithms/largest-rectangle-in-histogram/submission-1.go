func largestRectangleArea(heights []int) int {
    max := 0
    stack := []int{}
    
    for i := 0; i < len(heights); i++ {
        for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
            h := heights[stack[len(stack)-1]]
            stack = stack[:len(stack)-1]
            
            w := i
            if len(stack) > 0 {
                w = i - stack[len(stack)-1] - 1
            }
            if h*w > max {
                max = h * w
            }
        }
        stack = append(stack, i)
    }

    // Handle remaining stack
    for len(stack) > 0 {
        h := heights[stack[len(stack)-1]]
        stack = stack[:len(stack)-1]
        
        w := len(heights)
        if len(stack) > 0 {
            w = len(heights) - stack[len(stack)-1] - 1
        }
        if h*w > max {
            max = h * w
        }
    }
    return max
}

//Biggest trip ups - the width - I kept getting that off by one error, and still do not know why the additional 1 is really needed
//I think that is just a problem with my mental model, working while hammering is going on above my head