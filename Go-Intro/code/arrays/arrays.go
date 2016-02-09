package main

import "fmt"
func spacer() {
    fmt.Println("==============================")
}

func main() {
    var nums[5] float64 //Size and type definition 

    //Initialize values
    nums[0] = 163 
    nums[1] = 78557
    nums[2] = 132452634.9
    nums[3] = 3.14159
    nums[4] = 4

    //C style for loops
    for i := 0; i < 5; i++ {
        fmt.Printf("num[%d] is %f\n", i, nums[i]) //C Style printf
    }

    //Assign a 5 slot array of float64 with definition
    nums2 := [5]int {0,1,2,3,4}
    //The following will not work since it treats it as a code block with the {}
    //nums3 := {0,1,2,3,4}
    
    spacer()

    //Python style for loop
    //The first variable 'i' has the index
    //The second variable 'value' has the value
    for i, value := range nums2 {
        fmt.Printf("nums2[%d] is %d\n", i, value)
    }
    
    spacer()

    /*The '_' variable is special.  This variable cannot be used.
    This allows us to say that we don't care about the result.
    Kind of like a python 
    ```
        for item in items
    ```
    */

    for _, value := range nums2 {
        fmt.Printf("Ignore the iterator: %d\n", value)
    }

    spacer()

    //Slicing arrays and such
    numSlice1 := []int {5, 4, 3, 2, 1}
    numSlice2 := numSlice1[3:5] /*slice out position 3, 4, 5
                                  Since the 5th position doesn't exist,
                                  it ignores it, and there are only 2 items
                                  that will be printed out in the for loop.
                                */
    for i := 0; i < len(numSlice2); i++ {
        fmt.Printf("numSlice2[%d] = %d\n", i, numSlice2[i])
    }

    spacer()

    //Slicing arrays and such
    numSlice2 = numSlice1[:2] //slice out position from 0 up until 2
    for i := 0; i < len(numSlice2); i++ {
        fmt.Printf("numSlice2[%d] = %d\n", i, numSlice2[i])
    }

    spacer()

    //Slicing arrays and such
    numSlice2 = numSlice1[2:] //slice out from position 2 up until the end
    for i := 0; i < len(numSlice2); i++ {
        fmt.Printf("numSlice2[%d] = %d\n", i, numSlice2[i])
    }

    spacer()
                 /*
                 +->make function
                 |    +->Describe the type of slice that we want
                 |    |      +->Default value of 0 for the first 5 entries
                 |    |      |  +->Max size of 10 for the slice*/
    numSlice3 := make([]int, 5, 10)

    copy(numSlice3, numSlice1)
    for i := 0; i < len(numSlice3); i++ {
        fmt.Printf("numSlice3[%d] = %d\n", i ,numSlice3[i])
    }

    spacer()
    //appending arrays and stuff
    numSlice4 := append(numSlice3, 0, -1)
    for i := 0; i < len(numSlice4); i++ {
        fmt.Printf("numSlice4[%d] = %d\n", i ,numSlice4[i])
    }
    
    
}
