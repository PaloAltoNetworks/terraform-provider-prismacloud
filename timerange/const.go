package timerange

// These constants are implicitly used based on the time range value specified.
const (
	TypeRelative = "relative"
	TypeAbsolute = "absolute"
	TypeToNow    = "to_now"
)

// Valid values for Relative.Unit and ToNow.Unit.
const (
	Login = "login" // Valid for ToNow.Unit.
	Epoch = "epoch" // Valid for ToNow.Unit.
	Hour  = "hour"  // Valid for Relative.Unit.
	Day   = "day"
	Week  = "week"
	Month = "month"
	Year  = "year"
)
