package goodhosts

import (
	"fmt"
	"reflect"
	"testing"
)

func TestHostsLineIsComment(t *testing.T) {
	testWithAppendToLine(t, func(t *testing.T, appendToLine bool) {
		comment := "   # This is a comment   "
		line := NewHostsLine(comment)
		result := line.IsComment()
		if !result {
			t.Error(fmt.Sprintf("'%s' should be a comment", comment))
		}
	})
}

func TestNewHostsLineWithEmptyLine(t *testing.T) {
	testWithAppendToLine(t, func(t *testing.T, appendToLine bool) {
		line := NewHostsLine("")
		if line.Raw != "" {
			t.Error("Failed to load empty line.")
		}
	})
}

func TestHostsHas(t *testing.T) {
	testWithAppendToLine(t, func(t *testing.T, appendToLine bool) {
		hosts := new(Hosts)
		hosts.Lines = []HostsLine{
			NewHostsLine("127.0.0.1 yadda"), NewHostsLine("10.0.0.7 nada")}

		// We should find this entry.
		if !hosts.Has("10.0.0.7", "nada") {
			t.Error("Couldn't find entry in hosts file.")
		}

		// We shouldn't find this entry
		if hosts.Has("10.0.0.7", "shuda") {
			t.Error("Found entry that isn't in hosts file.")
		}
	})
}

func TestHostsHasDoesntFindMissingEntry(t *testing.T) {
	testWithAppendToLine(t, func(t *testing.T, appendToLine bool) {
		hosts := new(Hosts)
		hosts.Lines = []HostsLine{
			NewHostsLine("127.0.0.1 yadda"), NewHostsLine("10.0.0.7 nada")}

		if hosts.Has("10.0.0.7", "brada") {
			t.Error("Found missing entry.")
		}
	})
}

func TestHostsAddWhenIpHasOtherHosts(t *testing.T) {
	testWithAppendToLine(t, func(t *testing.T, appendToLine bool) {
		hosts := new(Hosts)
		hosts.Lines = []HostsLine{
			NewHostsLine("127.0.0.1 yadda"), NewHostsLine("10.0.0.7 nada yadda")}

		_ = hosts.Add("10.0.0.7", "brada", "yadda")

		var expectedLines []HostsLine
		if appendToLine {
			expectedLines = []HostsLine{
				NewHostsLine("127.0.0.1 yadda"),
				NewHostsLine("10.0.0.7 nada yadda brada")}
		} else {
			expectedLines = []HostsLine{
				NewHostsLine("127.0.0.1 yadda"),
				NewHostsLine("10.0.0.7 nada yadda"),
				NewHostsLine("10.0.0.7 brada")}
		}

		if !reflect.DeepEqual(hosts.Lines, expectedLines) {
			t.Error("Add entry failed to append entry.")
		}
	})
}

func TestHostsAddWhenIpDoesntExist(t *testing.T) {
	testWithAppendToLine(t, func(t *testing.T, appendToLine bool) {
		hosts := new(Hosts)
		hosts.Lines = []HostsLine{
			NewHostsLine("127.0.0.1 yadda")}

		_ = hosts.Add("10.0.0.7", "brada", "yadda")

		var expectedLines []HostsLine
		if appendToLine {
			expectedLines = []HostsLine{
				NewHostsLine("127.0.0.1 yadda"), NewHostsLine("10.0.0.7 brada yadda")}
		} else {
			expectedLines = []HostsLine{
				NewHostsLine("127.0.0.1 yadda"), NewHostsLine("10.0.0.7 brada"), NewHostsLine("10.0.0.7 yadda")}
		}

		if !reflect.DeepEqual(hosts.Lines, expectedLines) {
			t.Error("Add entry failed to append entry.")
		}
	})
}

func TestHostsRemoveWhenLastHostIpCombo(t *testing.T) {
	testWithAppendToLine(t, func(t *testing.T, appendToLine bool) {
		hosts := new(Hosts)
		hosts.Lines = []HostsLine{
			NewHostsLine("127.0.0.1 yadda"), NewHostsLine("10.0.0.7 nada")}

		_ = hosts.Remove("10.0.0.7", "nada")

		expectedLines := []HostsLine{NewHostsLine("127.0.0.1 yadda")}

		if !reflect.DeepEqual(hosts.Lines, expectedLines) {
			t.Error("Remove entry failed to remove entry.")
		}
	})
}

