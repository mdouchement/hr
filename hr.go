package main

import (
	"bufio"
	"flag"
	"fmt"
	hrlib "github.com/dustin/go-humanize"
	"os"
	"strconv"
	"strings"
)

const version = "1.0.0"

func main() {
	v := flag.Bool("v", false, "Print out the version")
	h := flag.Bool("h", false, "Print usage")
	flag.Parse()

	if *v {
		printVersion()
		os.Exit(1)
	}
	if *h {
		printUsage()
		os.Exit(1)
	}

	info, _ := os.Stdin.Stat()
	if (info.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
		printUsage()
		os.Exit(1)
	} else if info.Size() > 0 {
		reader := bufio.NewReader(os.Stdin)
		humanize(reader)
	}
}

func humanize(r *bufio.Reader) {
	input, err := r.ReadString('\n')
	check(err)
	input = strings.Trim(input, "\n")

	val, err := strconv.ParseUint(input, 10, 64)
	check(err)

	// Formating
	hr := hrlib.Bytes(val)
	hr = strings.Replace(hr, "B", "", -1)
	hr = strings.Replace(hr, " ", "", -1)

	fmt.Println(hr)
}

func printUsage() {
	printVersion()
	fmt.Println("The command is intended to work with pipes.")
	fmt.Println("Usage:")
	fmt.Println("  echo 500000 | hr")
	fmt.Println("")
	fmt.Println("Options:")
	fmt.Println("  -h      show help")
	fmt.Println("  -v      show version")
}

func printVersion() {
	fmt.Printf("hr version %s\n\n", version)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
