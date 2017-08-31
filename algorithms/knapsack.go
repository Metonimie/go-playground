package main

import (
    "fmt"
)

var weight []int
var n = 100
var k = 4090

func Knapsack(n int, k int) bool {
    if n == 0 {
        if weight[n] == k {
            return true
        } else {
            return false
        }
    }
    if Knapsack(n - 1, k) {
        return true
    } else {
        if weight[n] == k {
            return true
        } else if k - weight[n] >= 0 {
            return Knapsack(n - 1, k - weight[n])
        } else {
            return false
        }
    }
}

var isDone = make([][]bool, n)
var dasResult = make([][]bool, n)

func KnapsackM(n int, k int) bool  {
    if isDone[n][k] {
        return dasResult[n][k]
    } else {
        isDone[n][k] = true
        dasResult[n][k] = KnapsackM(n, k)
    }

    return false
}

func main() {
    for i := 0; i < n; i++ {
        isDone[i] = make([]bool, k)
    }

    for i := 0; i < n; i++ {
        dasResult[i] = make([]bool, k)
    }



    fmt.Println(KnapsackM(n, k))
}
