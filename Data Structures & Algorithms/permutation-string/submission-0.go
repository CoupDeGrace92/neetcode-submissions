import (
    "maps"
)

func checkInclusion(s1 string, s2 string) bool {
    
    if len(s1) > len(s2) {
        return false
    }
    
    s1map := make(map[rune]int)
    s2map := make(map[rune]int)
    
    for _, char := range s1 {
        s1map[char]++
    }
    
    lpoint := 0
    rpoint := 0
    l := len(s1)
    for i, char := range s2 {
        rpoint = i
        s2map[char]++
        if rpoint - lpoint + 1 == l {
            if maps.Equal(s1map, s2map) {
                return true
            } else {
                s2map[rune(s2[lpoint])]--
                if s2map[rune(s2[lpoint])] == 0 {
                    delete(s2map, rune(s2[lpoint]))
                }
                lpoint++
            }
        }
    }
    return false
}

// for space complexity of O(1), we can sort both strings and see if s2 contains s1
// unfortunately, for go, the built in mechanism for sorting strings would be to sort a slice of runes, then cast it back to strings

//This would not work anyway - duplicate characters problem