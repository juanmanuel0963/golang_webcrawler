// Package for saving a sitemap as a tree structure
package trees

import (
	"encoding/json"
	"fmt"
)

// Structure to save the tree nodes
type TreeNode struct {
	Name     string      `json:"name"`
	Children []*TreeNode `json:"children,omitempty"`
}

// Convert the tree nodes to JSON
func ConvertTreeToJson(root *TreeNode) ([]byte, error) {

	//Marshal the node to json
	jsonData, err := json.Marshal(root)

	//if error
	if err != nil {

		//return null
		return nil, err
	}

	//return the json object
	return jsonData, nil
}

// Function to recursively print the tree
func PrintTree(node *TreeNode, level int) {

	//Print the node
	fmt.Printf("%s%s\n", GetIndent(level), node.Name)

	//Recursively print the node sons
	for _, child := range node.Children {

		//Print the node recursively
		PrintTree(child, level+1)
	}
}

// Helper function to get the appropriate indentation
func GetIndent(level int) string {
	indent := ""

	//Move the position to the right
	for i := 0; i < level; i++ {
		indent += "-"
	}
	return indent
}
