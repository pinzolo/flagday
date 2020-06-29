package flagday

import (
	"errors"
	"sort"
	"time"
)

var cache = make(map[int][]Holiday)

// InYear returns holiadys in given year.
func InYear(year int) []Holiday {
	if dates, ok := cache[year]; ok {
		return dates
	}
	defs := DefsInYear(year)
	dates := Holidays(defs, year)
	sort.Slice(dates, func(i, j int) bool {
		return dates[i].Time().Unix() < dates[j].Time().Unix()
	})
	cache[year] = dates
	return dates
}

// InMonth returns holiadys in given year and month.
func InMonth(year, month int) []Holiday {
	ydates := InYear(year)
	var mdates []Holiday
	for _, d := range ydates {
		if d.Month() == month {
			mdates = append(mdates, d)
		}
	}
	return mdates
}

// Holidays for given definitions in given year.
func Holidays(defs []Definition, year int) []Holiday {
	var holidays []Holiday
	for _, def := range defs {
		if def.Func() == nil {
			continue
		}
		holiday := def.Func()(def, year)

		// national holiday
		if len(holidays) > 0 {
			if !holiday.Time().Before(NationalHolidayStartDate) {
				if last := holidays[len(holidays)-1]; last.Kind() != SubstituteHoliday && last.Kind() != NationalHoliday {
					if prev := holiday.Time().AddDate(0, 0, -2); last.Time().Equal(prev) {
						ndate := holiday.Time().AddDate(0, 0, -1)
						holidays = append(holidays, newNationalHoliday(year, int(ndate.Month()), ndate.Day()))
					}
				}
			}
		}

		holidays = append(holidays, holiday)

		// Substitute holiday
		substDate, err := substituteHolidayOf(holiday)
		if err == nil {
			holidays = append(holidays, substDate)
		}
	}
	return holidays
}

// PublicHolidayOf returns public holiday information of given date units.
func PublicHolidayOf(year, month, day int) (Holiday, error) {
	holidays := InYear(year)
	for _, d := range holidays {
		if d.Month() == month && d.Day() == day {
			return d, nil
		}
	}
	return nil, errors.New("not public holiday")
}

// PublicHolidayTimeOf returns public holiday information of given time.
func PublicHolidayTimeOf(tm time.Time) (Holiday, error) {
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
	cache = make(map[int][]Holiday)
}

// FixedDateHoliday returns fixed date holiday.
func FixedDateHoliday(def Definition, year int) Holiday {
	return newPublicHoliday(def, year, def.Day())
}

// WeekNumHolidayFunc returns function for getting nth weekday holiday.
func WeekNumHolidayFunc(weekday time.Weekday) func(def Definition, year int) Holiday {
	return func(def Definition, year int) Holiday {
		fday := firstDateOf(year, def.Month())
		var days int
		if int(fday.Weekday()) <= int(weekday) {
			days = (def.WeekNum()-1)*7 - int(fday.Weekday()) + int(weekday)
		} else {
			days = def.WeekNum()*7 - int(fday.Weekday()) + int(weekday)
		}
		tm := fday.AddDate(0, 0, days)
		return newPublicHoliday(def, year, tm.Day())

	}
}

func happyMonday(def Definition, year int) Holiday {
	return WeekNumHolidayFunc(time.Monday)(def, year)
}

func imperialRelatedHoliday(def Definition, year int) Holiday {
	return newImperialRelatedHoliday(def, year, def.Day())
}

func firstDateOf(year int, month int) time.Time {
	return time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
}

func substituteHolidayOf(holiday Holiday) (Holiday, error) {
	if holiday.Time().Before(SubstituteHolidayStartDate) {
		return nil, errors.New("before enforcement of law")
	}
	if holiday.Time().Weekday() != time.Sunday {
		return nil, errors.New("not Sunday")
	}
	day := holiday.Day() + 1
	// Golden week
	if holiday.Month() == 5 && (holiday.Day() == 3 || holiday.Day() == 4) {
		day = 6
	}
	return newSubstituteHoliday(holiday.Year(), holiday.Month(), day, holiday), nil
}
