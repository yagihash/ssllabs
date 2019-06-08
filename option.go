package ssllabs

type Option func(*Client)

// OptionBaseURL is the functional option to change the base URL for API access
func OptionBaseURL(baseURL string) func(*Client) {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}
