package stopwatch

import (
	"fmt"
	"time"
)

// Example showing most typical usage.
func ExampleStopwatch() {
	s := NewStopwatch()
	s.Start()
	// [measure some stuff]
	s.Lap("Mock1")
	s.Lap("Mock2")
	s.Stop("Mock3")

	// mock time for testing
	delay := 600 * time.Millisecond
	s.laps[0].duration = delay
	s.laps[1].duration = 2 * delay
	s.laps[2].duration = 2150 * delay

	fmt.Println(s)
	// Output:
	// Mock1: 600ms
	// Mock2: 1.2s
	// Mock3: 21m30s
}
