package store

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"time"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/zlypher/track/interrupt"
)

const (
	trackDirectory     = ".track"
	trackInterruptFile = "_interrupts"
)

// // Storeable defines operation to save/load data TODO store interface
// type Storeable interface {
// 	Save()
// 	Load()
// }

// SaveInterrupt saves a single interrupt entry
func SaveInterrupt(entry interrupt.Entry) {
	dir := ensureTrackFolder()
	file := ensureInterruptFile(dir)
	writeInterruptEntry(entry, file)
}

// SaveTrack saves a single track entry
func SaveTrack(entry interrupt.Entry) {
	dir := ensureTrackFolder()
	file := ensureTrackFile(dir)
	writeTrackEntry(entry, file)
}

// SaveStop saves a single stop entry
func SaveStop(entry interrupt.Entry) {
	dir := ensureTrackFolder()
	file := ensureTrackFile(dir)
	writeStopEntry(entry, file)
}

// LoadInterrupts loads the file and returns the loaded data
func LoadInterrupts() []interrupt.Entry {
	dir := ensureTrackFolder()
	interruptFilename := ensureInterruptFile(dir)

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
	return ensureTrackFolder()
}

func ensureTrackFolder() string {
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

func ensureInterruptFile(dir string) string {
	return path.Join(dir, trackInterruptFile)
}

func ensureTrackFile(dir string) string {
	filename := time.Now().Format("2006-01-02")
	return path.Join(dir, filename)
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

func writeInterruptEntry(entry interrupt.Entry, filename string) {
	writeToFile(filename, entry.String())
}

func writeTrackEntry(entry interrupt.Entry, filename string) {
	writeToFile(filename, entry.String())
}

func writeStopEntry(entry interrupt.Entry, filename string) {
	writeToFile(filename, entry.String())
}
