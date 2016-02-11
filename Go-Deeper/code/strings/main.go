package main

import (
	"fmt"
)

func main() {
	var runearray []rune = []rune("Hello!")
	var bytearray []byte = []byte("Hello!")
	var astring string = "Hello!"

    fmt.Println("\nUsing: " + astring + "\nCasted type\tByte Length\tValue")
    fmt.Printf("String\t\t%d bytes\t\t%s\n", len(astring), astring)
    fmt.Printf("Byte array\t%d bytes\t\t%s\n", len(bytearray), bytearray)
    fmt.Printf("Rune array\t%d bytes\t\t%s\n", len(runearray), runearray)

    // The rune array output looks a little funny. And look at that length. A
    // rune is a 4 byte UTF-8 value, but if you look at the output you will see
    // the rune is outputting the ASCII integer value, and is only taking up one
    // byte per rune. What gives??
    //
    // UTF-8 is a *variable length encoding*, so the rune only has to be as many
    // bytes as it takes to store the encoded character. In addition to that,
    // the first 7 bits of UTF-8 correspond exactly to the 7 bit ASCII space!
    // Lets try something a little more exotic

    const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
    fmt.Printf("\nUsing byte array: % x \n", sample)
    fmt.Printf("Length: %d bytes\n", len(sample))
    fmt.Printf("String output: %s\n", sample)

    // Interesting. Printf was given an 8 byte array, but only printed out 6
    // characters (One of them is a space so you cannot see it). This shows how
    // strings are just an interpretation of byte arrays. The last 3 bytes of
    // the array, 0xE2 0x8C 0x98, are the UTF-8 representation of ⌘

}
