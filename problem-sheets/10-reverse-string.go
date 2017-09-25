/*
Problem description

Write a function to reverse a string in Go.
*/

package main

import (
    "./util"
    "fmt"
)

func main(){
    fmt.Println(util.Reverse("Hello"))
    fmt.Println(util.Reverse("World"))
    fmt.Println(util.Reverse("This string is reversed!"))
    fmt.Println(util.Reverse(util.Reverse("Reversed Twice")))
}

/*
Output:

olleH
dlroW
!desrever si gnirts sihT
Reversed Twice
*/