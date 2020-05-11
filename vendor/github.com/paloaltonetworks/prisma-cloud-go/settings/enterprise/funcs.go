package enterprise

import (
	pc "github.com/paloaltonetworks/prisma-cloud-go"
)

// Get returns the enterprise settings for your tenant.
func Get(c pc.PrismaCloudClient) (Config, error) {
	c.Log(pc.LogAction, "(get) %s", singular)

	path := Suffix[:]

	var ans Config
	_, err := c.Communicate("GET", path, nil, nil, &ans)
	return ans, err
}

// Update configures enterprise settings for your tenant.
func Update(c pc.PrismaCloudClient, conf Config) error {
	c.Log(pc.LogAction, "(update) %s", singular)

	path := Suffix[:]
	_, err := c.Communicate("POST", path, nil, conf, nil)
	return err
}
