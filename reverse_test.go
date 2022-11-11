package goModuleTest

import (
	"fmt"
	"testing"
)

func TestReverseRunes11(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		got := ReverseRunes(c.in)
		if got != c.want {
			t.Errorf("ReverseRunes(%q) == %q, want %q", c.in, got, c.want)
		}
	}
	fmt.Println("======1======")
}

func TestReverseRunes22(t *testing.T) {
	fmt.Println("======2======")
}
