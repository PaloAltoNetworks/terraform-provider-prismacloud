package dataprofile

import (
	"fmt"
	"strings"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
)

// List returns a list of available custom data profiles.
func List(c pc.PrismaCloudClient) ([]Profile, error) {
	c.Log(pc.LogAction, "(get) list of %s", plural)

	var ans ListBody

	_, err := c.Communicate("GET", Suffix, nil, nil, &ans)
	return ans.Profiles, err
}

// Identify returns the ID for the given data profile.
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

// Get returns the data profile that has the specified ID.
func Get(c pc.PrismaCloudClient, id string) (Profile, error) {
	c.Log(pc.LogAction, "(get) %s id:%s", singular, id)

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	path = append(path, "id")
	path = append(path, id)

	var ans Profile
	_, err := c.Communicate("GET", path, nil, nil, &ans)

	return ans, err
}

// Create adds a new data profile.
func Create(c pc.PrismaCloudClient, profile Profile) error {
	return createUpdate(false, c, profile)
}

// Update modifies the existing data profile.
func Update(c pc.PrismaCloudClient, profile Profile) error {
	return createUpdate(true, c, profile)
}

// Delete removes a data profile using its ID.
func Delete(c pc.PrismaCloudClient, id string) error {
	c.Log(pc.LogAction, "(delete) %s id:%s", singular, id)

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	path = append(path, "id")
	path = append(path, id)
	_, err := c.Communicate("DELETE", path, nil, nil, nil)
	return err
}

func createUpdate(exists bool, c pc.PrismaCloudClient, profile Profile) error {
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
		fmt.Fprintf(&logMsg, ":%s", profile.Id)
	}

	c.Log(pc.LogAction, logMsg.String())

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	if exists {
		path = append(path, "id")
		path = append(path, profile.Id)
	}

	_, err := c.Communicate(method, path, nil, profile, nil)
	return err
}
