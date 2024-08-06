package taxcalculator

import (
	"fmt"
	"testing"
)

func TestCalOneMonth(t *testing.T) {
	cal, err := NewTaxCalculator(2024, RantingTag)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(cal.Cal(40000, 7))
	fmt.Println(cal.CalThisYear([]float64{6300, 6300, 6300, 6300, 6300, 6300, 30000, 36000}))
}
