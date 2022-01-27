package mipush

import (
	"log"

	"github.com/Scorpio69t/go-mipush/httpclient"
)

type MIPushClient struct {
	Client    *httpclient.HTTPClient
	AppSecret string
	PushPath  string
}

func NewMIPushClient(appSecret string) *MIPushClient {
	client, err := httpclient.NewHTTPClient()
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &MIPushClient{
		Client:    client,
		AppSecret: appSecret,
	}
}
