package main

import (
	"flag"
	"fmt"
	"github.com/lextoumbourou/goodhosts"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	showComments := flag.Bool("all", false, "Show comments when listing.")

	flag.Parse()

	args := flag.Args()

	if len(args) > 0 {
		command := args[0]
		hosts, err := goodhosts.NewHosts()
		check(err)

		switch command {
		case "list":
			total := 0
			for _, line := range hosts.Lines {
				var lineOutput string

				if line.IsComment() && !*showComments {
					continue
				}

				lineOutput = fmt.Sprintf("%s", line.Raw)
				if line.Err != nil {
					lineOutput = fmt.Sprintf("%s # <<< Malformated!", lineOutput)
				}
				total += 1

				fmt.Println(lineOutput)
			}

			fmt.Print("\nTotal: %d\n", total)

			return
		case "check":
			if len(os.Args) < 3 {
				fmt.Println("usage: goodhosts check 127.0.0.1 facebook.com")
				os.Exit(1)
			}

			ip := os.Args[2]
			host := os.Args[3]

			if !hosts.Has(ip, host) {
				fmt.Fprintln(os.Stderr, fmt.Sprintf("%s %s is not in the hosts file", ip, host))
				os.Exit(1)
			}

			return
		case "add":
			if len(os.Args) < 3 {
				fmt.Println("usage: goodhosts add 127.0.0.1 facebook.com")
				os.Exit(1)
			}

			ip := os.Args[2]
			inputHosts := os.Args[3:]

			if !hosts.IsWritable() {
				fmt.Fprintln(os.Stderr, "Host file not writable. Try running with elevated privileges.")
				os.Exit(1)
			}

			err = hosts.Add(ip, inputHosts...)
			if err != nil {
				fmt.Fprintln(os.Stderr, fmt.Sprintf("%s", err.Error()))
				os.Exit(2)
			}

			err = hosts.Flush()
			check(err)

			return
		case "rm", "remove":
			if len(os.Args) < 3 {
				fmt.Println("usage: goodhost remove 127.0.0.1 facebook.com")
				os.Exit(1)
			}

			ip := os.Args[2]
			inputHosts := os.Args[3:]

			if !hosts.IsWritable() {
				fmt.Fprintln(os.Stderr, "Host file not writable. Try running with elevated privileges.")
				os.Exit(1)
			}

			err = hosts.Remove(ip, inputHosts...)
			if err != nil {
				fmt.Fprintf(os.Stderr, fmt.Sprintf("%s\n", err.Error()))
				os.Exit(2)
			}

			err = hosts.Flush()
			check(err)

			return
		}
	}

	fmt.Println("Add --help for usage.")
	os.Exit(2)
}
