package larry

import (
	"encoding/json"
	"github.com/bxcodec/httpcache"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Client struct {
	HttpClient *http.Client
}

type ClientOption func(*Client)

func NewClient(accessToken, accessTokenSecret string, opts ...ClientOption) (*Client, error) {
	c := new(Client)

	credentials := &clientcredentials.Config{
		ClientID:     accessToken,
		ClientSecret: accessTokenSecret,
		TokenURL:     "https://api.twitter.com/oauth2/token",
		AuthStyle:    oauth2.AuthStyleInParams,
	}

	c.HttpClient = credentials.Client(oauth2.NoContext)

	for _, opt := range opts {
		opt(c)
	}

	return c, nil
}

func WithCache() ClientOption {
	return func(c *Client) {
		_, err := httpcache.NewWithInmemoryCache(c.HttpClient, true, time.Minute*10)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (c *Client) MakeRequest(req *Request, v interface{}) error {
	resp, err := c.HttpClient.Get(req.Path)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, v)
	if err != nil {
		return err
	}
	return nil
}
