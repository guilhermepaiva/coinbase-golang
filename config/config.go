package config

var (
	BaseUrl string
	Sandbox = false
)

func init() {
	BaseUrl = "https://api.coinbase.com/v2/"
	if Sandbox == true {
		BaseUrl = "https://api-public.sandbox.pro.coinbase.com/v2/"
	}
}
