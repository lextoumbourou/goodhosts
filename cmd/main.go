package main

import (
	"fmt"
	"github.com/lextoumbourou/goodhost"
	"os"
	"os/user"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	if len(os.Args) > 1 {
		command := os.Args[1]
		switch command {
		case "list":
			lines, err := goodhost.GetLines(false)
			check(err)

			for i := range lines {
				fmt.Printf("%s\n", lines[i])
			}
			return
		case "check":
			if len(os.Args) < 3 {
				fmt.Println("usage: goodhost check 127.0.0.1 salt")
				os.Exit(1)
			}

			ip := os.Args[2]
			host := os.Args[3]
			hasEntry, err := goodhost.HasEntry(ip, host)
			check(err)

			if !hasEntry {
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
			hasEntry, err := goodhost.HasEntry(ip, host)
			check(err)

			if hasEntry {
				fmt.Fprintf(os.Stderr, "Line already in host file. Nothing to do.\n")
				os.Exit(2)
			}

			if user == nil || user.Uid != "0" {
				fmt.Fprintf(os.Stderr, "Need to be root user. Try running with sudo.\n")
				os.Exit(1)
			}

			err = goodhost.AddEntry(ip, host)
			check(err)
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
			hasEntry, err := goodhost.HasEntry(ip, host)
			check(err)

			if !hasEntry {
				fmt.Fprintf(os.Stderr, "Line not in host file. Nothing to do.\n")
				os.Exit(3)
			}

			if user == nil || user.Uid != "0" {
				fmt.Fprintf(os.Stderr, "Need to be root user. Try running with sudo.\n")
				os.Exit(1)
			}

			err = goodhost.RemoveEntry(ip, host)
			check(err)
			return
		}
	}

	fmt.Println("Help should go here.")
}
