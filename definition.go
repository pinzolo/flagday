package flagday

import "math"

// DefType is type of definition.
type DefType int

const (
	// FixedDay means that public holiday has fixed day.
	FixedDay DefType = iota
	// HappyMonday means that public holiday is happy monday.
	HappyMonday
	// EquinoxDay means that public holiday is equinox day.
	EquinoxDay
)

// Definition is definition of public holiday in Japan.
type Definition struct {
	Type    DefType
	Name    string
	Month   int
	Day     int
	WeekNum int
	Func    func(Definition, int) Date
	Begin   int
	End     int
}

// DefsInYear returns definitions that are available in given year.
func DefsInYear(year int) []Definition {
	var defs []Definition
	for _, def := range AllDefs {
		if def.Begin <= year && (year <= def.End || def.End == 0) {
			defs = append(defs, def)
		}
	}
	return defs
}

// AllDefs are all definitions of public holidays.
var AllDefs = []Definition{
	Definition{
		Type:  FixedDay,
		Name:  "元日",
		Month: 1,
		Day:   1,
		Begin: 1949,
	},
	Definition{
		Type:  FixedDay,
		Name:  "成人の日",
		Month: 1,
		Day:   15,
		Begin: 1949,
		End:   1999,
	},
	Definition{
		Type:    HappyMonday,
		Name:    "成人の日",
		Month:   1,
		WeekNum: 2,
		Begin:   2000,
	},
	Definition{
		Type:  FixedDay,
		Name:  "建国記念の日",
		Month: 2,
		Day:   11,
		Begin: 1967,
	},
	Definition{
		Type:  EquinoxDay,
		Name:  "春分の日",
		Month: 3,
		Begin: 1949,
		Func: func(def Definition, year int) Date {
			day := 0.24242*float64(year) - math.Floor(float64(year)/4.0) + 35.84
			return newPublicHoliday(def, year, def.Month, int(day))
		},
	},
	Definition{
		Type:  FixedDay,
		Name:  "天皇誕生日",
		Month: 4,
		Day:   29,
		Begin: 1949,
		End:   1988,
	},
	Definition{
		Type:  FixedDay,
		Name:  "みどりの日",
		Month: 4,
		Day:   29,
		Begin: 1989,
		End:   2006,
	},
	Definition{
		Type:  FixedDay,
		Name:  "昭和の日",
		Month: 4,
		Day:   29,
		Begin: 2007,
	},
	Definition{
		Type:  FixedDay,
		Name:  "憲法記念日",
		Month: 5,
		Day:   3,
		Begin: 1967,
	},
	Definition{
		Type:  FixedDay,
		Name:  "みどりの日",
		Month: 5,
		Day:   4,
		Begin: 2007,
	},
	Definition{
		Type:  FixedDay,
		Name:  "こどもの日",
		Month: 5,
		Day:   5,
		Begin: 1949,
	},
	Definition{
		Type:  FixedDay,
		Name:  "海の日",
		Month: 7,
		Day:   20,
		Begin: 1996,
		End:   2002,
	},
	Definition{
		Type:    HappyMonday,
		Name:    "海の日",
		Month:   7,
		WeekNum: 3,
		Begin:   2003,
	},
	Definition{
		Type:  FixedDay,
		Name:  "山の日",
		Month: 8,
		Day:   11,
		Begin: 2016,
	},
	Definition{
		Type:  FixedDay,
		Name:  "敬老の日",
		Month: 9,
		Day:   15,
		Begin: 1966,
		End:   2002,
	},
	Definition{
		Type:    HappyMonday,
		Name:    "敬老の日",
		Month:   9,
		WeekNum: 3,
		Begin:   2003,
	},
	Definition{
		Type:  EquinoxDay,
		Name:  "秋分の日",
		Month: 9,
		Begin: 1948,
		Func: func(def Definition, year int) Date {
			day := 0.24204*float64(year) - math.Floor(float64(year)/4.0) + 39.01
			return newPublicHoliday(def, year, def.Month, int(day))
		},
	},
	Definition{
		Type:  FixedDay,
		Name:  "体育の日",
		Month: 10,
		Day:   10,
		Begin: 1966,
		End:   1999,
	},
	Definition{
		Type:    HappyMonday,
		Name:    "体育の日",
		Month:   10,
		WeekNum: 2,
		Begin:   2000,
	},
	Definition{
		Type:  FixedDay,
		Name:  "文化の日",
		Month: 11,
		Day:   3,
		Begin: 1948,
	},
	Definition{
		Type:  FixedDay,
		Name:  "勤労感謝の日",
		Month: 11,
		Day:   23,
		Begin: 1948,
	},
	Definition{
		Type:  FixedDay,
		Name:  "天皇誕生日",
		Month: 12,
		Day:   23,
		Begin: 1989,
	},
}
