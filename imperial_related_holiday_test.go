package flagday

import "testing"

func TestWeddingOfPrinceAkihito(t *testing.T) {
	_, err := PublicHolidayOf(1958, 4, 10)
	if err == nil {
		t.Errorf("1958/04/10 is not holiday")
	}
	d, err := PublicHolidayOf(1959, 4, 10)
	if err != nil {
		t.Errorf("1959/04/10 is not holiday")
	}
	if d.Name() != "皇太子明仁親王の結婚の儀" {
		t.Errorf("invalid holiday name, got %s", d.Name())
	}
	_, err = PublicHolidayOf(1960, 4, 10)
	if err == nil {
		t.Errorf("1960/04/10 is not holiday")
	}
}

func TestFuneralCeremonyOfEmperorShowa(t *testing.T) {
	_, err := PublicHolidayOf(1988, 2, 24)
	if err == nil {
		t.Errorf("1988/02/24 is not holiday")
	}
	d, err := PublicHolidayOf(1989, 2, 24)
	if err != nil {
		t.Errorf("1989/02/24 is not holiday")
	}
	if d.Name() != "昭和天皇の大喪の礼" {
		t.Errorf("invalid holiday name, got %s", d.Name())
	}
	_, err = PublicHolidayOf(1990, 2, 24)
	if err == nil {
		t.Errorf("1990/02/24 is not holiday")
	}
}

func TestCeremonyOfSokuiReiseiden(t *testing.T) {
	_, err := PublicHolidayOf(1989, 11, 12)
	if err == nil {
		t.Errorf("1988/02/24 is not holiday")
	}
	d, err := PublicHolidayOf(1990, 11, 12)
	if err != nil {
		t.Errorf("1989/02/24 is not holiday")
	}
	if d.Name() != "即位礼正殿の儀" {
		t.Errorf("invalid holiday name, got %s", d.Name())
	}
	_, err = PublicHolidayOf(1991, 11, 12)
	if err == nil {
		t.Errorf("1990/02/24 is not holiday")
	}
}

func TestWeddingOfPrinceNaruhito(t *testing.T) {
	_, err := PublicHolidayOf(1992, 6, 9)
	if err == nil {
		t.Errorf("1992/06/09 is not holiday")
	}
	d, err := PublicHolidayOf(1993, 6, 9)
	if err != nil {
		t.Errorf("1993/06/09 is not holiday")
	}
	if d.Name() != "皇太子徳仁親王の結婚の儀" {
		t.Errorf("invalid holiday name, got %s", d.Name())
	}
	_, err = PublicHolidayOf(1994, 6, 9)
	if err == nil {
		t.Errorf("1994/06/09 is not holiday")
	}
}
