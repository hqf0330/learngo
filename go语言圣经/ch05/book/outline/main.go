package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

/*
当outline递归调用自己的时候，被调用的函数会接收到栈的副本，尽管被调用者可能会对slice进行元素
的添加、修改甚至是创建新数组的操作，但它并不会修改调用者原来传递的元素，所以当被调函数返回的时候，
调用者的栈依旧保持原样
*/
func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)
}
