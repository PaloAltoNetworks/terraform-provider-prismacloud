package ip_address

import (
	"fmt"
	"strings"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
)

// List lists accessible Login-Ip-Allow List.
func List(c pc.PrismaCloudClient) ([]LoginIpAllow, error) {
	c.Log(pc.LogAction, "(get) list of %s", plural)

	var ans []LoginIpAllow
	if _, err := c.Communicate("GET", Suffix, nil, nil, &ans); err != nil {
		return nil, err
	}

	return ans, nil
}

// Get returns all information about an Login-Ip-Allow List using its ID.
func Get(c pc.PrismaCloudClient, id string) (LoginIpAllow, error) {

	var ans LoginIpAllow
	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	path = append(path, id)

	c.Log(pc.LogAction, "(get) %s: %s path: %s", singular, id, path)
	_, err := c.Communicate("GET", path, nil, nil, &ans)
	return ans, err
}

// Identify returns the ID for the given Login-Ip-Allow.
func Identify(c pc.PrismaCloudClient, name string) (string, error) {
	c.Log(pc.LogAction, "(get) id for %s: %s api: %s", singular, name, Suffix)

	var ans []LoginIpAllow
	if _, err := c.Communicate("GET", Suffix, nil, nil, &ans); err != nil {
		return "", err
	}

	for _, o := range ans {
		if o.Name == name {
			return o.Id, nil
		}
	}

	return "", pc.ObjectNotFoundError
}

// Create makes a new account LoginIpAllow on the Prisma Cloud platform.
func Create(c pc.PrismaCloudClient, loginIpAllow LoginIpAllow) error {
	return createUpdate(false, c, loginIpAllow)
}

// Update modifies information related to an existing account LoginIpAllow.
func Update(c pc.PrismaCloudClient, loginIpAllow LoginIpAllow) error {
	return createUpdate(true, c, loginIpAllow)
}

// Delete removes an existing account LoginIpAllow using its ID.
func Delete(c pc.PrismaCloudClient, id string) error {
	c.Log(pc.LogAction, "(delete) %s: %s", singular, id)

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	path = append(path, id)

	_, err := c.Communicate("DELETE", path, nil, nil, nil)
	return err
}

func createUpdate(exists bool, c pc.PrismaCloudClient, loginIpAllow LoginIpAllow) error {
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
		fmt.Fprintf(&logMsg, ": %s", loginIpAllow.Id)
	}

	c.Log(pc.LogAction, "Create Resource", logMsg.String())

	path := make([]string, 0, len(Suffix)+2)
	path = append(path, Suffix...)
	if exists {
		path = append(path, loginIpAllow.Id)
	}

	_, err := c.Communicate(method, path, nil, loginIpAllow, nil)
	return err
}
