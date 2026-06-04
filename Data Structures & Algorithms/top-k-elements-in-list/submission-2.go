import (
    "slices"
    "cmp"
)

func topKFrequent(nums []int, k int) []int {
    //First instinct - create a freq map
    //Create a inverse of the map then take the largest k

    //Now we want to keep track of the buckets
    //To do this we will create a pointer to the current max bucket

    numFreq := make(map[int]int)
    buckets := make(map[int][]int)   //(freq, list of numbers with freq)
    freqList := []int{}

    for _, num := range nums {
        numFreq[num] ++
        numList, ok := buckets[numFreq[num]]
        if ok {
            buckets[numFreq[num]] = append(numList, num)
        } else {
            buckets[numFreq[num]] = []int{num}
            freqList = append(freqList, numFreq[num])
        }
        if numFreq[num] > 1 {
            numIndex := slices.Index(buckets[numFreq[num] - 1], num)
            buckets[numFreq[num]-1] = slices.Delete(buckets[numFreq[num]-1], numIndex, numIndex + 1)
            if len(buckets[numFreq[num] - 1]) == 0 {
                delete(buckets, numFreq[num] - 1)
                lastIndex := slices.Index(freqList, numFreq[num]-1)
                freqList = slices.Delete(freqList, lastIndex, lastIndex + 1)
            }
        }
    }
    slices.SortFunc(freqList, func(a, b int) int {
        return cmp.Compare(b, a)
    })
    out := []int{}
    bucketIndex := 0
    freqIndex := 0
    for len(out) < k {
        freq := freqList[freqIndex]
        out = append(out, buckets[freq][bucketIndex])
        bucketIndex ++
        if bucketIndex >= len(buckets[freq]) {
            bucketIndex = 0
            freqIndex ++
        }
    }
    return out
}
