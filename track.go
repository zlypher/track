package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/mitchellh/go-homedir"
)

var interruptDateFormat = "2006-01-02 15:04:05"

type InterruptEntry struct {
	date  time.Time
	label string
}

func (entry InterruptEntry) String() string {
	return fmt.Sprintf("%s - %s", entry.date.Format(interruptDateFormat), entry.label)
}

func parse(line string) InterruptEntry {
	parts := strings.Split(line, " - ")
	dateStr := parts[0]
	labelStr := parts[1]

	parsedDate, _ := time.Parse(interruptDateFormat, dateStr)

	return InterruptEntry{label: labelStr, date: parsedDate}
}

const (
	listCommand      = "list"
	interruptCommand = "int"
	locationCommand  = "location"
)

func main() {
	if len(os.Args) < 2 {
		executeUsageCommand()
		return
	}

	// Usage:
	// track -i "XXX" # Track an interrupt with label "XXX"
	// track list
	// track "XXX" # Add track event with label "XXX"
	// track --stop # Add stop track event

	switch os.Args[1] {
	case listCommand:
		executeListCommand()
	case interruptCommand:
		executeInterruptCommand()
	case locationCommand:
		executeLocationCommand()
	default:
		executeUsageCommand()
	}
}

func executeUsageCommand() {
	fmt.Println("Usage ...")
}

func executeListCommand() {
	data := readInterruptData()

	for _, entry := range data {
		fmt.Println(entry.String())
	}
}

func readInterruptData() []InterruptEntry {
	dir := ensureTrackFolder()
	interruptFilename := ensureInterruptFile(dir)

	file, err := os.OpenFile(interruptFilename, os.O_CREATE|os.O_RDONLY, 0600)
	if err != nil {
		fmt.Println("oh no ", err)
	}

	defer file.Close()

	var entries []InterruptEntry

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		entry := parse(line)
		entries = append(entries, entry)
	}

	return entries
}

func executeInterruptCommand() {
	if len(os.Args) < 3 {
		fmt.Println("oh no not enough args")
		return
	}

	label := os.Args[2]
	createInterrupt(label)
}

func executeLocationCommand() {
	fmt.Println(ensureTrackFolder())
}

func executeGeneralCommand() {
	// isInterrupt := flag.Bool("i", false, "Create a new interrupt")
	// flag.Parse()

	// if *isInterrupt {
	// 	createInterrupt("test")
	// }
}

func createInterrupt(label string) {
	dir := ensureTrackFolder()
	trackFilename := ensureTrackFile(dir)
	interruptFilename := ensureInterruptFile(dir)
	entry := createInterruptEntry(label)

	writeInterruptEntry(entry, trackFilename)
	writeInterruptEntry(entry, interruptFilename)
}

func ensureTrackFolder() string {
	dir, err := homedir.Dir()
	if err != nil {
		fmt.Println("oh no ", err)
	}

	trackDir := path.Join(dir, ".track")
	err = os.MkdirAll(trackDir, os.ModePerm)
	if err != nil {
		fmt.Println("oh no ", err)
	}

	return trackDir
}

func ensureTrackFile(dir string) string {
	filename := time.Now().Format("2006-01-02")
	return path.Join(dir, filename)
}

func ensureInterruptFile(dir string) string {
	return path.Join(dir, "_interrupts")
}

func createInterruptEntry(label string) InterruptEntry {
	return InterruptEntry{label: label, date: time.Now()}
}

func writeInterruptEntry(entry InterruptEntry, filename string) {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = fmt.Fprintf(f, "%s\n", entry.String()); err != nil {
		panic(err)
	}
}
