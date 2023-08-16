package dataprofile

import (
	"fmt"
	"strings"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
)

// GetTenantId returns dlp tenant id.
func GetTenantId(c pc.PrismaCloudClient) (string, error) {
	c.Log(pc.LogAction, "(get) prisma id")

	var ans TenantInfo

	path := make([]string, 0, len(TenantSuffix)+1)
	path = append(path, TenantSuffix...)

	_, err := c.Communicate("GET", path, nil, nil, &ans)
	return ans.DlpTenantId, err
}

// List returns a list of available custom data profiles.
func List(c pc.PrismaCloudClient) ([]ListProfile, error) {
	c.Log(pc.LogAction, "(get) list of %s", plural)

	var ans ListBody
	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)

	dlpTenantId, err1 := GetTenantId(c)
	if err1 != nil {
		return nil, err1
	}

	path = append(path, dlpTenantId)
	_, err := c.Communicate("GET", path, nil, nil, &ans)
	return ans.Profiles, err
}

// Identify returns the ID for the given data profile.
func Identify(c pc.PrismaCloudClient, name string) (string, error) {
	c.Log(pc.LogAction, "(get) id for %s name:%s", singular, name)

	list, err := List(c)

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

// Get returns the data profile that has the specified ID.
func Get(c pc.PrismaCloudClient, id string) (Profile, error) {
	c.Log(pc.LogAction, "(get) %s id:%s", singular, id)

	var ans Profile

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)

	dlpTenantId, err1 := GetTenantId(c)
	if err1 != nil {
		return ans, err1
	}

	path = append(path, dlpTenantId)
	path = append(path, "id")
	path = append(path, id)

	_, err := c.Communicate("GET", path, nil, nil, &ans)

	return ans, err
}

// Create adds a new data profile.
func Create(c pc.PrismaCloudClient, profile Profile) error {
	return createUpdate(false, c, profile)
}

// Update modifies the existing data profile.
func Update(c pc.PrismaCloudClient, profile Profile) error {
	return createUpdate(true, c, profile)
}

// Delete removes a data profile using its ID.
func Delete(c pc.PrismaCloudClient, id string) error {
	c.Log(pc.LogAction, "(delete) %s id:%s", singular, id)

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)

	dlpTenantId, err1 := GetTenantId(c)
	if err1 != nil {
		return err1
	}

	path = append(path, dlpTenantId)
	path = append(path, "id")
	path = append(path, id)
	_, err := c.Communicate("DELETE", path, nil, nil, nil)
	return err
}

func createUpdate(exists bool, c pc.PrismaCloudClient, profile Profile) error {
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
	path = append(path, Suffix...)

	dlpTenantId, err1 := GetTenantId(c)
	if err1 != nil {
		return err1
	}

	path = append(path, dlpTenantId)
	if exists {
		path = append(path, "id")
		path = append(path, profile.Id)
	}

	_, err := c.Communicate(method, path, nil, profile, nil)
	return err
}
