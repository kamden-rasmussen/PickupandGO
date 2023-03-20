package calcs

import "time"

// need to offset by 7 dats to get the correct date

func GetOneWeekAgo() string {
	t := time.Now()
	OneWeekDate := t.AddDate(0, 0, 0).Format("2006-01-02")
	return OneWeekDate
}

func GetTwoWeeksAgo() string {
	t := time.Now()
	TwoWeeksDate := t.AddDate(0, 0, -7).Format("2006-01-02")
	return TwoWeeksDate
}

func GetOneMonthAgo() string {
	t := time.Now()
	OneMonthDate := t.AddDate(0, 0, -23).Format("2006-01-02")
	return OneMonthDate
}

func GetThreeMonthsAgo() string {
	t := time.Now()
	ThreeMonthsDate := t.AddDate(0, 0, -83).Format("2006-01-02")
	return ThreeMonthsDate
}

func GetSixMonthsAgo() string {
	t := time.Now()
	SixMonthsDate := t.AddDate(0, 0, -173).Format("2006-01-02")
	return SixMonthsDate
}

func GetOneYearAgo() string {
	t := time.Now()
	OneYearDate := t.AddDate(0, 0, -358).Format("2006-01-02")
	return OneYearDate
}
