package flagday

import "time"

var (
	// PublicHolidayStartDate is effective date of the law concerning national holidays.
	// 国民の祝日に関する法律が公布・施行された日
	PublicHolidayStartDate = time.Date(1948, 7, 20, 0, 0, 0, 0, jst())
	// SubstituteHolidayStartDate is effective date of the law concerning substitute holidays.
	// 振替休日に関する改正、および施行された日
	SubstituteHolidayStartDate = time.Date(1973, 4, 12, 0, 0, 0, 0, jst())
	// NationalHolidayStartDate is effective date of the law concerning national holidays.
	// 国民の休日に関する改正、および施行された日
	NationalHolidayStartDate = time.Date(1985, 12, 27, 0, 0, 0, 0, jst())
	// HappyMondayStartDate is effective date of the law concerning Happy Monday.
	// ハッピーマンデー制度が施行された日
	HappyMondayStartDate = time.Date(2000, 1, 1, 0, 0, 0, 0, jst())
	// HappyMondayAddedDate is effective date of the law where Happy Monday was added.
	// 海の日と敬老の日がハッピーマンデーとされた改正の施行日
	HappyMondayAddedDate = time.Date(2003, 1, 1, 0, 0, 0, 0, jst())
)
