package easytime

import (
    "strings"
    "time"
)

const (
    // 基础时间格式
    BaseFormat = "2006-01-02 15:04:05"
)

var (
    // 星期列表
    weekDayMap = map[string]int{
        "Monday":    1,
        "Tuesday":   2,
        "Wednesday": 3,
        "Thursday":  4,
        "Friday":    5,
        "Saturday":  6,
        "Sunday":    0,
    }
)

// 获取当前时间基准格式
func BaseFormatTime() string {
    return time.Now().Format(BaseFormat)
}

// 获取自定义格式时间
func FormatTime(format string, T time.Time) string {
    replacer := strings.NewReplacer("y", "2006", "m", "01", "d", "02", "h", "15", "i", "04", "s", "05")
    formatTime := replacer.Replace(format)
    return T.Format(formatTime)
}

// 获取指定时间的time对象
func assignTime(assignYear int, assignMonth int, assignDay int) time.Time {
    return time.Now().AddDate(assignYear, assignMonth, assignDay)
}

// 获取指定时间的时间戳
func AssignTimeForUnix(assignYear int, assignMonth int, assignDay int) int64 {
    return assignTime(assignYear, assignMonth, assignDay).Unix()
}

// 获取指定时间的日期
func AssignTimeForDate(assignYear int, assignMonth int, assignDay int) string {
    return assignTime(assignYear, assignMonth, assignDay).Format("20060102")
}

/**
 * 获取星期
 *
 * 1-星期一 2-星期二 3-星期三 4-星期死 5-星期五 6-星期六 0-星期日
 *
 */
// 获取当前时间
func NowWeekDay() int {
    wd := time.Now().Weekday().String()
    return weekDayMap[wd]
}

// 获取指定日期是星期几
func WeekDay(T time.Time) int {
    wd := T.Weekday().String()
    return weekDayMap[wd]
}
