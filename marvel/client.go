package marvel

import (
	"crypto/md5"
	"ioutil"
	"net/http"
)

const baseURL = "https://gateway.marvel.com"

// Client poops
type Client struct {
	httpclient *http.Client
	publicKey  string
	privateKey string
}

// NewClient constructs a new Client struct with the default HTTP client.
func NewClient(publicKey, privateKey string) Client {
	return Client{publicKey, privateKey, http.DefaultClient}
}

func (client Client) sendRequest(method, path string) ([]byte, error) {
	resp, err := http.Get()
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
