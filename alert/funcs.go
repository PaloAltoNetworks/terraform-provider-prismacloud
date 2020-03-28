package alert

import (
	"fmt"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
)

// List returns a list of alerts that match the constraints specified.
func List(c pc.PrismaCloudClient, req Request) (Response, error) {
	c.Log(pc.LogAction, "(get) list of %s", plural)

	var resp Response

	// Sanity check the time range.
	switch v := req.TimeRange.Value.(type) {
	case Absolute:
		req.TimeRange.Type = TimeAbsolute
	case Relative:
		req.TimeRange.Type = TimeRelative
	case ToNow:
		req.TimeRange.Type = TimeToNow
	case nil:
		return resp, fmt.Errorf("time range must be specified")
	default:
		return resp, fmt.Errorf("invalid time range type: %v", v)
	}

	_, err := c.Communicate("POST", []string{"v2", "alert"}, req, &resp, true)
	return resp, err
}

// Get returns information about an alert for the specified ID.
func Get(c pc.PrismaCloudClient, id string) (Alert, error) {
	c.Log(pc.LogAction, "(get) %s: %s", singular, id)

	var ans Alert

	_, err := c.Communicate("GET", []string{"alert", id}, nil, &ans, true)
	return ans, err
}
