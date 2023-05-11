package notification_template

import (
	"fmt"
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"strings"
)

func List(c pc.PrismaCloudClient) ([]NotificationTemplate, error) {
	c.Log(pc.LogAction, "(get) list of %s", plural)

	var templates []NotificationTemplate
	path := make([]string, 0, len(Suffix1)+len(Suffix2)+1)
	path = append(path, Suffix1...)
	prismaId, err := GetPrismaId(c)
	if err != nil {
		return nil, nil
	}
	path = append(path, prismaId)
	path = append(path, Suffix2...)
	if _, err := c.Communicate("GET", path, nil, nil, &templates); err != nil {
		return nil, err
	}
	return templates, nil
}

// Get returns the NotificationTemplate.
func Get(c pc.PrismaCloudClient, id string) (NotificationTemplate, error) {
	c.Log(pc.LogAction, "(get) %s id:%s", singular, id)
	path := make([]string, 0, len(Suffix1)+len(Suffix2)+1)
	path = append(path, Suffix1...)
	prismaId, err := GetPrismaId(c)
	if err != nil {
		return NotificationTemplate{}, nil
	}
	path = append(path, prismaId)
	path = append(path, Suffix2...)
	path = append(path, id)
	var template NotificationTemplate
	if _, err := c.Communicate("GET", path, nil, nil, &template); err != nil {
		return NotificationTemplate{}, nil
	}
	return template, nil
}

// Create adds a new NotificationTemplate.
func Create(c pc.PrismaCloudClient, template NotificationTemplateRequest) (NotificationTemplate, error) {
	return createUpdate(false, c, template, "")
}
func Delete(c pc.PrismaCloudClient, id string) error {
	c.Log(pc.LogAction, "(delete) %s id:%s", singular, id)
	path := make([]string, 0, len(Suffix1)+len(Suffix2)+1)
	path = append(path, Suffix1...)

	prismaId, err := GetPrismaId(c)
	if err != nil {
		return nil
	}
	path = append(path, prismaId)
	path = append(path, Suffix2...)
	path = append(path, id)
	_, err = c.Communicate("GET", path, nil, nil, nil)

	if err != nil {
		return err
	}
	_, err = c.Communicate("DELETE", path, nil, nil, nil)
	if err != nil {
		return err

	}
	return err
}

// Update modifies the existing NotificationTemplate.
func Update(c pc.PrismaCloudClient, template NotificationTemplateRequest, id string) (NotificationTemplate, error) {
	return createUpdate(true, c, template, id)
}
func createUpdate(exists bool, c pc.PrismaCloudClient, template NotificationTemplateRequest, id string) (NotificationTemplate, error) {
	var (
		logMsg strings.Builder
		method string
	)
	logMsg.Grow(30)
	logMsg.WriteString("(")
	if exists {
		logMsg.WriteString("update")
		method = "PATCH"
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

	path := make([]string, 0, len(Suffix1)+len(Suffix2)+1)
	path = append(path, Suffix1...)
	prismaId, err := GetPrismaId(c)
	if err != nil {
		return NotificationTemplate{}, nil
	}
	path = append(path, prismaId)
	path = append(path, Suffix2...)
	if exists {
		path = append(path, id)
	}
	var templateRes NotificationTemplate
	if exists {
		_, err = c.Communicate(method, path, nil, template, nil)
	} else {
		_, err = c.Communicate(method, path, nil, template, &templateRes)
	}
	return templateRes, err
}

func Identify(c pc.PrismaCloudClient, templateName string) (string, error) {
	c.Log(pc.LogAction, "(get) id for %s templateName:%s", singular, templateName)
	ans, err := List(c)
	if err != nil {
		return "", err
	}
	for _, o := range ans {
		if o.Name == templateName {
			return o.Id, nil
		}
	}
	return "", pc.ObjectNotFoundError
}

// GetPrismaId returns prisma id.
func GetPrismaId(c pc.PrismaCloudClient) (string, error) {
	c.Log(pc.LogAction, "(get) prisma id")
	var ans LicenseInfo
	path := make([]string, 0, len(LicenseSuffix)+1)
	path = append(path, LicenseSuffix...)
	_, err := c.Communicate("GET", path, nil, nil, &ans)
	return ans.PrismaId, err
}
