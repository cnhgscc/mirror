package httpd

import "net/http"

var (
	client *Client
)

func init() {
	client = NewClient()
}

func NewURL(req *Request, header ...Header) *Resp {
	// req.Request.Method = strings.ToUpper(method)
	for _, h := range header {
		h(&req.Header)
	}
	return client.NewURL(req)
}

type Header func(header *http.Header)
