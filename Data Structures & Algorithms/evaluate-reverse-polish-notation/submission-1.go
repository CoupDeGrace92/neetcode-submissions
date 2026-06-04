
func evalRPN(tokens []string) int {
    stack := []int{}

    for i, c := range tokens {
        cint, err := strconv.Atoi(c)
        if err != nil && i==0 {
            fmt.Println("Invalid tokens: ", tokens)
            return 0
        } else if err != nil {
            //operation logic here
            if len(stack) < 2 {
                fmt.Println("Invalid tokens - not enough ints for operator: ", tokens)
                return 0
            }
            switch c {
            case "+":
                n := stack[len(stack)-2] + stack[len(stack)-1]
                stack = append(stack[:len(stack)-2], n)
            case "-":
                n := stack[len(stack)-2] - stack[len(stack)-1]
                stack = append(stack[:len(stack)-2], n)
            case "*":
                n := stack[len(stack)-2] * stack[len(stack)-1]
                stack = append(stack[:len(stack)-2], n)
            case "/":
                n := stack[len(stack)-2] / stack[len(stack)-1]
                stack = append(stack[:len(stack)-2], n)
            default:
                fmt.Println("Unrecognized operand: ", c)
                return 0
            }
            continue
        }
        stack = append(stack, cint)
    }
    if len(stack) == 1 {
        return stack[0]
    }
    fmt.Println("Stack has too many entries: ", stack)
    return 0
}