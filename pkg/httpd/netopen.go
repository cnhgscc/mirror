package httpd

import (
	"net/http"
)

var (
	client *Client
)

func init() {
	client = NewClient()
}

func NewURL(req *Requ, header ...Header) *Resp {
	for _, h := range header {
		h(&req.Header)
	}
	return client.NewURL(req)
}

// Header add header for http.Request
type Header func(header *http.Header)
