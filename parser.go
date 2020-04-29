package newick

import (
	"fmt"
	"regexp"
	"strconv"
)

const (
	JSON = "json"
)

type tree *node

type node struct {
	Name      string  `json:"name"`
	Length    float64 `json:"length"`
	BranchSet []*node `json:"branchset"`
}

func Parse(s string) tree {
	ancestors := make([]*node, 0)
	t := &node{}

	tokens := make([]string, 0)
	r := regexp.MustCompile(`(\(|\)|;|:|,|\d+\.\d+|\w+)`)
	matches := r.FindAllStringSubmatch(s, -1)
	for _, match := range matches {
		tokens = append(tokens, match[0])
	}

	for i, token := range tokens {
		switch token {
		case "(":
			subtree := &node{}
			t.BranchSet = []*node{subtree}
			ancestors = append(ancestors, t)
			t = subtree
		case ",":
			subtree := &node{}
			ancestors[len(ancestors)-1].BranchSet = append(ancestors[len(ancestors)-1].BranchSet, subtree)
			t = subtree
		case ")":
			t = ancestors[len(ancestors)-1]
			ancestors = ancestors[0 : len(ancestors)-1]
		case ":":
			break
		default:
			x := tokens[i-1]
			if x == ")" || x == "(" || x == "," {
				t.Name = token
			} else if x == ":" {
				f, _ := strconv.ParseFloat(token, 64)
				fmt.Println("float: ", token)
				t.Length = f
			}
		}
	}

	return t
}
