package rrule

import "time"

type RRuleSet interface {
	All() []time.Time
	Between(after time.Time, before time.Time, include bool) []time.Time
	Before(dateBefore time.Time, include bool) []time.Time
	After(date time.Time, include bool) []time.Time
	Count() int
	RRule(rrule RRule)
	ExRule(rrule RRule)
	RDate(date time.Time)
	ExDate(date time.Time)
	ValueOf() string
	ToString() string
}
