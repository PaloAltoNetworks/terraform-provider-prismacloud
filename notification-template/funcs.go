package notification_template

import (
	pc "github.com/paloaltonetworks/prisma-cloud-go"
)

func List(c pc.PrismaCloudClient) ([]NotificationTemplate, error) {
	c.Log(pc.LogAction, "(get) list of %s", plural)

	var ans []NotificationTemplate
	if _, err := c.Communicate("GET", Suffix, nil, nil, &ans); err != nil {
		return nil, err
	}

	return ans, nil
}
