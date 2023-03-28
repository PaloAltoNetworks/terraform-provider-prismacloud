package permission_group

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

	var ans []NameId
	if _, err := c.Communicate("GET", path, nil, nil, &ans); err != nil {
		return "", err
	}

	for _, o := range ans {
		if o.Name == name {
			return o.Id, nil
		}
	}

	return "", pc.ObjectNotFoundError
}
func List(c pc.PrismaCloudClient) ([]PermissionGroup, error) {
	c.Log(pc.LogAction, "(get) list of %s", plural)

	var ans []PermissionGroup
	if _, err := c.Communicate("GET", Suffix, nil, nil, &ans); err != nil {
		return nil, err
	}

	return ans, nil
}

// Get returns all information about an user role using its ID.
func Get(c pc.PrismaCloudClient, id string) (PermissionGroup, error) {
	c.Log(pc.LogAction, "(get) %s id:%s", singular, id)

	ans := PermissionGroup{}

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	path = append(path, id)
	if _, err := c.Communicate("GET", path, nil, nil, &ans); err != nil {
		return ans, err
	}

	return ans, nil
}

// Create makes a new user role on the Prisma Cloud platform.
func Create(c pc.PrismaCloudClient, obj PermissionGroup) error {
	return createUpdate(false, c, obj)
}

// Update modifies information related to an existing user role.
func Update(c pc.PrismaCloudClient, obj PermissionGroup) error {
	return createUpdate(true, c, obj)
}

// Delete removes an existing user role using its ID.
func Delete(c pc.PrismaCloudClient, id string) error {
	c.Log(pc.LogAction, "(delete) %s id:%s", singular, id)

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	path = append(path, id)
	_, err := c.Communicate("DELETE", path, nil, nil, nil)
	return err
}

func createUpdate(exists bool, c pc.PrismaCloudClient, obj PermissionGroup) error {
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
