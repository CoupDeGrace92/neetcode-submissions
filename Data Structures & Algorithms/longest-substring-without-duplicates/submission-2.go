func lengthOfLongestSubstring(s string) int {
    leftmost := make(map[rune]int) //char to index
    longest := 0
    lpoint := 0

    for i, char := range s {
        idx, exists := leftmost[char]
        if exists && idx >= lpoint {
            lpoint = idx + 1
        } 

        leftmost[char] = i

        if i - lpoint + 1> longest {
            longest = i-lpoint+1
        }
    }



    return longest
}
