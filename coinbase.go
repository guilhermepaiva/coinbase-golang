package coinbase

type Client struct {
	rpc rpc
}

func ApiKeyClient(key string, secret string) Client {
	c := Client{
		rpc: rpc{
			auth: apiKeyAuth(key, secret),
			mock: false,
		},
	}
	return c
}
