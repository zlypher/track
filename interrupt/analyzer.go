package interrupt

import (
	"fmt"
	"math"
	"sort"
	"time"
)

// Statistic represents an interruption statistic
type Statistic struct {
	numberOfInterrupts int
	avgPerDay          float32
	perPerson          map[string]int
	firstDate          time.Time
	entries            []Entry
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
		perPerson:          calculatePerPerson(entries),
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
	for person, count := range statistic.perPerson {
		fmt.Printf("%s: %d\n", person, count)
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

func calculatePerPerson(entries []Entry) map[string]int {
	m := make(map[string]int)

	for _, entry := range entries {
		_, ok := m[entry.label]
		if !ok {
			m[entry.label] = 1
		} else {
			m[entry.label] = m[entry.label] + 1
		}
	}

	return m
}
