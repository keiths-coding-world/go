package main

import (
"fmt"
"bufio"
)

func main() {
    //reading an integer
    var age int
    fmt.Println("What is your age?")
    _, err: fmt.Scan(&age)

    //reading a string
    reader := bufio.newReader(os.Stdin)
    var name string
    fmt.Println("What is your name?")
    name, _ := reader.readString("\n")

    fmt.Println("Your name is ", name, " and you are age ", age)
}
