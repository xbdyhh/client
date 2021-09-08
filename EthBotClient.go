package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type  EthInteractiveBotClient struct {
	http.Client
	Request *http.Request
	Query url.Values//url中的query
	ResponseBody []byte

}

func NewEthINteractiveBotClient() *EthInteractiveBotClient {
	return &EthInteractiveBotClient{
		ResponseBody: nil,
	}
}

func (ecli *EthInteractiveBotClient)InitUrl(host,scheme,path,fragment string,Query map[string][]string)  {
	termUrl := new(url.URL)
	termUrl.Host =  host
	termUrl.Scheme = scheme
	termUrl.Path = path
	termUrl.Fragment = fragment
	ecli.Request.URL =termUrl
	ecli.Query = Query
}


func (ECli *EthInteractiveBotClient)Get() error{
	ECli.Request.URL.RawQuery = ECli.Query.Encode()
	resp,err := ECli.Do(ECli.Request)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return err
	}
	ECli.ResponseBody = body
	if resp.Body == nil{
		fmt.Print("body is nil")
	}else {
		fmt.Printf("body is %v",string(body))
	}
	return nil
}
