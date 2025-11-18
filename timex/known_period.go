package timex

import (
	"time"
)

// 包内统一使用 [start, end) 边界
const (
	boundary = IncludeStartExcludeEnd // "[)"
)

// 获取今天零点
func today() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

// TodayPeriod 返回今日区间 [00:00, 明日00:00)
func TodayPeriod() *Period {
	start := today()
	end := start.Add(24 * time.Hour)
	v := NewPeriod(start, end, boundary)
	return &v
}

// YesterdayPeriod 返回昨日区间
func YesterdayPeriod() *Period {
	end := today()
	start := end.AddDate(0, 0, -1)
	v := NewPeriod(start, end, boundary)
	return &v
}

// LastNDaysPeriod 返回最近 N 天区间（例如 N=7 表示最近7天）
func LastNDaysPeriod(n int) *Period {
	end := today()
	start := end.AddDate(0, 0, -n)
	v := NewPeriod(start, end, boundary)
	return &v
}

// Last7DaysPeriod 最近7天
func Last7DaysPeriod() *Period {
	return LastNDaysPeriod(7)
}

// Last30DaysPeriod 最近30天
func Last30DaysPeriod() *Period {
	return LastNDaysPeriod(30)
}

// Last90DaysPeriod 最近90天
func Last90DaysPeriod() *Period {
	return LastNDaysPeriod(90)
}

// Last180DaysPeriod 最近180天
func Last180DaysPeriod() *Period {
	return LastNDaysPeriod(180)
}

// ThisWeekPeriod 返回本周区间（周一算第一天）
func ThisWeekPeriod() *Period {
	now := time.Now()
	weekday := int(now.Weekday())
	if weekday == 0 {
		weekday = 7 // 周日修正为 7
	}
	start := today().AddDate(0, 0, -(weekday - 1))
	end := start.AddDate(0, 0, 7)
	v := NewPeriod(start, end, boundary)
	return &v
}

// LastWeekPeriod 返回上周区间
func LastWeekPeriod() *Period {
	thisWeek := ThisWeekPeriod()
	start := thisWeek.GetStartDate().AddDate(0, 0, -7)
	end := thisWeek.GetStartDate()
	v := NewPeriod(start, end, boundary)
	return &v
}

// ThisMonthPeriod 返回本月区间
func ThisMonthPeriod() *Period {
	now := time.Now()
	start := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	end := start.AddDate(0, 1, 0)
	v := NewPeriod(start, end, boundary)
	return &v
}

// LastMonthPeriod 返回上个月区间
func LastMonthPeriod() *Period {
	thisMonth := ThisMonthPeriod()
	start := thisMonth.GetStartDate().AddDate(0, -1, 0)
	end := thisMonth.GetStartDate()
	v := NewPeriod(start, end, boundary)
	return &v
}

// 本季度
func ThisQuarter() *Period {
	now := time.Now()
	y, m, _ := now.Date()
	q := (int(m)-1)/3 + 1
	startMonth := time.Month((q-1)*3 + 1)
	start := time.Date(y, startMonth, 1, 0, 0, 0, 0, now.Location())
	end := start.AddDate(0, 3, 0)
	v := NewPeriod(start, end, boundary)
	return &v
}

// 本年
func ThisYear() *Period {
	now := time.Now()
	y := now.Year()
	start := time.Date(y, 1, 1, 0, 0, 0, 0, now.Location())
	end := start.AddDate(1, 0, 0)
	v := NewPeriod(start, end, boundary)
	return &v
}

// 去年
func LastYear() *Period {
	now := time.Now()
	y := now.Year() - 1

	start := time.Date(y, 1, 1, 0, 0, 0, 0, now.Location())
	end := start.AddDate(1, 0, 0)

	v := NewPeriod(start, end, boundary)
	return &v
}
