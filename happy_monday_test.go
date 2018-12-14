package flagday

import "testing"

func TestHappyMondayBefore(t *testing.T) {
	testdata := []expected{
		{1, 1, "元日", PublicHoliday},
		{1, 15, "成人の日", PublicHoliday},
		{2, 11, "建国記念の日", PublicHoliday},
		{3, 21, "春分の日", PublicHoliday},
		{3, 22, "振替休日", SubstituteHoliday},
		{4, 29, "みどりの日", PublicHoliday},
		{5, 3, "憲法記念日", PublicHoliday},
		{5, 4, "国民の休日", NationalHoliday},
		{5, 5, "こどもの日", PublicHoliday},
		{7, 20, "海の日", PublicHoliday},
		{9, 15, "敬老の日", PublicHoliday},
		{9, 23, "秋分の日", PublicHoliday},
		{10, 10, "体育の日", PublicHoliday},
		{10, 11, "振替休日", SubstituteHoliday},
		{11, 3, "文化の日", PublicHoliday},
		{11, 23, "勤労感謝の日", PublicHoliday},
		{12, 23, "天皇誕生日", PublicHoliday},
	}
	year := 1999
	dates := InYear(year)
	check(t, year, dates, testdata)
}

func TestHappyMonday1st(t *testing.T) {
	testdata := []expected{
		{1, 1, "元日", PublicHoliday},
		{1, 10, "成人の日", PublicHoliday}, // changed
		{2, 11, "建国記念の日", PublicHoliday},
		{3, 20, "春分の日", PublicHoliday},
		{4, 29, "みどりの日", PublicHoliday},
		{5, 3, "憲法記念日", PublicHoliday},
		{5, 4, "国民の休日", NationalHoliday},
		{5, 5, "こどもの日", PublicHoliday},
		{7, 20, "海の日", PublicHoliday},
		{9, 15, "敬老の日", PublicHoliday},
		{9, 23, "秋分の日", PublicHoliday},
		{10, 9, "体育の日", PublicHoliday}, // changed
		{11, 3, "文化の日", PublicHoliday},
		{11, 23, "勤労感謝の日", PublicHoliday},
		{12, 23, "天皇誕生日", PublicHoliday},
	}
	year := 2000
	dates := InYear(year)
	check(t, year, dates, testdata)
}

func TestHappyMonday2nd(t *testing.T) {
	testdata := []expected{
		{1, 1, "元日", PublicHoliday},
		{1, 13, "成人の日", PublicHoliday},
		{2, 11, "建国記念の日", PublicHoliday},
		{3, 21, "春分の日", PublicHoliday},
		{4, 29, "みどりの日", PublicHoliday},
		{5, 3, "憲法記念日", PublicHoliday},
		{5, 4, "国民の休日", NationalHoliday},
		{5, 5, "こどもの日", PublicHoliday},
		{7, 21, "海の日", PublicHoliday}, // changed
		{9, 15, "敬老の日", PublicHoliday},
		{9, 23, "秋分の日", PublicHoliday},
		{10, 13, "体育の日", PublicHoliday},
		{11, 3, "文化の日", PublicHoliday},
		{11, 23, "勤労感謝の日", PublicHoliday},
		{11, 24, "振替休日", SubstituteHoliday},
		{12, 23, "天皇誕生日", PublicHoliday},
	}
	year := 2003
	dates := InYear(year)
	check(t, year, dates, testdata)

	testdata = []expected{
		{1, 1, "元日", PublicHoliday},
		{1, 12, "成人の日", PublicHoliday},
		{2, 11, "建国記念の日", PublicHoliday},
		{3, 20, "春分の日", PublicHoliday},
		{4, 29, "みどりの日", PublicHoliday},
		{5, 3, "憲法記念日", PublicHoliday},
		{5, 4, "国民の休日", NationalHoliday},
		{5, 5, "こどもの日", PublicHoliday},
		{7, 19, "海の日", PublicHoliday},
		{9, 20, "敬老の日", PublicHoliday}, // changed
		{9, 23, "秋分の日", PublicHoliday},
		{10, 11, "体育の日", PublicHoliday},
		{11, 3, "文化の日", PublicHoliday},
		{11, 23, "勤労感謝の日", PublicHoliday},
		{12, 23, "天皇誕生日", PublicHoliday},
	}
	year = 2004
	dates = InYear(year)
	check(t, year, dates, testdata)
}
