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
type Holiday interface {
	// Def returns definition of holiday.
	Def() *Definition
	// Year of holiday.
	Year() int
	// Month of holiday.
	Month() int
	// Day of holiday.
	Day() int
	// Name of holiday.
	Name() string
	// Kind of holiday.
	Kind() HolidayKind
	// Time is time.Time instance of holiday (JST)
	Time() time.Time
}

type holiday struct {
	def   *Definition
	year  int
	month int
	day   int
	name  string
	kind  HolidayKind
	time  time.Time
}

func (d holiday) Def() *Definition {
	return d.def
}

func (d holiday) Year() int {
	return d.year
}

func (d holiday) Month() int {
	return d.month
}

func (d holiday) Day() int {
	return d.day
}

func (d holiday) Name() string {
	return d.name
}

func (d holiday) Kind() HolidayKind {
	return d.kind
}

func (d holiday) Time() time.Time {
	return d.time
}

// NewHoliday returns new date with time.
func NewHoliday(def *Definition, year, month, day int, name string, kind HolidayKind) Holiday {
	return holiday{
		def:   def,
		year:  year,
		month: month,
		day:   day,
		name:  name,
		kind:  kind,
		time:  timeFrom(year, month, day),
	}
}

func newPublicHoliday(def Definition, year, day int) Holiday {
	return NewHoliday(&def, year, def.Month(), day, def.Name(), PublicHoliday)
}

func newNationalHoliday(year, month, day int) Holiday {
	return NewHoliday(nil, year, month, day, "国民の休日", NationalHoliday)
}

func newSubstituteHoliday(year, month, day int) Holiday {
	return NewHoliday(nil, year, month, day, "振替休日", SubstituteHoliday)
}

func newImperialRelatedHoliday(def Definition, year, day int) Holiday {
	return NewHoliday(&def, year, def.Month(), day, def.Name(), ImperialRelated)
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

func timeFrom(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, jst())
}
