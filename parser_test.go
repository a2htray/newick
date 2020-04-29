package newick

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	s := "(A,B,((C,Y)c,D)e)f;"
	tree := Parse(s)

	output, err := json.MarshalIndent(tree, "", " ")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(output))
}

func TestNode_JSON(t *testing.T) {
	s := "(A,B,((C,Y)c,D)e)f;"
	tree := Parse(s)

	fmt.Println((*tree).JSON())
}

func TestNode_Depth(t *testing.T) {
	s := "(A,B,((C,Y)c,D)e)f;"
	tree := Parse(s)

	fmt.Println(tree.Depth())
}