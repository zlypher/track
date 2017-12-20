package interrupt

import (
	"reflect"
	"testing"
	"time"
)

func TestDoAnalyze(t *testing.T) {
	entries := []Entry{
		Entry{Label: "a", Date: time.Now()},
		Entry{Label: "b", Date: time.Now()},
		Entry{Label: "a", Date: time.Now()},
	}

	expected := Statistic{
		numberOfInterrupts: 3,
		avgPerDay:          3,
		perPerson:          map[string]int{"a": 2, "b": 1},
		firstDate:          time.Now(), // TODO
		entries:            entries,
	}

	stats := DoAnalyze(entries)
	checkDoAnalyzeResult(t, stats, expected)
}

func TestDoAnalyzeEmpty(t *testing.T) {
	entries := []Entry{}

	expected := Statistic{
		numberOfInterrupts: 0,
		avgPerDay:          0,
		perPerson:          map[string]int{},
		firstDate:          time.Now(), // TODO
		entries:            entries,
	}

	stats := DoAnalyze(entries)
	checkDoAnalyzeResult(t, stats, expected)
}

func TestDoAnalyzeOne(t *testing.T) {
	entries := []Entry{
		Entry{Label: "a", Date: time.Now()},
	}

	expected := Statistic{
		numberOfInterrupts: 1,
		avgPerDay:          1,
		perPerson:          map[string]int{"a": 1},
		firstDate:          time.Now(), // TODO
		entries:            entries,
	}

	stats := DoAnalyze(entries)
	checkDoAnalyzeResult(t, stats, expected)
}

func checkDoAnalyzeResult(t *testing.T, actual Statistic, expected Statistic) {
	if actual.numberOfInterrupts != expected.numberOfInterrupts {
		t.Errorf("numberOfInterrupts from DoAnalyze() was incorrect, got: %d, want: %d.", actual.numberOfInterrupts, expected.numberOfInterrupts)
	}
	if actual.avgPerDay != expected.avgPerDay {
		t.Errorf("avgPerDay from DoAnalyze() was incorrect, got: %v, want: %v.", actual.avgPerDay, expected.avgPerDay)
	}
	if !reflect.DeepEqual(actual.entries, expected.entries) {
		t.Errorf("entries from DoAnalyze() was incorrect, got: %v, want: %v.", actual.entries, expected.entries)
	}

}
