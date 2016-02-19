package main

import (
    "fmt"
    "os/signal"
    "os"
    "time"
)

func main() {
    sig := make(chan os.Signal)
    signal.Notify(sig, os.Interrupt)

    go func() {
        for {
            <-sig
            fmt.Println("I stop when I want to")
        }
    }()

    for i:=0; i<6; i++{
        fmt.Println("Loops are fun")
        time.Sleep(1 * time.Second)
    }
}
