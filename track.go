package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/mitchellh/go-homedir"
)

type InterruptEntry struct {
	date  time.Time
	label string
}

func (entry InterruptEntry) String() string {
	return fmt.Sprintf("%s: %s", entry.date.Format("2006-01-02 15:04:05"), entry.label)
}

func main() {
	fmt.Printf("Hello World")

	if len(os.Args) < 2 {
		fmt.Println("Invalid number of arguments")
	}

	var isInterrupt = flag.Bool("i", false, "Create a new interrupt")
	flag.Parse()

	fmt.Println("interrupt has value ", *isInterrupt)

	if *isInterrupt {
		createInterrupt("test")
	}

	// Usage:
	// track -i "XXX" # Track an interrupt with label "XXX"
	// track "XXX" # Add track event with label "XXX"
	// track --stop # Add stop track event
}

func createInterrupt(label string) {
	var dir = ensureTrackFolder()
	var filename = ensureTrackFile(dir)
	var entry = createInterruptEntry(label)

	fmt.Println("Create interrupt ", label, dir, entry)

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = fmt.Fprintf(f, "%s\n", entry.String()); err != nil {
		panic(err)
	}
}

func ensureTrackFolder() string {
	var dir, err = homedir.Dir()
	if err != nil {
		fmt.Println("oh no ", err)
	}

	var trackDir = path.Join(dir, ".track")
	err = os.MkdirAll(trackDir, os.ModePerm)
	if err != nil {
		fmt.Println("oh no ", err)
	}

	fmt.Println("ensure track file ", trackDir)
	return trackDir
}

func ensureTrackFile(dir string) string {
	var filename = time.Now().Format("2006-01-02")
	return path.Join(dir, filename)
}

func createInterruptEntry(label string) InterruptEntry {
	return InterruptEntry{label: label, date: time.Now()}
}
