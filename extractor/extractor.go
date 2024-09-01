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
}

func New() *JsonBoy {
	return &JsonBoy{
		stack:     stack.New[byte](),
		initCap:   0,
		jsonStart: -1,
	}
}

func (j JsonBoy) HasJson(s string) bool {
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

func (j JsonBoy) GetJson(s string) []string {
	jsons := make([]string, 0)
	for i, v := range s {
		switch v {
		case '{':
			{
				if j.stack.IsEmpty() {
					j.jsonStart = i
				}
				j.stack.Push('{')
			}
		case '}':
			{
				if j.stack.IsEmpty() || j.stack.Top() != '{' {
					j.reset()
				} else {
					j.stack.Pop()
					if j.stack.IsEmpty() {
						jsons = append(jsons, s[j.jsonStart:i+1])
						j.reset()
					}
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
	}
}
