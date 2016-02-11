package main 

import (
    "fmt"
    "strings"
)
func spacer() {
    fmt.Println(strings.Repeat("=", 80))
}

func testSquare(r * Rectangle) {
    if r.IsSqaure() {
        fmt.Println("It's a Square!")
    } else {
        fmt.Println("It's just a Rectangle")
    }
}

func main() {
    spacer()
    //Definition of variable beforehand
    fmt.Println("Creating rect1")
    var rect1 Rectangle
    rect1 = Rectangle{3, 4}
    fmt.Println(rect1)
    testSquare(&rect1)

    spacer()
    fmt.Println("Creating rect2")
    //Usage of the assignment operator 
    rect2:= new(Rectangle)
    rect2.length = 4
    rect2.width = 4
    fmt.Println(*rect2)
    testSquare(rect2)

    spacer()
    fmt.Println("Creating circ1")
    var circ1 Circle
    circ1 = Circle{3}
    fmt.Println(circ1)
    fmt.Printf("circ1's area: %f\n", circ1.Area())

    spacer()
    fmt.Println("Creating circ2")
    circ2 := Circle{4}
    fmt.Println(circ2)
    fmt.Printf("circ2's area: %f\n", circ2.Area())
}
