package goodhosts

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const commentChar string = "#"

type HostsLine struct {
	Ip    string
	Hosts []string
	Raw   string
}

// Return ```true``` if the line is a comment.
func (l HostsLine) IsComment() bool {
	trimLine := strings.TrimSpace(l.Raw)
	isComment := strings.HasPrefix(trimLine, commentChar)
	return isComment
}

// Create a new instance of ```HostsLine```.
func NewHostsLine(raw string) HostsLine {
	fields := strings.Fields(raw)
	if len(fields) == 0 {
		return HostsLine{Raw: raw}
	}

	return HostsLine{Ip: fields[0], Hosts: fields[1:], Raw: raw}
}

type Hosts struct {
	Path  string
	Lines []HostsLine
}

// Load the hosts file into ``l.Lines``.
func (h *Hosts) Load() error {
	var lines []HostsLine

	file, err := os.Open(h.Path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := NewHostsLine(scanner.Text())
		if err != nil {
			return err
		}

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	h.Lines = lines

	return nil
}

// Flush any changes made to hosts file.
func (h Hosts) Flush() error {
	file, err := os.Create(h.Path)
	if err != nil {
		return err
	}

	w := bufio.NewWriter(file)

	for _, line := range h.Lines {
		fmt.Fprintln(w, line.Raw)
	}

	return w.Flush()
}

// Add an entry to the hosts file.
func (h *Hosts) AddEntry(ip string, host string) {
	position := h.getIpPosition(ip)
	if position == -1 {
		// Ip line is not in file, so we just append our new line.
		endLine := NewHostsLine(fmt.Sprintf("%s %s", ip, host))
		h.Lines = append(h.Lines, endLine)
	} else {
		// Otherwise, we replace the line in the correct position
		endLine := NewHostsLine(fmt.Sprintf("%s %s", h.Lines[position].Raw, host))
		h.Lines[position] = endLine
	}
}

// Return a bool if ip/host combo in hosts file.
func (h Hosts) HasEntry(ip string, host string) (bool, error) {
	pos, err := h.getHostPosition(ip, host)

	return pos != -1, err
}

// Remove an entry from the hosts file.
func (h *Hosts) RemoveEntry(ip string, host string) error {
	pos := h.getIpPosition(ip)
	line := h.Lines[pos]

	hostPos := -1
	for i := range line.Hosts {
		if line.Hosts[i] == host {
			hostPos = i
		}
	}

	newHosts := append(line.Hosts[:hostPos], line.Hosts[hostPos+1:]...)
	if len(newHosts) == 0 {
		// Just remove the line if there's no new hosts.
		h.Lines = append(h.Lines[:pos], h.Lines[pos+1:]...)
	} else {
		newLineRaw := line.Ip
		for i := range newHosts {
			newLineRaw = fmt.Sprintf("%s %s", newLineRaw, newHosts[i])
		}
		newLine := NewHostsLine(newLineRaw)
		h.Lines[pos] = newLine
	}

	return nil
}

func (h Hosts) getHostPosition(ip string, host string) (int, error) {
	for i := range h.Lines {
		line := h.Lines[i]
		if !line.IsComment() && line.Raw != "" {
			if ip == line.Ip && itemInSlice(host, line.Hosts) {
				return i, nil
			}
		}
	}

	return -1, nil
}

func (h Hosts) getIpPosition(ip string) int {
	for i := range h.Lines {
		line := h.Lines[i]
		if !line.IsComment() && line.Raw != "" {
			if line.Ip == ip {
				return i
			}
		}
	}

	return -1
}

func NewHosts() Hosts {
	// To do: add Windows support.
	path := "/etc/hosts"
	host := Hosts{Path: path}

	return host
}
