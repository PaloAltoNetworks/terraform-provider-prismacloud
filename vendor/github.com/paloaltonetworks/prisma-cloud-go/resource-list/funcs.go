package resource_list

import (
	"fmt"
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"strings"
)

func List(c pc.PrismaCloudClient) ([]ResourceList, error) {
	c.Log(pc.LogAction, "(get) list of %s", plural)

	var resourceLists []ResourceList
	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	if _, err := c.Communicate("GET", path, nil, nil, &resourceLists); err != nil {
		return nil, err
	}
	return resourceLists, nil
}

// Get returns the ResourceList.
func Get(c pc.PrismaCloudClient, id string) (ResourceList, error) {
	c.Log(pc.LogAction, "(get) %s id:%s", singular, id)
	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	path = append(path, id)
	var template ResourceList
	if _, err := c.Communicate("GET", path, nil, nil, &template); err != nil {
		return ResourceList{}, err
	}
	return template, nil
}

// Create adds a new ResourceList.
func Create(c pc.PrismaCloudClient, template ResourceListRequest) (ResourceList, error) {
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

// Update modifies the existing ResourceListTemplate.
func Update(c pc.PrismaCloudClient, template ResourceListRequest, id string) (ResourceList, error) {
	return createUpdate(true, c, template, id)
}
func createUpdate(exists bool, c pc.PrismaCloudClient, template ResourceListRequest, id string) (ResourceList, error) {
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
	var templateRes ResourceList
	var err error
	if exists {
		_, err = c.Communicate(method, path, nil, template, nil)
	} else {
		_, err = c.Communicate(method, path, nil, template, &templateRes)
	}
	return templateRes, err
}
