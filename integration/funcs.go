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

// GetPrismaId returns prisma id.
func GetPrismaId(c pc.PrismaCloudClient) (string, error) {
	c.Log(pc.LogAction, "(get) prisma id")

	var ans LicenseInfo

	path := make([]string, 0, len(LicenseSuffix)+1)
	path = append(path, LicenseSuffix...)

	_, err := c.Communicate("GET", path, nil, nil, &ans)
	return ans.PrismaId, err
}

// List returns all your integrations, optionally filtered by type.
func List(c pc.PrismaCloudClient, t string, prismaIdRequired bool) ([]Integration, error) {
	c.Log(pc.LogAction, "(get) list of %s", plural)

	var query url.Values
	if t != "" {
		query = url.Values{}
		query.Add("type", t)
	}

	var ans []Integration

	path := make([]string, 0, len(v1Suffix)+len(Suffix)+1)
	if prismaIdRequired {
		prismaId, err := GetPrismaId(c)
		if err != nil {
			return nil, err
		}

		path = append(path, v1Suffix...)
		path = append(path, prismaId)
	}

	path = append(path, Suffix...)
	if _, err := c.Communicate("GET", path, query, nil, &ans); err != nil {
		return nil, err
	}

	return ans, nil
}

// Identify returns the ID for the given integration name.
func Identify(c pc.PrismaCloudClient, name string, prismaIdRequired bool) (string, error) {
	c.Log(pc.LogAction, "(get) id for %s: %s", singular, name)

	list, err := List(c, "", prismaIdRequired)
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
func Get(c pc.PrismaCloudClient, id string, prismaIdRequired bool) (Integration, error) {
	c.Log(pc.LogAction, "(get) %s: %s", singular, id)

	var ans Integration

	path := make([]string, 0, len(v1Suffix)+len(Suffix)+1)
	if prismaIdRequired {
		prismaId, err := GetPrismaId(c)
		if err != nil {
			return ans, err
		}

		path = append(path, v1Suffix...)
		path = append(path, prismaId)
	}

	path = append(path, Suffix...)
	path = append(path, id)

	_, err := c.Communicate("GET", path, nil, nil, &ans)
	return ans, err
}

// Create adds an integration with the specified external system.
func Create(c pc.PrismaCloudClient, obj Integration, prismaIdRequired bool) error {
	return createUpdate(false, c, obj, prismaIdRequired)
}

// Update modifies the specified integration.
func Update(c pc.PrismaCloudClient, obj Integration, prismaIdRequired bool) error {
	return createUpdate(true, c, obj, prismaIdRequired)
}

// Delete removes the integration for the specified ID.
func Delete(c pc.PrismaCloudClient, id string, prismaIdRequired bool) error {
	c.Log(pc.LogAction, "(delete) %s: %s", singular, id)

	path := make([]string, 0, len(v1Suffix)+len(Suffix)+1)
	if prismaIdRequired {
		prismaId, err := GetPrismaId(c)
		if err != nil {
			return err
		}

		path = append(path, v1Suffix...)
		path = append(path, prismaId)
	}

	path = append(path, Suffix...)
	path = append(path, id)

	_, err := c.Communicate("DELETE", path, nil, nil, nil)
	return err
}

func createUpdate(exists bool, c pc.PrismaCloudClient, obj Integration, prismaIdRequired bool) error {
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

	path := make([]string, 0, len(v1Suffix)+len(Suffix)+1)
	if prismaIdRequired {
		prismaId, err := GetPrismaId(c)
		if err != nil {
			return err
		}

		path = append(path, v1Suffix...)
		path = append(path, prismaId)
	}

	path = append(path, Suffix...)
	if exists {
		path = append(path, obj.Id)
	}

	_, err := c.Communicate(method, path, nil, obj, nil)
	return err
}
