package standard

import (
	"fmt"
	"strings"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
)

// List returns all system supported and custom compliance standards.
func List(c pc.PrismaCloudClient) ([]Standard, error) {
	c.Log(pc.LogAction, "(get) list of %s", plural)

	var ans []Standard
	if _, err := c.Communicate("GET", Suffix, nil, nil, &ans); err != nil {
		return nil, err
	}

	return ans, nil
}

// Identify returns the ID for the given compliance standard name.
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

// Get returns the compliance standard for the specified ID.
func Get(c pc.PrismaCloudClient, id string) (Standard, error) {
	c.Log(pc.LogAction, "(get) %s id:%s", singular, id)

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	path = append(path, id)

	var ans Standard
	_, err := c.Communicate("GET", path, nil, nil, &ans)
	return ans, err
}

// Create allows for the creation of a custom compliance standard.
func Create(c pc.PrismaCloudClient, cs Standard) error {
	return createUpdate(false, c, cs)
}

// Update updates an existing compliance standard.
func Update(c pc.PrismaCloudClient, cs Standard) error {
	return createUpdate(true, c, cs)
}

// Delete removes the compliance standard for the specified ID.
func Delete(c pc.PrismaCloudClient, id string) error {
	c.Log(pc.LogAction, "(delete) %s id:%s", singular, id)

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	path = append(path, id)
	_, err := c.Communicate("DELETE", path, nil, nil, nil)
	return err
}

func createUpdate(exists bool, c pc.PrismaCloudClient, cs Standard) error {
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
		fmt.Fprintf(&logMsg, ": %s", cs.Id)
	}

	c.Log(pc.LogAction, logMsg.String())

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	if exists {
		path = append(path, cs.Id)
	}

	_, err := c.Communicate(method, path, nil, cs, nil)
	return err
}
