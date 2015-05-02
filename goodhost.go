package goodhost

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const commentChar string = "#"

// To do: add Windows support.
const hostsFilePath string = "/etc/hosts"

type hostsFileLine struct {
	ip    string
	hosts []string
}

func isComment(line string) bool {
	trimLine := strings.TrimSpace(line)
	isComment := strings.HasPrefix(trimLine, commentChar)
	return isComment
}

func parseLine(line string) hostsFileLine {
	var output hostsFileLine

	fields := strings.Fields(line)
	if len(fields) == 0 {
		log.Fatal(fmt.Sprintf("Unable to parse line: %q", line))
	}

	output = hostsFileLine{ip: fields[0], hosts: fields[1:]}
	return output
}

func itemInSlice(item string, list []string) bool {
	for _, i := range list {
		if i == item {
			return true
		}
	}

	return false
}

func GetLines(includeComments bool) ([]string, error) {
	var output []string

	hostsFile, err := os.Open(hostsFilePath)
	if err != nil {
		return output, err
	}

	defer hostsFile.Close()

	scanner := bufio.NewScanner(hostsFile)
	for scanner.Scan() {
		line := scanner.Text()
		if isComment(line) && !includeComments {
			continue
		}

		output = append(output, line)
	}

	if err := scanner.Err(); err != nil {
		return output, err
	}

	return output, nil
}

func getPosition(ip string, host string) (int, error) {
	var line hostsFileLine

	lines, err := GetLines(true)
	if err != nil {
		return -1, err
	}

	for i := range lines {
		if !isComment(lines[i]) && lines[i] != "" {
			line = parseLine(lines[i])
			if ip == line.ip && itemInSlice(host, line.hosts) {
				return i, nil
			}
		}
	}

	return -1, nil
}

func rewriteHosts(lines []string) error {
	file, err := os.Create(hostsFilePath)
	if err != nil {
		return err
	}

	w := bufio.NewWriter(file)

	for _, line := range lines {
		fmt.Fprintln(w, line)
	}

	return w.Flush()
}

func getIpPosition(ip string, lines []string) int {
	var line hostsFileLine

	for i := range lines {
		if !isComment(lines[i]) && lines[i] != "" {
			line = parseLine(lines[i])
			if line.ip == ip {
				return i
			}
		}
	}

	return -1
}

func HasEntry(ip string, host string) (bool, error) {
	pos, err := getPosition(ip, host)
	return pos != -1, err
}

func AddEntry(ip string, host string) error {
	lines, err := GetLines(true)
	if err != nil {
		return err
	}

	position := getIpPosition(ip, lines)

	if position == -1 {
		// Ip line is not in file, so we just append our new line.
		endLine := fmt.Sprintf("%s %s", ip, host)
		lines = append(lines, endLine)
	} else {
		// Otherwise, we replace the line in the correct position
		lines[position] = fmt.Sprintf("%s %s", lines[position], host)
	}

	err = rewriteHosts(lines)
	return err
}

func RemoveEntry(ip string, host string) error {
	lines, err := GetLines(true)
	if err != nil {
		return err
	}

	pos := getIpPosition(ip, lines)
	line := parseLine(lines[pos])

	hostPos := -1
	for i := range line.hosts {
		if line.hosts[i] == host {
			hostPos = i
		}
	}

	newHosts := append(line.hosts[:hostPos], line.hosts[hostPos+1:]...)
	if len(newHosts) == 0 {
		// Just remove the line if there's no new hosts.
		lines = append(lines[:pos], lines[pos+1:]...)
	} else {
		newLine := line.ip
		for i := range newHosts {
			newLine = fmt.Sprintf("%s %s", newLine, newHosts[i])
		}
		lines[pos] = newLine
	}

	err = rewriteHosts(lines)
	return err
}
