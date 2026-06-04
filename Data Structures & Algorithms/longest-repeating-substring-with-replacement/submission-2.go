func characterReplacement(s string, k int) int {
    if len(s) <= k+1 {
        return len(s)
    }
    
    lpoint := 0 
    maxOccurance := 0
    total := 0
    tracking := make(map[rune]int)
    max := 0

    for _, char := range s {
        total ++
        tracking[char]++
        if tracking[char] > maxOccurance {
            maxOccurance = tracking[char]
        }
        

        for total - maxOccurance > k {
            tracking[rune(s[lpoint])]--
            total--
            lpoint++
        }

        if total > max {
            max = total
        }
    }
    return max
}


//hash table with number of occurences - while second highest >= k, the substring is fine.  When we violate, move lpointer right
//decriment element in hash table, then repeat until hash table is valid (remove the offending element)
//then we move the right pointer and continue

//the above does not quite work - it puts bounds on a singular replacement at k, not all replacements

//Instead, keep a total, a max letter struct