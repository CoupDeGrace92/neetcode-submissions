func isPalindrome(s string) bool {
    //We will do a two pointers approach instead of my naive two pass approach 
    left := 0
    right := len(s) - 1
    for left < right {
        if !isAlphanumeric(s[left]) {
            left ++
            continue
        }
        if !isAlphanumeric(s[right]) {
            right --
            continue
        }

        if charToLower(s[left]) != charToLower(s[right]) {
            return false
        }
        left ++
        right --
    }
    return true
}

func isAlphanumeric(b byte) bool {
    return (b>= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9')
}

func charToLower(b byte) byte {
    if (b >= 'A' && b <= 'Z') {
        return b + 32
    }
    return b
}