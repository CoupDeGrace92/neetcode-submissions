func isValidSudoku(board [][]byte) bool {
    //Brute force is check for no duplicates in the 9x9
    //However, we want to try not to check pairs multiple times.

    //Hash rows/columns

    //Thought 1: add pair to hashmap once checked - then we check hashmap before searching the board

    // 1. Check for any two points, if i = i, the values are not the same,
    // 2. Check for any two points, if j = j, the values are not the same,
    // 3. Check the squares - we need to check the squares but have to define those squares elegantly
    //          All combinations of i, j where for a i1, j1: 
    //          a. We can manually define buckets
    //          b. Use floor division and have square indicies
    //              i/3, j/3 would index the square that it is in


    //With all of this, instead of having a single hash maps with all values, we can create a hash map for rows, columns, and squares
    //We only need to check the portion of the square we have not already checked

    rowMaps := make(map[int]map[byte]struct{})
    colMaps := make(map[int]map[byte]struct{})
    sqMap := make(map[point]map[byte]struct{})

    //Check maps as we go, appending what we need
    for i := 0; i < 9; i++ {
        for j := 0; j<9; j++{
            val := board[i][j]

            row, ok := rowMaps[i]
            if !ok {
                rowMaps[i] = make(map[byte]struct{})
                rowMaps[i][val] = struct{}{}
            } else {
                _, exists := row[val]
                if exists && val != '.' {
                    return false
                }
                row[val] = struct{}{}
            }

            col, ok := colMaps[j]
            if !ok {
                colMaps[j] = make(map[byte]struct{})
                colMaps[j][val] = struct{}{}
            } else {
                _, exists := col[val]
                if exists && val != '.' {
                    return false
                }
                col[val] = struct{}{}
            }

            p := point {
                x: i/3,
                y: j/3,
            }
            square, ok := sqMap[p]
            if !ok {
                sqMap[p] = make(map[byte]struct{})
                sqMap[p][val] = struct{}{}
            } else {
                _, exists := square[val]
                if exists && val != '.' {
                    return false
                }
                square[val] = struct{}{}
            }
        }
    }
    return true
}

type point struct{
    x int
    y int
}
