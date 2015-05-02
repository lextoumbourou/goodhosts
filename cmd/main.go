package main

import (
	"fmt"
	"github.com/lextoumbourou/goodhosts"
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
		hosts := goodhosts.NewHosts()

		switch command {
		case "list":
			// To do: make this a flag
			hideComments := true

			err := hosts.Load()
			check(err)

			for _, line := range hosts.Lines {
				if line.IsComment() && hideComments {
					continue
				}

				fmt.Printf("%s\n", line.Raw)
			}

			fmt.Printf("\nTotal: %d\n", len(hosts.Lines))

			return
		case "check":
			if len(os.Args) < 3 {
				fmt.Println("usage: goodhosts check 127.0.0.1 facebook.com")
				os.Exit(1)
			}

			err := hosts.Load()
			check(err)

			ip := os.Args[2]
			host := os.Args[3]
			hasEntry, err := hosts.HasEntry(ip, host)
			check(err)

			if !hasEntry {
				fmt.Fprintf(os.Stderr, "%s %s is not in the hosts file\n", ip, host)
				os.Exit(1)
			}
			return
		case "add":
			if len(os.Args) < 3 {
				fmt.Println("usage: goodhosts add 127.0.0.1 facebook.com")
				os.Exit(1)
			}
			user, err := user.Current()
			check(err)

			err = hosts.Load()
			check(err)

			ip := os.Args[2]
			host := os.Args[3]
			hasEntry, err := hosts.HasEntry(ip, host)
			check(err)

			if hasEntry {
				fmt.Fprintf(os.Stderr, "Line already in host file. Nothing to do.\n")
				os.Exit(2)
			}

			if user == nil || user.Uid != "0" {
				fmt.Fprintf(os.Stderr, "Need to be root user. Try running with sudo.\n")
				os.Exit(1)
			}

			hosts.AddEntry(ip, host)

			hosts.Flush()
			return
		case "remove":
			if len(os.Args) < 3 {
				fmt.Println("usage: goodhost remove 127.0.0.1 facebook.com")
				os.Exit(1)
			}
			user, err := user.Current()
			check(err)

			err = hosts.Load()
			check(err)

			ip := os.Args[2]
			host := os.Args[3]
			hasEntry, err := hosts.HasEntry(ip, host)
			check(err)

			if !hasEntry {
				fmt.Fprintf(os.Stderr, "Line not in host file. Nothing to do.\n")
				os.Exit(3)
			}

			if user == nil || user.Uid != "0" {
				fmt.Fprintf(os.Stderr, "Need to be root user. Try running with sudo.\n")
				os.Exit(1)
			}

			err = hosts.RemoveEntry(ip, host)
			check(err)

			hosts.Flush()
			return
		}
	}

	fmt.Println("Help should go here.")
}
