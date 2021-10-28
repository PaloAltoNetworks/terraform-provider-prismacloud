package datapattern

import (
	"fmt"
	"strings"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
)

// List returns a list of available data patterns
func List(c pc.PrismaCloudClient) ([]Pattern, error) {
	c.Log(pc.LogAction, "(get) list of %s", plural)

	var ans ListBody

	_, err := c.Communicate("PUT", Suffix, nil, listBody, &ans)
	return ans.Patterns, err
}

// Identify returns the ID for the given data pattern.
func Identify(c pc.PrismaCloudClient, name string) (string, error) {
	c.Log(pc.LogAction, "(get) id for %s name:%s", singular, name)

	list, err := List(c)
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

// Get returns the data pattern that has the specified ID.
func Get(c pc.PrismaCloudClient, id string) (Pattern, error) {
	c.Log(pc.LogAction, "(get) %s id:%s", singular, id)

	var ans Pattern
	list, err := List(c)
	if err != nil {
		return ans, err
	}

	for _, o := range list {
		if o.Id == id {
			ans = o
		}
	}
	return ans, err
}

// Create adds a new data pattern.
func Create(c pc.PrismaCloudClient, pattern Pattern) error {
	return createUpdate(false, c, pattern)
}

// Update modifies the existing data pattern.
func Update(c pc.PrismaCloudClient, pattern Pattern) error {
	return createUpdate(true, c, pattern)
}

// Delete removes a data pattern using its ID.
func Delete(c pc.PrismaCloudClient, id string) error {
	c.Log(pc.LogAction, "(delete) %s id:%s", singular, id)

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	path = append(path, id)
	_, err := c.Communicate("DELETE", path, nil, nil, nil)
	return err
}

func createUpdate(exists bool, c pc.PrismaCloudClient, pattern Pattern) error {
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
		fmt.Fprintf(&logMsg, ":%s", pattern.Id)
	}

	c.Log(pc.LogAction, logMsg.String())

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	if exists {
		path = append(path, pattern.Id)
	}

	_, err := c.Communicate(method, path, nil, pattern, nil)
	return err
}
