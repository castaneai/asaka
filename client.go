package asaka

import (
	"context"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

type Client struct {
	hc   *http.Client
	opts *ClientOption
}

type ClientOption struct {
	Headers map[string]string
	Cookies map[string]http.Cookie
}

func NewClient(hc *http.Client, opts *ClientOption) (*Client, error) {
	if opts == nil {
		opts = &ClientOption{
			Headers: make(map[string]string),
			Cookies: make(map[string]http.Cookie),
		}
	}
	return &Client{hc: hc, opts: opts}, nil
}

func (c *Client) newGETRequest(ctx context.Context, url string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	for key, value := range c.opts.Headers {
		req.Header.Set(key, value)
	}
	for _, cookie := range c.opts.Cookies {
		req.AddCookie(&cookie)
	}
	return req, nil
}

func (c *Client) GetDoc(ctx context.Context, url string) (*goquery.Document, error) {
	req, err := c.newGETRequest(ctx, url)
	if err != nil {
		return nil, err
	}
	res, err := c.hc.Do(req)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		return nil, err
	}
	return doc, nil
}
