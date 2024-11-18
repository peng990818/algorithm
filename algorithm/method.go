package algorithm

import (
    "math/rand/v2"
)

// 生成随机数组（1～10000个）
func GenerateRandomArray() []int {
    // 生成 1 到 10000 之间的随机数
    randomNumber := rand.IntN(10000) + 1
    res := make([]int, 0, randomNumber)
    for i := 0; i < randomNumber; i++ {
        res = append(res, rand.IntN(10000)+1)
    }
    return res
}
