package interrupt

import (
	"fmt"
	"strings"
	"time"
)

// interruptDateFormat represents the format for interruption dates
var interruptDateFormat = "2006-01-02 15:04:05"

// Entry represents an interruption entry
type Entry struct {
	Date  time.Time
	Label string
}

func (entry Entry) String() string {
	return fmt.Sprintf("%s - %s", entry.Date.Format(interruptDateFormat), entry.Label)
}

// ParseEntry parses a given string into a new interruption entry
func ParseEntry(line string) Entry {
	parts := strings.Split(line, " - ")
	dateStr := parts[0]
	labelStr := parts[1]

	parsedDate, _ := time.Parse(interruptDateFormat, dateStr)

	return Entry{Label: labelStr, Date: parsedDate}
}
