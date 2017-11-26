package flagday

import (
	"time"
)

var jstLocation *time.Location

// HolidayKind is kind of holiday.
type HolidayKind int

const (
	// PublicHoliday means that holiday is public holiday.
	PublicHoliday HolidayKind = iota
	// NationalHoliday means that holiday is national holiday.
	NationalHoliday
	// SubstituteHoliday means that holiday is substitute holiday.
	SubstituteHoliday
)

// Date public holiday information.
type Date struct {
	year  int
	month int
	day   int
	name  string
	kind  HolidayKind
	time  time.Time
}

// Year of holiday
func (d Date) Year() int {
	return d.year
}

// Month of holiday
func (d Date) Month() int {
	return d.month
}

// Day of holiday
func (d Date) Day() int {
	return d.day
}

// Name of holiday
func (d Date) Name() string {
	return d.name
}

// Kind of holiday
func (d Date) Kind() HolidayKind {
	return d.kind
}

// Time returns time.Time converted in JST.
func (d Date) Time() time.Time {
	return d.time
}

// NewDate returns new date with time.
func NewDate(year, month, day int, name string, kind HolidayKind) Date {
	return Date{
		year:  year,
		month: month,
		day:   day,
		name:  name,
		kind:  kind,
		time:  getTime(year, month, day),
	}
}

func newPublicHoliday(def Definition, year, day int) Date {
	return NewDate(year, def.Month, day, def.Name, PublicHoliday)
}

func newNationalHoliday(year, month, day int) Date {
	return NewDate(year, month, day, "国民の休日", NationalHoliday)
}

func newSubstituteHoliday(year, month, day int) Date {
	return NewDate(year, month, day, "振替休日", SubstituteHoliday)
}

func jst() *time.Location {
	if jstLocation == nil {
		loc, err := time.LoadLocation("Asia/Tokyo")
		if err != nil {
			loc = time.FixedZone("Asia/Tokyo", 9*60*60)
		}
		jstLocation = loc
	}
	return jstLocation
}

func getTime(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, jst())
}
