package flagday

import "testing"

func TestSubstituteHolidayBefore(t *testing.T) {
	testdata := []expected{
		{1, 1, "元日", PublicHoliday},
		{1, 15, "成人の日", PublicHoliday},
		{2, 11, "建国記念の日", PublicHoliday},
		{3, 20, "春分の日", PublicHoliday},
		{4, 29, "天皇誕生日", PublicHoliday},
		{5, 3, "憲法記念日", PublicHoliday},
		{5, 5, "こどもの日", PublicHoliday}, // Sunday
		{9, 15, "敬老の日", PublicHoliday},
		{9, 23, "秋分の日", PublicHoliday},
		{10, 10, "体育の日", PublicHoliday},
		{11, 03, "文化の日", PublicHoliday},
		{11, 23, "勤労感謝の日", PublicHoliday},
	}
	year := 1968
	dates := InYear(year)
	check(t, year, dates, testdata)
}

func TestSubstituteHolidayGoldenWeekSunday3(t *testing.T) {
	testdata := []expected{
		{1, 1, "元日", PublicHoliday},
		{1, 12, "成人の日", PublicHoliday},
		{2, 11, "建国記念の日", PublicHoliday},
		{3, 20, "春分の日", PublicHoliday},
		{4, 29, "昭和の日", PublicHoliday},
		{5, 3, "憲法記念日", PublicHoliday}, // Sunday
		{5, 4, "みどりの日", PublicHoliday},
		{5, 5, "こどもの日", PublicHoliday},
		{5, 6, "振替休日", SubstituteHoliday},
		{7, 20, "海の日", PublicHoliday},
		{9, 21, "敬老の日", PublicHoliday},
		{9, 22, "国民の休日", NationalHoliday},
		{9, 23, "秋分の日", PublicHoliday},
		{10, 12, "体育の日", PublicHoliday},
		{11, 3, "文化の日", PublicHoliday},
		{11, 23, "勤労感謝の日", PublicHoliday},
		{12, 23, "天皇誕生日", PublicHoliday},
	}
	year := 2009
	dates := InYear(year)
	check(t, year, dates, testdata)
}

func TestSubstituteHolidayGoldenWeekSunday4(t *testing.T) {
	testdata := []expected{
		{1, 1, "元日", PublicHoliday},
		{1, 14, "成人の日", PublicHoliday},
		{2, 11, "建国記念の日", PublicHoliday},
		{3, 20, "春分の日", PublicHoliday},
		{4, 29, "昭和の日", PublicHoliday},
		{5, 3, "憲法記念日", PublicHoliday},
		{5, 4, "みどりの日", PublicHoliday}, // Sunday
		{5, 5, "こどもの日", PublicHoliday},
		{5, 6, "振替休日", SubstituteHoliday},
		{7, 21, "海の日", PublicHoliday},
		{9, 15, "敬老の日", PublicHoliday},
		{9, 23, "秋分の日", PublicHoliday},
		{10, 13, "体育の日", PublicHoliday},
		{11, 3, "文化の日", PublicHoliday},
		{11, 23, "勤労感謝の日", PublicHoliday},
		{11, 24, "振替休日", SubstituteHoliday},
		{12, 23, "天皇誕生日", PublicHoliday},
	}
	year := 2008
	dates := InYear(year)
	check(t, year, dates, testdata)
}

func TestSubstituteHolidayGoldenWeekSunday5(t *testing.T) {
	testdata := []expected{
		{1, 1, "元日", PublicHoliday},
		{1, 14, "成人の日", PublicHoliday},
		{2, 11, "建国記念の日", PublicHoliday},
		{3, 20, "春分の日", PublicHoliday},
		{4, 29, "昭和の日", PublicHoliday},
		{5, 3, "憲法記念日", PublicHoliday},
		{5, 4, "みどりの日", PublicHoliday},
		{5, 5, "こどもの日", PublicHoliday}, // Sunday
		{5, 6, "振替休日", SubstituteHoliday},
		{7, 15, "海の日", PublicHoliday},
		{9, 16, "敬老の日", PublicHoliday},
		{9, 23, "秋分の日", PublicHoliday},
		{10, 14, "体育の日", PublicHoliday},
		{11, 3, "文化の日", PublicHoliday},
		{11, 4, "振替休日", SubstituteHoliday},
		{11, 23, "勤労感謝の日", PublicHoliday},
		{12, 23, "天皇誕生日", PublicHoliday},
	}
	year := 2013
	dates := InYear(year)
	check(t, year, dates, testdata)
}
