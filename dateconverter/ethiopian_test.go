package dateconverter

import (
	"reflect"
	"testing"
)

func TestEthiopianDate(t *testing.T) {
	var expectedEthio, expectedGrego string

	ethiopianDate := "2015-01-18 00:00:00 +0000 UTC"
	gregorianDate := "2022-09-28 00:00:00 +0000 UTC"

	time, err := Ethiopian(2022, 9, 28)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	expectedEthio = time.String()
	time, err = Gregorian(2015, 1, 18)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	expectedGrego = time.String()

	if !reflect.DeepEqual(ethiopianDate, expectedEthio) && !reflect.DeepEqual(gregorianDate, expectedGrego) {
		t.Errorf("Date conversion not work got %v want %v", expectedEthio, ethiopianDate)
	}
}
