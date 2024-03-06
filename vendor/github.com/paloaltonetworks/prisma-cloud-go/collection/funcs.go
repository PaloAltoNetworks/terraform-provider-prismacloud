package collection

import (
	"fmt"
	"strings"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
)

func List(c pc.PrismaCloudClient) (Response, error) {
	c.Log(pc.LogAction, "(get) list of %s", plural)

	var res Response
	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	if _, err := c.Communicate("GET", path, nil, nil, &res); err != nil {
		return res, err
	}
	return res, nil
}

// Get returns the Collection.
func Get(c pc.PrismaCloudClient, id string) (Collection, error) {
	c.Log(pc.LogAction, "(get) %s id:%s", singular, id)
	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	path = append(path, id)
	var template Collection
	if _, err := c.Communicate("GET", path, nil, nil, &template); err != nil {
		return Collection{}, err
	}
	return template, nil
}

// Create adds a new ResourceList.
func Create(c pc.PrismaCloudClient, template CollectionRequest) (Collection, error) {
	return createUpdate(false, c, template, "")
}
func Delete(c pc.PrismaCloudClient, id string) error {
	c.Log(pc.LogAction, "(delete) %s id:%s", singular, id)
	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	path = append(path, id)
	_, err := c.Communicate("GET", path, nil, nil, nil)
	if err != nil {
		return err
	}
	_, err = c.Communicate("DELETE", path, nil, nil, nil)
	if err != nil {
		return err

	}
	return err
}

// Update modifies the existing CollectionTemplate.
func Update(c pc.PrismaCloudClient, template CollectionRequest, id string) (Collection, error) {
	return createUpdate(true, c, template, id)
}
func createUpdate(exists bool, c pc.PrismaCloudClient, template CollectionRequest, id string) (Collection, error) {
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
		fmt.Fprintf(&logMsg, ":%s", template.Name)
	}
	c.Log(pc.LogAction, logMsg.String())

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	if exists {
		path = append(path, id)
	}
	var templateRes Collection
	var err error
	if exists {
		_, err = c.Communicate(method, path, nil, template, nil)
	} else {
		_, err = c.Communicate(method, path, nil, template, &templateRes)
	}
	return templateRes, err
}
