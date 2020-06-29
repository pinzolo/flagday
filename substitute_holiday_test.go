package flagday_test

import (
	"testing"

	"github.com/pinzolo/flagday"
)

type expectedSubstitute struct {
	holiday  expected
	original expected
}

func TestSubstituteHolidayBefore(t *testing.T) {
	testdata := []expected{
		{1, 1, "元日", flagday.PublicHoliday},
		{1, 15, "成人の日", flagday.PublicHoliday},
		{2, 11, "建国記念の日", flagday.PublicHoliday},
		{3, 20, "春分の日", flagday.PublicHoliday},
		{4, 29, "天皇誕生日", flagday.PublicHoliday},
		{5, 3, "憲法記念日", flagday.PublicHoliday},
		{5, 5, "こどもの日", flagday.PublicHoliday}, // Sunday
		{9, 15, "敬老の日", flagday.PublicHoliday},
		{9, 23, "秋分の日", flagday.PublicHoliday},
		{10, 10, "体育の日", flagday.PublicHoliday},
		{11, 03, "文化の日", flagday.PublicHoliday},
		{11, 23, "勤労感謝の日", flagday.PublicHoliday},
	}
	year := 1968
	dates := flagday.InYear(year)
	testHolidays(t, year, dates, testdata)
}

func TestSubstituteHolidayGoldenWeekSunday3(t *testing.T) {
	testdata := []expected{
		{1, 1, "元日", flagday.PublicHoliday},
		{1, 12, "成人の日", flagday.PublicHoliday},
		{2, 11, "建国記念の日", flagday.PublicHoliday},
		{3, 20, "春分の日", flagday.PublicHoliday},
		{4, 29, "昭和の日", flagday.PublicHoliday},
		{5, 3, "憲法記念日", flagday.PublicHoliday}, // Sunday
		{5, 4, "みどりの日", flagday.PublicHoliday},
		{5, 5, "こどもの日", flagday.PublicHoliday},
		{5, 6, "振替休日", flagday.SubstituteHoliday},
		{7, 20, "海の日", flagday.PublicHoliday},
		{9, 21, "敬老の日", flagday.PublicHoliday},
		{9, 22, "国民の休日", flagday.NationalHoliday},
		{9, 23, "秋分の日", flagday.PublicHoliday},
		{10, 12, "体育の日", flagday.PublicHoliday},
		{11, 3, "文化の日", flagday.PublicHoliday},
		{11, 23, "勤労感謝の日", flagday.PublicHoliday},
		{12, 23, "天皇誕生日", flagday.PublicHoliday},
	}
	year := 2009
	dates := flagday.InYear(year)
	testHolidays(t, year, dates, testdata)
	substituteTestdata := []expectedSubstitute{
		{
			original: testdata[5], // 5/3
			holiday:  testdata[8], // 5/6
		},
	}
	testSubstitute(t, year, dates, substituteTestdata)
}

func TestSubstituteHolidayGoldenWeekSunday4(t *testing.T) {
	testdata := []expected{
		{1, 1, "元日", flagday.PublicHoliday},
		{1, 14, "成人の日", flagday.PublicHoliday},
		{2, 11, "建国記念の日", flagday.PublicHoliday},
		{3, 20, "春分の日", flagday.PublicHoliday},
		{4, 29, "昭和の日", flagday.PublicHoliday},
		{5, 3, "憲法記念日", flagday.PublicHoliday},
		{5, 4, "みどりの日", flagday.PublicHoliday}, // Sunday
		{5, 5, "こどもの日", flagday.PublicHoliday},
		{5, 6, "振替休日", flagday.SubstituteHoliday},
		{7, 21, "海の日", flagday.PublicHoliday},
		{9, 15, "敬老の日", flagday.PublicHoliday},
		{9, 23, "秋分の日", flagday.PublicHoliday},
		{10, 13, "体育の日", flagday.PublicHoliday},
		{11, 3, "文化の日", flagday.PublicHoliday},
		{11, 23, "勤労感謝の日", flagday.PublicHoliday}, // Sunday
		{11, 24, "振替休日", flagday.SubstituteHoliday},
		{12, 23, "天皇誕生日", flagday.PublicHoliday},
	}
	year := 2008
	dates := flagday.InYear(year)
	testHolidays(t, year, dates, testdata)
	substituteTestdata := []expectedSubstitute{
		{
			original: testdata[6], // 5/4
			holiday:  testdata[8], // 5/6
		},
		{
			original: testdata[14], // 11/23
			holiday:  testdata[15], // 11/24
		},
	}
	testSubstitute(t, year, dates, substituteTestdata)
}

