package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const commentChar string = "#"

type hostsFileLine struct {
	ip     string
	hosts  []string
	lineNo int
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

func itemInSlice(item string, list []string) bool {
	for _, i := range list {
		if i == item {
			return true
		}
	}

	return false
}

func getHostLines() []hostsFileLine {
	var output []hostsFileLine

	hostsFile, err := os.Open("/etc/hosts")
	check(err)
	defer hostsFile.Close()

	lineNo := 0

	scanner := bufio.NewScanner(hostsFile)
	for scanner.Scan() {
		line := scanner.Text()
		if !isComment(line) {
			fields := strings.Fields(line)
			if len(fields) > 0 {
				line := hostsFileLine{ip: fields[0], hosts: fields[1:], lineNo: lineNo}
				output = append(output, line)
			}
		}
		lineNo += 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return output
}

func list() {
	lines := getHostLines()
	for i := range lines {
		line := lines[i]
		fmt.Printf("%s %s\n", line.ip, line.hosts)
	}
}

func lineInHosts(ip string, host string) bool {
	lines := getHostLines()
	for i := range lines {
		if ip == lines[i].ip {
			if itemInSlice(host, lines[i].hosts) {
				return true
			}
		}
	}

	return false
}

func addLine(ip string, host string) {
	fmt.Printf("About to add %s and %s to the hosts file.\n", ip, host)
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
			if !lineInHosts(ip, host) {
				fmt.Fprintf(os.Stderr, "%s %s is not in the hosts file\n", ip, host)
				os.Exit(1)
			}
			return
		case "add":
			if len(os.Args) < 3 {
				fmt.Println("usage: goodhost add 127.0.0.1 myhost")
				os.Exit(1)
			}
			ip := os.Args[2]
			host := os.Args[3]

			addLine(ip, host)
			return
		}
	}

	fmt.Println("Help should go here.")
}
