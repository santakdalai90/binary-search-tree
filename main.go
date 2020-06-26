package main

import (
    "fmt"
    "main/bst"
)

func main() {
    bst := bst.NewBST()
    bst.Insert(2)
    bst.Insert(1)
    bst.Insert(3)
    bst.Print()
    fmt.Println("Hello World")
}