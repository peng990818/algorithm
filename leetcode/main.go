package main

import "fmt"

func main() {
    words := []string{"bella", "label", "roller"}
    fmt.Println(commonChars(words))
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
