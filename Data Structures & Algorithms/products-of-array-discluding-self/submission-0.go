func productExceptSelf(nums []int) []int {
    //We are going to go for the division method - compute the whole product, then divide by the indexed value
    //I am going to have to think through the time complexity of the hinted at solution more

    //The one issue we will run into with the division approach - zero.  What happens when we run into zero in the array

    fullProd := 1
    for _, i := range nums {
        fullProd = fullProd * i
    }

    out := []int{}
    for idxi, i := range nums {
        if i != 0 {
            out = append(out, fullProd/i)
        } else {
            partialProd := 1
            for idxj, j := range nums {
                if idxi != idxj {
                    partialProd = partialProd * j
                }
            }
            out = append(out, partialProd)
        }
    }

    return out
}
