package flagday_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/pinzolo/flagday"
)

func TestTime(t *testing.T) {
	def, err := getDef("勤労感謝の日", 2017)
	if err != nil {
		t.Error(err)
	}
	d := flagday.FixedDateHoliday(def, 2017)
	tm := d.Time()
	if d.Def() == nil {
		t.Error("holiday should have own definition on public holiday")
	}
	if !isSameDef(*d.Def(), def) {
		t.Error("holiday should have same definition")
	}
	testTimeFields(t, tm, d)
}

func testTimeFields(t *testing.T, tm time.Time, d flagday.Holiday) {
	t.Helper()
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

func getDef(name string, year int) (flagday.Definition, error) {
	for _, def := range flagday.DefsInYear(year) {
		if def.Name() == name {
			return def, nil
		}
	}
	return nil, fmt.Errorf("definition not found: %s", name)
}

func isSameDef(def1 flagday.Definition, def2 flagday.Definition) bool {
	if def1.Type() != def2.Type() {
		return false
	}
	if def1.Name() != def2.Name() {
		return false
	}
	if def1.Month() != def2.Month() {
		return false
	}
	if def1.Day() != def2.Day() {
		return false
	}
	if def1.WeekNum() != def2.WeekNum() {
		return false
	}
	if def1.Begin() != def2.Begin() {
		return false
	}
	if def1.End() != def2.End() {
		return false
	}
	return true
}
