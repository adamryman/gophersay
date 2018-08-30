package gopher

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"time"

	"github.com/adamryman/gophersay/gopherart"
)

var (
	sayings   []string
	gopherArt string
)

func init() {
	// Load in gopher ascii art from go-bindata
	gopherArtBytes, err := gopherart.Asset("gopherart/gopher.ascii")
	if err != nil {
		log.Fatal(err)
	}
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

func Proverb(w io.Writer) {
	rand.Seed(time.Now().UnixNano())
	saying := sayings[rand.Intn(len(sayings))]
	Say(w, saying)
}

func Say(w io.Writer, saying string) {
	fmt.Fprintf(w, " ------------------------\n%s\n%s",
		saying,
		gopherArt,
	)
}
