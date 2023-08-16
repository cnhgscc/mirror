package httpd

var (
	client *Client
)

func init() {
	client = NewClient()
}

func NewURL(req *Request) *Resp {
	return client.NewURL(req)
}
