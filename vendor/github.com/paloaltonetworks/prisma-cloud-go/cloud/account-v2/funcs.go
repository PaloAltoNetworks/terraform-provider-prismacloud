package accountv2

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

func ListGcp(c pc.PrismaCloudClient) ([]GcpAccountResponse, error) {
	c.Log(pc.LogAction, "(get) list of %s", plural)

	var ansgcp []GcpAccountResponse
	if _, err := c.Communicate("GET", ListSuffixGcp, nil, nil, &ansgcp); err != nil {
		return nil, err
	}

	return ansgcp, nil
}

func ListIbm(c pc.PrismaCloudClient) ([]IbmAccountResponse, error) {
	c.Log(pc.LogAction, "(get) list of %s", plural)
	path := make([]string, 0, len(ListSuffix)+1)
	path = append(path, ListSuffix...)
	path = append(path, "accounts?cloudTypes=ibm")

	var ansib []IbmAccountResponse
	if _, err := c.Communicate("GET", path, nil, nil, &ansib); err != nil {
		return nil, err
	}

	return ansib, nil
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
	if strings.EqualFold("gcp", cloudType) {
		ansgcp, err := ListGcp(c)
		if err != nil {
			return "", err
		}

		for _, o := range ansgcp {
			if strings.EqualFold(o.CloudType, cloudType) && o.Name == name {
				return o.AccountId, nil
			}

		}
	}
	if strings.EqualFold("ibm", cloudType) {
		ansib, err := ListIbm(c)
		if err != nil {
			return "", err
		}

		for _, o := range ansib {
			if strings.EqualFold(o.CloudType, cloudType) && o.Name == name {
				return o.AccountId, nil
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

	if cloudType == "gcp" {
		path := make([]string, 0, len(ListSuffixGcp)+1)
		path = append(path, ListSuffixGcp...)
		path = append(path, id)
		var ans interface{}

		switch cloudType {
		case TypeGcp:
			ans = &GcpV2{}
		default:
			return nil, fmt.Errorf("Invalid cloud type: %s", cloudType)
		}
		_, err := c.Communicate("GET", path, nil, nil, ans)

		switch cloudType {
		case TypeGcp:
			return *ans.(*GcpV2), err
		}
	} else if cloudType == "ibm" {
		path := make([]string, 0, len(ListSuffixIbm)+1)
		path = append(path, ListSuffixIbm...)
		path = append(path, id)

		var ans interface{}

		switch cloudType {
		case TypeIbm:
			ans = &IbmV2{}
		default:
			return nil, fmt.Errorf("Invalid cloud type: %s", cloudType)
		}
		_, err := c.Communicate("GET", path, nil, nil, ans)

		switch cloudType {
		case TypeIbm:
			return *ans.(*IbmV2), err
		}
	} else {
		path := make([]string, 0, len(ListSuffix)+1)
		path = append(path, ListSuffix...)
		path = append(path, cloud, id)
		var ans interface{}

		switch cloudType {
		case TypeAws:
			ans = &AwsV2{}
		case TypeAzure:
			ans = &AzureV2{}
		default:
			return nil, fmt.Errorf("Invalid cloud type: %s", cloudType)
		}
		_, err := c.Communicate("GET", path, nil, nil, ans)

		switch cloudType {
		case TypeAws:
			return *ans.(*AwsV2), err
		case TypeAzure:
			return *ans.(*AzureV2), err
		}

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
	case Aws:
		logMsg.WriteString("aws")
		cloudType = TypeAws
		id = v.AccountId
	case Azure:
		logMsg.WriteString("azure")
		cloudType = TypeAzure
		id = v.CloudAccountAzure.AccountId
	case Gcp:
		logMsg.WriteString("gcp")
		cloudType = TypeGcp
		id = v.CloudAccountGcp.AccountId
	case Ibm:
		logMsg.WriteString("ibm")
		cloudType = TypeIbm
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
	if cloudType == "ibm" {
		path := make([]string, 0, len(Suffix)+1)
		path = append(path, Suffix...)
		path = append(path, "cloud-type")
		path = append(path, "ibm")
		path = append(path, "account")
		if exists {
			path = append(path, id)
		}
		_, err := c.Communicate(method, path, nil, account, nil)
		return err
	} else {
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
}