func TestSubstituteHolidayGoldenWeekSunday5(t *testing.T) {
	testdata := []expected{
		{1, 1, "元日", flagday.PublicHoliday},
		{1, 14, "成人の日", flagday.PublicHoliday},
		{2, 11, "建国記念の日", flagday.PublicHoliday},
		{3, 20, "春分の日", flagday.PublicHoliday},
		{4, 29, "昭和の日", flagday.PublicHoliday},
		{5, 3, "憲法記念日", flagday.PublicHoliday},
		{5, 4, "みどりの日", flagday.PublicHoliday},
		{5, 5, "こどもの日", flagday.PublicHoliday}, // Sunday
		{5, 6, "振替休日", flagday.SubstituteHoliday},
		{7, 15, "海の日", flagday.PublicHoliday},
		{9, 16, "敬老の日", flagday.PublicHoliday},
		{9, 23, "秋分の日", flagday.PublicHoliday},
		{10, 14, "体育の日", flagday.PublicHoliday},
		{11, 3, "文化の日", flagday.PublicHoliday}, // Sunday
		{11, 4, "振替休日", flagday.SubstituteHoliday},
		{11, 23, "勤労感謝の日", flagday.PublicHoliday},
		{12, 23, "天皇誕生日", flagday.PublicHoliday},
	}
	year := 2013
	dates := flagday.InYear(year)
	testHolidays(t, year, dates, testdata)
	substituteTestdata := []expectedSubstitute{
		{
			original: testdata[7], // 5/5
			holiday:  testdata[8], // 5/6
		},
		{
			original: testdata[13], // 11/3
			holiday:  testdata[14], // 11/4
		},
	}
	testSubstitute(t, year, dates, substituteTestdata)
}

func testSubstitute(t *testing.T, year int, dates []flagday.Holiday, testdata []expectedSubstitute) {
	t.Helper()
	testSubstituteNoOriginal(t, dates)
	testSubstituteIsSubstituted(t, dates)
	for _, td := range testdata {
		testSubstituteInvalidOriginal(t, year, dates, td)
	}
}

func testSubstituteNoOriginal(t *testing.T, dates []flagday.Holiday) {
	for _, date := range dates {
		if date.Kind() == flagday.SubstituteHoliday {
			if date.Original() == nil {
				s := date.Time().Format("2006/1/2")
				t.Errorf("Substitute holiday (%s) should have original holiday", s)
			}
		}
	}
}

func testSubstituteIsSubstituted(t *testing.T, dates []flagday.Holiday) {
	for _, date := range dates {
		if date.Kind() == flagday.SubstituteHoliday {
			if !date.Original().IsSubstituted() {
				s := date.Original().Time().Format("2006/1/2")
				t.Errorf("Substituted original holiday (%s) should have substituted flag", s)
			}
		}
	}
}

func testSubstituteInvalidOriginal(t *testing.T, year int, dates []flagday.Holiday, td expectedSubstitute) {
	found := false
	for _, date := range dates {
		if date.Month() == td.holiday.month && date.Day() == td.holiday.day {
			found = true
			testDate(t, year, date.Original(), td.original)
		}
	}
	if !found {
		t.Errorf("substitute holiday not found: %d/%d/%d", year, td.holiday.month, td.holiday.day)
	}
}
