package httpd

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type Request struct {
	*http.Request

	ReqCtx    context.Context
	ReqURL    string
	ReqMethod string
	ReqBody   io.Reader
	ReqErr    error

	ReqHeader http.Header

	BasicAuth []string
}

func (r *Request) Error() error {
	return r.ReqErr
}

func NewRequest(url string, opts ...NewReqOption) *Request {
	req := &Request{ReqURL: url, ReqMethod: "GET"}
	for _, opt := range opts {
		opt(req)
	}
	if req.ReqCtx == nil {
		req.ReqCtx = context.Background()
	}

	hreq, err := http.NewRequestWithContext(req.ReqCtx, req.ReqMethod, req.ReqURL, req.ReqBody)
	if err != nil {
		req.ReqErr = err
	}

	req.Request = hreq

	if header := req.ReqHeader.Get("Content-Type"); header != "" {
		req.Request.Header.Set("Content-Type", header)
	}

	if req.BasicAuth != nil {
		req.Request.SetBasicAuth(req.BasicAuth[0], req.BasicAuth[1])
	}

	return req
}

type NewReqOption func(req *Request)

func WithContext(ctx context.Context) NewReqOption {
	return func(req *Request) {
		req.ReqCtx = ctx
	}
}

func WithMethod(method string) NewReqOption {
	return func(req *Request) {
		req.ReqMethod = method
	}
}

func WithJSONBody(payload any) NewReqOption {
	return func(req *Request) {
		tmp, err := json.Marshal(payload)
		if err != nil {
			req.ReqErr = err
			return
		}
		req.ReqHeader = http.Header{}
		req.ReqHeader.Set("Content-Type", "application/json")
		req.ReqBody = bytes.NewBuffer(tmp)
	}
}

func WithBasicAuth(username, password string) NewReqOption {
	return func(req *Request) {
		req.BasicAuth = []string{username, password}
	}
}
