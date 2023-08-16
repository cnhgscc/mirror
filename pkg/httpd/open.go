package httpd

var (
	client *Client
)

func init() {
	client = NewClient()
}

func NewURL(req *Request) *Resp {
	// req.Request.Method = strings.ToUpper(method)
	return client.NewURL(req)
}
