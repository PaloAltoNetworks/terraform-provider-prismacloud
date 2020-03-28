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

// Valid values for Relative.Unit and ToNow.Unit.
const (
    TimeLogin = "login" // Valid for ToNow.Unit.
    TimeEpoch = "epoch" // Valid for ToNow.Unit.
	TimeHour  = "hour" // Valid for Relative.Unit.
	TimeDay   = "day"
	TimeWeek  = "week"
	TimeMonth = "month"
	TimeYear  = "year"
)
