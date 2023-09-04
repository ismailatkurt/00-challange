package app

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"
)

// used to compare delivery hours accurately
func convertTo24HourFormat(s string) uint8 {
	f := s[len(s)-2:]
	n, _ := strconv.Atoi(s[:len(s)-2])

	if f == "AM" {
		if n == 12 {
			return 0
		}
		return uint8(n)
	}

	if n == 12 {
		return 12
	}

	return uint8(n + 12)
}

// used when printing the final result to stdout
func convertTo12HourFormat(i uint8) string {
	switch {
	case i == 0:
		return "12AM"
	case i == 12:
		return "12PM"
	case i < 12:
		return strconv.Itoa(int(i)) + "AM"
	default:
		return strconv.Itoa(int(i-12)) + "PM"
	}
}

// PrintExecutionDetails Printing Execution details time and memory usage to "stderr"
// Should be accessible outside of app package, called by main
func PrintExecutionDetails(duration time.Duration) {
	_, _ = fmt.Fprintf(os.Stderr, "Execution took: %s\n", duration)
	_, _ = fmt.Fprintf(os.Stderr, "Memory usage: %s\n", printMemUsage())
}

func printMemUsage() string {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	s := fmt.Sprintf("Alloc = %v MiB", bToMb(m.Alloc))
	s += fmt.Sprintf("\tCumulative Memory Allocation = %v MiB", bToMb(m.TotalAlloc))
	s += fmt.Sprintf("\tSys = %v MiB", bToMb(m.Sys))
	s += fmt.Sprintf("\tNum of Garbage Collector Cycles = %v\n", m.NumGC)

	return s
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
