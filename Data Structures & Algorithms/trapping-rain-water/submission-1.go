func trap(height []int) int {
    if len(height) < 3 {
        return 0
    }
    l := 0
    r := len(height) - 1

    total := 0
    tempH := min(height[l], height[r])

    for l < r {
        switch {
            case height[l] > height[r]:
                r--
                if height[r] > tempH {
                    tempH = min(height[r], height[l])
                } else {
                    total += tempH - height[r]
                }
            default:
                l++
                if height[l] > tempH {
                    tempH = min(height[r], height[l])
                } else {
                    total += tempH - height[l]
                }
        }
    }
    return total
}

//My idea - find max, then expand out with 2 pointers - keep track of strictly decreasing points
//s.t. each point is greater than every point between it and the one on the right
//This recursive solution - we call trap on the excluded areas until nothing remains

//Is there a way we can do two pointers outside working in - 
//If len(height) < 3 {}

//go right to left, keeping track of the max we have seen - hold 2 values, 1 if max is highest we find, 1 if we can find something higher than left


//two pointers - calc water between, move lower pointer in, recalc until lower pointer is higher, then, we have a pool to add to final