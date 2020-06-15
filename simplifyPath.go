package main 

import (
	"fmt"
	"strings"
)

type Stack struct {
	data []string
}

func (s *Stack) push(x string) {
	s.data = append(s.data, x)
}

func (s *Stack) pop() string {
	if len(s.data) == 0 {
		return "" 
	}
	var x string
	x, s.data = s.data[len(s.data)-1], s.data[0: len(s.data)-1]
	return x
}

func (s *Stack) toString() string {
	return "/" + strings.Join(s.data, "/")
}

func simplifyPath(path string) string {
	segs := strings.Split(path, "/")
	fmt.Println(segs)
	stack := Stack{data: make([]string, 0, 10)}
	for _, s := range segs {
		switch s {
		case "..":
			stack.pop()
		case ".", "":
			 break
		default:
			stack.push(s)
		}

	}
	return stack.toString()
}

func main() {
	fmt.Println(simplifyPath("/a//b////c/d//././/.."))
}

