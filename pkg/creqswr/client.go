package creqswr

import (
	"net/http"
)

type Client struct {
	http.Client
}

func NewClient() *Client {
	return &Client{}
}

func (cli *Client) NewURL(req *Request) *Resp {
	resp := &Resp{}
	if req.ReqErr != nil {
		resp.RespErr = req.ReqErr
		return resp
	}

	resp.Response, resp.RespErr = cli.Do(req.Request)
	return resp
}
