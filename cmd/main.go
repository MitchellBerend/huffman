package main

import (
	"huffman"
	"huffman/node"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "codec"}

	var encodeCmd = &cobra.Command{
		Use:   "encode [input]",
		Short: "Encode a string into a file",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			inputFilePath := args[0]

			inputFileContent, err := os.ReadFile(inputFilePath)
			if err != nil {
				log.Fatal(err)
			}

			tree := node.BuildTree(inputFileContent)
			encodedData := huffman.Encode(inputFileContent, tree)

			_ = os.WriteFile(inputFilePath+".compressed", encodedData, 0666)
			_ = os.WriteFile(inputFilePath+".tree", tree.ToBinary(), 0666)

			return nil
		},
	}

	var decodeCmd = &cobra.Command{
		Use:   "decode [compressed path] [tree path] [output path]",
		Short: "Decode a compressed file with a tree file",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			compressedPath := args[0]
			treePath := args[1]
			outputFilePath := args[2]

			compressedContent, err := os.ReadFile(compressedPath)
			if err != nil {
				log.Fatal(err)
			}

			treeContent, err := os.ReadFile(treePath)
			if err != nil {
				log.Fatal(err)
			}

			decodedData := huffman.Decode(compressedContent, node.BuildTreeFromBinary(treeContent))
			_ = os.WriteFile(outputFilePath, decodedData, 0666)

			return nil
		},
	}

	rootCmd.AddCommand(encodeCmd, decodeCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
