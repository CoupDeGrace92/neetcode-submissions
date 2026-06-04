type Solution struct{}

func (s *Solution) Encode(strs []string) string {
    //We are going to include lengths, seperated by a delim, with one final delim at the end
    if len(strs) == 0 {
        return ""
    }
    
    delimString := ""
    smashed := ""
    for _, str := range strs {
        l := strconv.Itoa(len(str))
        delimString = fmt.Sprintf("%v%v/", delimString, l)
        smashed = smashed + str
    }
    delimString = delimString[:len(delimString)-1] //Check to see if the answers are ASCII only or if we need to think about UTF-8
    delimString = fmt.Sprintf("%s#%s", delimString, smashed)
    return delimString
}

func (s *Solution) Decode(encoded string) []string {
    out := []string{}
    delimIndex := strings.Index(encoded, "#")
    if delimIndex == -1 {
        return out
    }
    delim := encoded[: delimIndex]
    enc := encoded[delimIndex+1:]
    
    lens := strings.Split(delim, "/")
    for _, l := range lens {
        lint, _ := strconv.Atoi(l)
        out = append(out, enc[:lint])
        enc = enc[lint:]
    }
    return out
}
