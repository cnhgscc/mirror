package httpd

import (
	"fmt"
	"testing"
)

func TestClient_NewRequest(t *testing.T) {

	req := NewRequest("https://www.baidu.com", WithMethod("POST"))
	resp := NewClient().NewURL(req)

	fmt.Println(resp.Status)

}

func TestNewURL(t *testing.T) {

	req := NewRequest("https://www.baidu.com", WithMethod("POST"))
	resp := NewURL(req)
	fmt.Println(resp.Status)

}
