package bst
import (
    "fmt"
)

type ADT interface {
    Insert(interface{})error
    Update(interface{})error
    Delete(interface{})error
}

type Node struct {
    Left *Node
    Value int
    Right *Node
}

type BST struct {
    root *Node
}

func findInsertPos(root *Node, val int) *Node {
    for val <= root.Value {
        if root.Left != nil {
            return findInsertPos(root.Left, val)
        } else {
            return root
        }
    }
    for val > root.Value {
        if root.Right != nil {
            return findInsertPos(root.Right, val)
        } else {
            return root
        }
    }
    return nil
}

func NewBST() *BST {
    return &BST{root:nil}
}

func (b *BST) Insert(data int) {
    n := Node{
        Left:nil,
        Value: data,
        Right:nil,
    }
    if b.root == nil {
        b.root = &n
        return
    }

    pos := findInsertPos(b.root, data)
    if data > pos.Value {
        pos.Right = &n
    } else {
        pos.Left = &n
    }
}

func preorderTraverse(root *Node) {
    if root == nil {
        return
    }
    fmt.Println(root.Value, "left:", root.Left, "right:", root.Right)
    preorderTraverse(root.Left)
    preorderTraverse(root.Right)
}

func(b *BST) Print() {
    preorderTraverse(b.root)
}