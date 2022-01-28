package profile

import (
	"fmt"
	"net/url"
	"strings"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
)

// List returns a list of available users and service accounts
func List(c pc.PrismaCloudClient) ([]Profile, error) {
	c.Log(pc.LogAction, "(get) list of %s", plural)

	var ans []Profile
	if _, err := c.Communicate("GET", Suffix, nil, nil, &ans); err != nil {
		return nil, err
	}

	return ans, nil
}

// Get returns the user profile or service account profile that has the specified ID.
func Get(c pc.PrismaCloudClient, id string) (Profile, error) {
	c.Log(pc.LogAction, "(get) %s id:%s", singular, id)

	var ans Profile

	list, err := List(c)
	if err != nil {
		return ans, err
	}

	for _, o := range list {
		if o.Username == id {
			return o, nil
		}
	}

	return ans, pc.ObjectNotFoundError
}

// Create adds a user profile or a service account profile.
func Create(c pc.PrismaCloudClient, profile Profile) ([]byte, error) {
	return createUpdate(false, c, profile)
}

// Update modifies the existing user profile.
func Update(c pc.PrismaCloudClient, profile Profile) ([]byte, error) {
	return createUpdate(true, c, profile)
}

// Delete removes a user profile or service account profile using its ID.
func Delete(c pc.PrismaCloudClient, id string, accountType string) error {
	c.Log(pc.LogAction, "(delete) %s id:%s", singular, id)

	path := make([]string, 0, len(Suffix))
	if accountType == TypeServiceAccount {
		path = append(path, Suffix[1], id)
	} else {
		path = append(path, Suffix[1], url.QueryEscape(id))
	}

	_, err := c.Communicate("DELETE", path, nil, nil, nil)
	return err
}

func createUpdate(exists bool, c pc.PrismaCloudClient, profile Profile) ([]byte, error) {
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

	if exists {
		path = append(path, "v2", Suffix[1], url.QueryEscape(profile.Id))
	} else {
		path = append(path, Suffix...)
	}

	response, err := c.Communicate(method, path, nil, profile, nil)
	return response, err
}
