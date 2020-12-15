package search

import (
	_ "net/url"
	_ "strconv"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
)

// ConfigSearch performs a config RQL search.
func ConfigSearch(c pc.PrismaCloudClient, req ConfigRequest) (ConfigResponse, error) {
	c.Log(pc.LogAction, "(get) performing %s", configSingular)

	var resp ConfigResponse

	// Sanity check the time range.
	if err := req.TimeRange.SetType(); err != nil {
		return resp, err
	}

	path := make([]string, 0, len(BaseSuffix)+len(ConfigSuffix))
	path = append(path, BaseSuffix...)
	path = append(path, ConfigSuffix...)

	_, err := c.Communicate("POST", path, nil, req, &resp)
	return resp, err
}
