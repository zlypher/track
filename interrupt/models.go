package interrupt

import (
	"fmt"
	"strings"
	"time"
)

// InterruptDateFormat represents the format for interruption dates
var InterruptDateFormat = "2006-01-02 15:04:05"

// Entry represents an interruption entry
type Entry struct {
	date  time.Time
	label string
}

func (entry Entry) String() string {
	return fmt.Sprintf("%s - %s", entry.date.Format(InterruptDateFormat), entry.label)
}

// ParseEntry parses a given string into a new interruption entry
func ParseEntry(line string) Entry {
	parts := strings.Split(line, " - ")
	dateStr := parts[0]
	labelStr := parts[1]

	parsedDate, _ := time.Parse(InterruptDateFormat, dateStr)

	return Entry{label: labelStr, date: parsedDate}
}

// CreateEntry creates a new interruption entry with a given label
func CreateEntry(label string) Entry {
	return Entry{label: label, date: time.Now()}
}
