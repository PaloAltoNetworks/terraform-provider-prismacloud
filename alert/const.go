package alert

const (
	singular = "alert"
	plural   = "alerts"
)

// These are used by List().
const (
	TimeRelative = "relative"
	TimeAbsolute = "absolute"
	TimeToNow    = "to_now"
)

// Valid values for Relative.Unit.
const (
	TimeHour  = "hour"
	TimeDay   = "day"
	TimeWeek  = "week"
	TimeMonth = "month"
	TimeYear  = "year"
)
