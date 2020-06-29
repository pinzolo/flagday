package flagday

import (
	"testing"
)

type expected struct {
	month int
	day   int
	name  string
	kind  HolidayKind
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

func testHolidays(t *testing.T, year int, dates []Holiday, testdata []expected) {
	t.Helper()
	if len(testdata) != len(dates) {
		t.Errorf("holiday count is not match, expected %d but got %d", len(testdata), len(dates))
		return
	}
	for i, td := range testdata {
		date := dates[i]
		testDate(t, year, date, td)
	}
}

func testDate(t *testing.T, year int, date Holiday, td expected) {
	t.Helper()
	s := date.Time().Format("2006/1/2")
	if date.Year() != year {
		t.Errorf("year of %s is not match, expected %d but got %d", s, year, date.Year())
	}
	if date.Month() != td.month {
		t.Errorf("month of %s is not match, expected %d but got %d", s, td.month, date.Month())
	}
	if date.Day() != td.day {
		t.Errorf("day is of %s not match, expected %d/%d but got %d/%d", s, td.month, td.day, date.Month(), date.Day())
	}
	if date.Name() != td.name {
		t.Errorf("holiday name of %s is not match, expected %s but got %s", s, td.name, date.Name())
	}
	if date.Kind() != td.kind {
		t.Errorf("kind of %s is not match, expected %v but got %v", s, td.kind, date.Kind())
	}
}

func TestInYear(t *testing.T) {
	testdata := []expected{
		{1, 1, "元日", PublicHoliday},
		{1, 2, "振替休日", SubstituteHoliday},
		{1, 9, "成人の日", PublicHoliday},
		{2, 11, "建国記念の日", PublicHoliday},
		{3, 20, "春分の日", PublicHoliday},
		{4, 29, "昭和の日", PublicHoliday},
		{5, 3, "憲法記念日", PublicHoliday},
		{5, 4, "みどりの日", PublicHoliday},
		{5, 5, "こどもの日", PublicHoliday},
		{7, 17, "海の日", PublicHoliday},
		{8, 11, "山の日", PublicHoliday},
		{9, 18, "敬老の日", PublicHoliday},
		{9, 23, "秋分の日", PublicHoliday},
		{10, 9, "体育の日", PublicHoliday},
		{11, 3, "文化の日", PublicHoliday},
		{11, 23, "勤労感謝の日", PublicHoliday},
		{12, 23, "天皇誕生日", PublicHoliday},
	}
	year := 2017
	dates := InYear(year)
	testHolidays(t, year, dates, testdata)
}

func TestIn2018(t *testing.T) {
	testdata := []expected{
		{1, 1, "元日", PublicHoliday},
		{1, 8, "成人の日", PublicHoliday},
		{2, 11, "建国記念の日", PublicHoliday},
		{2, 12, "振替休日", SubstituteHoliday},
		{3, 21, "春分の日", PublicHoliday},
		{4, 29, "昭和の日", PublicHoliday},
		{4, 30, "振替休日", SubstituteHoliday},
		{5, 3, "憲法記念日", PublicHoliday},
		{5, 4, "みどりの日", PublicHoliday},
		{5, 5, "こどもの日", PublicHoliday},
		{7, 16, "海の日", PublicHoliday},
		{8, 11, "山の日", PublicHoliday},
		{9, 17, "敬老の日", PublicHoliday},
		{9, 23, "秋分の日", PublicHoliday},
		{9, 24, "振替休日", SubstituteHoliday},
		{10, 8, "体育の日", PublicHoliday},
		{11, 3, "文化の日", PublicHoliday},
		{11, 23, "勤労感謝の日", PublicHoliday},
		{12, 23, "天皇誕生日", PublicHoliday},
		{12, 24, "振替休日", SubstituteHoliday},
	}
	year := 2018
	dates := InYear(year)
	testHolidays(t, year, dates, testdata)
}

func TestIn2019(t *testing.T) {
	testdata := []expected{
		{1, 1, "元日", PublicHoliday},
		{1, 14, "成人の日", PublicHoliday},
		{2, 11, "建国記念の日", PublicHoliday},
		{3, 21, "春分の日", PublicHoliday},
		{4, 29, "昭和の日", PublicHoliday},
		{4, 30, "国民の休日", NationalHoliday},
		{5, 1, "即位の礼", ImperialRelated},
		{5, 2, "国民の休日", NationalHoliday},
		{5, 3, "憲法記念日", PublicHoliday},
		{5, 4, "みどりの日", PublicHoliday},
		{5, 5, "こどもの日", PublicHoliday},
		{5, 6, "振替休日", SubstituteHoliday},
		{7, 15, "海の日", PublicHoliday},
		{8, 11, "山の日", PublicHoliday},
		{8, 12, "振替休日", SubstituteHoliday},
		{9, 16, "敬老の日", PublicHoliday},
		{9, 23, "秋分の日", PublicHoliday},
		{10, 14, "体育の日", PublicHoliday},
		{10, 22, "即位礼正殿の儀", ImperialRelated},
		{11, 3, "文化の日", PublicHoliday},
		{11, 4, "振替休日", SubstituteHoliday},
		{11, 23, "勤労感謝の日", PublicHoliday},
	}
	year := 2019
	dates := InYear(year)
	testHolidays(t, year, dates, testdata)
}

func TestIn2020(t *testing.T) {
	testdata := []expected{
		{1, 1, "元日", PublicHoliday},
		{1, 13, "成人の日", PublicHoliday},
		{2, 11, "建国記念の日", PublicHoliday},
		{2, 23, "天皇誕生日", PublicHoliday},
		{2, 24, "振替休日", SubstituteHoliday},
		{3, 20, "春分の日", PublicHoliday},
		{4, 29, "昭和の日", PublicHoliday},
		{5, 3, "憲法記念日", PublicHoliday},
		{5, 4, "みどりの日", PublicHoliday},
		{5, 5, "こどもの日", PublicHoliday},
		{5, 6, "振替休日", SubstituteHoliday},
		{7, 23, "海の日", PublicHoliday},
		{7, 24, "スポーツの日", PublicHoliday},
		{8, 10, "山の日", PublicHoliday},
		{9, 21, "敬老の日", PublicHoliday},
		{9, 22, "秋分の日", PublicHoliday},
		{11, 3, "文化の日", PublicHoliday},
		{11, 23, "勤労感謝の日", PublicHoliday},
	}
	year := 2020
	dates := InYear(year)
	testHolidays(t, year, dates, testdata)
}

func TestIn2021(t *testing.T) {
	testdata := []expected{
		{1, 1, "元日", PublicHoliday},
		{1, 11, "成人の日", PublicHoliday},
		{2, 11, "建国記念の日", PublicHoliday},
		{2, 23, "天皇誕生日", PublicHoliday},
		{3, 20, "春分の日", PublicHoliday},
		{4, 29, "昭和の日", PublicHoliday},
		{5, 3, "憲法記念日", PublicHoliday},
		{5, 4, "みどりの日", PublicHoliday},
		{5, 5, "こどもの日", PublicHoliday},
		{7, 19, "海の日", PublicHoliday},
		{8, 11, "山の日", PublicHoliday},
		{9, 20, "敬老の日", PublicHoliday},
		{9, 23, "秋分の日", PublicHoliday},
		{10, 11, "スポーツの日", PublicHoliday},
		{11, 3, "文化の日", PublicHoliday},
		{11, 23, "勤労感謝の日", PublicHoliday},
	}
	year := 2021
	dates := InYear(year)
	testHolidays(t, year, dates, testdata)
}

func TestInMonth(t *testing.T) {
	testdata := []expected{
		{1, 1, "元日", PublicHoliday},
		{1, 2, "振替休日", SubstituteHoliday},
		{1, 9, "成人の日", PublicHoliday},
	}
	year := 2017
	dates := InMonth(year, 1)
	testHolidays(t, year, dates, testdata)

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
	testDate(t, 2017, d, expected{11, 23, "勤労感謝の日", PublicHoliday})
	_, err = PublicHolidayOf(2017, 11, 24)
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
	d, err := PublicHolidayTimeOf(timeFrom(2017, 11, 23))
	if err != nil {
		t.Error(err)
	}
	testDate(t, 2017, d, expected{11, 23, "勤労感謝の日", PublicHoliday})
	_, err = PublicHolidayTimeOf(timeFrom(2017, 11, 24))
	if err == nil {
		t.Errorf("2017/11/24 is not public holiday, PublicHolidayOf should return error")
	}
}

func TestIsPublicHolidayTime(t *testing.T) {
	if !IsPublicHolidayTime(timeFrom(2017, 11, 23)) {
		t.Errorf("2017/11/23 is public holiday, but IsPublicHoliday returns false")
	}
	if IsPublicHolidayTime(timeFrom(2017, 11, 24)) {
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

func logDates(t *testing.T, dates []Holiday) {
	for _, d := range dates {
		t.Logf("%s %s", d.Time().Format("2006/01/02"), d.Name())
	}
}
