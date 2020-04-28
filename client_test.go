package prismacloud

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"
)

var TestPath = []string{"test", "path"}

var GenericAuthFile = []byte(`{
    "url": "api.prismacloud.io",
    "username": "defaultUser",
    "password": "defaultPassword",
    "customer_name": "defaultCustomer",
    "skip_ssl_cert_verification": true,
    "logging": {"quiet": true}
}
`)

var LoginBody = `{
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

func OkResponse(body string) *http.Response {
	return &http.Response{
		StatusCode: 200,
		Body: ioutil.NopCloser(
			strings.NewReader(body),
		),
	}
}

func ErrorResponse(code int, body string, e interface{}) *http.Response {
	ans := &http.Response{
		StatusCode: code,
		Body: ioutil.NopCloser(
			strings.NewReader(body),
		),
	}

	if e != nil {
		errBody, _ := json.Marshal(e)
		ans.Header = map[string][]string{
			"X-Redlock-Status": []string{
				"[" + string(errBody) + "]",
			},
		}
	}

	return ans
}

func MockClient(responses []*http.Response) Client {
	c := Client{
		pcAuthFileContent: GenericAuthFile,
	}

	c.pcResponses = make([]*http.Response, 0, len(responses)+1)
	c.pcResponses = append(c.pcResponses, OkResponse(LoginBody))
	c.pcResponses = append(c.pcResponses, responses...)

	if len(responses) > 0 {
		_ = c.Initialize("test")
	}

	return c
}

func TestLogin(t *testing.T) {
	c := MockClient(nil)
	err := c.Initialize("test")

	if err != nil {
		t.Fail()
	}
}

func TestInitializeSetsJwt(t *testing.T) {
	c := MockClient(nil)
	_ = c.Initialize("test")

	if c.JsonWebToken == "" {
		t.Fail()
	}
}

func TestInitializeReadsDefaults(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer rl()

	c := MockClient(nil)
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

	s := buf.String()
	if s == "" {
		t.Fail()
	}
}

func TestLogAction(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer rl()

	c := MockClient(nil)
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

	c := MockClient(nil)
	c.Logging = map[string]bool{LogQuiet: true}
	c.Log(LogAction, "ok")
	s := buf.String()
	if s != "" {
		t.Fail()
	}
}

func TestPathWithQuery(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer rl()

	c := MockClient(
		[]*http.Response{
			OkResponse(""),
		},
	)
	c.Logging[LogPath] = true

	v := url.Values{}
	v.Set("foo", "bar")

	_, err := c.Communicate("GET", TestPath, v, nil, nil)
	if err != nil {
		t.Fail()
	}

	expected := "path: https://api.prismacloud.io/test/path?foo=bar\n"
	if buf.String() != expected {
		t.Errorf("expected:%q  got:%q", expected, buf.String())
	}
}

func TestPathWithoutQuery(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer rl()

	c := MockClient(
		[]*http.Response{
			OkResponse(""),
		},
	)
	c.Logging[LogPath] = true

	_, err := c.Communicate("GET", TestPath, nil, nil, nil)
	if err != nil {
		t.Fail()
	}

	expected := "path: https://api.prismacloud.io/test/path\n"
	if buf.String() != expected {
		t.Errorf("expected:%q  got:%q", expected, buf.String())
	}
}

func TestReauthenticateHappens(t *testing.T) {
	expected := "okay"
	c := MockClient(
		[]*http.Response{
			ErrorResponse(http.StatusUnauthorized, "", nil),
			OkResponse(LoginBody),
			OkResponse(expected),
		},
	)
	c.JsonWebToken = "oldToken"

	body, err := c.Communicate("GET", TestPath, nil, nil, nil)
	if err != nil {
		t.Errorf("Failed communicate: %s", err)
	}

	if string(body) != expected {
		t.Errorf("expected %q, got %q", expected, string(body))
	}

	if c.JsonWebToken == "oldToken" {
		t.Fail()
	}
}

func TestAlreadyExists(t *testing.T) {
	c := MockClient(
		[]*http.Response{
			ErrorResponse(
				http.StatusTeapot,
				"",
				PrismaCloudError{
					Message:  "teapot_already_exists",
					Severity: "error",
				},
			),
		},
	)

	_, err := c.Communicate("GET", TestPath, nil, nil, nil)
	if err == nil {
		t.Fail()
	}

	if err != AlreadyExistsError {
		t.Errorf("error is %s, not %s", err, AlreadyExistsError)
	}
}

func TestInvalidCredentialsError(t *testing.T) {
	c := MockClient(
		[]*http.Response{
			ErrorResponse(http.StatusUnauthorized, "", nil),
			ErrorResponse(http.StatusUnauthorized, "", nil),
		},
	)

	_, err := c.Communicate("GET", TestPath, nil, nil, nil)
	if err == nil {
		t.Fail()
	}

	if err != InvalidCredentialsError {
		t.Errorf("error is %s, not %s", err, InvalidCredentialsError)
	}
}

func TestObjectNotFoundError(t *testing.T) {
	c := MockClient(
		[]*http.Response{
			ErrorResponse(
				http.StatusTeapot,
				"",
				PrismaCloudError{
					Message:  "not_found",
					Severity: "error",
				},
			),
		},
	)

	_, err := c.Communicate("GET", TestPath, nil, nil, nil)
	if err == nil {
		t.Fail()
	}

	if err != ObjectNotFoundError {
		t.Errorf("error is %s, not %s", err, ObjectNotFoundError)
	}
}

func TestInvalidResponse(t *testing.T) {
	expectedHtml := `<html>
  <body>
    <p>nope</p>
  </body>
</html>`
	c := MockClient(
		[]*http.Response{
			ErrorResponse(http.StatusTeapot, expectedHtml, nil),
		},
	)

	data, err := c.Communicate("GET", TestPath, nil, nil, nil)
	if string(data) != expectedHtml {
		t.Errorf("Expected:%s got:%s", expectedHtml, data)
	}

	if err == nil {
		t.Errorf("No error returned")
	}

	if !strings.HasSuffix(err.Error(), expectedHtml) {
		t.Errorf("Error doesn't have expected suffix:\n%s", err)
	}
}
