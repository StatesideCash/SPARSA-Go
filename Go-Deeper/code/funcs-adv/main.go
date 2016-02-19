package main

import (
    "fmt"
    "os/exec"
)

func main() {
    // Normally you would have to write and call a function to do this.
    // However, since it is only used once and is relatively simple, we can
    // use a closure to write the function inline with the go subroutine.
    defer func() {
        fmt.Println("Entering closure")
        fmt.Println("Pinging Google...")
        command := exec.Command("/bin/ping", "-c 3", "8.8.8.8")
        err := command.Run()
        if err != nil {
            fmt.Println("Oh noes something broke or Google is down :O")
        } else {
            fmt.Println("Google is still up!")
        }
    }() // This is where arguments to the closure would go if we had any

    // We can also make a closure and assign it to a variable!
    doSomeMath := func(x, y int) int {
        return x * y
    }
    someNumber := doSomeMath(4, 5)
    fmt.Printf("%d * %d = %d\n", 4, 5, someNumber)

}
