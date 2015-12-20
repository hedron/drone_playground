package main

import "fmt"

func main() {
    fmt.Println(fib(50))
}


func fib(n int) int {
    current := 0
    previous := 1
    tmp := 0
    for i := 0; i < n; i++ {
        tmp = current
        current += previous
        previous = tmp
    }
    return current
}
