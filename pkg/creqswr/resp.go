package creqswr

import (
	"io"
	"net/http"
)

type Resp struct {
	*http.Response

	RespErr error
}

func (rep *Resp) Text() string {
	tmp, err := io.ReadAll(rep.Response.Body)
	if err != nil {
		return ""
	}
	return string(tmp)
}

func (rep *Resp) Byte() []byte {
	tmp, err := io.ReadAll(rep.Response.Body)
	if err != nil {
		return []byte{}
	}
	return tmp
}
