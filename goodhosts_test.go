package goodhosts

import (
	"fmt"
	"testing"
)

func TestHostsLineIsComment(t *testing.T) {
	comment := "   # This is a comment   "
	line := NewHostsLine(comment)
	result := line.IsComment()
	if !result {
		t.Error(fmt.Sprintf("'%s' should be a comment"), comment)
	}
}

func TestHostsAddEntryWhenIpMissing(t *testing.T) {
	inputIp := "10.0.0.7"
	hosts := NewHosts()

	line1 := NewHostsLine("127.0.0.1 yadda")
	line2 := NewHostsLine("10.0.0.5 nada")

	hosts.Lines = []HostsLine{line1, line2}

	hosts.RemoveEntry(inputIp, "nada")
	if len(hosts.Lines) > 1 {
		t.Error("Remove entry failed to remove entry.")
	}
}
