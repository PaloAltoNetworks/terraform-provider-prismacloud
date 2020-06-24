package alert

import (
	pc "github.com/paloaltonetworks/prisma-cloud-go"
)

// List returns a list of alerts that match the constraints specified.
func List(c pc.PrismaCloudClient, req Request) (Response, error) {
	c.Log(pc.LogAction, "(get) list of %s", plural)

	var resp Response

	// Sanity check the time range.
	if err := req.TimeRange.SetType(); err != nil {
		return resp, err
	}

	_, err := c.Communicate("POST", []string{"v2", "alert"}, nil, req, &resp)
	return resp, err
}

// Get returns information about an alert for the specified ID.
func Get(c pc.PrismaCloudClient, id string) (Alert, error) {
	c.Log(pc.LogAction, "(get) %s: %s", singular, id)

	var ans Alert

	_, err := c.Communicate("GET", []string{"alert", id}, nil, nil, &ans)
	return ans, err
}
