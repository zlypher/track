package store

import (
	"bufio"
	"fmt"
	"os"
	"path"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/zlypher/track/interrupt"
)

const (
	trackDirectory     = ".track"
	trackFile          = "_tasks"
	trackInterruptFile = "_interrupts"
)

// SaveInterrupt saves a single interrupt entry
func SaveInterrupt(entry interrupt.Entry) {
	dir := createAndGetTrackFolder()
	file := path.Join(dir, trackInterruptFile)
	writeToFile(file, entry.String())
}

// SaveTrack saves a single track entry
func SaveTrack(entry interrupt.Entry) {
	dir := createAndGetTrackFolder()
	file := path.Join(dir, trackFile)
	writeToFile(file, entry.String())
}

// SaveStop saves a single stop entry
func SaveStop(entry interrupt.Entry) {
	dir := createAndGetTrackFolder()
	file := path.Join(dir, trackFile)
	writeToFile(file, entry.String())
}

// LoadInterrupts loads the file and returns the loaded data
func LoadInterrupts() []interrupt.Entry {
	dir := createAndGetTrackFolder()
	interruptFilename := path.Join(dir, trackInterruptFile)

	file, err := os.OpenFile(interruptFilename, os.O_CREATE|os.O_RDONLY, 0600)
	if err != nil {
		fmt.Println("oh no ", err)
	}

	defer file.Close()

	var entries []interrupt.Entry

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		entry := interrupt.ParseEntry(line)
		entries = append(entries, entry)
	}

	return entries
}

// Location returns the current location of the stored data
func Location() string {
	return createAndGetTrackFolder()
}

func createAndGetTrackFolder() string {
	dir, err := homedir.Dir()
	if err != nil {
		fmt.Println("oh no ", err)
	}

	trackDir := path.Join(dir, trackDirectory)
	err = os.MkdirAll(trackDir, os.ModePerm)
	if err != nil {
		fmt.Println("oh no ", err)
	}

	return trackDir
}

func writeToFile(file string, content string) {
	f, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = fmt.Fprintf(f, "%s\n", content); err != nil {
		panic(err)
	}
}
