package trees

import (
	"testing"
)

func TestConvertTreeToJson(t *testing.T) {
	// Test case 1: Empty tree
	emptyTree := &TreeNode{Name: "root"}
	jsonData, err := ConvertTreeToJson(emptyTree)
	if err != nil {
		t.Errorf("Unexpected error while converting empty tree to JSON: %v", err)
	}
	expectedJSON := `{"name":"root"}`
	if string(jsonData) != expectedJSON {
		t.Errorf("Expected JSON: %s, but got: %s", expectedJSON, string(jsonData))
	}

	// Test case 2: Tree with children
	root := &TreeNode{Name: "root"}
	child1 := &TreeNode{Name: "child1"}
	child2 := &TreeNode{Name: "child2"}
	root.Children = []*TreeNode{child1, child2}
	jsonData, err = ConvertTreeToJson(root)
	if err != nil {
		t.Errorf("Unexpected error while converting tree with children to JSON: %v", err)
	}
	expectedJSON = `{"name":"root","children":[{"name":"child1"},{"name":"child2"}]}`
	if string(jsonData) != expectedJSON {
		t.Errorf("Expected JSON: %s, but got: %s", expectedJSON, string(jsonData))
	}

	// Test case 3: Error case (unmarshalable structure)
	invalidTree := &TreeNode{Name: "invalidNode"}
	invalidTree.Children = append(invalidTree.Children, invalidTree) // Circular reference
	_, err = ConvertTreeToJson(invalidTree)
	if err == nil {
		t.Error("Expected an error for an unmarshalable structure, but got none")
	}
}
