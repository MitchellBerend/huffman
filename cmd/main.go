package main

import (
	"fmt"
	"huffman"
	"huffman/node"
	"log"
	"os"
)

func main() {
	content, err := os.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)

	}
	// content := []byte("Lorem ipsum dolor sit amet consectetur adipiscing elit. Quisque faucibus ex sapien vitae pellentesque sem placerat. In id cursus mi pretium tellus duis convallis. Tempus leo eu aenean sed diam urna tempor. Pulvinar vivamus fringilla lacus nec metus bibendum egestas. Iaculis massa nisl malesuada lacinia integer nunc posuere. Ut hendrerit semper vel class aptent taciti sociosqu. Ad litora torquent per conubia nostra inceptos himenaeos.")

	fullTree := node.BuildTree(content)

	compressed := huffman.Encode(content, fullTree)

	_ = os.WriteFile("compressed", compressed, 0666)

	decoded := huffman.Decode(compressed, fullTree)

	if !(string(content) == string(decoded)) {
		// node.PrintTree(fullTree, "    ", false)
		fmt.Printf("encoded:\n|%s|\n", content)
		fmt.Printf("decoded:\n|%s|\n", decoded)
	} else {
		fmt.Println("correct round trip")
	}
}
