package httpd

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type Requ struct {
	*http.Request

	ReqCtx    context.Context
	ReqURL    string
	ReqMethod string
	ReqBody   io.Reader
	ReqErr    error

	ReqHeader http.Header

	BasicAuth []string
}

func (r *Requ) Error() error {
	return r.ReqErr
}

func NewRequest(url string, opts ...NewReqOption) *Requ {
	requ := &Requ{ReqURL: url, ReqMethod: "GET"}
	for _, opt := range opts {
		opt(requ)
	}
	if requ.ReqCtx == nil {
		requ.ReqCtx = context.Background()
	}

	req, err := http.NewRequestWithContext(requ.ReqCtx, requ.ReqMethod, requ.ReqURL, requ.ReqBody)
	if err != nil {
		requ.ReqErr = err
	}

	requ.Request = req

	if header := requ.ReqHeader.Get("Content-Type"); header != "" {
		requ.Request.Header.Set("Content-Type", header)
	}

	if requ.BasicAuth != nil {
		requ.Request.SetBasicAuth(requ.BasicAuth[0], requ.BasicAuth[1])
	}

	return requ
}

type NewReqOption func(requ *Requ)

func WithContext(ctx context.Context) NewReqOption {
	return func(req *Requ) {
		req.ReqCtx = ctx
	}
}

func WithMethod(method string) NewReqOption {
	return func(req *Requ) {
		req.ReqMethod = method
	}
}

func WithJSONBody(payload any) NewReqOption {
	return func(req *Requ) {
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
	return func(req *Requ) {
		req.BasicAuth = []string{username, password}
	}
}
