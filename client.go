package figma

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"runtime"
)

var baseURL = "https://api.figma.com"
var version = "0.0.1"
var userAgent = fmt.Sprintf("FigmaGoClient/%s (%s)", version, runtime.Version())

type Client struct {
	// Auth token
	token string

	// http client
	httpclient *http.Client
}

// NewClient builds a API client from the provided token and options.
func NewClient(token string) *Client {
	client := &Client{
		token:      token,
		httpclient: http.DefaultClient,
	}

	return client
}

func (c *Client) newRequest(ctx context.Context, method, spath string, body io.Reader) (*http.Request, error) {
	u, err := url.Parse(baseURL)

	if err != nil {
		return nil, err
	}

	u.Path = path.Join(u.Path, spath)

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	req.Header.Set("X-Figma-Token", c.token)
	req.Header.Set("User-Agent", userAgent)

	return req, nil
}

func (c *Client) send(ctx context.Context, request *http.Request) ([]byte, error) {
	resp, _ := c.httpclient.Do(request)
	defer resp.Body.Close()

	byteArray, err := ioutil.ReadAll(resp.Body)
	return byteArray, err
}
