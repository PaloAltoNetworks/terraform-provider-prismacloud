package org

import (
	"fmt"
	"strings"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
)

// List lists all cloud accounts onboarded onto the Prisma Cloud platform.
func List(c pc.PrismaCloudClient) ([]OrgAccount, error) {
	c.Log(pc.LogAction, "(get) list of %s", plural)

	var ans []OrgAccount
	if _, err := c.Communicate("GET", Suffix, nil, nil, &ans); err != nil {
		return nil, err
	}

	return ans, nil
}

// Names returns the name listing for cloud accounts.
func Names(c pc.PrismaCloudClient) ([]NameTypeId, error) {
	c.Log(pc.LogAction, "(get) %s names", singular)

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	path = append(path, "name")

	var ans []NameTypeId
	_, err := c.Communicate("GET", path, nil, nil, &ans)

	return ans, err
}

// Identify returns the ID for the given cloud type and name.
func Identify(c pc.PrismaCloudClient, cloudType, name string) (string, error) {
	c.Log(pc.LogAction, "(get) id for %s type:%s name:%s", singular, cloudType, name)

	ans, err := Names(c)
	if err != nil {
		return "", err
	}

	for _, o := range ans {
		if o.CloudType == cloudType && o.Name == name {
			return o.AccountId, nil
		}
	}

	return "", pc.ObjectNotFoundError
}

/*
Get returns top level information about the cloud account.

The interface returned will be one of the following:
- account.Aws
- account.Azure
- account.Gcp
- account.Alibaba
- nil
*/
func Get(c pc.PrismaCloudClient, cloudType, id string) (interface{}, error) {
	c.Log(pc.LogAction, "(get) %s type:%s id:%s", singular, cloudType, id)

	path := make([]string, 0, len(Suffix)+2)
	path = append(path, Suffix...)
	path = append(path, cloudType, id)

	var ans interface{}

	switch cloudType {
	case TypeAwsOrg:
		ans = &AwsOrg{}
	case TypeAzureOrg:
		ans = &AzureOrg{}
	case TypeGcpOrg:
		ans = &GcpOrg{}
	case TypeOci:
		ans = &Oci{}
	default:
		return nil, fmt.Errorf("Invalid cloud type: %s", cloudType)
	}

	_, err := c.Communicate("GET", path, nil, nil, ans)

	// Can't just return ans here or it won't be right, so re-cast it back
	// to the appropriate specific object type.
	switch cloudType {
	case TypeAwsOrg:
		return *ans.(*AwsOrg), err
	case TypeAzureOrg:
		return *ans.(*AzureOrg), err
	case TypeGcpOrg:
		return *ans.(*GcpOrg), err
	case TypeOci:
		return *ans.(*Oci), err
	}

	return nil, fmt.Errorf("Invalid cloud type: %s", cloudType)
}

// Create onboards a new cloud account onto the Prisma Cloud platform.
func Create(c pc.PrismaCloudClient, account interface{}) error {
	return createUpdate(false, c, account)
}

// Update modifies information related to a cloud account.
func Update(c pc.PrismaCloudClient, account interface{}) error {
	return createUpdate(true, c, account)
}

// Delete removes an onboarded cloud account using the cloud account ID.
func Delete(c pc.PrismaCloudClient, cloudType, id string) error {
	c.Log(pc.LogAction, "(delete) %s type:%s id:%s", singular, cloudType, id)

	//path := strings.Join([]string{Suffix, cloudType, id}, "/")
	path := make([]string, 0, len(Suffix)+2)
	path = append(path, Suffix...)
	path = append(path, cloudType, id)
	_, err := c.Communicate("DELETE", path, nil, nil, nil)
	return err
}

func createUpdate(exists bool, c pc.PrismaCloudClient, account interface{}) error {
	var (
		cloudType string
		logMsg    strings.Builder
		id        string
		method    string
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

	switch v := account.(type) {
	case nil:
		return fmt.Errorf("Cloud account specified")
	case AwsOrg:
		logMsg.WriteString("aws")
		cloudType = TypeAwsOrg
		id = v.AccountId
	case AzureOrg:
		logMsg.WriteString("azure")
		cloudType = TypeAzureOrg
		id = v.Account.AccountId
	case GcpOrg:
		logMsg.WriteString("gcp")
		cloudType = TypeGcpOrg
		id = v.Account.AccountId
	case Oci:
		logMsg.WriteString("oci")
		cloudType = TypeOci
		id = v.AccountId
	default:
		return fmt.Errorf("invalid account type %v", v)
	}

	logMsg.WriteString(" ")
	logMsg.WriteString(singular)
	if exists {
		fmt.Fprintf(&logMsg, ": %s", id)
	}

	c.Log(pc.LogAction, logMsg.String())

	path := make([]string, 0, len(Suffix)+2)
	path = append(path, Suffix...)
	path = append(path, cloudType)
	if exists {
		path = append(path, id)
	}

	_, err := c.Communicate(method, path, nil, account, nil)
	return err
}
