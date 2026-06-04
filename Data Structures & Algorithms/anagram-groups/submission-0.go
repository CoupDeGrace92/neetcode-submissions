import (
    "slices"
)

func groupAnagrams(strs []string) [][]string {
    //Here we can sort the string.  Store the sorted value in a hash map/check hash map.
    //Compose the solution from the values of the hash map

    //this does require conversion to a slice of runes then back

    anagrams := make(map[string][]string)
    for _, str := range strs {
        runes := []rune(str)
        slices.Sort(runes)
        sorted := string(runes)

        list, ok := anagrams[sorted]
        if !ok {
            anagrams[sorted] = []string{str}
        } else {
            anagrams[sorted] = append(list, str)
        }
    }
    out := [][]string{}
    for _, strList := range anagrams {
        out = append(out, strList)
    }
    return out
}
