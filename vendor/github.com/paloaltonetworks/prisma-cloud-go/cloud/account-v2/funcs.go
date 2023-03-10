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
	if _, err := c.Communicate("GET", ListSuffix, nil, nil, &ans); err != nil {
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

	ans, err := List(c)
	if err != nil {
		return "", err
	}

	for _, o := range ans {
		if strings.EqualFold(o.CloudAccountResp.CloudType, cloudType) && o.CloudAccountResp.Name == name {
			return o.CloudAccountResp.AccountId, nil
		}
	}

	return "", pc.ObjectNotFoundError
}

func Get(c pc.PrismaCloudClient, cloudType, id string) (interface{}, error) {
	c.Log(pc.LogAction, "(get) %s type:%s id:%s", singular, cloudType, id)

	path := make([]string, 0, len(ListSuffix)+1)
	path = append(path, ListSuffix...)
	path = append(path, id)

	var ans interface{}

	if cloudType == TypeAws {
		ans = &AwsV2{}
	} else {
		return nil, fmt.Errorf("Invalid cloud type: %s", cloudType)
	}

	_, err := c.Communicate("GET", path, nil, nil, ans)

	if cloudType == TypeAws {
		return *ans.(*AwsV2), err
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
		logMsg strings.Builder
		id     string
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

	switch v := account.(type) {
	case nil:
		return fmt.Errorf("Cloud account specified")
	case Aws:
		logMsg.WriteString("aws")
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

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	if exists {
		path = append(path, id)
	}
	_, err := c.Communicate(method, path, nil, account, nil)
	return err
}
