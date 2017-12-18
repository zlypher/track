package main

import (
	"fmt"
	"os"

	"github.com/zlypher/track/interrupt"
	"github.com/zlypher/track/store"
)

const (
	listCommand      = "list"
	interruptCommand = "int"
	locationCommand  = "location"
	versionCommand   = "version"
	startCommand     = "start"
	stopCommand      = "stop"
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
	case startCommand:
		executeTrackCommand()
	case stopCommand:
		executeStopCommand()
	case locationCommand:
		executeLocationCommand()
	case versionCommand:
		executeVersionCommand()
	default:
		executeUsageCommand()
	}
}

func executeUsageCommand() {
	fmt.Println("Usage ...")
}

func executeVersionCommand() {
	fmt.Printf("Version: %s\n", Version)
}

func executeListCommand() {
	data := store.LoadInterrupts()
	statistic := interrupt.DoAnalyze(data)
	interrupt.PrintStatistic(statistic)
}

func executeInterruptCommand() {
	if len(os.Args) < 3 {
		fmt.Println("oh no not enough args")
		return
	}

	label := os.Args[2]
	entry := interrupt.CreateEntry(label)
	store.SaveInterrupt(entry)
}

func executeLocationCommand() {
	fmt.Println(store.Location())
}

func executeTrackCommand() {
	if len(os.Args) < 3 {
		fmt.Println("oh no not enough args")
		return
	}

	label := os.Args[2]
	entry := interrupt.CreateEntry(label)
	store.SaveTrack(entry)
}

func executeStopCommand() {
	entry := interrupt.CreateEntry("stop")
	store.SaveStop(entry)
}
