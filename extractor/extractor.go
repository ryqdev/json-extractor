package extractor

import (
	"github.com/ryqdev/golang_utils/stack"
)

type JsonMaster interface {
	hasJson(s string) bool
	getJson(s string) []string
	reset()
}

type JsonBoy struct {
	stack     stack.Stack[byte]
	initCap   int
	jsonStart int
	jsonEnd   int
}

func New() *JsonBoy {
	return &JsonBoy{
		stack:     stack.New[byte](),
		initCap:   0,
		jsonStart: -1,
		jsonEnd:   -1,
	}
}

func (j JsonBoy) hasJson(s string) bool {
	// MVP version: only detect braces: { and }
	for _, char := range s {
		switch char {
		case '{':
			{
				j.stack.Push('{')
			}
		case '}':
			{
				if j.stack.IsEmpty() {
					return false
				} else {
					j.stack.Pop()
				}
			}
		}
	}
	return j.stack.IsEmpty()
}

func (j JsonBoy) getJson(s string) []string {
	jsons := make([]string, 0)
	// TODO: how to record the json efficiently
	for _, char := range s {
		switch char {
		case '{':
			{
				j.stack.Push('{')
			}
		case '}':
			{
				if j.stack.IsEmpty() || j.stack.Top() != '}' {
					j.reset()
				}
			}
		case '"':
			{
				if j.stack.IsEmpty() || j.stack.Top() == '"' {
					j.stack.Pop()
				} else if j.stack.Top() == '{' {
					j.stack.Push('"')
				} else {
					j.reset()
				}
			}
		}
	}

	return jsons
}

func (j JsonBoy) reset() {
	j = JsonBoy{
		stack:     stack.New[byte](),
		initCap:   0,
		jsonStart: -1,
		jsonEnd:   -1,
	}
}
