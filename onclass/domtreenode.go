package main

import "text/template/parse"

type Node struct {
	tag string
	id string
	text string
	children []*Node
}

func main()  {
	p := Node{
		tag: 	"p",
		text: 	"This is a simple HTML document.",
		id: 	"foo",
	}

	h1 := Node{
		tag: "h1",
		text: "Hello, World!",
	}

	html := Node{
		tag: "html",
		children: []*Node{&p, &h1},
	}
}

func findById(root *Node, id string) *Node {
	//
}

func findByIdDFS(node *Node, id string) *Node  {
	if node.id == id {
		return node
	}

	if len(node.children) > 0 {
		for _, child := range node.children {
			findByIdDFS(child, id)
		}
	}
	return nil
}
