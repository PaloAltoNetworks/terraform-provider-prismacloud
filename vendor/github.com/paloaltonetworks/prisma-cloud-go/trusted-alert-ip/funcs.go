package trustedalertip

import (
	"fmt"
	"strings"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
)

// Identify returns the ID for the given account group.
func Identify(c pc.PrismaCloudClient, name string) (string, error) {
	c.Log(pc.LogAction, "(get) id for %s: %s", singular, name)

	listing, err := List(c)
	if err != nil {
		return "", err
	}

	for _, o := range listing {
		if o.Name == name {
			return o.UUID, nil
		}
	}

	return "", pc.ObjectNotFoundError
}

// List returns a list of available users and service accounts
func List(c pc.PrismaCloudClient) ([]TrustedAlertIP, error) {
	c.Log(pc.LogAction, "(get) list of %s", plural)

	var ans []TrustedAlertIP
	if _, err := c.Communicate("GET", Suffix, nil, nil, &ans); err != nil {
		return nil, err
	}

	return ans, nil
}

func Get(c pc.PrismaCloudClient, id string) (TrustedAlertIP, error) {
	c.Log(pc.LogAction, "(get) %s id:%s", singular, id)

	var ans TrustedAlertIP

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	path = append(path, id)

	_, err := c.Communicate("GET", path, nil, nil, &ans)

	return ans, err
}

func Create(c pc.PrismaCloudClient, trustedAlertIP TrustedAlertIP) ([]byte, error) {
	return createUpdate(false, c, trustedAlertIP)
}

func Update(c pc.PrismaCloudClient, trustedAlertIP TrustedAlertIP) ([]byte, error) {
	return createUpdate(true, c, trustedAlertIP)
}

func Delete(c pc.PrismaCloudClient, id string, uuid string) error {
	c.Log(pc.LogAction, "(delete) %s id:%s", singular, id)

	path := make([]string, 0, len(Suffix)+5)
	path = append(path, Suffix...)
	path = append(path, id)
	path = append(path, "cidr")
	path = append(path, uuid)
	_, err := c.Communicate("DELETE", path, nil, nil, nil)
	return err
}

func createUpdate(exists bool, c pc.PrismaCloudClient, trustedAlertIP TrustedAlertIP) ([]byte, error) {
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
		fmt.Fprintf(&logMsg, ":%s", trustedAlertIP.UUID)
	}

	c.Log(pc.LogAction, logMsg.String())

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)

	response, err := c.Communicate(method, path, nil, trustedAlertIP, nil)
	return response, err
}

func CreateCIDR(c pc.PrismaCloudClient, trustedAlertIP CIDRS, id string) ([]byte, error) {
	var (
		logMsg strings.Builder
		method string
	)
	method = "POST"
	logMsg.Grow(30)
	logMsg.WriteString("(")
	logMsg.WriteString(method)
	logMsg.WriteString(") ")
	logMsg.WriteString(singular)

	c.Log(pc.LogAction, logMsg.String())

	path := make([]string, 0, len(Suffix)+4)
	path = append(path, Suffix...)
	path = append(path, id)
	path = append(path, "cidr")
	response, err := c.Communicate(method, path, nil, trustedAlertIP, nil)
	return response, err
}

func UpdateCIDR(c pc.PrismaCloudClient, trustedAlertIP CIDRS, id string, uuid string) ([]byte, error) {
	var (
		logMsg strings.Builder
		method string
	)
	method = "PUT"

	c.Log(pc.LogAction, logMsg.String())

	path := make([]string, 0, len(Suffix)+5)
	path = append(path, Suffix...)
	path = append(path, id)
	path = append(path, "cidr")
	path = append(path, uuid)
	response, err := c.Communicate(method, path, nil, trustedAlertIP, nil)
	return response, err
}

func DeleteCIDRFromTrustedAlertIp(c pc.PrismaCloudClient, id string, uuid string) ([]byte, error) {
	var (
		logMsg strings.Builder
		method string
	)
	method = "DELETE"

	c.Log(pc.LogAction, logMsg.String())

	path := make([]string, 0, len(Suffix)+5)
	path = append(path, Suffix...)
	path = append(path, id)
	path = append(path, "cidr")
	path = append(path, uuid)
	response, err := c.Communicate(method, path, nil, nil, nil)
	return response, err
}
