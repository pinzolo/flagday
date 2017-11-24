package flagday

import "testing"

func TestNationalHoliday(t *testing.T) {
	testdata := []expected{
		{1, 1, "元日"},
		{1, 12, "成人の日"},
		{2, 11, "建国記念の日"},
		{3, 21, "春分の日"},
		{4, 29, "昭和の日"},
		{5, 3, "憲法記念日"},
		{5, 4, "みどりの日"},
		{5, 5, "こどもの日"},
		{5, 6, "振替休日"},
		{7, 20, "海の日"},
		{9, 21, "敬老の日"},
		{9, 22, "国民の休日"},
		{9, 23, "秋分の日"},
		{10, 12, "体育の日"},
		{11, 3, "文化の日"},
		{11, 23, "勤労感謝の日"},
		{12, 23, "天皇誕生日"},
	}
	year := 2015
	dates := InYear(year)
	check(t, year, dates, testdata)
}

func TestNationalHolidayBefore(t *testing.T) {
	testdata := []expected{
		{1, 1, "元日"},
		{1, 15, "成人の日"},
		{2, 11, "建国記念の日"},
		{3, 21, "春分の日"},
		{3, 22, "振替休日"},
		{4, 29, "天皇誕生日"},
		{5, 3, "憲法記念日"},
		// 5/4 is not holiday
		{5, 5, "こどもの日"},
		{9, 15, "敬老の日"},
		{9, 23, "秋分の日"},
		{10, 10, "体育の日"},
		{10, 11, "振替休日"},
		{11, 03, "文化の日"},
		{11, 23, "勤労感謝の日"},
	}
	year := 1982
	dates := InYear(year)
	check(t, year, dates, testdata)
}
