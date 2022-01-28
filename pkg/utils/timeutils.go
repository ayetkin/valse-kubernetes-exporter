package utils

import (
	"fmt"
	"time"
)

// Age calculates kubernetes resources age while using CreationTimestamp field
func Age(creationTime time.Time) string {
	return formatSince(time.Since(creationTime).Round(time.Second))
}

func formatSince(t time.Duration) string {
	const (
		DeciSecond = 100 * time.Millisecond
		Day        = 24 * time.Hour
	)
	ts := t
	sign := time.Duration(1)
	if ts < 0 {
		sign = -1
		ts = -ts
	}
	ts += +DeciSecond / 2
	d := sign * (ts / Day)
	ts = ts % Day
	h := ts / time.Hour
	ts = ts % time.Hour
	m := ts / time.Minute
	ts = ts % time.Minute
	s := ts / time.Second
	ts = ts % time.Second

	return formatTime(d, h, m, s)
}

func formatTime(d, h, m, s time.Duration) string {
	if d > 2 {
		return fmt.Sprintf("%dd", d)
	} else if d > 0 {
		return fmt.Sprintf("%dd%dh", d, h)
	}

	if h > 2 {
		return fmt.Sprintf("%dh", h)
	} else if h > 0 {
		return fmt.Sprintf("%dh%dm", h, m)
	}

	if m > 2 {
		return fmt.Sprintf("%dm", m)
	} else if m > 0 {
		return fmt.Sprintf("%dm%ds", m, s)
	}

	return fmt.Sprintf("%ds", s)
}
