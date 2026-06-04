func maxProfit(prices []int) int {
    if len(prices) < 2 {
        return 0
    }
    buy := 100 //set by constraints, otherwise we can use max floating point: math.MaxInt
    sell := 0
    temp := 0

    for i := 0; i < len(prices); i++ {
        switch {
            case prices[i] < buy:
                temp = max(sell-buy, temp)
                buy = prices[i]
                sell = 0
            case prices[i] > sell:
                sell = prices[i]
        }
    }

    if sell - buy > temp {
        temp = sell - buy
    }

    return temp

}
