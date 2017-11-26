package flagday

import (
	"errors"
	"sort"
	"time"
)

var cache = make(map[int][]Date)

// InYear returns holiadys in given year.
func InYear(year int) []Date {
	if dates, ok := cache[year]; ok {
		return dates
	}
	defs := DefsInYear(year)
	dates := PublicHolidays(defs, year)
	sort.Slice(dates, func(i, j int) bool {
		return dates[i].Time().Unix() < dates[j].Time().Unix()
	})
	cache[year] = dates
	return dates
}

// InMonth returns holiadys in given year and month.
func InMonth(year, month int) []Date {
	ydates := InYear(year)
	var mdates []Date
	for _, d := range ydates {
		if d.Month() == month {
			mdates = append(mdates, d)
		}
	}
	return mdates
}

// PublicHolidays for given definitions in given year.
func PublicHolidays(defs []Definition, year int) []Date {
	var dates []Date
	for _, def := range defs {
		if def.Func() == nil {
			continue
		}
		date := def.Func()(def, year)

		// natinal holiday
		if len(dates) > 0 {
			if !date.Time().Before(NationalHolidayStartDate) {
				prev := date.Time().AddDate(0, 0, -2)
				last := dates[len(dates)-1]
				if last.Time().Equal(prev) && last.Kind() == PublicHoliday {
					ndate := date.Time().AddDate(0, 0, -1)
					dates = append(dates, newNationalHoliday(year, int(ndate.Month()), ndate.Day()))
				}
			}
		}

		dates = append(dates, date)

		// Substitute holiday
		substDate, err := substituteHolidayOf(date)
		if err == nil {
			dates = append(dates, substDate)
		}
	}
	return dates
}

// PublicHolidayOf returns public holiday information of given date units.
func PublicHolidayOf(year, month, day int) (Date, error) {
	dates := InYear(year)
	for _, d := range dates {
		if d.Month() == month && d.Day() == day {
			return d, nil
		}
	}
	return Date{}, errors.New("not public holiday")
}

// PublicHolidayTimeOf returns public holiday information of given time.
func PublicHolidayTimeOf(tm time.Time) (Date, error) {
	return PublicHolidayOf(tm.Year(), int(tm.Month()), tm.Day())
}

// IsPublicHoliday returns true if given date units are public holiday.
func IsPublicHoliday(year, month, day int) bool {
	_, err := PublicHolidayOf(year, month, day)
	return err == nil
}

// IsPublicHolidayTime returns true if given time is public holiday.
func IsPublicHolidayTime(tm time.Time) bool {
	return IsPublicHoliday(tm.Year(), int(tm.Month()), tm.Day())
}

// ClearCache clear internal cache.
func ClearCache() {
	cache = make(map[int][]Date)
}

func fixedPublicHoliday(def Definition, year int) Date {
	return newPublicHoliday(def, year, def.Day())
}

func happyMonday(def Definition, year int) Date {
	fday := firstDate(year, def.Month())
	var days int
	if int(fday.Weekday()) <= int(time.Monday) {
		days = (def.WeekNum()-1)*7 - int(fday.Weekday()) + int(time.Monday)
	} else {
		days = def.WeekNum()*7 - int(fday.Weekday()) + int(time.Monday)
	}
	tm := fday.AddDate(0, 0, days)
	return newPublicHoliday(def, year, tm.Day())
}

func imperialRelatedHoliday(def Definition, year int) Date {
	return newImperialRelatedHoliday(def, year, def.Day())
}

func firstDate(year int, month int) time.Time {
	return time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
}

func substituteHolidayOf(date Date) (Date, error) {
	if date.Time().Before(SubstituteHolidayStartDate) {
		return Date{}, errors.New("before enforcement of law")
	}
	if date.Time().Weekday() != time.Sunday {
		return Date{}, errors.New("not Sunday")
	}
	day := date.Day() + 1
	// Golden week
	if date.Month() == 5 && (date.Day() == 3 || date.Day() == 4) {
		day = 6
	}
	return newSubstituteHoliday(date.Year(), date.Month(), day), nil
}
