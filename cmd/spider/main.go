package main

import (
	"fmt"

	"github.com/cnhgscc/mirror/pkg/httpd"
)

func main() {

	resp := map[string]any{}
	reply, err := httpd.NewURL(httpd.NewRequest("https://www.baidu.com")).JSONRespRender(&resp)
	if err != nil {
		return
	}
	fmt.Println(reply)

}
