package interrupt

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"time"
)

// Statistic represents an interruption statistic
type Statistic struct {
	numberOfInterrupts int
	avgPerDay          float32
	firstDate          time.Time
	entries            []Entry
}

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

// DoAnalyze analyzes the given list of interruption entries and generates
// an interruption statistic.
func DoAnalyze(entries []Entry) Statistic {
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].date.Before(entries[j].date)
	})

	return Statistic{
		numberOfInterrupts: len(entries),
		avgPerDay:          calculateAveragePerDay(entries),
		firstDate:          time.Now(),
		entries:            entries}
}

// PrintStatistic prints the given interruption statistic
func PrintStatistic(statistic Statistic) {
	if statistic.numberOfInterrupts == 0 {
		fmt.Println("You weren't interrupted yet!")
		return
	}

	fmt.Printf("You were interrupted %v times since %v\n", statistic.numberOfInterrupts, statistic.firstDate.Format("2006-01-02"))
	fmt.Printf("On average %v times per day\n", statistic.avgPerDay)

	fmt.Println("-----------------------------------")
	for _, entry := range statistic.entries {
		fmt.Println(entry.String())
	}
	fmt.Println("-----------------------------------")
}

func calculateAveragePerDay(entries []Entry) float32 {
	numEntries := len(entries)
	if numEntries == 0 {
		return 0
	}

	if numEntries == 1 {
		return 1
	}

	first := entries[0]
	last := entries[numEntries-1]
	numDays := float32(math.Abs(first.date.Sub(last.date).Hours() / 24))
	return float32(numEntries) / numDays
}
