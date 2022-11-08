package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Stack struct {
	data []int
}

func main() {
	s := &Stack{}
	input := "1 2 + 3 4 - *"
	inputArray := strings.Split(input, " ")

	for _, e := range inputArray {
		if isNumber(e) {
			num, _ := strconv.Atoi(e)
			s.push(num)
		} else {
			s.calcNum(e)
		}
	}
	fmt.Println(s.pop())
}

// 構造体Stackのdataプロパティの中で、最後の要素を取り出して返す
func (s *Stack) pop() int {
	if s.isEmpty() {
		return 0
	}

	last := len(s.data) - 1
	popData := s.data[last]
	s.data = s.data[:last]
	return popData
}

// 構造体Stackのdataプロパティの末尾に新たな要素を追加する
func (s *Stack) push(num int) {
	if s.isFull() {
		return
	}

	s.data = append(s.data, num)
}

// 構造体Stackのdataプロパティの要素の数を検証
func (s *Stack) isEmpty() bool {
	return len(s.data) == 0
}

// 構造体Stackのdataプロパティの要素の数を検証
func (s *Stack) isFull() bool {
	return len(s.data) == 99
}

func isNumber(num string) bool {
	if "0" <= num && num <= "9" {
		return true
	}

	return false
}

func (s *Stack) calcNum(e string) {
	num1 := s.pop()
	num2 := s.pop()

	switch e {
	case "+":
		s.push(num2 + num1)
	case "-":
		s.push(num2 - num1)
	case "*":
		s.push(num2 * num1)
	default:
		fmt.Println("不正な値です。")
	}
}
