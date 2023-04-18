package org

import (
	"fmt"
	"strings"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
)

// List lists all cloud accounts onboarded onto the Prisma Cloud platform.
func List(c pc.PrismaCloudClient) ([]AccountResponse, error) {
	c.Log(pc.LogAction, "(get) list of %s", plural)

	var ans []AccountResponse
	if _, err := c.Communicate("GET", ListSuffixAws, nil, nil, &ans); err != nil {
		return nil, err
	}

	return ans, nil
}

func ListAzure(c pc.PrismaCloudClient) ([]AzureAccountResponse, error) {
	c.Log(pc.LogAction, "(get) list of %s", plural)

	var ansaz []AzureAccountResponse
	if _, err := c.Communicate("GET", ListSuffixAzure, nil, nil, &ansaz); err != nil {
		return nil, err
	}

	return ansaz, nil
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
	if cloudType == "aws" {
		ans, err := List(c)
		if err != nil {
			return "", err
		}

		for _, o := range ans {
			if strings.EqualFold(o.CloudAccountResp.CloudType, cloudType) && o.CloudAccountResp.Name == name {
				return o.CloudAccountResp.AccountId, nil
			}

		}
	}
	if strings.EqualFold("azure", cloudType) {
		ansaz, err := ListAzure(c)
		if err != nil {
			return "", err
		}

		for _, o := range ansaz {
			if strings.EqualFold(o.CloudAccountAzureResp.CloudType, cloudType) && o.CloudAccountAzureResp.Name == name {
				return o.CloudAccountAzureResp.AccountId, nil
			}

		}
	}

	return "", pc.ObjectNotFoundError
}

func Get(c pc.PrismaCloudClient, cloudType, id string) (interface{}, error) {
	var cloud string
	cloud = cloudType
	cloud = cloud + "Accounts"
	c.Log(pc.LogAction, "(get) %s type:%s id:%s", singular, cloudType, id)

	path := make([]string, 0, len(ListSuffix)+1)
	path = append(path, ListSuffix...)
	path = append(path, cloud, id)

	var ans interface{}

	switch cloudType {
	case TypeAwsOrg:
		ans = &AwsOrgV2{}
	case TypeAzureOrg:
		ans = &AzureOrgV2{}
	default:
		return nil, fmt.Errorf("Invalid cloud type: %s", cloudType)
	}
	_, err := c.Communicate("GET", path, nil, nil, ans)

	switch cloudType {
	case TypeAwsOrg:
		return *ans.(*AwsOrgV2), err
	case TypeAzureOrg:
		return *ans.(*AzureOrgV2), err
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

	path := make([]string, 0, len(DeleteSuffix)+2)
	path = append(path, DeleteSuffix...)
	path = append(path, cloudType, id)
	_, err := c.Communicate("DELETE", path, nil, nil, nil)
	return err
}

func DisableCloudAccount(c pc.PrismaCloudClient, accountId string) error {
	var (
		method string
	)
	method = "PATCH"

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, "cloud")
	path = append(path, accountId)
	path = append(path, "status")
	path = append(path, "false")

	_, err := c.Communicate(method, path, nil, nil, nil)
	return err
}

func createUpdate(exists bool, c pc.PrismaCloudClient, account interface{}) error {
	var (
		logMsg    strings.Builder
		id        string
		method    string
		cloudType string
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
		id = v.OrgAccountAzure.AccountId
	default:
		return fmt.Errorf("invalid account type %v", v)
	}

	logMsg.WriteString(" ")
	logMsg.WriteString(singular)
	if exists {
		fmt.Fprintf(&logMsg, ": %s", id)
	}

	c.Log(pc.LogAction, logMsg.String())

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	cloudType = cloudType + "_account"
	path = append(path, cloudType)
	if exists {
		path = append(path, id)
	}
	_, err := c.Communicate(method, path, nil, account, nil)
	return err
}
