package main

import "fmt"

func safeDiv(num1, num2 int) int {
    defer func() {
        //Recover allows you to recover from a fatal error
        //in this example it catches a divide by 0
        //Printing the recover statement gives an error message and continues
        //executing the program
        fmt.Println(recover())
    }()

    solution := num1 / num2
    return solution
}

func demPanic() {
    defer func() {
        fmt.Println(recover())
    }()

    panic("PANIC")
}

func main() {
    fmt.Println(safeDiv(4,0))
    fmt.Println(safeDiv(4,2))
    demPanic()
}
