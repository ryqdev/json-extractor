package tests

import (
	"github.com/ryqdev/json_extractor/extractor"
	"testing"
)

func TestHasJson(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"{}", true},
		{"bb{a:1}", true},
		{"{a:{b:2}}cc", true},
		{"{a:1", false},
		{"{a:1}}", false},
		{"{a:{b:2}}}", false},
	}

	for _, test := range tests {
		j := extractor.New()
		result := j.HasJson(test.input)
		if result != test.expected {
			t.Errorf("hasJson(%q) = %v; expected %v", test.input, result, test.expected)
		}
	}
}

func TestGetJson(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"{}", []string{"{}"}},
		{"aa,{a:1}", []string{"{a:1}"}},
		{"{a:{b:2}}bbb", []string{"{a:{b:2}}"}},
		{"{a:1}{b:2}", []string{"{a:1}", "{b:2}"}},
		{"{a:1}{b:2}{c:3}laksjf[]11", []string{"{a:1}", "{b:2}", "{c:3}"}},
		{"{a:1}{b:2{c:3}}", []string{"{a:1}", "{b:2{c:3}}"}},
	}

	for _, test := range tests {
		j := extractor.New()
		result := j.GetJson(test.input)
		if !equal(result, test.expected) {
			t.Errorf("getJson(%q) = %v; expected %v", test.input, result, test.expected)
		}
	}
}

// Helper function to compare slices of strings
func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
