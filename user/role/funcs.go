package role

import (
	"fmt"
	"strings"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
)

// Identify returns the ID associated with the specified user role name.
func Identify(c pc.PrismaCloudClient, name string) (string, error) {
	c.Log(pc.LogAction, "(get) id for %s: %s", singular, name)

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	path = append(path, "name")

	var ans []NameId
	if _, err := c.Communicate("GET", path, nil, &ans, true); err != nil {
		return "", err
	}

	for _, o := range ans {
		if o.Name == name {
			return o.Id, nil
		}
	}

	return "", pc.ObjectNotFoundError
}

// List returns the user roles.
func List(c pc.PrismaCloudClient) ([]Role, error) {
	c.Log(pc.LogAction, "(get) list of %s", plural)

	var ans []Role
	if _, err := c.Communicate("GET", Suffix, nil, &ans, true); err != nil {
		return nil, err
	}

	return ans, nil
}

// Get returns all information about an user role using its ID.
func Get(c pc.PrismaCloudClient, id string) (Role, error) {
	c.Log(pc.LogAction, "(get) %s id:%s", singular, id)

	ans := Role{}

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	path = append(path, id)

	if _, err := c.Communicate("GET", path, nil, &ans, true); err != nil {
		return ans, err
	}

	return ans, nil
}

// Update modifies information related to an existing user role.
func Update(c pc.PrismaCloudClient, role Role) error {
	return createUpdate(false, c, role)
}

// Create makes a new user role on the Prisma Cloud platform.
func Create(c pc.PrismaCloudClient, role Role) error {
	return createUpdate(true, c, role)
}

// Delete removes an existing user role using its ID.
func Delete(c pc.PrismaCloudClient, id string) error {
	c.Log(pc.LogAction, "(delete) %s id:%s", singular, id)

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	path = append(path, id)
	_, err := c.Communicate("DELETE", path, nil, nil, true)
	return err
}

func createUpdate(exists bool, c pc.PrismaCloudClient, role Role) error {
	var (
		logMsg strings.Builder
		id     string
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
		fmt.Fprintf(&logMsg, ": %s", id)
	}

	c.Log(pc.LogAction, logMsg.String())

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	if exists {
		path = append(path, id)
	}

	_, err := c.Communicate(method, path, role, nil, true)
	return err
}
