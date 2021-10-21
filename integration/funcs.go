package integration

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
)

// Types returns a list of all supported integration types.
func Types(c pc.PrismaCloudClient) ([]string, error) {
	c.Log(pc.LogAction, "(get) list of %s types", singular)

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	path = append(path, "type")

	var ans []string
	_, err := c.Communicate("GET", path, nil, nil, &ans)
	return ans, err
}

// List returns all your integrations, optionally filtered by type.
func List(c pc.PrismaCloudClient, t string) ([]Integration, error) {
	c.Log(pc.LogAction, "(get) list of %s", plural)

	var query url.Values
	if t != "" {
		query = url.Values{}
		query.Add("type", t)
	}

	var ans []Integration
	if _, err := c.Communicate("GET", Suffix, query, nil, &ans); err != nil {
		return nil, err
	}

	return ans, nil
}

// Identify returns the ID for the given integration name.
func Identify(c pc.PrismaCloudClient, name string) (string, error) {
	c.Log(pc.LogAction, "(get) id for %s: %s", singular, name)

	list, err := List(c, "")
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

// jira auth url give the authenticalurl used to get token
func JiraAuthurl(c pc.PrismaCloudClient, url AuthUrl) (string, error) {
	c.Log(pc.LogAction, "(get) %s:", JiraAuthUrl)

	path := make([]string, 0, len(JiraUrlSuffix)+1)
	path = append(path, JiraUrlSuffix...)
	authurlresponse, err := c.Communicate("POST", path, nil, &url, nil)
	return string(authurlresponse), err
}

// Jira_SecretKey returns link with auth token
func JiraSecretKey(c pc.PrismaCloudClient, data SecretKeyJira, authurl string) (string, error) {

	cjira := c.(*pc.Client)
	prismaurl := cjira.Url
	cjira.Url = strings.Split(authurl, "https://")[1]
	suffix := make([]string, 0, len(JiraSecretKeySuffix)+1)
	suffix = append(suffix, JiraSecretKeySuffix...)
	data_x_www_form := url.Values{}
	data_x_www_form.Set("oauth_token", data.OauthToken)
	data_x_www_form.Set("approve", "Allow")
	data_x_www_form.Set("jira_username", data.JiraUserName)
	data_x_www_form.Set("jira_password", data.JiraPassword)
	encodedData := data_x_www_form.Encode()
	var path strings.Builder
	path.Grow(30)
	fmt.Fprintf(&path, "%s://%s", cjira.Protocol, cjira.Url)
	if cjira.Port != 0 {
		fmt.Fprintf(&path, ":%d", cjira.Port)
	}
	for _, v := range suffix {
		path.WriteString("/")
		path.WriteString(v)
	}

	req, err := http.NewRequest("POST", path.String(), strings.NewReader(encodedData))
	if err != nil {
		return "", err
	}
	jira_auth_details := data.JiraUserName + ":" + data.JiraPassword
	encodedjiraauthdetails := base64.StdEncoding.EncodeToString([]byte(jira_auth_details))
	AuthorizationHeadervalue := "Basic " + encodedjiraauthdetails
	req.Header.Set("Authorization", AuthorizationHeadervalue)
	req.Header.Add("X-Atlassian-Token", "no-check")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := cjira.DoJira(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	htmlcode := string(body)

	switch resp.StatusCode {
	case http.StatusOK, http.StatusNoContent:
		// Alert rule deletion returns StatusNoContent
	case http.StatusUnauthorized:
		log.Printf("Status is unauthorized")
	default:
		log.Printf("X-Redlock-Status error")
	}

	secretcodeline := strings.Split(htmlcode, "Your verification code is ")[1]
	secretcode := strings.Split(secretcodeline, ". You will need")[0]
	substr := strings.Split(secretcode, "&#39;")[1]

	cjira.Url = prismaurl
	return substr, err
}

// Jira_Oauth Token returns link with auth token
func JiraOauthToken(c pc.PrismaCloudClient, oauthtoken OauthTokenJira) (string, error) {
	c.Log(pc.LogAction, "(get) %s:", JiraOauthtoken)
	path := make([]string, 0, len(JiraTokenSuffix)+1)
	path = append(path, JiraTokenSuffix...)
	tokenresponse, err := c.Communicate("POST", path, nil, &oauthtoken, nil)

	return string(tokenresponse), err
}

// Get returns integration details for the specified ID.
func Get(c pc.PrismaCloudClient, id string) (Integration, error) {
	c.Log(pc.LogAction, "(get) %s: %s", singular, id)

	var ans Integration
	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	path = append(path, id)

	_, err := c.Communicate("GET", path, nil, nil, &ans)
	return ans, err
}

// Create adds an integration with the specified external system.
func Create(c pc.PrismaCloudClient, obj Integration) error {
	return createUpdate(false, c, obj)
}

// Update modifies the specified integration.
func Update(c pc.PrismaCloudClient, obj Integration) error {
	return createUpdate(true, c, obj)
}

// Delete removes the integration for the specified ID.
func Delete(c pc.PrismaCloudClient, id string) error {
	c.Log(pc.LogAction, "(delete) %s: %s", singular, id)

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	path = append(path, id)

	_, err := c.Communicate("DELETE", path, nil, nil, nil)
	return err
}

func createUpdate(exists bool, c pc.PrismaCloudClient, obj Integration) error {
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
		fmt.Fprintf(&logMsg, ": %s", obj.Id)
	}

	c.Log(pc.LogAction, logMsg.String())

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	if exists {
		path = append(path, obj.Id)
	}

	_, err := c.Communicate(method, path, nil, obj, nil)
	return err
}
