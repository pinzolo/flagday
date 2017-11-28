package flagday

import (
	"testing"
)

type expected struct {
	month int
	day   int
	name  string
}

func BenchmarkInYear(b *testing.B) {
	begin := 1948
	end := 2017
	ClearCache()
	for i := 0; i < b.N; i++ {
		for j := begin; j <= end; j++ {
			InYear(j)
		}
	}
}

func BenchmarkInYearWithoutCache(b *testing.B) {
	begin := 1948
	end := 2017
	for i := 0; i < b.N; i++ {
		ClearCache()
		for j := begin; j <= end; j++ {
			InYear(j)
		}
	}
}

func check(t *testing.T, year int, dates []Date, testdata []expected) {
	for i, td := range testdata {
		date := dates[i]
		checkDate(t, year, date, td)
	}
}

func checkDate(t *testing.T, year int, date Date, td expected) {
	if date.Year() != year {
		t.Errorf("year is not match, expected %d but got %d", year, date.Year())
	}
	if date.Month() != td.month {
		t.Errorf("month is not match, expected %d but got %d", td.month, date.Month())
	}
	if date.Day() != td.day {
		t.Errorf("day is not match, expected %d/%d but got %d/%d", td.month, td.day, date.Month(), date.Day())
	}
	if date.Name() != td.name {
		t.Errorf("holiday name is not match, expected %s but got %s", td.name, date.Name())
	}
	if td.name == "振替休日" {
		if date.Kind() != SubstituteHoliday {
			t.Errorf("kind mismatch, expected %v but got %v", SubstituteHoliday, date.Kind())
		}
	} else if td.name == "国民の休日" {
		if date.Kind() != NationalHoliday {
			t.Errorf("kind mismatch, expected %v but got %v", NationalHoliday, date.Kind())
		}
	} else {
		if date.Kind() != PublicHoliday {
			t.Errorf("kind mismatch, expected %v but got %v", PublicHoliday, date.Kind())
		}
	}

}

func TestInYear(t *testing.T) {
	testdata := []expected{
		{1, 1, "元日"},
		{1, 2, "振替休日"},
		{1, 9, "成人の日"},
		{2, 11, "建国記念の日"},
		{3, 20, "春分の日"},
		{4, 29, "昭和の日"},
		{5, 3, "憲法記念日"},
		{5, 4, "みどりの日"},
		{5, 5, "こどもの日"},
		{7, 17, "海の日"},
		{8, 11, "山の日"},
		{9, 18, "敬老の日"},
		{9, 23, "秋分の日"},
		{10, 9, "体育の日"},
		{11, 03, "文化の日"},
		{11, 23, "勤労感謝の日"},
		{12, 23, "天皇誕生日"},
	}
	year := 2017
	dates := InYear(year)
	check(t, year, dates, testdata)
}

func TestInMonth(t *testing.T) {
	testdata := []expected{
		{1, 1, "元日"},
		{1, 2, "振替休日"},
		{1, 9, "成人の日"},
	}
	year := 2017
	dates := InMonth(year, 1)
	check(t, year, dates, testdata)

	dates = InMonth(year, 6)
	if len(dates) != 0 {
		t.Error("should return empty slice")
	}
}

func TestPublicHolidayOf(t *testing.T) {
	d, err := PublicHolidayOf(2017, 11, 23)
	if err != nil {
		t.Error(err)
	}
	checkDate(t, 2017, d, expected{11, 23, "勤労感謝の日"})
	d, err = PublicHolidayOf(2017, 11, 24)
	if err == nil {
		t.Errorf("2017/11/24 is not public holiday, PublicHolidayOf should return error")
	}
}

func TestIsPublicHoliday(t *testing.T) {
	if !IsPublicHoliday(2017, 11, 23) {
		t.Errorf("2017/11/23 is public holiday, but IsPublicHoliday returns false")
	}
	if IsPublicHoliday(2017, 11, 24) {
		t.Errorf("2017/11/24 is public holiday, but IsPublicHoliday returns true")
	}
}

func TestPublicHolidayTimeOf(t *testing.T) {
	d, err := PublicHolidayTimeOf(getTime(2017, 11, 23))
	if err != nil {
		t.Error(err)
	}
	checkDate(t, 2017, d, expected{11, 23, "勤労感謝の日"})
	d, err = PublicHolidayTimeOf(getTime(2017, 11, 24))
	if err == nil {
		t.Errorf("2017/11/24 is not public holiday, PublicHolidayOf should return error")
	}
}

func TestIsPublicHolidayTime(t *testing.T) {
	if !IsPublicHolidayTime(getTime(2017, 11, 23)) {
		t.Errorf("2017/11/23 is public holiday, but IsPublicHoliday returns false")
	}
	if IsPublicHolidayTime(getTime(2017, 11, 24)) {
		t.Errorf("2017/11/24 is public holiday, but IsPublicHoliday returns true")
	}
}

func TestClearCache(t *testing.T) {
	InYear(2017)
	if len(cache) == 0 {
		t.Error("cache length should be zero")
	}
	ClearCache()
	if len(cache) != 0 {
		t.Error("cache length should be zero after ClearCache called")
	}
}
