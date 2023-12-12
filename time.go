package skmk

import "time"

func convertRomanMonth() string {
	currentMonth := time.Now().Month()
	// Konversi bulan ke format Romawi
	romanMonths := map[time.Month]string{
		time.January:   "I",
		time.February:  "II",
		time.March:     "III",
		time.April:     "IV",
		time.May:       "V",
		time.June:      "VI",
		time.July:      "VII",
		time.August:    "VIII",
		time.September: "IX",
		time.October:   "X",
		time.November:  "XI",
		time.December:  "XII",
	}

	return romanMonths[currentMonth]
}

func GetCurrentYear() int {
	currentYear := time.Now().Year()
	return currentYear
}
