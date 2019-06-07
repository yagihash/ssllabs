package ssllabs

type Option func(*Client)

func OptionBaseURL(baseURL string) func(*Client) {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}
