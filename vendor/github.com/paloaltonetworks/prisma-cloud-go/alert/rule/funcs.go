package rule

import (
	"fmt"
	"strings"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
)

// Identify returns the ID associated with the specified alert rule name.
func Identify(c pc.PrismaCloudClient, name string) (string, error) {
	c.Log(pc.LogAction, "(get) id for %s: %s", singular, name)

	list, err := List(c)
	if err != nil {
		return "", err
	}

	for _, o := range list {
		if o.Name == name {
			return o.PolicyScanConfigId, nil
		}
	}

	return "", pc.ObjectNotFoundError
}

// List returns a list of alerts that match the constraints specified.
func List(c pc.PrismaCloudClient) ([]Rule, error) {
	c.Log(pc.LogAction, "(get) list of %s", plural)

	var ans []Rule

	_, err := c.Communicate("GET", []string{"v2", "alert", "rule"}, nil, nil, &ans)
	return ans, err
}

// Get returns information about an alert for the specified ID.
func Get(c pc.PrismaCloudClient, id string) (Rule, error) {
	c.Log(pc.LogAction, "(get) %s: %s", singular, id)

	var ans Rule

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	path = append(path, id)

	_, err := c.Communicate("GET", path, nil, nil, &ans)
	return ans, err
}

// Create makes a new alert rule.
func Create(c pc.PrismaCloudClient, rule Rule) error {
	return createUpdate(false, c, rule)
}

// Update modifies information about the alert rule that has the specified ID.
func Update(c pc.PrismaCloudClient, rule Rule) error {
	return createUpdate(true, c, rule)
}

// Delete removes the alert rule that has the specified ID.
func Delete(c pc.PrismaCloudClient, id string) error {
	c.Log(pc.LogAction, "(delete) %s id:%s", singular, id)

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	path = append(path, id)

	_, err := c.Communicate("DELETE", path, nil, nil, nil)
	return err
}

func createUpdate(exists bool, c pc.PrismaCloudClient, rule Rule) error {
	var (
		logMsg strings.Builder
		method string
	)

	logMsg.Grow(30)
	logMsg.WriteString("(")
	if exists {
		logMsg.WriteString("update")
		method = "PUT"
	} else {
		logMsg.WriteString("create")
		method = "POST"
	}
	logMsg.WriteString(") ")

	logMsg.WriteString(" ")
	logMsg.WriteString(singular)
	if exists {
		fmt.Fprintf(&logMsg, ": %s", rule.PolicyScanConfigId)
	}

	c.Log(pc.LogAction, logMsg.String())

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	if exists {
		path = append(path, rule.PolicyScanConfigId)
	}

	_, err := c.Communicate(method, path, nil, rule, nil)
	return err
}
