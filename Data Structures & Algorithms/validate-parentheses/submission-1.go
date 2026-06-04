func isValid(s string) bool {
    stack := []rune{}
    
    for _, char := range s {
        switch char{
        case '(':
            stack = append(stack, ')')
        case '{':
            stack = append(stack, '}')
        case '[':
            stack = append(stack, ']')
        case ')', '}', ']':
            if len(stack) == 0 {
                return false
            }
            if char != stack[len(stack)-1] {
                return false
            }
            stack = stack[:len(stack)-1] //Quick note - if we had been using pointers, those pointers would still exist in the underlying slice
        }
    }

    if len(stack) != 0 {
        return false
    }
    return true
}

//Reason to use the pointer struct representation of a stack, or just an array?
