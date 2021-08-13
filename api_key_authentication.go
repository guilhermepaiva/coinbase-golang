package coinbasegolang

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"strconv"
	"time"

	"github.com/fabioberger/coinbase-go/config"
)

type apiKeyAuthentication struct {
	Key     string
	Secret  string
	BaseUrl string
	Client  http.Client
}

func apiKeyAuth(key string, secret string) *apiKeyAuthentication {
	a := apiKeyAuthentication{
		Key:     key,
		Secret:  secret,
		BaseUrl: config.BaseUrl,
		Client: http.Client{
			Transport: &http.Transport{
				Dial: dialTimeout,
			},
		},
	}
	return &a
}

func (a apiKeyAuthentication) authenticate(req *http.Request, endpoint string, params []byte) error {

	nonce := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	message := nonce + endpoint + string(params) //As per Coinbase Documentation

	req.Header.Set("ACCESS_KEY", a.Key)

	h := hmac.New(sha256.New, []byte(a.Secret))
	h.Write([]byte(message))

	signature := hex.EncodeToString(h.Sum(nil))

	req.Header.Set("ACCESS_SIGNATURE", signature)
	req.Header.Set("ACCESS_NONCE", nonce)

	return nil
}

func (a apiKeyAuthentication) getBaseUrl() string {
	return a.BaseUrl
}

func (a apiKeyAuthentication) getClient() *http.Client {
	return &a.Client
}
