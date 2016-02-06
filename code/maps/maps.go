package main

import "fmt"

func spacer() {
    fmt.Println("============================================")
}

func main() {
    
    /*     +-> make a thing
           |    +-> Specify that we want a map
           |    |    +-> Keyed on strings
           |    |    |      +-> Integers are the values
           |    |    |      |*/
    age := make(map[string] int)

    age["Teddy"] = 42
    fmt.Printf("%s is %d years old\n", "Teddy", age["Teddy"])
    
    age["John"] = 43
    fmt.Println(age)

    delete(age, "Teddy")
    fmt.Println((age))
}
