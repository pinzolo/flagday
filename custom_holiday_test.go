package flagday_test

import (
	"testing"
	"time"

	"github.com/pinzolo/flagday"
)

var custom flagday.DefType = 99
var secondSaturday flagday.HolidayKind = 99

type customDef struct {
	name    string
	month   int
	weekNum int
	fn      func(def flagday.Definition, year int) flagday.Holiday
}

func (def customDef) Type() flagday.DefType {
	return custom
}

func (def customDef) Name() string {
	return "Second Saturday"
}

func (def customDef) Month() int {
	return def.month
}

func (def customDef) Day() int {
	return 0
}

func (def customDef) WeekNum() int {
	return def.weekNum
}

func (def customDef) Func() func(def flagday.Definition, year int) flagday.Holiday {
	return def.fn
}

func (def customDef) Begin() int {
	return 2000
}

func (def customDef) End() int {
	return 0
}

func TestCustomHoliday(t *testing.T) {
	defs := make([]flagday.Definition, 12)
	for i := 0; i <= 11; i++ {
		def := customDef{
			month:   i + 1,
			weekNum: 2,
			fn:      flagday.WeekNumHolidayFunc(time.Saturday),
		}
		defs[i] = def
	}
	year := 2017
	dates := flagday.Holidays(defs, year)
	testdata := []expected{
		{1, 14, "Second Saturday", flagday.PublicHoliday},
		{2, 11, "Second Saturday", flagday.PublicHoliday},
		{3, 11, "Second Saturday", flagday.PublicHoliday},
		{4, 8, "Second Saturday", flagday.PublicHoliday},
		{5, 13, "Second Saturday", flagday.PublicHoliday},
		{6, 10, "Second Saturday", flagday.PublicHoliday},
		{7, 8, "Second Saturday", flagday.PublicHoliday},
		{8, 12, "Second Saturday", flagday.PublicHoliday},
		{9, 9, "Second Saturday", flagday.PublicHoliday},
		{10, 14, "Second Saturday", flagday.PublicHoliday},
		{11, 11, "Second Saturday", flagday.PublicHoliday},
		{12, 9, "Second Saturday", flagday.PublicHoliday},
	}
	testHolidays(t, year, dates, testdata)
}

func TestNoFuncDefinition(t *testing.T) {
	def := customDef{
		month:   1,
		weekNum: 2,
	}
	defs := []flagday.Definition{def}
	dates := flagday.Holidays(defs, 2017)
	if len(dates) != 0 {
		t.Error("should return date when given no func definition")
	}
}
