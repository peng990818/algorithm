package main

import (
    "fmt"
)

func main() {
    // grid := [][]int{{2, 2}, {3, 3}}
    // fmt.Println(maxCount(3, 3, grid))
    // defer func() {
    //     fmt.Println("test1")
    // }()
    // defer func() {
    //     test()
    // }()
    // panic("test")
    // go func() {
    //     var b chan string
    //     <-b
    // }()
    // go func() {
    //     for {
    //         fmt.Println("aaa")
    //         time.Sleep(time.Second)
    //     }
    // }()
    // select {}

    // ch := make(chan int, 1)
    //
    // go func() {
    //     // 向无缓冲 channel 发送数据
    //     ch <- 42
    //     fmt.Println("Sent value")
    // }()
    //
    // time.Sleep(time.Second) // 模拟一些延迟
    //
    // // 从无缓冲 channel 接收数据
    // value := <-ch
    // fmt.Println("Received value:", value)
    // select {}

    tmp := [600]int{}
    tmp1 := tmp[:600:600]
    for i := 0; i < len(tmp1); i++ {
        fmt.Println(tmp1[i])
    }
}

func test() {
    fmt.Println("test2")
    recover()
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
