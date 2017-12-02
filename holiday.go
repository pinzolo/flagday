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
	// ImperialRelated means that holiday is Imperial related holiday.
	ImperialRelated
)

// Holiday holds public holiday information.
type Holiday struct {
	def   *Definition
	year  int
	month int
	day   int
	name  string
	kind  HolidayKind
	time  time.Time
}

// Def returns definition of date.
func (d Holiday) Def() *Definition {
	return d.def
}

// Year of holiday
func (d Holiday) Year() int {
	return d.year
}

// Month of holiday
func (d Holiday) Month() int {
	return d.month
}

// Day of holiday
func (d Holiday) Day() int {
	return d.day
}

// Name of holiday
func (d Holiday) Name() string {
	return d.name
}

// Kind of holiday
func (d Holiday) Kind() HolidayKind {
	return d.kind
}

// Time returns time.Time converted in JST.
func (d Holiday) Time() time.Time {
	return d.time
}

// NewDate returns new date with time.
func NewDate(def *Definition, year, month, day int, name string, kind HolidayKind) Holiday {
	return Holiday{
		def:   def,
		year:  year,
		month: month,
		day:   day,
		name:  name,
		kind:  kind,
		time:  getTime(year, month, day),
	}
}

func newPublicHoliday(def Definition, year, day int) Holiday {
	return NewDate(&def, year, def.Month(), day, def.Name(), PublicHoliday)
}

func newNationalHoliday(year, month, day int) Holiday {
	return NewDate(nil, year, month, day, "国民の休日", NationalHoliday)
}

func newSubstituteHoliday(year, month, day int) Holiday {
	return NewDate(nil, year, month, day, "振替休日", SubstituteHoliday)
}

func newImperialRelatedHoliday(def Definition, year, day int) Holiday {
	return NewDate(&def, year, def.Month(), day, def.Name(), ImperialRelated)
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
