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
	isComplete bool
	stack      stack.Stack[byte]
	initCap    int
	jsonStart  int
	jsonEnd    int
}

func New() *JsonBoy {
	return &JsonBoy{
		isComplete: true,
		stack:      stack.New[byte](),
		initCap:    0,
		jsonStart:  -1,
		jsonEnd:    -1,
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
		case '[':
			{
				// TODO
			}
		case ']':
			{
				// TODO
			}
		case ':':
			{
				// TODO
			}
		}
	}

	return jsons
}

func (j JsonBoy) reset() {
	j = JsonBoy{
		isComplete: true,
		stack:      stack.New[byte](),
		initCap:    0,
		jsonStart:  -1,
		jsonEnd:    -1,
	}
}
