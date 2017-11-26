package flagday

import (
	"fmt"
	"testing"
)

func TestTime(t *testing.T) {
	def, err := getDef("勤労感謝の日", 2017)
	if err != nil {
		t.Error(err)
	}
	d := FixedDateHoliday(def, 2017)
	tm := d.Time()
	if tm.Year() != d.Year() {
		t.Errorf("converted time should have %d as year, but got %d", d.Year(), tm.Year())
	}
	if int(tm.Month()) != d.Month() {
		t.Errorf("converted time should have %d as month, but got %d", d.Month(), tm.Month())
	}
	if tm.Day() != d.Day() {
		t.Errorf("converted time should have %d as day, but got %d", d.Day(), tm.Day())
	}
	if tm.Hour() != 0 {
		t.Errorf("converted time should have 0 as hour, but got %d", tm.Hour())
	}
	if tm.Minute() != 0 {
		t.Errorf("converted time should have 0 as minute, but got %d", tm.Minute())
	}
	if tm.Second() != 0 {
		t.Errorf("converted time should have 0 as second, but got %d", tm.Second())
	}
	if tm.Location().String() != "Asia/Tokyo" {
		t.Errorf("converted time should have JST location, but got %s", tm.Location())
	}
}

func getDef(name string, year int) (Definition, error) {
	for _, def := range DefsInYear(year) {
		if def.Name() == name {
			return def, nil
		}
	}
	return nil, fmt.Errorf("definition not found: %s", name)
}
