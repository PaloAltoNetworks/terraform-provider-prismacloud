package anomalyTrustedList

import (
	"fmt"
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"strconv"
	"strings"
)

// List returns a list of all entries in the anomaly trusted list.
func List(c pc.PrismaCloudClient) ([]AnomalyTrustedList, error) {
	c.Log(pc.LogAction, "(get) list of %s", plural)

	var ans []AnomalyTrustedList

	_, err := c.Communicate("GET", Suffix, nil, nil, &ans)
	return ans, err
}

// Identify returns the ID for the given anomaly trusted list.
func Identify(c pc.PrismaCloudClient, id string) (string, error) {
	c.Log(pc.LogAction, "(get) id for %s id:%s", singular, id)

	list, err := List(c)
	if err != nil {
		return "", err
	}

	for _, o := range list {
		if strconv.Itoa(o.Atl_Id) == id {
			return id, nil
		}
	}

	return "", pc.ObjectNotFoundError
}

// Get returns anomaly trusted list for the specified ID
func Get(c pc.PrismaCloudClient, id string) (AnomalyTrustedList, error) {
	c.Log(pc.LogAction, "(get) %s id:%s", singular, id)
	var ans AnomalyTrustedList
	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	path = append(path, id)
	_, err := c.Communicate("GET", path, nil, nil, &ans)
	return ans, err
}

// Create creates a new anomaly trusted list.
func Create(c pc.PrismaCloudClient, anomalyTrustedList AnomalyTrustedList) (int, error) {
	return createUpdate(false, c, anomalyTrustedList)
}

// Update modifies information related to an anomaly trusted list
func Update(c pc.PrismaCloudClient, anomalyTrustedList AnomalyTrustedList) (int, error) {
	return createUpdate(true, c, anomalyTrustedList)
}

// Delete removes an anomaly trusted list using its ID.
func Delete(c pc.PrismaCloudClient, id string) error {
	c.Log(pc.LogAction, "(delete) %s id:%s", singular, id)

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	path = append(path, id)
	_, err := c.Communicate("DELETE", path, nil, nil, nil)
	return err
}

func createUpdate(exists bool, c pc.PrismaCloudClient, anomalyTrustedList AnomalyTrustedList) (int, error) {
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

	id := strconv.Itoa(anomalyTrustedList.Atl_Id)
	if exists {
		fmt.Fprintf(&logMsg, ":%d", anomalyTrustedList.Atl_Id)
	}

	c.Log(pc.LogAction, logMsg.String())

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	if exists {
		path = append(path, id)
	}

	var ans [1]int //returns integer as response if succeeds
	_, err := c.Communicate(method, path, nil, anomalyTrustedList, &ans)
	return ans[0], err
}
