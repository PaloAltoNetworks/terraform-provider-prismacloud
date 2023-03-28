package prismacloud

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Client is a client connection to Prisma Cloud.
type Client struct {
	// Properties.
	Url                     string          `json:"url"`
	Username                string          `json:"username"`
	Password                string          `json:"password"`
	CustomerName            string          `json:"customer_name"`
	Protocol                string          `json:"protocol"`
	Port                    int             `json:"port"`
	Timeout                 int             `json:"timeout"`
	SkipSslCertVerification bool            `json:"skip_ssl_cert_verification"`
	Logging                 map[string]bool `json:"logging"`
	DisableReconnect        bool            `json:"disable_reconnect"`

	// Advanced user config.
	Transport *http.Transport `json:"-"`

	// Set at runtime.
	JsonWebToken string `json:"json_web_token"`
	con          *http.Client

	// Variables for testing.
	pcAuthFileContent []byte
	pcResponses       []*http.Response
	pcResponseIndex   int
}

/*
Initialize prepares the client connection and attempts a login.

This function can be passed a credentials file to read in settings
that will act as defaults if the given variables are currently unset.
*/
func (c *Client) Initialize(filename string) error {
	c2 := Client{}

	if filename != "" {
		var (
			b   []byte
			err error
		)

		if len(c.pcResponses) == 0 {
			b, err = ioutil.ReadFile(filename)
		} else {
			b, err = c.pcAuthFileContent, nil
		}

		if err != nil {
			return err
		}

		if err = json.Unmarshal(b, &c2); err != nil {
			return err
		}
	}

	if len(c.Logging) == 0 {
		if len(c2.Logging) > 0 {
			c.Logging = make(map[string]bool)
			for key, val := range c2.Logging {
				c.Logging[key] = val
			}
		} else {
			c.Logging = map[string]bool{LogAction: true}
		}
	}

	var tout time.Duration
	if c.Timeout == 0 {
		if c2.Timeout > 0 {
			c.Timeout = c2.Timeout
		} else {
			c.Timeout = 90
		}
	}
	if c.Timeout < 0 {
		return fmt.Errorf("Invalid timeout")
	}
	tout = time.Duration(time.Duration(c.Timeout) * time.Second)

	if c.Port == 0 {
		if c2.Port != 0 {
			c.Port = c2.Port
		}
	}
	if c.Port > 65535 || c.Port < 0 {
		return fmt.Errorf("Invalid port number")
	}

	if c.Protocol == "" {
		if c2.Protocol != "" {
			c.Protocol = c2.Protocol
		} else {
			c.Protocol = "https"
		}
	}
	if c.Protocol != "http" && c.Protocol != "https" {
		return fmt.Errorf("Invalid protocol")
	}

	if c.Url == "" && c2.Url != "" {
		c.Url = c2.Url
	}
	if strings.HasPrefix(c.Url, "http://") || strings.HasPrefix(c.Url, "https://") {
		return fmt.Errorf("Specify protocol using the Protocol param, not as the URL")
	}
	c.Url = strings.TrimRight(c.Url, "/")
	if c.Url == "" {
		return fmt.Errorf("Prisma Cloud URL is not set")
	}

	if c.Username == "" && c2.Username != "" {
		c.Username = c2.Username
	}

	if c.Password == "" && c2.Password != "" {
		c.Password = c2.Password
	}

	if c.CustomerName == "" && c2.CustomerName != "" {
		c.CustomerName = c2.CustomerName
	}

	if c.Transport == nil {
		c.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: c.SkipSslCertVerification,
			},
			Proxy: http.ProxyFromEnvironment,
		}
	}

	c.con = &http.Client{
		Transport: c.Transport,
		Timeout:   tout,
	}

	if c.JsonWebToken == "" && c2.JsonWebToken != "" {
		c.JsonWebToken = c2.JsonWebToken
		return nil
	}

	return c.Authenticate()
}

// Authenticate retrieves and saves a JSON web token from Prisma Cloud.
func (c *Client) Authenticate() error {
	var err error

	type initial struct {
		Username     string `json:"username"`
		Password     string `json:"password"`
		CustomerName string `json:"customerName,omitempty"`
	}

	ans := AuthResponse{}

	/*
	   Ideally we would do a JSON web token refresh if we got one
	   previously, but just straight up performing a re-authentication
	   always works.  So just do a full login again if we have the
	   username and password, falling back to a token refresh.
	*/
	if c.Username != "" && c.Password != "" {
		c.Log(LogAction, "(auth) retrieving jwt")
		req := initial{c.Username, c.Password, c.CustomerName}
		_, err = c.communicate("POST", []string{"login"}, nil, &req, &ans, false)
	} else if c.JsonWebToken != "" {
		c.Log(LogAction, "(auth) refreshing jwt")
		_, err = c.communicate("GET", []string{"auth_token", "extend"}, nil, nil, &ans, false)
	} else {
		return fmt.Errorf("no authentication params given")
	}

	if err != nil {
		return err
	}

	c.JsonWebToken = ans.Token
	return nil
}

