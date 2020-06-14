package main 

import (
	"strings"
	"fmt"
)
const (
	QUEEN = 'Q'
	SPACE = '.'
)

type Board struct {
	arr [][]byte
	x1, x2, colum []bool
}

func (b *Board) size() int {
	return len(b.arr)
}

func (b *Board) put(c byte, i, j int) {
	b.arr[i][j] = c
	b.colum[j] = true
	b.x1[i+j] = true
	b.x2[i-j+b.size()-1] = true
}

func (b *Board) clear(i, j int) {
	b.arr[i][j] = '.'
	b.colum[j] = false
	b.x1[i+j] = false
	b.x2[i-j+b.size()-1] = false 
}

func (b *Board) checkQ(i, j int) bool {
	return !b.colum[j] && !b.x1[i+j] && !b.x2[i-j+b.size()-1]
}

func (b *Board) toStringArray() []string {
	var sb strings.Builder
	ret := make([]string, 0, b.size())
	for _, arr := range b.arr {
		sb.Reset()
		for _, c := range arr {
			sb.WriteByte(c)
		}
		ret = append(ret, sb.String())
	}
	return ret
}

func solve(b *Board, r int, ret *[][]string) {
	if r == b.size() {
		*ret = append(*ret, b.toStringArray())
		return
	}
	for i :=0 ; i < b.size(); i++ {
		if b.checkQ(r, i) {
			b.put(QUEEN, r, i)
			solve(b, r+1, ret)
			b.clear(r, i)
		}
	}
}

func solveNQueens(n int) [][]string {
	array := make([]([]byte), n)
	for i, _ := range array {
		array[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			array[i][j] = SPACE
		}
	}
	board := &Board{arr: array, colum: make([]bool, n), x1 : make([]bool, 2*n-1), x2: make([]bool, 2*n-1)}
	ret := make([][]string, 0, 100)
	solve(board, 0, &ret)
	printNQueens(ret)
	return ret
}

func printNQueens(ans [][]string) {
	for _, board := range ans {
		for index, s := range board {
			fmt.Print(s)
			if index != len(board) {
				fmt.Println(",")
			} else {
				fmt.Println()
			}
		}
		fmt.Println()
	}
}

func main() {
	solveNQueens(4)
}

