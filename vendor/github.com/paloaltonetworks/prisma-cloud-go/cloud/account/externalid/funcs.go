package external_id

import (
	"fmt"
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"strings"
)

//// Identify returns the ID associated with the specified user role name.
//func Identify(c pc.PrismaCloudClient, name string) (string, error) {
//	c.Log(pc.LogAction, "(get) id for account name:%s", name)
//
//	list, err := List(c)
//	if err != nil {
//		return "", err
//	}
//
//	for _, o := range list {
//		if o.Name == name {
//			return o.Id, nil
//		}
//	}
//
//	return "", pc.ObjectNotFoundError
//}

// List returns the user roles.
//func List(c pc.PrismaCloudClient) ([]Role, error) {
//	c.Log(pc.LogAction, "(get) list of %s", plural)
//
//	var ans []Role
//	if _, err := c.Communicate("GET", Suffix, nil, nil, &ans); err != nil {
//		return nil, err
//	}
//
//	return ans, nil
//}

// Get returns all information about an user role using its ID.
//func Get(c pc.PrismaCloudClient, id string) (ExternalId, error) {
//	c.Log(pc.LogAction, "(get) %s id:%s", singular, id)
//
//	ans := Role{}
//
//	path := make([]string, 0, len(Suffix)+1)
//	path = append(path, Suffix...)
//	path = append(path, id)
//
//	if _, err := c.Communicate("GET", path, nil, nil, &ans); err != nil {
//		return ans, err
//	}
//
//	return ans, nil
//}

func createUpdate(exists bool, c pc.PrismaCloudClient, obj ExternalId) error {
	var (
		logMsg strings.Builder
		method string
	)

	logMsg.Grow(30)
	logMsg.WriteString("(")

	logMsg.WriteString("create")
	method = "POST"

	logMsg.WriteString(") ")

	logMsg.WriteString(" ")
	logMsg.WriteString("external id")
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