/*
Communicate handles basic communication with Prisma Cloud.

If a non-nil interface is given as the "ans" variable, then this function
will unmarshal the returned JSON into it, and you can safely discard the
slice of bytes returned.
*/
func (c *Client) Communicate(method string, suffix []string, query, data interface{}, ans interface{}) ([]byte, error) {
	return c.communicate(method, suffix, query, data, ans, true)
}

// Log logs a message to the user if the appropriate style is enabled.
func (c *Client) Log(flag, msg string, i ...interface{}) {
	if c.Logging[flag] {
		log.Printf(msg, i...)
	}
}

func (c *Client) do(req *http.Request) (*http.Response, error) {
	if len(c.pcResponses) == 0 {
		return c.con.Do(req)
	}

	resp := c.pcResponses[c.pcResponseIndex]
	c.pcResponseIndex++
	return resp, nil
}

// logSendReceive outputs raw data sent and received, but tries to
// remove sensitive information from being printed.
func (c *Client) logSendReceive(logFlag string, code int, b []byte) {
	var desc string

	switch logFlag {
	case LogSend:
		desc = "sending:\n"
	case LogReceive:
		desc = fmt.Sprintf("received (%d):", code)
	default:
		return
	}

	if !c.Logging[logFlag] {
		return
	} else if len(b) == 0 {
		log.Printf("%s", desc)
		return
	}

	var ti interface{}
	if err := json.Unmarshal(b, &ti); err != nil {
		log.Printf("failed to unmarshal %s: %s", logFlag, err)
		log.Printf("%s\n%s", desc, scrubSensitiveData(b))
		return
	}

	b2, _ := json.MarshalIndent(ti, "", "    ")
	log.Printf("%s\n%s", desc, scrubSensitiveData(b2))
}

func (c *Client) communicate(method string, suffix []string, query, data interface{}, ans interface{}, allowRetry bool) ([]byte, error) {
	var err error
	var buf bytes.Buffer

	if data != nil {
		b, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		buf = *bytes.NewBuffer(b)
		c.logSendReceive(LogSend, 0, b)
	}

	var path strings.Builder
	path.Grow(30)
	fmt.Fprintf(&path, "%s://%s", c.Protocol, c.Url)
	if c.Port != 0 {
		fmt.Fprintf(&path, ":%d", c.Port)
	}
	for _, v := range suffix {
		path.WriteString("/")
		path.WriteString(v)
	}
	if query != nil {
		qv := query.(url.Values)
		path.WriteString("?")
		path.WriteString(qv.Encode())
	}
	c.Log(LogPath, "path: %s", path.String())

	req, err := http.NewRequest(method, path.String(), &buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if c.JsonWebToken != "" {
		req.Header.Set("x-redlock-auth", c.JsonWebToken)
	}

	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	c.logSendReceive(LogReceive, resp.StatusCode, []byte(body))

	switch resp.StatusCode {
	case http.StatusOK, http.StatusNoContent, http.StatusCreated:
		// Alert rule deletion returns StatusNoContent
	case http.StatusUnauthorized:
		if !c.DisableReconnect && allowRetry {
			if err = c.Authenticate(); err == nil {
				return c.communicate(method, suffix, query, data, ans, false)
			}
		}
		return body, InvalidCredentialsError
	default:
		errLocation := "X-Redlock-Status"
		if _, ok := resp.Header[errLocation]; !ok {
			return body, fmt.Errorf("%d error without the %q header - returned HTML:\n%s", resp.StatusCode, errLocation, body)
		}
		pcel := PrismaCloudErrorList{
			Method:     method,
			StatusCode: resp.StatusCode,
			Path:       path.String(),
		}
		info := resp.Header[errLocation][0]
		if err = json.Unmarshal([]byte(info), &pcel.Errors); err != nil {
			return body, fmt.Errorf("%d error, and could not unmarshal header %q: %s", resp.StatusCode, info, err)
		}
		if ce := pcel.GenericError(); ce != nil {
			return body, ce
		}
		return body, pcel
	}

	if ans != nil {
		if err = json.Unmarshal(body, ans); err != nil {
			return body, err
		}
	}

	return body, nil
}
