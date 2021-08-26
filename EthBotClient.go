package go_http_client

import (
	"net/http"
	"net/url"
)

type  EthInteractiveBotClient struct {
	Url string
	UrlData map[string]string
	Data map[interface{}]interface{}
	Response *http.Response
}

func NewEthINteractiveBotClient() *EthInteractiveBotClient {
	return &EthInteractiveBotClient{
		Url:      "",
		UrlData: nil,
		Data:     nil,
		Response: nil,
	}
}


func (ECli *EthInteractiveBotClient)Get() error{
	var data  url.Values
	for k,v := range ECli.UrlData{
		data.Set(k,v)
	}
	rawQuery := data.Encode()
	u, err := url.ParseRequestURI(ECli.Url)
	if err != nil{
		return err
	}
	u.RawQuery = rawQuery
	resp, err := http.Get(u.String())
	if err != nil{
		return err
	}
	ECli.Response = resp
	return nil
}