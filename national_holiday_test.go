package flagday_test

import (
	"testing"

	"github.com/pinzolo/flagday"
)

func TestNationalHoliday(t *testing.T) {
	testdata := []expected{
		{1, 1, "元日", flagday.PublicHoliday},
		{1, 12, "成人の日", flagday.PublicHoliday},
		{2, 11, "建国記念の日", flagday.PublicHoliday},
		{3, 21, "春分の日", flagday.PublicHoliday},
		{4, 29, "昭和の日", flagday.PublicHoliday},
		{5, 3, "憲法記念日", flagday.PublicHoliday},
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
	year := 2015
	dates := flagday.InYear(year)
	testHolidays(t, year, dates, testdata)
}

func TestNationalHolidayBefore(t *testing.T) {
	testdata := []expected{
		{1, 1, "元日", flagday.PublicHoliday},
		{1, 15, "成人の日", flagday.PublicHoliday},
		{2, 11, "建国記念の日", flagday.PublicHoliday},
		{3, 21, "春分の日", flagday.PublicHoliday},
		{3, 22, "振替休日", flagday.SubstituteHoliday},
		{4, 29, "天皇誕生日", flagday.PublicHoliday},
		{5, 3, "憲法記念日", flagday.PublicHoliday},
		// 5/4 is not holiday
		{5, 5, "こどもの日", flagday.PublicHoliday},
		{9, 15, "敬老の日", flagday.PublicHoliday},
		{9, 23, "秋分の日", flagday.PublicHoliday},
		{10, 10, "体育の日", flagday.PublicHoliday},
		{10, 11, "振替休日", flagday.SubstituteHoliday},
		{11, 03, "文化の日", flagday.PublicHoliday},
		{11, 23, "勤労感謝の日", flagday.PublicHoliday},
	}
	year := 1982
	dates := flagday.InYear(year)
	testHolidays(t, year, dates, testdata)
}
