package ns

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
)

// NetScaler represents a NetScaler instance
type NetScaler struct {
	Server, User, Password, SessionID string
}

// APIResp is an API response from NetScaler API
type APIResp struct {
	RespBody   string
	StatusCode int
}

// API sends an API call to a NetScaler instance
func (n *NetScaler) API(endpoint, action string, body []byte) (APIResp, error) {
	// Create HTTP client and disable TLS verification for self-signed certs
	jar, err := cookiejar.New(nil)
	if err != nil {
		return APIResp{}, fmt.Errorf("creating cookie jar - %s", err)
	}
	client := &http.Client{Jar: jar, Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}

	// Create the HTTP request with the content-type header
	req, err := http.NewRequest(action, fmt.Sprintf("https://%s/nitro/v1/config/%s", n.Server, endpoint), bytes.NewBuffer(body))
	if err != nil {
		return APIResp{}, fmt.Errorf("creating http request - %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	if n.SessionID != "" {
		req.AddCookie(&http.Cookie{Name: "NITRO_AUTH_TOKEN", Value: n.SessionID})
	}

	// Make the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		return APIResp{StatusCode: resp.StatusCode}, fmt.Errorf("making http request - %s", err)
	}

	// Process response
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return APIResp{}, fmt.Errorf("processing response - %s", err)
	}

	return APIResp{StatusCode: resp.StatusCode, RespBody: string(data[:])}, nil
}
