package calcs

import "time"


func GetOneWeekAgo() string {
	t := time.Now()
	OneWeekDate := t.AddDate(0, 0, 7).Format("2006-01-02")
	return OneWeekDate
}

func GetTwoWeeksAgo() string {
	t := time.Now()
	TwoWeeksDate := t.AddDate(0, 0, 14).Format("2006-01-02")
	return TwoWeeksDate
}

func GetOneMonthAgo() string {
	t := time.Now()
	OneMonthDate := t.AddDate(0, 0, 30).Format("2006-01-02")
	return OneMonthDate
}

func GetThreeMonthsAgo() string {
	t := time.Now()
	ThreeMonthsDate := t.AddDate(0, 0, 90).Format("2006-01-02")
	return ThreeMonthsDate
}

func GetSixMonthsAgo() string {
	t := time.Now()
	SixMonthsDate := t.AddDate(0, 0, 180).Format("2006-01-02")
	return SixMonthsDate
}

func GetOneYearAgo() string {
	t := time.Now()
	OneYearDate := t.AddDate(0, 0, 365).Format("2006-01-02")
	return OneYearDate
}
