package httpd

import (
	"encoding/json"
	"io"
	"net/http"
)

type Resp struct {
	*http.Response

	RespErr error
}

func (rep *Resp) Str() string {
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

func (rep *Resp) JSONRespRender(v any) string {
	if rep.RespErr != nil {
		return ""
	}
	s := rep.Byte()
	rep.RespErr = json.Unmarshal(s, v)
	return string(s)
}
