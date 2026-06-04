func isAnagram(s string, t string) bool {
    //We can sort the strings
    //Or we can put them in a hash map, then compare.
    //One thing to note - when we iterate over the characters in a string, we are iterating over runes:

    m := make(map[rune]int)
    for _, char := range s {
        m[char] ++
    }

    //we can now create a second map, OR decriment the first map, if value = 0, delete the entry from the map
    //Finally check the length of the map


    for _, char := range t {
        val, exists := m[char]
        if !exists {
            return false
        }
        if val == 1 {
            delete(m, char)
        } else {
            m[char] --
        }
    }
    if len(m) > 0 {
        return false
    }

    return true
}
