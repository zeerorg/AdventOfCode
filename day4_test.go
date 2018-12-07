package AdventOfCode

import (
	"testing"
	"time"
)

func TestGetTotalSleep(t *testing.T) {
	var cycles []SleepWake
	dur, _ := time.ParseDuration("0m")
	if GetTotalSleep(cycles) != dur {
		t.Fatal(dur, GetTotalSleep(cycles))
	}
}
