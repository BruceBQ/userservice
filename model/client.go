package model

const (
	HEADER_BEARER = "BEARER"
	HEADER_AUTH   = "Authorization"

	API_URL_SUFFIX_V1 = "/api/v1"
	API_URL_SUFFIX    = API_URL_SUFFIX_V1
)

type Client struct {
	Url        string
	ApiUrl     string
	AuthToken  string
	HttpHeader map[string]string
}

func (c *Client) SetToken(token string) {
	c.AuthToken = token
}
