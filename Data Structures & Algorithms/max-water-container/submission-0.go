func maxArea(heights []int) int {
    l := 0
    r := len(heights) - 1
    max := 0

    if len(heights) == 0 {
        return 0
    }

    for l < r {
        c := heights[r] - heights[l]
        water := min(heights[r], heights[l]) * (r - l)
        if water > max {
            max = water
        }

        switch {
            case c>0:
                l++
            case c<0:
                r--
            case c==0:
                if heights[r-1] > heights[l+1] {
                    r--
                } else {
                    l++
                }
        }
    }

    return max

}

//So this problem - I think we start with two pointers on the l, r.  
//Formula is min(l, r) * distance between
//
//Because its the min, we can move the lower of the two pointers