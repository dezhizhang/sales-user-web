package main

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
)

//func main() {
//
//	// Create a new Node with a Node number of 1
//	node, err := snowflake.NewNode(1)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	// Generate a snowflake ID.
//	id := node.Generate()
//
//	// Print out the ID in a few different ways.
//	fmt.Printf("Int64  ID: %d\n", id)
//	fmt.Printf("String ID: %s\n", id)
//}

//func main() {
//	node, err := snowflake.NewNode(1)
//	if err != nil {
//		panic(err)
//	}
//	id := node.Generate()
//	fmt.Println(id)
//}

//func main() {
//	node, err := snowflake.NewNode(1)
//	if err != nil {
//		panic(err)
//	}
//	id := node.Generate()
//	fmt.Println(id)
//}

func main() {
	node, err := snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}
	id := node.Generate().String()
	fmt.Println(id)
}
