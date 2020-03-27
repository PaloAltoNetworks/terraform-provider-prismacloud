package account

import (
	"fmt"
	"strings"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
)

// List lists all cloud accounts onboarded onto the Prisma Cloud platform.
func List(c pc.PrismaCloudClient) ([]Account, error) {
	c.Log(pc.LogAction, "(get) list of %s", plural)

	var ans []Account
	if _, err := c.Communicate("GET", Suffix, nil, &ans, true); err != nil {
		return nil, err
	}

	return ans, nil
}

// Identify returns the ID for the given cloud type and name.
func Identify(c pc.PrismaCloudClient, cloudType, name string) (string, error) {
	c.Log(pc.LogAction, "(get) id for %s type:%s name:%s", singular, cloudType, name)

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	path = append(path, "name")

	var ans []NameTypeId
	if _, err := c.Communicate("GET", path, nil, &ans, true); err != nil {
		return "", err
	}

	for _, o := range ans {
		if o.CloudType == cloudType && o.Name == name {
			return o.AccountId, nil
		}
	}

	return "", pc.ObjectNotFoundError
}

// Get returns top level information about the cloud account.
func Get(c pc.PrismaCloudClient, cloudType, id string) (Account, error) {
	c.Log(pc.LogAction, "(get) %s type:%s id:%s", singular, cloudType, id)

	ans := AccountAndCredentials{}

	//path := strings.Join([]string{Suffix, cloudType, id}, "/")
	path := make([]string, 0, len(Suffix)+2)
	path = append(path, Suffix...)
	path = append(path, cloudType, id)

	if _, err := c.Communicate("GET", path, nil, &ans, true); err != nil {
		return ans.Account, err
	}

	return ans.Account, nil
}

// Update modifies information related to a cloud account.
func Update(c pc.PrismaCloudClient, account interface{}) error {
	return createUpdate(false, c, account)
}

// Create onboards a new cloud account onto the Prisma Cloud platform.
func Create(c pc.PrismaCloudClient, account interface{}) error {
	return createUpdate(true, c, account)
}

// Delete removes an onboarded cloud account using the cloud account ID.
func Delete(c pc.PrismaCloudClient, cloudType, id string) error {
	c.Log(pc.LogAction, "(delete) %s type:%s id:%s", singular, cloudType, id)

	//path := strings.Join([]string{Suffix, cloudType, id}, "/")
	path := make([]string, 0, len(Suffix)+2)
	path = append(path, Suffix...)
	path = append(path, cloudType, id)
	_, err := c.Communicate("DELETE", path, nil, nil, true)
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
	case AwsAccount:
		logMsg.WriteString("aws")
		cloudType = TypeAws
		id = v.AccountId
	case AzureAccount:
		logMsg.WriteString("azure")
		cloudType = TypeAzure
		id = v.Account.AccountId
	case GcpAccount:
		logMsg.WriteString("gcp")
		cloudType = TypeGcp
		id = v.Account.AccountId
	case AlibabaAccount:
		logMsg.WriteString("alibaba")
		cloudType = TypeAlibaba
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

	_, err := c.Communicate(method, path, account, nil, true)
	return err
}
