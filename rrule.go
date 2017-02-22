package rrule

import "time"

type RRule interface {
	All() []time.Time
	Between(after time.Time, before time.Time, include bool) []time.Time
	Before(dateBefore time.Time, include bool) []time.Time
	After(date time.Time, include bool) []time.Time
	Count() int
	FromString(string string) RRule
	ToString() string
}