func TestHostsRemoveWhenIpHasOtherHosts(t *testing.T) {
	testWithAppendToLine(t, func(t *testing.T, appendToLine bool) {
		hosts := new(Hosts)

		hosts.Lines = []HostsLine{
			NewHostsLine("127.0.0.1 yadda"), NewHostsLine("10.0.0.7 nada brada")}

		_ = hosts.Remove("10.0.0.7", "nada")

		expectedLines := []HostsLine{
			NewHostsLine("127.0.0.1 yadda"), NewHostsLine("10.0.0.7 brada")}

		if !reflect.DeepEqual(hosts.Lines, expectedLines) {
			t.Error("Remove entry failed to remove entry.")
		}
	})
}

func TestHostsRemoveMultipleEntries(t *testing.T) {
	testWithAppendToLine(t, func(t *testing.T, appendToLine bool) {
		hosts := new(Hosts)
		hosts.Lines = []HostsLine{
			NewHostsLine("127.0.0.1 yadda nadda prada")}

		_ = hosts.Remove("127.0.0.1", "yadda", "prada")
		if hosts.Lines[0].Raw != "127.0.0.1 nadda" {
			t.Error("Failed to remove multiple entries.")
		}
	})
}

func TestHostsLine_IsComment(t *testing.T) {
	tests := []struct {
		name string
		l    HostsLine
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.IsComment(); got != tt.want {
				t.Errorf("HostsLine.IsComment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewHostsLine(t *testing.T) {
	type args struct {
		raw string
	}
	tests := []struct {
		name string
		args args
		want HostsLine
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHostsLine(tt.args.raw); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHostsLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHosts_IsWritable(t *testing.T) {
	tests := []struct {
		name string
		h    *Hosts
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.IsWritable(); got != tt.want {
				t.Errorf("Hosts.IsWritable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHosts_Load(t *testing.T) {
	tests := []struct {
		name    string
		h       *Hosts
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.Load(); (err != nil) != tt.wantErr {
				t.Errorf("Hosts.Load() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHosts_Flush(t *testing.T) {
	tests := []struct {
		name    string
		h       Hosts
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.Flush(); (err != nil) != tt.wantErr {
				t.Errorf("Hosts.Flush() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHosts_Add(t *testing.T) {
	type args struct {
		ip    string
		hosts []string
	}
	tests := []struct {
		name    string
		h       *Hosts
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.Add(tt.args.ip, tt.args.hosts...); (err != nil) != tt.wantErr {
				t.Errorf("Hosts.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHosts_Has(t *testing.T) {
	type args struct {
		ip   string
		host string
	}
	tests := []struct {
		name string
		h    Hosts
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Has(tt.args.ip, tt.args.host); got != tt.want {
				t.Errorf("Hosts.Has() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHosts_Remove(t *testing.T) {
	type args struct {
		ip    string
		hosts []string
	}
	tests := []struct {
		name    string
		h       *Hosts
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.Remove(tt.args.ip, tt.args.hosts...); (err != nil) != tt.wantErr {
				t.Errorf("Hosts.Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHosts_getHostPosition(t *testing.T) {
	type args struct {
		ip   string
		host string
	}
	tests := []struct {
		name string
		h    Hosts
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.getHostPosition(tt.args.ip, tt.args.host); got != tt.want {
				t.Errorf("Hosts.getHostPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHosts_getIpPosition(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name string
		h    Hosts
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.getIpPosition(tt.args.ip); got != tt.want {
				t.Errorf("Hosts.getIpPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewHosts(t *testing.T) {
	tests := []struct {
		name    string
		want    Hosts
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewHosts()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewHosts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHosts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func testWithAppendToLine(t *testing.T, testFunc func(t *testing.T, appendToLine bool)) {
	for _, testCase := range []struct {
		appendToLine bool
	}{{true}, {false}} {
		t.Run(fmt.Sprintf("appendToLine=%v", testCase.appendToLine), func(t *testing.T) {
			appendToLine = testCase.appendToLine
			testFunc(t, testCase.appendToLine)
		})
	}
}
