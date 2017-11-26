package flagday

import (
	"math"
)

// DefType is type of definition.
type DefType int

const (
	// FixedDay means that public holiday has fixed day.
	FixedDay DefType = iota
	// HappyMonday means that public holiday is happy monday.
	HappyMonday
	// EquinoxDay means that public holiday is equinox day.
	EquinoxDay
	// ImperialRelatedHoliday means that public holiday is imperial related holiday (that year's limit).
	ImperialRelatedHoliday
)

// Definition is definition of public holiday in Japan.
type Definition interface {
	Type() DefType
	Name() string
	Month() int
	Day() int
	WeekNum() int
	Func() func(Definition, int) Date
	Begin() int
	End() int
}

type definition struct {
	defType DefType
	name    string
	month   int
	day     int
	weekNum int
	fn      func(def Definition, year int) Date
	begin   int
	end     int
}

func (def definition) Type() DefType {
	return def.defType
}

func (def definition) Name() string {
	return def.name
}

func (def definition) Month() int {
	return def.month
}

func (def definition) Day() int {
	return def.day
}

func (def definition) WeekNum() int {
	return def.weekNum
}

func (def definition) Func() func(def Definition, year int) Date {
	return def.fn
}

func (def definition) Begin() int {
	return def.begin
}

func (def definition) End() int {
	return def.end
}

// DefsInYear returns definitions that are available in given year.
func DefsInYear(year int) []Definition {
	var defs []Definition
	for _, def := range AllDefs {
		if def.Begin() <= year && (year <= def.End() || def.End() == 0) {
			defs = append(defs, def)
		}
	}
	return defs
}

// AllDefs are all definitions of public holidays.
var AllDefs = []Definition{
	definition{
		defType: FixedDay,
		name:    "元日",
		month:   1,
		day:     1,
		fn:      FixedDateHoliday,
		begin:   1949,
	},
	definition{
		defType: FixedDay,
		name:    "成人の日",
		month:   1,
		day:     15,
		fn:      FixedDateHoliday,
		begin:   1949,
		end:     1999,
	},
	definition{
		defType: HappyMonday,
		name:    "成人の日",
		month:   1,
		weekNum: 2,
		fn:      happyMonday,
		begin:   2000,
	},
	definition{
		defType: FixedDay,
		name:    "建国記念の日",
		month:   2,
		day:     11,
		fn:      FixedDateHoliday,
		begin:   1967,
	},
	definition{
		defType: ImperialRelatedHoliday,
		name:    "昭和天皇の大喪の礼",
		month:   2,
		day:     24,
		fn:      imperialRelatedHoliday,
		begin:   1989,
		end:     1989,
	},
	definition{
		defType: EquinoxDay,
		name:    "春分の日",
		month:   3,
		fn: func(def Definition, year int) Date {
			day := 0.24242*float64(year) - math.Floor(float64(year)/4.0) + 35.84
			return newPublicHoliday(def, year, int(day))
		},
		begin: 1949,
	},
	definition{
		defType: ImperialRelatedHoliday,
		name:    "皇太子明仁親王の結婚の儀",
		month:   4,
		day:     10,
		fn:      imperialRelatedHoliday,
		begin:   1959,
		end:     1959,
	},
	definition{
		defType: FixedDay,
		name:    "天皇誕生日",
		month:   4,
		day:     29,
		fn:      FixedDateHoliday,
		begin:   1949,
		end:     1988,
	},
	definition{
		defType: FixedDay,
		name:    "みどりの日",
		month:   4,
		day:     29,
		fn:      FixedDateHoliday,
		begin:   1989,
		end:     2006,
	},
	definition{
		defType: FixedDay,
		name:    "昭和の日",
		month:   4,
		day:     29,
		fn:      FixedDateHoliday,
		begin:   2007,
	},
	definition{
		defType: FixedDay,
		name:    "憲法記念日",
		month:   5,
		day:     3,
		fn:      FixedDateHoliday,
		begin:   1967,
	},
	definition{
		defType: FixedDay,
		name:    "みどりの日",
		month:   5,
		day:     4,
		fn:      FixedDateHoliday,
		begin:   2007,
	},
	definition{
		defType: FixedDay,
		name:    "こどもの日",
		month:   5,
		day:     5,
		fn:      FixedDateHoliday,
		begin:   1949,
	},
	definition{
		defType: ImperialRelatedHoliday,
		name:    "皇太子徳仁親王の結婚の儀",
		month:   6,
		day:     9,
		fn:      imperialRelatedHoliday,
		begin:   1993,
		end:     1993,
	},
	definition{
		defType: FixedDay,
		name:    "海の日",
		month:   7,
		day:     20,
		fn:      FixedDateHoliday,
		begin:   1996,
		end:     2002,
	},
	definition{
		defType: HappyMonday,
		name:    "海の日",
		month:   7,
		weekNum: 3,
		fn:      happyMonday,
		begin:   2003,
	},
	definition{
		defType: FixedDay,
		name:    "山の日",
		month:   8,
		day:     11,
		fn:      FixedDateHoliday,
		begin:   2016,
	},
	definition{
		defType: FixedDay,
		name:    "敬老の日",
		month:   9,
		day:     15,
		fn:      FixedDateHoliday,
		begin:   1966,
		end:     2002,
	},
	definition{
		defType: HappyMonday,
		name:    "敬老の日",
		month:   9,
		weekNum: 3,
		fn:      happyMonday,
		begin:   2003,
	},
	definition{
		defType: EquinoxDay,
		name:    "秋分の日",
		month:   9,
		begin:   1948,
		fn: func(def Definition, year int) Date {
			day := 0.24204*float64(year) - math.Floor(float64(year)/4.0) + 39.01
			return newPublicHoliday(def, year, int(day))
		},
	},
	definition{
		defType: FixedDay,
		name:    "体育の日",
		month:   10,
		day:     10,
		fn:      FixedDateHoliday,
		begin:   1966,
		end:     1999,
	},
	definition{
		defType: HappyMonday,
		name:    "体育の日",
		month:   10,
		weekNum: 2,
		fn:      happyMonday,
		begin:   2000,
	},
	definition{
		defType: FixedDay,
		name:    "文化の日",
		month:   11,
		day:     3,
		fn:      FixedDateHoliday,
		begin:   1948,
	},
	definition{
		defType: ImperialRelatedHoliday,
		name:    "即位礼正殿の儀",
		month:   11,
		day:     12,
		fn:      imperialRelatedHoliday,
		begin:   1990,
		end:     1990,
	},
	definition{
		defType: FixedDay,
		name:    "勤労感謝の日",
		month:   11,
		day:     23,
		fn:      FixedDateHoliday,
		begin:   1948,
	},
	definition{
		defType: FixedDay,
		name:    "天皇誕生日",
		month:   12,
		day:     23,
		fn:      FixedDateHoliday,
		begin:   1989,
	},
}
