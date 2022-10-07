package anomalySettings

import (
	"fmt"
	"net/url"
	"strings"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
)

// List returns a list of available anomalySettings settings
func List(c pc.PrismaCloudClient, t string) (map[string]interface{}, error) {
	c.Log(pc.LogAction, "(get) list of %s", plural)

	var ans map[string]interface{}
	var query url.Values
	if t != "" {
		query = url.Values{}
		query.Add("type", t)
	}
	if _, err := c.Communicate("GET", Suffix, query, nil, &ans); err != nil {
		return ans, err
	}

	return ans, nil
}

// Get returns anomaly settings for the specified policy ID
func Get(c pc.PrismaCloudClient, id string) (AnomalySettings, error) {
	c.Log(pc.LogAction, "(get) %s id:%s", singular, id)

	var ans AnomalySettings
	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	path = append(path, id)

	_, err := c.Communicate("GET", path, nil, nil, &ans)
	return ans, err
}

// Create modifies the existing anomaly setting.
func Create(c pc.PrismaCloudClient, anomalySettings AnomalySettings) error {
	return createUpdate(false, c, anomalySettings)
}

// Update modifies the existing anomaly setting.
func Update(c pc.PrismaCloudClient, anomalySettings AnomalySettings) error {
	return createUpdate(false, c, anomalySettings)
}

func createUpdate(exists bool, c pc.PrismaCloudClient, anomalySettings AnomalySettings) error {
	var (
		logMsg strings.Builder
		method string
	)

	logMsg.Grow(30)
	logMsg.WriteString("(")
	if exists {
		logMsg.WriteString("update")
		method = "POST"
	} else {
		logMsg.WriteString("create")
		method = "POST"
	}
	logMsg.WriteString(") ")

	logMsg.WriteString(singular)
	if exists {
		fmt.Fprintf(&logMsg, ":%s", anomalySettings.PolicyId)
	}

	c.Log(pc.LogAction, logMsg.String())

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	path = append(path, anomalySettings.PolicyId)
	_, err := c.Communicate(method, path, nil, anomalySettings, nil)
	return err
}
