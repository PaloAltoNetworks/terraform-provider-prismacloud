package section

import (
	"fmt"
	"strings"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
)

// List returns a list of all compliance requirements sections for the specified compliance requirement ID.
func List(c pc.PrismaCloudClient, rid string) ([]Section, error) {
	c.Log(pc.LogAction, "(get) list of %s", plural)

	var ans []Section
	if _, err := c.Communicate("GET", RequirementSuffix(rid), nil, nil, &ans); err != nil {
		return nil, err
	}

	return ans, nil
}

// GetId returns the compliance requirement section data for the specified ID.
func GetId(c pc.PrismaCloudClient, rid, id string) (Section, error) {
	list, err := List(c, rid)
	if err != nil {
		return Section{}, err
	}

	for _, o := range list {
		if o.Id == id {
			return o, nil
		}
	}

	return Section{}, pc.ObjectNotFoundError
}

// Get returns the compliance requirement section data for the specified requirements section ID.
func Get(c pc.PrismaCloudClient, rid, sectionId string) (Section, error) {
	list, err := List(c, rid)
	if err != nil {
		return Section{}, err
	}

	for _, o := range list {
		if o.SectionId == sectionId {
			return o, nil
		}
	}

	return Section{}, pc.ObjectNotFoundError
}

// Create allows for the creation of a custom compliance standard.
func Create(c pc.PrismaCloudClient, s Section) error {
	return createUpdate(false, c, s)
}

// Update updates an existing compliance standard.
func Update(c pc.PrismaCloudClient, s Section) error {
	return createUpdate(true, c, s)
}

// Delete removes the compliance standard for the specified ID.
func Delete(c pc.PrismaCloudClient, id string) error {
	c.Log(pc.LogAction, "(delete) %s id:%s", singular, id)

	path := SectionSuffix(id)
	_, err := c.Communicate("DELETE", path, nil, nil, nil)
	return err
}

func createUpdate(exists bool, c pc.PrismaCloudClient, s Section) error {
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
		path = SectionSuffix(s.Id)
	} else {
		logMsg.WriteString("create")
		method = "POST"
		path = RequirementSuffix(s.RequirementId)
	}
	logMsg.WriteString(") ")

	logMsg.WriteString(singular)
	if exists {
		fmt.Fprintf(&logMsg, ": %s", s.Id)
	}

	c.Log(pc.LogAction, logMsg.String())

	_, err := c.Communicate(method, path, nil, s, nil)
	return err
}
