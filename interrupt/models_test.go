package interrupt

import (
	"testing"
	"time"
)

func TestParseEntry(t *testing.T) {
	tests := []struct {
		line  string
		date  time.Time
		label string
	}{
		{"2017-12-13 19:48:46 - test", time.Date(2017, 12, 13, 19, 48, 46, 0, time.UTC), "test"},
		{"2015-01-02 19:48:46 - Hello World", time.Date(2015, 01, 02, 19, 48, 46, 0, time.UTC), "Hello World"},
	}

	for _, test := range tests {
		entry := ParseEntry(test.line)
		if entry.Label != test.label {
			t.Errorf("Parsed entry of (%s) was incorrect, got: %s, want: %s.", test.line, entry.Label, test.label)
		}
		if entry.Date != test.date {
			t.Errorf("Parsed entry of (%s) was incorrect, got: %s, want: %s.", test.line, entry.Date, test.date)
		}
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		entry    Entry
		expected string
	}{
		{Entry{Label: "Test", Date: time.Date(2015, 01, 02, 19, 48, 46, 0, time.UTC)}, "2015-01-02 19:48:46 - Test"},
		{Entry{Label: "Hello World", Date: time.Date(2002, 04, 03, 19, 0, 0, 0, time.UTC)}, "2002-04-03 19:00:00 - Hello World"},
	}

	for _, test := range tests {
		stringified := test.entry.String()
		if stringified != test.expected {
			t.Errorf("Entry.String() was incorrect, got: %s, want: %s.", stringified, test.expected)
		}
	}
}
