package meerkat

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

const (
	// LibraryVersion represents this library version
	LibraryVersion = "0.1"

	// UserAgent represents this client User-Agent
	UserAgent = "go-meerkat v" + LibraryVersion
)

// A Client manages communication with the Meerkat API.
type Client struct {
	// HTTP client used to communicate with the API.
	client *http.Client

	// UserAgent agent used when communicating with Meerkat API.
	UserAgent string

	Profiles   *ProfileService
	Broadcasts *BroadcastService
}

// Response specifies Meerkat's response structure.
type Response struct {
	Response        *http.Response    // HTTP response
	Result          interface{}       `json:"result,omitempty"`
	FollowupActions map[string]string `json:"followupActions,omitempty"`
}

// NewClient returns a new Meerkat API client. if a nil httpClient is
// provided, http.DefaultClient will be used.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	c := &Client{
		client:    httpClient,
		UserAgent: UserAgent,
	}

	c.Profiles = &ProfileService{Client: c}
	c.Broadcasts = &BroadcastService{Client: c}

	return c
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash. If
// specified
func (c *Client) NewRequest(method, urlStr string, body string) (*http.Request, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	// q := u.Query()
	// u.RawQuery = q.Encode()

	req, err := http.NewRequest(method, u.String(), bytes.NewBufferString(body))
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", c.UserAgent)
	return req, nil
}

// Do sends an API request and returns the API response. The API response is
// decoded and stored in the value pointed to by v, or returned as an error if
// an API error has occurred.
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// err = CheckResponse(resp)
	// if err != nil {
	// 	return resp, err
	// }

	r := &Response{Response: resp}
	if v != nil {
		r.Result = v
		err = json.NewDecoder(resp.Body).Decode(r)
		// c.Response = r
	}
	return resp, err
}
