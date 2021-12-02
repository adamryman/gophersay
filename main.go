package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/adamryman/gophersay/gopher"
	"golang.org/x/term"
)

func main() {
	// If there are any command line arguments, join them together with spaces
	// and output them as the saying
	if len(os.Args) > 1 {
		gopher.Say(os.Stdout, strings.Join(os.Args[1:], " "))
		return
	}

	// If there is no data on stdin
	if term.IsTerminal(int(os.Stdin.Fd())) {
		gopher.Proverb(os.Stdout)
		return
	}

	// We found data on stdin
	data := os.Stdin
	scan := bufio.NewScanner(data)
	var stdInSlice []string
	for scan.Scan() {
		stdInSlice = append(stdInSlice, scan.Text()+"\n")
	}

	// Take the newline off the last line
	lastLine := stdInSlice[len(stdInSlice)-1]
	lastLine = strings.TrimSuffix(lastLine, "\n")
	stdInSlice[len(stdInSlice)-1] = lastLine

	// Join the lines
	gopher.Say(os.Stdout, strings.Join(stdInSlice, ""))
}
