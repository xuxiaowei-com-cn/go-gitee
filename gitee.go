package gitee

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/go-retryablehttp"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

const (
	defaultBaseURL = "https://gitee.com"
	apiVersionPath = "api/v5/"
	userAgent      = "go-gitee"
)

type Client struct {
	client *retryablehttp.Client

	baseURL *url.URL

	token string

	UserAgent string

	GetV5UserRepos *GetV5UserReposService
}

type ListOptions struct {
	Page    int `url:"page,omitempty" json:"page,omitempty"`         // 当前的页码
	PerPage int `url:"per_page,omitempty" json:"per_page,omitempty"` // 每页的数量，最大为 100
}

func NewClient(token string) (*Client, error) {
	client, err := newClient()
	if err != nil {
		return nil, err
	}
	client.token = token
	return client, nil
}

func newClient() (*Client, error) {
	c := &Client{UserAgent: userAgent}

	c.client = &retryablehttp.Client{
		CheckRetry:   c.retryHTTPCheck,
		ErrorHandler: retryablehttp.PassthroughErrorHandler,
		HTTPClient:   cleanhttp.DefaultPooledClient(),
		RetryWaitMin: 100 * time.Millisecond,
		RetryWaitMax: 400 * time.Millisecond,
		RetryMax:     5,
	}

	err := c.setBaseURL(defaultBaseURL)
	if err != nil {
		return nil, err
	}

	c.GetV5UserRepos = &GetV5UserReposService{client: c}

	return c, nil
}

func (c *Client) retryHTTPCheck(ctx context.Context, resp *http.Response, err error) (bool, error) {
	if ctx.Err() != nil {
		return false, ctx.Err()
	}
	if err != nil {
		return false, err
	}
	return false, nil
}

func (c *Client) BaseURL() *url.URL {
	u := *c.baseURL
	return &u
}

func (c *Client) setBaseURL(urlStr string) error {
	if !strings.HasSuffix(urlStr, "/") {
		urlStr += "/"
	}

	baseURL, err := url.Parse(urlStr)
	if err != nil {
		return err
	}

	if !strings.HasSuffix(baseURL.Path, apiVersionPath) {
		baseURL.Path += apiVersionPath
	}

	c.baseURL = baseURL

	return nil
}

func (c *Client) NewRequest(method, path string, request interface{}) (*retryablehttp.Request, error) {
	u := *c.baseURL
	unescaped, err := url.PathUnescape(path)
	if err != nil {
		return nil, err
	}

	u.RawPath = c.baseURL.Path + path
	u.Path = c.baseURL.Path + unescaped

	reqHeaders := make(http.Header)
	reqHeaders.Set("Accept", "application/json")
	reqHeaders.Set("Content-Type", "application/json")

	if c.UserAgent != "" {
		reqHeaders.Set("User-Agent", c.UserAgent)
	}

	var body interface{}

	q, err := query.Values(request)
	if err != nil {
		return nil, err
	}

	q.Set("access_token", c.token)

	u.RawQuery = q.Encode()

	req, err := retryablehttp.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	for k, v := range reqHeaders {
		req.Header[k] = v
	}

	return req, nil
}

type Response struct {
	*http.Response
}

func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	return response
}

func (c *Client) Do(req *retryablehttp.Request, v interface{}) (*Response, error) {

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	defer io.Copy(io.Discard, resp.Body)

	response := newResponse(resp)

	err = CheckResponse(resp)
	if err != nil {
		return response, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
		}
	}

	return response, err
}

type ErrorResponse struct {
	Body     []byte
	Response *http.Response
	Message  string
}

func (e *ErrorResponse) Error() string {
	path, _ := url.QueryUnescape(e.Response.Request.URL.Path)
	u := fmt.Sprintf("%s://%s%s", e.Response.Request.URL.Scheme, e.Response.Request.URL.Host, path)
	return fmt.Sprintf("%s %s: %d %s", e.Response.Request.Method, u, e.Response.StatusCode, e.Message)
}

func CheckResponse(r *http.Response) error {
	switch r.StatusCode {
	case 200, 201, 202, 204, 304:
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	data, err := io.ReadAll(r.Body)
	if err == nil && data != nil {
		errorResponse.Body = data

		var raw interface{}
		if err := json.Unmarshal(data, &raw); err != nil {
			errorResponse.Message = fmt.Sprintf("failed to parse unknown error format: %s", data)
		} else {
			errorResponse.Message = parseError(raw)
		}
	}

	return errorResponse
}

func parseError(raw interface{}) string {
	switch raw := raw.(type) {
	case string:
		return raw

	case []interface{}:
		var errs []string
		for _, v := range raw {
			errs = append(errs, parseError(v))
		}
		return fmt.Sprintf("[%s]", strings.Join(errs, ", "))

	case map[string]interface{}:
		var errs []string
		for k, v := range raw {
			errs = append(errs, fmt.Sprintf("{%s: %s}", k, parseError(v)))
		}
		sort.Strings(errs)
		return strings.Join(errs, ", ")

	default:
		return fmt.Sprintf("failed to parse unexpected error type: %T", raw)
	}
}
