package main

import (
	"fmt"

	"github.com/cnhgscc/mirror/pkg/httpd"
)

func main() {
	resp := httpd.NewHttpd(httpd.NewRequest("https://www.baidu.com"))
	fmt.Println(resp.Text())
}
