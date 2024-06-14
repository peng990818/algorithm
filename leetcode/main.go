package main

import "fmt"

func main() {
    grid := [][]int{{2, 2}, {3, 3}}
    fmt.Println(maxCount(3, 3, grid))
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

func maxCount(m int, n int, ops [][]int) int {
    res := make([][]int, m)
    for i := range res {
        res[i] = make([]int, n)
    }
    max := 0
    for _, op := range ops {
        for i := 0; i < op[0]; i++ {
            for j := 0; j < op[1]; j++ {
                res[i][j]++
                if res[i][j] > max {
                    max = res[i][j]
                }
            }
        }
    }
    cnt := 0
    for _, v := range res {
        for _, val := range v {
            if val == max {
                cnt++
            }
        }
    }
    return cnt
}
