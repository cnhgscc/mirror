package main

import (
	"fmt"
	"net/http"

	"github.com/cnhgscc/mirror/pkg/httpd"
)

func main() {

	reply := map[string]any{}

	req := httpd.NewRequest("http://127.0.0.1:8003/api/v1/platform/games", httpd.WithMethod("POST"), httpd.WithJSONBody("3123123"))

	resp := httpd.NewURL(req, func(header *http.Header) {
		header.Set("TX-Client-Version", "1.0.0")
	})

	_, _ = resp.JSONRespRender(&reply)

	fmt.Println(resp.Status)
	fmt.Println(resp, reply)

}
