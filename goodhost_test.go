package goodhost

import (
	"fmt"
	"testing"
)

func TestIsComment(t *testing.T) {
	comment := "   # This is a comment   "
	result := isComment(comment)
	if !result {
		t.Error(fmt.Sprintf("'%s' is a comment"))
	}
}

func TestItemInSlice(t *testing.T) {
	item := "this"
	list := []string{"hello", "brah"}
	result := itemInSlice("goodbye", list)
	if result {
		t.Error(fmt.Sprintf("'%' should not have been found in slice.", item))
	}

	item = "hello"
	result = itemInSlice(item, list)
	if !result {
		t.Error(fmt.Sprintf("'%' should have been found in slice.", item))
	}
}
