package report

import (
	"fmt"
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"net/url"
	"strings"
)

// List returns a list of available alert and compliance reports
func List(c pc.PrismaCloudClient) ([]Report, error) {
	c.Log(pc.LogAction, "(get) list of %s", plural)

	var alertReportList []Report
	qv := url.Values{}
	key := "report_view"
	for _, val := range alertReportTypes {
		qv.Add(key, val)
	}

	if _, err := c.Communicate("GET", Suffix, qv, nil, &alertReportList); err != nil {
		return nil, err
	}

	var complianceReportList []Report

	if _, err := c.Communicate("GET", Suffix, nil, nil, &complianceReportList); err != nil {
		return nil, err
	}

	var ans []Report
	ans = append(alertReportList, complianceReportList...)
	return ans, nil
}

// Identify returns the ID for the given report name.
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

// Get returns the report that has the specified ID.
func Get(c pc.PrismaCloudClient, id string) (Report, error) {
	c.Log(pc.LogAction, "(get) %s id:%s", singular, id)

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	path = append(path, id)

	var ans Report
	_, err := c.Communicate("GET", path, nil, nil, &ans)

	return ans, err
}

// Create adds a new report.
func Create(c pc.PrismaCloudClient, report Report) error {
	return createUpdate(false, c, report)
}

// Update modifies the existing report.
func Update(c pc.PrismaCloudClient, report Report) error {
	return createUpdate(true, c, report)
}

// Delete removes a report using its ID.
func Delete(c pc.PrismaCloudClient, id string) error {
	c.Log(pc.LogAction, "(delete) %s id:%s", singular, id)

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	path = append(path, id)
	_, err := c.Communicate("DELETE", path, nil, nil, nil)
	return err
}

func createUpdate(exists bool, c pc.PrismaCloudClient, report Report) error {
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
		fmt.Fprintf(&logMsg, ":%s", report.Id)
	}

	c.Log(pc.LogAction, logMsg.String())

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	if exists {
		path = append(path, report.Id)
	}

	_, err := c.Communicate(method, path, nil, report, nil)
	return err
}
