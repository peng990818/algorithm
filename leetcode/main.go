package main

import "fmt"

func main() {
    grid := [][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}
    fmt.Println(islandPerimeter(grid))
}

func islandPerimeter(grid [][]int) int {
    var res int
    i, j := 0, 0
    for i < len(grid) && j < len(grid[0]) {
        if grid[i][j] == 1 {
            res += 4
            if i < len(grid)-1 && grid[i+1][j] == 1 {
                res -= 1
            }
            if j < len(grid[0])-1 && grid[i][j+1] == 1 {
                res -= 1
            }
            if i > 0 && grid[i-1][j] == 1 {
                res -= 1
            }
            if j > 0 && grid[i][j-1] == 1 {
                res -= 1
            }
        }
        j++
        if j == len(grid[0]) {
            i++
            j = 0
        }
    }
    return res
}

func commonChars(words []string) []string {
    mp := make(map[rune]int, 26)
    for _, word := range words {
        for _, b := range word {
            _, ok := mp[b]
            if ok {
                mp[b]++
            } else {
                mp[b] = 1
            }
        }
    }
    res := make([]string, 0)
    for k, v := range mp {
        if v >= len(words) {
            for v >= len(words) {
                res = append(res, string(k))
                v--
            }
        }
    }
    return res
}
