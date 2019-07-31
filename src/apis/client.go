package apis

import (
	"github.com/wwleak/marvel/src/config"
	"net/http"
	"time"
)

type Client struct {
	BaseURL   string

	HttpClient *http.Client
}


func New() *Client{
	conf := config.New()

	return &Client{
		BaseURL: conf.Marvel.Host,
		HttpClient: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}