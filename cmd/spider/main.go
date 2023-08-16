package main

import (
	"fmt"

	"github.com/cnhgscc/mirror/pkg/httpd"
)

func main() {
	resp := httpd.NewURL(httpd.NewRequest("https://www.baidu.com", httpd.WithMethod("POST")))
	fmt.Println(resp.Str())
}
