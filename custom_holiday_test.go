package flagday

import (
	"testing"
	"time"
)

var custom DefType = 99
var secondSaturday HolidayKind = 99

type customDef struct {
	name    string
	month   int
	weekNum int
	fn      func(def Definition, year int) Date
}

func (def customDef) Type() DefType {
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

func (def customDef) Func() func(def Definition, year int) Date {
	return def.fn
}

func (def customDef) Begin() int {
	return 2000
}

func (def customDef) End() int {
	return 0
}

func TestCustomHoliday(t *testing.T) {
	defs := make([]Definition, 12)
	for i := 0; i <= 11; i++ {
		def := customDef{
			month:   i + 1,
			weekNum: 2,
			fn:      WeekNumHolidayFunc(time.Saturday),
		}
		defs[i] = def
	}
	year := 2017
	dates := PublicHolidays(defs, year)
	testdata := []expected{
		{1, 14, "Second Saturday"},
		{2, 11, "Second Saturday"},
		{3, 11, "Second Saturday"},
		{4, 8, "Second Saturday"},
		{5, 13, "Second Saturday"},
		{6, 10, "Second Saturday"},
		{7, 8, "Second Saturday"},
		{8, 12, "Second Saturday"},
		{9, 9, "Second Saturday"},
		{10, 14, "Second Saturday"},
		{11, 11, "Second Saturday"},
		{12, 9, "Second Saturday"},
	}
	check(t, year, dates, testdata)
}

func TestNoFuncDefinition(t *testing.T) {
	def := customDef{
		month:   1,
		weekNum: 2,
	}
	defs := []Definition{def}
	dates := PublicHolidays(defs, 2017)
	if len(dates) != 0 {
		t.Error("should return date when given no func definition")
	}
}
