package alert

import (
	"fmt"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
)

// List returns a list of alerts that match the constraints specified.
func List(c pc.PrismaCloudClient, req Request) ([]Alert, error) {
	c.Log(pc.LogAction, "(get) list of %s", plural)

	// Sanity check the time range.
	switch v := req.TimeRange.Value.(type) {
	case Absolute:
		req.TimeRange.Type = TimeAbsolute
	case Relative:
		req.TimeRange.Type = TimeRelative
	case ToNow:
		req.TimeRange.Type = TimeToNow
	case nil:
		return nil, fmt.Errorf("time range must be specified")
	default:
		return nil, fmt.Errorf("invalid time range type: %v", v)
	}

	var resp Response
	if _, err := c.Communicate("POST", []string{"v2", "alert"}, req, &resp, true); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

// Get returns information about an alert for the specified ID.
func Get(c pc.PrismaCloudClient, id string) (Alert, error) {
	c.Log(pc.LogAction, "(get) %s: %s", singular, id)

	var ans Alert
	if _, err := c.Communicate("GET", []string{"alert", id}, nil, &ans, true); err != nil {
		return ans, err
	}

	return ans, nil
}
