package goodhosts

import (
	"fmt"
	"reflect"
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

func TestNewHostsLineWithEmptyLine(t *testing.T) {
	line := NewHostsLine("")
	if line.Raw != "" {
		t.Error("Failed to load empty line.")
	}
}

func TestHostsHasEntryFindsEntry(t *testing.T) {
	hosts := new(Hosts)
	hosts.Lines = []HostsLine{
		NewHostsLine("127.0.0.1 yadda"), NewHostsLine("10.0.0.7 nada")}
	if !hosts.HasEntry("10.0.0.7", "nada") {
		t.Error("Failed to find entry.")
	}
}

func TestHostsHasEntryDoesntFindMissingEntry(t *testing.T) {
	hosts := new(Hosts)
	hosts.Lines = []HostsLine{
		NewHostsLine("127.0.0.1 yadda"), NewHostsLine("10.0.0.7 nada")}

	if hosts.HasEntry("10.0.0.7", "brada") {
		t.Error("Found missing entry.")
	}
}

func TestHostsAddEntryWhenIpHasOtherHosts(t *testing.T) {
	hosts := new(Hosts)
	hosts.Lines = []HostsLine{
		NewHostsLine("127.0.0.1 yadda"), NewHostsLine("10.0.0.7 nada")}

	hosts.AddEntry("10.0.0.7", "brada")

	expectedLines := []HostsLine{
		NewHostsLine("127.0.0.1 yadda"), NewHostsLine("10.0.0.7 nada brada")}

	if !reflect.DeepEqual(hosts.Lines, expectedLines) {
		t.Error("Add entry failed to append entry.")
	}
}

func TestHostsAddEntryWhenIpDoesntExist(t *testing.T) {
	hosts := new(Hosts)
	hosts.Lines = []HostsLine{
		NewHostsLine("127.0.0.1 yadda")}

	hosts.AddEntry("10.0.0.7", "brada")

	expectedLines := []HostsLine{
		NewHostsLine("127.0.0.1 yadda"), NewHostsLine("10.0.0.7 brada")}

	if !reflect.DeepEqual(hosts.Lines, expectedLines) {
		t.Error("Add entry failed to append entry.")
	}
}

func TestHostsAddMultipleEntries(t *testing.T) {
	hosts := new(Hosts)
	hosts.Lines = []HostsLine{
		NewHostsLine("127.0.0.1 yadda")}

	hosts.AddEntry("127.0.0.1", "brada", "prada", "nada")
	if hosts.Lines[0].Raw != "127.0.0.1 yadda brada prada nada" {
		t.Error("Failed to add multiple entries.")
	}
}

func TestHostsRemoveEntryWhenLastHostIpCombo(t *testing.T) {
	hosts := new(Hosts)
	hosts.Lines = []HostsLine{
		NewHostsLine("127.0.0.1 yadda"), NewHostsLine("10.0.0.7 nada")}

	hosts.RemoveEntry("10.0.0.7", "nada")

	expectedLines := []HostsLine{NewHostsLine("127.0.0.1 yadda")}

	if !reflect.DeepEqual(hosts.Lines, expectedLines) {
		t.Error("Remove entry failed to remove entry.")
	}
}

func TestHostsRemoveEntryWhenIpHasOtherHosts(t *testing.T) {
	hosts := new(Hosts)

	hosts.Lines = []HostsLine{
		NewHostsLine("127.0.0.1 yadda"), NewHostsLine("10.0.0.7 nada brada")}

	hosts.RemoveEntry("10.0.0.7", "nada")

	expectedLines := []HostsLine{
		NewHostsLine("127.0.0.1 yadda"), NewHostsLine("10.0.0.7 brada")}

	if !reflect.DeepEqual(hosts.Lines, expectedLines) {
		t.Error("Remove entry failed to remove entry.")
	}
}

func TestHostsRemoveMultipleEntries(t *testing.T) {
	hosts := new(Hosts)
	hosts.Lines = []HostsLine{
		NewHostsLine("127.0.0.1 yadda nadda prada")}

	hosts.RemoveEntry("127.0.0.1", "yadda", "prada")
	if hosts.Lines[0].Raw != "127.0.0.1 nadda" {
		t.Error("Failed to remove multiple entries.")
	}
}
