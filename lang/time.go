package lang

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// convert time.Now() to *time.Time
func NowToPtr() *time.Time {
	return TimeToPtr(time.Now())
}

// convert time to *time.Time
func TimeToPtr(t time.Time) *time.Time {
	return &t
}

// parse time string to time using China timeZone
func ParseTimeToChinaTimezone(layoutList []string, timeString string) (*time.Time, error) {
	if len(layoutList) <= 0 {
		return nil, errors.New("layoutList不能为空,或者长度不能小于0")
	}
	for _, eachLayout := range layoutList {
		timeValue, err := time.ParseInLocation(eachLayout, timeString, ChinaTimezone)
		if err == nil {
			return &timeValue, nil
		}
	}
	return nil, fmt.Errorf("无效的时间,时间格式必须是%s", strings.Join(layoutList, ","))
}

// parse time string to time used available layout
// return when any layout not error
func ParseTimeUseLayout(s string, layout []string, loc *time.Location) *time.Time {
	cLoc := loc
	if cLoc == nil {
		cLoc = time.Local
	}
	for _, eachLayout := range layout {
		t, err := time.ParseInLocation(eachLayout, s, cLoc)
		if err == nil {
			return &t
		}
	}
	return nil
}

// parse a string list to time list
func ParseStringListToTimeList(timeStringList []string) ([]*time.Time, error) {
	timeList := make([]*time.Time, 0)
	if len(timeStringList) <= 0 {
		return timeList, nil
	}
	supportLayoutList := []string{
		time.RFC3339,
		DefaultTimeLayout,
	}
	for _, eachS := range timeStringList {
		t := ParseTimeUseLayout(eachS, supportLayoutList, ChinaTimezone)
		if t == nil {
			return make([]*time.Time, 0), fmt.Errorf("invalid time format:%s", eachS)
		}
		timeList = append(timeList, t)
	}
	return timeList, nil
}

func DurationHumanText(duration time.Duration) string {
	if duration <= 0 {
		return "0 seconds"
	}
	remainDuration := duration
	days := duration / (24 * time.Hour)
	remainDuration = remainDuration - days*(24*time.Hour)

	hours := remainDuration / time.Hour
	remainDuration = remainDuration - hours*time.Hour

	minutes := remainDuration / time.Minute
	remainDuration = remainDuration - minutes*time.Minute

	seconds := remainDuration / time.Second
	// remainDuration = remainDuration - seconds*time.Second
	return fmt.Sprintf("%d days %d hours %d minutes %d seconds",
		days,
		hours,
		minutes,
		seconds)
}

func IsUTC(t *time.Time) bool {
	return t.Location().String() == time.UTC.String()
}

// 获取时间的utc时区unix时间戳，以秒为单位
func UnixTime(t *time.Time) int64 {
	var sec int64
	if IsUTC(t) {
		sec = t.Unix()
	} else {
		sec = t.UTC().Unix()
	}
	return sec
}

// 获取时间的utc时区unix时间戳，以分钟为单位
func UnixMinute(t *time.Time) int64 {
	return UnixTime(t) / 60
}

// 获取时间的utc时区unix时间戳，以小时为单位
func UnixHour(t *time.Time) int64 {
	return UnixTime(t) / 3600
}

// 获取时间的utc时区unix时间戳，以天为单位
func UnixDay(t *time.Time) int64 {
	return UnixTime(t) / 86400
}

// utc0 today range
func UTC0TodayRange() (start, end time.Time) {
	now := time.Now().UTC()

	start = time.Date(
		now.Year(), now.Month(), now.Day(),
		0, 0, 0, 0,
		time.UTC,
	)

	end = start.Add(24 * time.Hour).Add(-time.Nanosecond)
	return
}

// utc yesterday range
func UTC0YesterdayRange() (start, end time.Time) {
	todayStart, _ := UTC0TodayRange()

	start = todayStart.Add(-24 * time.Hour)
	end = todayStart.Add(-time.Nanosecond)
	return
}

// UTC+8 today
func UTC8TodayRange() (start, end time.Time) {
	now := time.Now()

	start = utc8DayStartInUTC(now)
	end = start.Add(24 * time.Hour).Add(-time.Nanosecond)
	return
}

// UTC+8 yesterday
func UTC8YesterdayRange() (start, end time.Time) {
	todayStart, _ := UTC8TodayRange()

	start = todayStart.Add(-24 * time.Hour)
	end = todayStart.Add(-time.Nanosecond)
	return
}

var utc8 = time.FixedZone("UTC+8", 8*3600)

func utc8DayStartInUTC(t time.Time) time.Time {
	local := t.In(utc8)

	startLocal := time.Date(
		local.Year(), local.Month(), local.Day(),
		0, 0, 0, 0,
		utc8,
	)

	return startLocal.UTC()
}
