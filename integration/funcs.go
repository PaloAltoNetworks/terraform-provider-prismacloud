package integration

import (
	"fmt"
	"net/url"
	"strings"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
)

// Types returns a list of all supported integration types.
func Types(c pc.PrismaCloudClient) ([]string, error) {
	c.Log(pc.LogAction, "(get) list of %s types", singular)

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	path = append(path, "type")

	var ans []string
	_, err := c.Communicate("GET", path, nil, nil, &ans)
	return ans, err
}

// List returns all your integrations, optionally filtered by type.
func List(c pc.PrismaCloudClient, t string) ([]Integration, error) {
	c.Log(pc.LogAction, "(get) list of %s", plural)

	var query url.Values
	if t != "" {
		query = url.Values{}
		query.Add("type", t)
	}

	var ans []Integration
	if _, err := c.Communicate("GET", Suffix, query, nil, &ans); err != nil {
		return nil, err
	}

	return ans, nil
}

// Identify returns the ID for the given integration name.
func Identify(c pc.PrismaCloudClient, name string) (string, error) {
	c.Log(pc.LogAction, "(get) id for %s: %s", singular, name)

	list, err := List(c, "")
	if err != nil {
		return "", err
	}

	for _, o := range list {
		if o.Name == name {
			return o.Id, nil
		}
	}

	return "", pc.ObjectNotFoundError
}

// Get returns integration details for the specified ID.
func Get(c pc.PrismaCloudClient, id string) (Integration, error) {
	c.Log(pc.LogAction, "(get) %s: %s", singular, id)

	var ans Integration
	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	path = append(path, id)

	_, err := c.Communicate("GET", path, nil, nil, &ans)
	return ans, err
}

// Create adds an integration with the specified external system.
func Create(c pc.PrismaCloudClient, obj Integration) error {
	return createUpdate(false, c, obj)
}

// Update modifies the specified integration.
func Update(c pc.PrismaCloudClient, obj Integration) error {
	return createUpdate(true, c, obj)
}

// Delete removes the integration for the specified ID.
func Delete(c pc.PrismaCloudClient, id string) error {
	c.Log(pc.LogAction, "(delete) %s: %s", singular, id)

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	path = append(path, id)

	_, err := c.Communicate("DELETE", path, nil, nil, nil)
	return err
}

func createUpdate(exists bool, c pc.PrismaCloudClient, obj Integration) error {
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

	logMsg.WriteString(singular)
	if exists {
		fmt.Fprintf(&logMsg, ": %s", obj.Id)
	}

	c.Log(pc.LogAction, logMsg.String())

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	if exists {
		path = append(path, obj.Id)
	}

	_, err := c.Communicate(method, path, nil, obj, nil)
	return err
}
