package main

import "fmt"

func subtractEm(args ... int) int {
    final := 0
    for i, val := range args {
        if i == 0 {
            final = args[0]
        } else {
            final -= val
        }
    }
    return final
}

func SwapInt(x int, y int) (int, int) {
             /*Specify both parameter types*/
                            /*specify both return types*/
    return y, x
}

func SwapStr(x, y string) (string, string) {
            /*abbreviated only once if there are duplicates*/
    return y, x
}

func Add(x, y int) int {
    return x + y
}

func factorial(x int) int {
    if x == 0 {
        return 1
    }
    return x * factorial(x - 1)
}

func printOne() {fmt.Println(1)}
func printTwo() {fmt.Println(2)}


func main() {
    var int1 int = 4
    int2 := 44

    var str1 string = "four"
    str2 := "fourty four"

    fmt.Println("int1:", int1)
    fmt.Println("int2:", int2)
    fmt.Println("str1:", str1)
    fmt.Println("str2:", str2)

    int1, int2 = SwapInt(int1, int2)
    str1, str2 = SwapStr(str1, str2)

    fmt.Println("int1:", int1)
    fmt.Println("int2:", int2)
    fmt.Println("str1:", str1)
    fmt.Println("str2:", str2)

    fmt.Println("Result:", Add(int1, int2))
    below()
    fmt.Printf("%d - %d: %d\n", int1, int2, subtractEm(int1, int2))

    spacer()
    fmt.Println("Creating a function within a function, AKA, a Closure.")
    doubleInt1 := func() int {
        return int1 * 2

    }
    fmt.Println("Int1:", int1)
    fmt.Println("doubleInt1:", doubleInt1())
    
    spacer()

    fmt.Println(factorial(int1))

    spacer()
    fmt.Println("Using the defer statement.")
    /*The defer statement tells the compiler to run this function when the 
    function completes.  This allows you to properly clean something up, like
    closing a file
    ``` pseudo code
    func stuff
        open(file)
        defer close(file)
        readFrom(file)
        writeTo(file)
        return 0
    // The deferred function will now run here
    ```*/
    defer printTwo()
    printOne()


}

func spacer() {
    fmt.Println("==================================================")
}
func below() {
    fmt.Println("You can even call functions below the main function.")
}
