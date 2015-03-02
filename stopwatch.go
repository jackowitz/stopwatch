// Package stopwatch implements a simple set of functions for measurement
// and display of coarse-grained timings. It is not intended to replace
// profiling.
package stopwatch

import (
	"fmt"
	"strings"
	"time"
)

// Internal representation of a single measurement.
type lap struct {
	label    string
	duration time.Duration
}

// Format the lap using the given format string.
// XXX: Validate the format.
func (l *lap) Format(format string) string {
	return fmt.Sprintf(format, l.label, l.duration)
}

type Stopwatch struct {
	last time.Time
	laps []lap
}

func NewStopwatch() *Stopwatch {
	return &Stopwatch{time.Time{}, make([]lap, 0)}
}

func (s *Stopwatch) Start() {
	s.last = time.Now()
}

func (s *Stopwatch) Stop(label string) {
	l := lap{label, time.Since(s.last)}
	s.laps = append(s.laps, l)
}

// Helper for the common case of back-to-back intervals.
func (s *Stopwatch) Lap(label string) {
	s.Stop(label)
	s.Start()
}

// Format the measurements using format for each lap and sep as the
// separator between laps. Values will be present in the order in which
// they were recorded.
func (s *Stopwatch) Format(format string, sep string) string {
	tokens := make([]string, 0, len(s.laps))
	for _, lap := range s.laps {
		tokens = append(tokens, lap.Format(format))
	}
	return strings.Join(tokens, sep)
}

// Default formatting string.
func (s *Stopwatch) String() string {
	return s.Format("%s: %s", "\n")
}
