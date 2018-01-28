//go:generate go-bindata -o gopherart/gopherart.go -pkg gopherart gopherart/gopher.ascii
package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/adamryman/gophersay/gopherart"
)

var sayings []string
var gopherArt string
var stdIn bool

var isWindows bool

func init() {
	flag.BoolVar(&isWindows, "w", false, "Use to make stdin work on windows")

	// Load in gopher ascii art from go-bindata
	gopherArtBytes, _ := gopherart.Asset("gopherart/gopher.ascii")
	gopherArt = string(gopherArtBytes)
	sayings = []string{
		"Don't communicate by sharing memory, share memory by communicating.", "Concurrency is not parallelism.",
		"Channels orchestrate; mutexes serialize.",
		"The bigger the interface, the weaker the abstraction.",
		"Make the zero value useful.",
		"interface{} says nothing.",
		"Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.",
		"A little copying is better than a little dependency.",
		"Syscall must always be guarded with build tags.",
		"Cgo must always be guarded with build tags.",
		"Cgo is not Go.",
		"With the unsafe package there are no guarantees.",
		"Clear is better than clever.",
		"Reflection is never clear.",
		"Errors are values.",
		"Don't just check errors, handle them gracefully.",
		"Design the architecture, name the components, document the details.",
		"Documentation is for users.",
		"Don't panic.",
	}
}

func main() {
	flag.Parse()
	var saying string

	// If there are any command line arguments, join them together with spaces
	// and output them as the saying
	// otherwise output a saying from the list
	if isStdIn() {
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
		saying = strings.Join(stdInSlice, "")

	} else if len(flag.Args()) > 0 {
		saying = strings.Join(flag.Args(), " ")
	} else {
		// Seed rand for psudo random numbers
		rand.Seed(time.Now().UnixNano())

		saying = sayings[rand.Intn(len(sayings))]
	}

	fmt.Println(" ------------------------")
	fmt.Println(saying)
	fmt.Print(gopherArt)
}

func isStdIn() bool {
	if isWindows {
		if terminal.IsTerminal(int(os.Stdin.Fd())) {
			return false
		}
		return true
	}

	stat, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		return true
	}
	return false
}
