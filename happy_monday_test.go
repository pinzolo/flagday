package flagday_test

import (
	"testing"

	"github.com/pinzolo/flagday"
)

func TestHappyMondayBefore(t *testing.T) {
	testdata := []expected{
		{1, 1, "元日", flagday.PublicHoliday},
		{1, 15, "成人の日", flagday.PublicHoliday},
		{2, 11, "建国記念の日", flagday.PublicHoliday},
		{3, 21, "春分の日", flagday.PublicHoliday},
		{3, 22, "振替休日", flagday.SubstituteHoliday},
		{4, 29, "みどりの日", flagday.PublicHoliday},
		{5, 3, "憲法記念日", flagday.PublicHoliday},
		{5, 4, "国民の休日", flagday.NationalHoliday},
		{5, 5, "こどもの日", flagday.PublicHoliday},
		{7, 20, "海の日", flagday.PublicHoliday},
		{9, 15, "敬老の日", flagday.PublicHoliday},
		{9, 23, "秋分の日", flagday.PublicHoliday},
		{10, 10, "体育の日", flagday.PublicHoliday},
		{10, 11, "振替休日", flagday.SubstituteHoliday},
		{11, 3, "文化の日", flagday.PublicHoliday},
		{11, 23, "勤労感謝の日", flagday.PublicHoliday},
		{12, 23, "天皇誕生日", flagday.PublicHoliday},
	}
	year := 1999
	dates := flagday.InYear(year)
	testHolidays(t, year, dates, testdata)
}

func TestHappyMonday1st(t *testing.T) {
	testdata := []expected{
		{1, 1, "元日", flagday.PublicHoliday},
		{1, 10, "成人の日", flagday.PublicHoliday}, // changed
		{2, 11, "建国記念の日", flagday.PublicHoliday},
		{3, 20, "春分の日", flagday.PublicHoliday},
		{4, 29, "みどりの日", flagday.PublicHoliday},
		{5, 3, "憲法記念日", flagday.PublicHoliday},
		{5, 4, "国民の休日", flagday.NationalHoliday},
		{5, 5, "こどもの日", flagday.PublicHoliday},
		{7, 20, "海の日", flagday.PublicHoliday},
		{9, 15, "敬老の日", flagday.PublicHoliday},
		{9, 23, "秋分の日", flagday.PublicHoliday},
		{10, 9, "体育の日", flagday.PublicHoliday}, // changed
		{11, 3, "文化の日", flagday.PublicHoliday},
		{11, 23, "勤労感謝の日", flagday.PublicHoliday},
		{12, 23, "天皇誕生日", flagday.PublicHoliday},
	}
	year := 2000
	dates := flagday.InYear(year)
	testHolidays(t, year, dates, testdata)
}

func TestHappyMonday2nd(t *testing.T) {
	testdata := []expected{
		{1, 1, "元日", flagday.PublicHoliday},
		{1, 13, "成人の日", flagday.PublicHoliday},
		{2, 11, "建国記念の日", flagday.PublicHoliday},
		{3, 21, "春分の日", flagday.PublicHoliday},
		{4, 29, "みどりの日", flagday.PublicHoliday},
		{5, 3, "憲法記念日", flagday.PublicHoliday},
		{5, 4, "国民の休日", flagday.NationalHoliday},
		{5, 5, "こどもの日", flagday.PublicHoliday},
		{7, 21, "海の日", flagday.PublicHoliday}, // changed
		{9, 15, "敬老の日", flagday.PublicHoliday},
		{9, 23, "秋分の日", flagday.PublicHoliday},
		{10, 13, "体育の日", flagday.PublicHoliday},
		{11, 3, "文化の日", flagday.PublicHoliday},
		{11, 23, "勤労感謝の日", flagday.PublicHoliday},
		{11, 24, "振替休日", flagday.SubstituteHoliday},
		{12, 23, "天皇誕生日", flagday.PublicHoliday},
	}
	year := 2003
	dates := flagday.InYear(year)
	testHolidays(t, year, dates, testdata)

	testdata = []expected{
		{1, 1, "元日", flagday.PublicHoliday},
		{1, 12, "成人の日", flagday.PublicHoliday},
		{2, 11, "建国記念の日", flagday.PublicHoliday},
		{3, 20, "春分の日", flagday.PublicHoliday},
		{4, 29, "みどりの日", flagday.PublicHoliday},
		{5, 3, "憲法記念日", flagday.PublicHoliday},
		{5, 4, "国民の休日", flagday.NationalHoliday},
		{5, 5, "こどもの日", flagday.PublicHoliday},
		{7, 19, "海の日", flagday.PublicHoliday},
		{9, 20, "敬老の日", flagday.PublicHoliday}, // changed
		{9, 23, "秋分の日", flagday.PublicHoliday},
		{10, 11, "体育の日", flagday.PublicHoliday},
		{11, 3, "文化の日", flagday.PublicHoliday},
		{11, 23, "勤労感謝の日", flagday.PublicHoliday},
		{12, 23, "天皇誕生日", flagday.PublicHoliday},
	}
	year = 2004
	dates = flagday.InYear(year)
	testHolidays(t, year, dates, testdata)
}
