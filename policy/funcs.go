package policy

import (
	"fmt"
	"net/url"
	"strings"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
)

// List returns a list of available policies, both system default and custom.
func List(c pc.PrismaCloudClient, query map[string]string) ([]Policy, error) {
	c.Log(pc.LogAction, "(get) list of %s", plural)

	qv := url.Values{}
	for k, v := range query {
		qv.Set(k, v)
	}

	var ans []Policy
	_, err := c.Communicate("GET", []string{"v2", "policy"}, qv, nil, &ans)

	return ans, err
}

// Identify returns the ID for the given policy name.
func Identify(c pc.PrismaCloudClient, name string) (string, error) {
	c.Log(pc.LogAction, "(get) id for %s name:%s", singular, name)

	ans, err := List(c, map[string]string{"policy.name": name})
	if err != nil {
		return "", err
	}

	switch len(ans) {
	case 0:
		return "", pc.ObjectNotFoundError
	case 1:
		return ans[0].PolicyId, nil
	}

	return "", fmt.Errorf("Got %d results back not 1", len(ans))
}

// Get returns the policy that has the specified ID.
func Get(c pc.PrismaCloudClient, id string) (Policy, error) {
	c.Log(pc.LogAction, "(get) %s id:%s", singular, id)

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	path = append(path, id)

	var ans Policy
	_, err := c.Communicate("GET", path, nil, nil, &ans)

	return ans, err
}

// Create adds a new policy.
func Create(c pc.PrismaCloudClient, policy Policy) error {
	return createUpdate(false, c, policy)
}

// Update modifies the existing policy.
func Update(c pc.PrismaCloudClient, policy Policy) error {
	return createUpdate(true, c, policy)
}

// Delete removes a policy using its ID.
func Delete(c pc.PrismaCloudClient, id string) error {
	c.Log(pc.LogAction, "(delete) %s id:%s", singular, id)

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	path = append(path, id)
	_, err := c.Communicate("DELETE", path, nil, nil, nil)
	return err
}

func createUpdate(exists bool, c pc.PrismaCloudClient, policy Policy) error {
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
		fmt.Fprintf(&logMsg, ":%s", policy.PolicyId)
	}

	c.Log(pc.LogAction, logMsg.String())

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	if exists {
		path = append(path, policy.PolicyId)
	}

	_, err := c.Communicate(method, path, nil, policy, nil)
	return err
}
