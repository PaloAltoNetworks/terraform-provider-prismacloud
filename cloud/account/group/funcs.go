package group

import (
	"fmt"
	"strings"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
)

// List lists accessible account groups.
func List(c pc.PrismaCloudClient) ([]Group, error) {
	c.Log(pc.LogAction, "(get) list of %s", plural)

	var ans []Group
	if _, err := c.Communicate("GET", Suffix, nil, nil, &ans); err != nil {
		return nil, err
	}

	return ans, nil
}

// Identify returns the ID for the given account group.
func Identify(c pc.PrismaCloudClient, name string) (string, error) {
	c.Log(pc.LogAction, "(get) id for %s: %s", singular, name)

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	path = append(path, "name")

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

// Get returns all information about an account group using its ID.
func Get(c pc.PrismaCloudClient, id string) (Group, error) {
	c.Log(pc.LogAction, "(get) %s: %s", singular, id)

	var ans Group
	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	path = append(path, id)

	_, err := c.Communicate("GET", path, nil, nil, &ans)
	return ans, err
}

// Create makes a new account group on the Prisma Cloud platform.
func Create(c pc.PrismaCloudClient, group Group) error {
	return createUpdate(false, c, group)
}

// UpdateUsingLiveAccountIds copies the AccountIds param from what's
// live, but otherwise with the group definition provided.  The problem is
// that the API endpoint can change the accounts associated with this group,
// but in Terraform the group membership is a read-only attribute.
func UpdateUsingLiveAccountIds(c pc.PrismaCloudClient, group Group) error {
	lg, err := Get(c, group.Id)
	if err != nil {
		return err
	}
	group.AccountIds = lg.AccountIds

	return Update(c, group)
}

// Update modifies information related to an existing account group.
func Update(c pc.PrismaCloudClient, group Group) error {
	return createUpdate(true, c, group)
}

// Delete removes an existing account group using its ID.
func Delete(c pc.PrismaCloudClient, id string) error {
	c.Log(pc.LogAction, "(delete) %s: %s", singular, id)

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	path = append(path, id)

	_, err := c.Communicate("DELETE", path, nil, nil, nil)
	return err
}

func createUpdate(exists bool, c pc.PrismaCloudClient, group Group) error {
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
		fmt.Fprintf(&logMsg, ": %s", group.Id)
	}

	c.Log(pc.LogAction, logMsg.String())

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	if exists {
		path = append(path, group.Id)
	}

	_, err := c.Communicate(method, path, nil, group, nil)
	return err
}
