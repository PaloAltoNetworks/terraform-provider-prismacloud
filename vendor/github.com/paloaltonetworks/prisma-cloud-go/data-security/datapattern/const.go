package datapattern

const (
	singular = "data pattern"
	plural   = "data patterns"
)

var Suffix = []string{"pcds", "config", "v3", "dss-api", "data-pattern", "dssTenantId"}

var TenantSuffix = []string{"api", "v1", "provision", "dlp", "status"}

var listBody = ListBody{ModeFilter: []string{"predefined", "custom"}}
