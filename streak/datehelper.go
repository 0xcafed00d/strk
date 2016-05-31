package streak

import "time"

func fromYearDay(year, yearDay int) time.Time {
	return time.Date(year, 1, yearDay, 0, 0, 0, 0, time.Local)
}

func isLeap(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
