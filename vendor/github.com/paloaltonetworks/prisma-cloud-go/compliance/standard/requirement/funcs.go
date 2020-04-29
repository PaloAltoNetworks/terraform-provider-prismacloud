package requirement

import (
	"fmt"
	"strings"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
)

// List returns a list of all compliance requirements for the specified compliance standard ID.
func List(c pc.PrismaCloudClient, cid string) ([]Requirement, error) {
	c.Log(pc.LogAction, "(get) list of %s", plural)

	var ans []Requirement
	if _, err := c.Communicate("GET", ComplianceSuffix(cid), nil, nil, &ans); err != nil {
		return nil, err
	}

	return ans, nil
}

// Identify returns the ID for the given compliance standard requirement name.
func Identify(c pc.PrismaCloudClient, cid, name string) (string, error) {
	c.Log(pc.LogAction, "(get) id for %s name:%s", singular, name)

	list, err := List(c, cid)
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

// Get returns the compliance requirement data for the specified requirements ID.
func Get(c pc.PrismaCloudClient, id string) (Requirement, error) {
	c.Log(pc.LogAction, "(get) %s id:%s", singular, id)

	var ans Requirement
	path := RequirementSuffix(id)

	_, err := c.Communicate("GET", path, nil, nil, &ans)
	return ans, err
}

// Create allows for the creation of a custom compliance standard.
func Create(c pc.PrismaCloudClient, req Requirement) error {
	return createUpdate(false, c, req)
}

// Update updates an existing compliance standard.
func Update(c pc.PrismaCloudClient, req Requirement) error {
	return createUpdate(true, c, req)
}

// Delete removes the compliance standard for the specified ID.
func Delete(c pc.PrismaCloudClient, id string) error {
	c.Log(pc.LogAction, "(delete) %s id:%s", singular, id)

	path := RequirementSuffix(id)
	_, err := c.Communicate("DELETE", path, nil, nil, nil)
	return err
}

func createUpdate(exists bool, c pc.PrismaCloudClient, req Requirement) error {
	var (
		logMsg strings.Builder
		method string
		path   []string
	)

	logMsg.Grow(30)
	logMsg.WriteString("(")
	if exists {
		logMsg.WriteString("update")
		method = "PUT"
		path = RequirementSuffix(req.Id)
	} else {
		logMsg.WriteString("create")
		method = "POST"
		path = ComplianceSuffix(req.ComplianceId)
	}
	logMsg.WriteString(") ")

	logMsg.WriteString(singular)
	if exists {
		fmt.Fprintf(&logMsg, ": %s", req.Id)
	}

	c.Log(pc.LogAction, logMsg.String())

	_, err := c.Communicate(method, path, nil, req, nil)
	return err
}
