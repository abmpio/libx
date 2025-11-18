package timex

import "sync"

// PeriodFactory 是一个返回 *Period 的函数
type PeriodFactory func() *Period
type PeriodKey string

// 全局 registry
var (
	periodRegistry = make(map[PeriodKey]PeriodFactory)
	mu             sync.RWMutex
)

const (
	// 今日 / 昨日
	KnownPeriodKey_Today     PeriodKey = "today"
	KnownPeriodKey_Yesterday PeriodKey = "yesterday"

	// 最近 N 天
	KnownPeriodKey_Last7Days   PeriodKey = "last7Days"
	KnownPeriodKey_Last30Days  PeriodKey = "last30Days"
	KnownPeriodKey_Last90Days  PeriodKey = "last90Days"
	KnownPeriodKey_Last180Days PeriodKey = "last180Days"

	// 周相关
	KnownPeriodKey_ThisWeek PeriodKey = "thisWeek"
	KnownPeriodKey_LastWeek PeriodKey = "lastWeek"

	// 月相关
	KnownPeriodKey_ThisMonth PeriodKey = "thisMonth"
	KnownPeriodKey_LastMonth PeriodKey = "lastMonth"

	// 季度
	KnownPeriodKey_ThisQuarter PeriodKey = "thisQuarter"

	// 年相关
	KnownPeriodKey_ThisYear PeriodKey = "thisYear"
	KnownPeriodKey_LastYear PeriodKey = "lastYear"
)

// RegisterPeriod 用于注册一个 key → 生成函数
func RegisterPeriod(key PeriodKey, fn PeriodFactory) {
	mu.Lock()
	defer mu.Unlock()
	periodRegistry[key] = fn
}

// GetPeriod 根据 key 获取一个 Period
// 找不到时返回 nil
func GetPeriod(key PeriodKey) *Period {
	mu.RLock()
	fn := periodRegistry[key]
	mu.RUnlock()

	if fn == nil {
		return nil
	}
	return fn()
}

func init() {

	// 今日 / 昨⽇
	RegisterPeriod(KnownPeriodKey_Today, TodayPeriod)
	RegisterPeriod(KnownPeriodKey_Yesterday, YesterdayPeriod)

	// 最近 N 天
	RegisterPeriod(KnownPeriodKey_Last7Days, Last7DaysPeriod)
	RegisterPeriod(KnownPeriodKey_Last30Days, Last30DaysPeriod)
	RegisterPeriod(KnownPeriodKey_Last90Days, Last90DaysPeriod)
	RegisterPeriod(KnownPeriodKey_Last180Days, Last180DaysPeriod)

	// 周
	RegisterPeriod(KnownPeriodKey_ThisWeek, ThisWeekPeriod)
	RegisterPeriod(KnownPeriodKey_LastWeek, LastWeekPeriod)

	// 月
	RegisterPeriod(KnownPeriodKey_ThisMonth, ThisMonthPeriod)
	RegisterPeriod(KnownPeriodKey_LastMonth, LastMonthPeriod)

	// 季度
	RegisterPeriod(KnownPeriodKey_ThisQuarter, ThisQuarter)

	// 年
	RegisterPeriod(KnownPeriodKey_ThisYear, ThisYear)
	RegisterPeriod(KnownPeriodKey_LastYear, LastYear)
}
