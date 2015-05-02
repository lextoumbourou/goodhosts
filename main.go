package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/user"
	"strings"
)

const commentChar string = "#"

// To do: add Windows support.
const hostsFilePath string = "/etc/hosts"

type hostsFileLine struct {
	ip    string
	hosts []string
}

func check(err error) {
	if err != nil {
		panic(err)
	}
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

func getHostLines(includeComments bool) []string {
	var output []string

	hostsFile, err := os.Open(hostsFilePath)
	check(err)
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
		log.Fatal(err)
	}

	return output
}

func list() {
	lines := getHostLines(false)
	for i := range lines {
		fmt.Printf("%s\n", lines[i])
	}
}

func lineInHosts(ip string, host string) int {
	var line hostsFileLine

	lines := getHostLines(true)
	for i := range lines {
		if !isComment(lines[i]) && lines[i] != "" {
			line = parseLine(lines[i])
			if ip == line.ip {
				if itemInSlice(host, line.hosts) {
					return i
				}
			}
		}
	}

	return -1
}

func rewriteHosts(lines []string) error {
	file, err := os.Create(hostsFilePath)
	check(err)

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

func addLine(ip string, host string) {
	lines := getHostLines(true)
	position := getIpPosition(ip, lines)

	if position == -1 {
		// Ip line is not in file, so we just append our new line.
		endLine := fmt.Sprintf("%s %s", ip, host)
		lines = append(lines, endLine)
	} else {
		// Otherwise, we replace the line in the correct position
		lines[position] = fmt.Sprintf("%s %s", lines[position], host)
	}

	err := rewriteHosts(lines)
	check(err)
}

func removeLine(ip string, host string) {
	lines := getHostLines(true)
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

	err := rewriteHosts(lines)
	check(err)
}

func main() {
	if len(os.Args) > 1 {
		command := os.Args[1]
		switch command {
		case "list":
			list()
			return
		case "check":
			if len(os.Args) < 3 {
				fmt.Println("usage: goodhost check 127.0.0.1 salt")
				os.Exit(1)
			}

			ip := os.Args[2]
			host := os.Args[3]
			if lineInHosts(ip, host) == -1 {
				fmt.Fprintf(os.Stderr, "%s %s is not in the hosts file\n", ip, host)
				os.Exit(1)
			}
			return
		case "add":
			if len(os.Args) < 3 {
				fmt.Println("usage: goodhost add 127.0.0.1 myhost")
				os.Exit(1)
			}
			user, err := user.Current()
			check(err)

			ip := os.Args[2]
			host := os.Args[3]
			if lineInHosts(ip, host) != -1 {
				fmt.Fprintf(os.Stderr, "Line already in host file. Nothing to do.\n")
				os.Exit(2)
			}

			if user == nil || user.Uid != "0" {
				fmt.Fprintf(os.Stderr, "Need to be root user. Try running with sudo.\n")
				os.Exit(1)
			}

			addLine(ip, host)
			return
		case "remove":
			if len(os.Args) < 3 {
				fmt.Println("usage: goodhost add 127.0.0.1 myhost")
				os.Exit(1)
			}
			user, err := user.Current()
			check(err)

			ip := os.Args[2]
			host := os.Args[3]
			if lineInHosts(ip, host) == -1 {
				fmt.Fprintf(os.Stderr, "Line not in host file. Nothing to do.\n")
				os.Exit(3)
			}

			if user == nil || user.Uid != "0" {
				fmt.Fprintf(os.Stderr, "Need to be root user. Try running with sudo.\n")
				os.Exit(1)
			}

			removeLine(ip, host)
			return
		}
	}

	fmt.Println("Help should go here.")
}
