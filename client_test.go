package prismacloud

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
)

var GenericAuthFile = []byte(`{
    "url": "api.prismacloud.io",
    "username": "defaultUser",
    "password": "defaultPassword",
    "customer-name": "defaultCustomer",
    "skip-ssl-cert-verification": true,
    "logging": {"send": true, "receive": true}
}
`)

var LoginResponse = `{
    "customerNames": [
        {
            "customerName": "RedLockDemo",
            "tosAccepted": true
        },
        {
            "customerName": "SESandBox",
            "tosAccepted": true
        }
    ],
    "message": "login_successful",
    "token": "testJwtToken"
}`

func init() {
	log.SetFlags(0)
}

func rl() {
	log.SetOutput(os.Stderr)
}

func MockClient(responses []*http.Response, withLogin bool) Client {
	c := Client{
		pcAuthFileContent: GenericAuthFile,
	}

	if withLogin {
		rlist := make([]*http.Response, 0, len(responses)+1)
		rlist = append(rlist, &http.Response{
			StatusCode: 200,
			Body: ioutil.NopCloser(
				strings.NewReader(LoginResponse),
			),
		})
		rlist = append(rlist, responses...)
		c.pcResponses = rlist
	} else {
		c.pcResponses = responses
	}

	return c
}

func TestLogin(t *testing.T) {
	c := MockClient(nil, true)
	err := c.Initialize("test")

	if err != nil {
		t.Fail()
	}
}

func TestInitializeSetsJwt(t *testing.T) {
	c := MockClient(nil, true)
	_ = c.Initialize("test")

	if c.JsonWebToken == "" {
		t.Fail()
	}
}

func TestInitializeReadsDefaults(t *testing.T) {
	c := MockClient(nil, true)
	c.Url = "testurl"
	c.Username = "user"
	c.Password = "secret"
	_ = c.Initialize("")

	if c.Protocol == "" {
		t.Fail()
	}

	if c.Timeout == 0 {
		t.Fail()
	}

	if len(c.Logging) == 0 {
		t.Fail()
	}
}

func TestLogAction(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer rl()

	c := MockClient(nil, true)
	c.Logging = map[string]bool{LogAction: true}
	c.Log(LogAction, "ok")
	s := buf.String()
	if s != "ok\n" {
		t.Fail()
	}
}

func TestLogActionDisabled(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer rl()

	c := MockClient(nil, true)
	c.Logging = map[string]bool{LogQuiet: true}
	c.Log(LogAction, "ok")
	s := buf.String()
	if s != "" {
		t.Fail()
	}
}
