rrule.go
========

**Go library for working with recurrence rules for calendar dates.**

Inspired by the amazing [rrule.js](https://github.com/jkbrzt/rrule).

* * * * *

### Quick Start

#### Installation

The only requirement is the [Go Programming Language](https://golang.org/dl/), at least 1.8

```bash
$ go get github.com/JulienBreux/rrule
```

#### Overview

##### RRule

```go
// Create a rule
rule := rrule.RRule{
	freq:      rrule.WEEKLY,
	interval:  5,
	byweekday: [rrule.MO, rrule.FR],
	dtstart:   time.Date(2017, time.January, 1, 10, 30, 0, 0, time.UTC),
	until:     time.Date(2017, time.December, 31, 10, 30, 0, 0, time.UTC),
}

// Get all occurrence dates (Date instances)
rule.All()

// Get a slice of occurrence dates (Date instances)
rule.Between(
	time.Date(2017, time.August, 1, 10, 30, 0, 0, time.UTC),
	time.Date(2017, time.September, 1, 10, 30, 0, 0, time.UTC),
)

// Get an iCalendar RRULE string representation:
// The output can be used with RRule.FromString("...").
rule.ToString()
```

##### RRuleSet

```go
// Create a rule set
set := rrule.RRuleSet{}

// Add a rule to set
set.RRule(
	rrule.RRule{
		freq:    rrule.MONTHLY,
		count:   2,
		dtstart: time.Date(2017, time.January, 1, 10, 30, 0, 0, time.UTC),
	}
)

// Add an exclusion rule to set
set.ExRule(
	rrule.RRule{
		freq:    rrule.MONTHLY,
		count:   2,
		dtstart: time.Date(2017, time.February, 1, 10, 30, 0, 0, time.UTC),
	}
)

// Add a date to set
set.RDate(time.Date(2017, time.June, 1, 10, 30, 0, 0, time.UTC))

// Add an another date to set
set.RDate(time.Date(2017, time.June, 2, 10, 30, 0, 0, time.UTC))

// Add an exclusion date to set
set.ExDate(time.Date(2017, time.May, 1, 10, 30, 0, 0, time.UTC))

// Get all occurrence dates (Date instances)
set.All()

// Get a slice of occurrence dates (Date instances)
set.Between(
	time.Date(2017, time.February, 1, 10, 30, 0, 0, time.UTC),
	time.Date(2017, time.June, 2, 10, 30, 0, 0, time.UTC),
)

// Array of strings (e.g. ["RRULE:...", "RDATE:...", "EXRULE:...", "EXDATE:..."])
set.ValueOf()

// To string (array stringified) (e.g. "["RRULE:...", "RDATE:...", "EXRULE:...", "EXDATE:..."]")
set.ToString()
```

##### RRule parser

```go
// Parse a RRule string, to a RRule instance
rrule.StrToRule("RRULE:FREQ=MONTHLY;COUNT=5;DTSTART=20170201T023000Z")

// Parse a RRule string, to a RRuleSet instance
rrule.StrToRuleSet("RRULE:FREQ=MONTHLY;COUNT=5;DTSTART=20170201T023000Z")

// Parse a RRuleSet string, to a RRuleSet instance
rrule.StrToRuleSet(
	"RRULE:FREQ=MONTHLY;COUNT=5;DTSTART=20170201T023000Z\n"+
		"RDATE:20170701T023000Z,20170702T023000Z\n"+
		"EXRULE:FREQ=MONTHLY;COUNT=2;DTSTART=20170301T023000Z\n"+
		"EXDATE:20170601T023000Z")
```

#### Authors

* [Julien Breux](https://julienbreux.uk/) ([@JulienBreux](http://twitter.com/JulienBreux))
