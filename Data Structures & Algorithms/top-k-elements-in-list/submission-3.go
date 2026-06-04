func topKFrequent(nums []int, k int) []int {
    numFreq := make(map[int]int)
    for _, num := range nums {
        numFreq[num]++
    }

    buckets := make([][]int, len(nums) + 1)
    for num, freq := range numFreq {
        buckets[freq] = append(buckets[freq], num)
    }

    out := []int{}
    for i := len(buckets) - 1; len(out) < k; i-- {
        out = append(out, buckets[i]...) //this only works because we guarentee a unique solution
        //Otherwise we could include a guard and slice the bucket before appending:
        /*
        remaining := k-len(out)
        bucket := buckets[i]
        if len(bucket) > remaining {
            bucket = bucket[:remaining]
        }
        out = append(out, bucket)
        */
    }
    return out
}
