package goodhosts

import (
	"fmt"
	"testing"
)

func TestItemInSlice(t *testing.T) {
	item := "this"
	list := []string{"hello", "brah"}
	result := itemInSlice("goodbye", list)
	if result {
		t.Error(fmt.Sprintf("'%s' should not have been found in slice.", item))
	}

	item = "hello"
	result = itemInSlice(item, list)
	if !result {
		t.Error(fmt.Sprintf("'%s' should have been found in slice.", item))
	}
}

func Test_itemInSlice(t *testing.T) {
	type args struct {
		item string
		list []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := itemInSlice(tt.args.item, tt.args.list); got != tt.want {
				t.Errorf("itemInSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildRawLine(t *testing.T) {
	type args struct {
		ip    string
		hosts []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildRawLine(tt.args.ip, tt.args.hosts); got != tt.want {
				t.Errorf("buildRawLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
