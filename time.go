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

func translateMonthToIndonesian(month time.Month) string {
	translations := map[time.Month]string{
		time.January:   "Januari",
		time.February:  "Februari",
		time.March:     "Maret",
		time.April:     "April",
		time.May:       "Mei",
		time.June:      "Juni",
		time.July:      "Juli",
		time.August:    "Agustus",
		time.September: "September",
		time.October:   "Oktober",
		time.November:  "November",
		time.December:  "Desember",
	}

	return translations[month]
}

// func GetCurrentDate() string {
// 	currentTime := time.Now()
// 	day := currentTime.Format("02")
// 	month := translateMonthToIndonesian(currentTime.Month())
// 	year := currentTime.Format("2006")

// 	return fmt.Sprintf("%s %s %s", day, month, year)
// }
