package main

import (
    "fmt"
    "sync"
    "strconv"
)

type Person interface {
    Eat()
    Drink()
}

type Basic struct {
}

func (b *Basic) Eat() {
    fmt.Println("basic eat")
}

func (b *Basic) Drink() {
    fmt.Println("basic drink")
}

type Student struct {
    Basic
}

func (s *Student) Eat() {
    fmt.Println("student eat")
}

type Teacher struct {
    Basic
}

func (t *Teacher) Eat() {
    fmt.Println("teacher eat")
}

func (t *Teacher) Drink() {
    fmt.Println("teacher drink")
}

func MyFmt(f func()) {
    f()
    fmt.Println("aaa")
}

func main() {
    wg := sync.WaitGroup{}
    ch1 := make(chan int)
    ch2 := make(chan string)
    wg.Add(1)
    go func() {
        for {
            num, ok := <-ch1
            if ok {
                fmt.Println(num)
                ch2 <- "string" + strconv.Itoa(num)
            } else {
                close(ch2)
                wg.Done()
                break
            }
        }

    }()

    wg.Add(1)
    go func() {
        for {
            num, ok := <-ch2
            if ok {
                fmt.Println(num)
            } else {
                wg.Done()
                break
            }
        }
    }()
    for i := 0; i < 10; i++ {
        ch1 <- i
    }
    close(ch1)
    wg.Wait()
    fmt.Println("Done")
}

// var i int
// ch := make(chan int)
// go func() {
//     for {
//         num := <-ch
//         fmt.Println("go1")
//         fmt.Println(num)
//         num++
//         ch <- num
//     }
// }()
//
// go func() {
//     for {
//         ch <- i
//         i = <-ch
//         fmt.Println("go2")
//         fmt.Println(i)
//         i++
//     }
// }()
// time.Sleep(1 * time.Millisecond)

// var i int
// ch := make(chan int, 1) // 创建一个带缓冲的通道
//
// go func() {
//     for {
//         ch <- i                            // 将 i 发送到通道
//         i = <-ch                           // 从通道接收值并赋值给 i
//         fmt.Println("go1")                 // 打印 go1
//         fmt.Println(i)                     // 打印接收到的值
//         i++                                // 递增 i
//         time.Sleep(100 * time.Millisecond) // 为了演示，让协程稍微休眠一下
//     }
// }()
//
// go func() {
//     for {
//         num := <-ch                        // 从通道接收值
//         fmt.Println("go2")                 // 打印 go2
//         fmt.Println(num)                   // 打印接收到的值
//         num++                              // 递增接收到的值
//         ch <- num                          // 将递增后的值发送回通道
//         time.Sleep(100 * time.Millisecond) // 为了演示，让协程稍微休眠一下
//     }
// }()
//
// time.Sleep(1 * time.Second) // 等待一段时间，确保协程执行
