package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type IReader struct {
	*strings.Reader
}

type  EthInteractiveBotClient struct {
	data url.Values
	http.Client
	Request *http.Request
	Query url.Values//url中的query
	ResponseBody []byte

}

// NewEthINteractiveBotClient 初始化一个http访问客户端
func NewEthINteractiveBotClient() *EthInteractiveBotClient {
	return &EthInteractiveBotClient{
		Client: http.Client{},
		Query: url.Values{},
		Request: &http.Request{},
		ResponseBody: nil,
	}
}

// InitUrl 初始化一个url，从前到后分别为域名 方法 路径 锚点 查询参数
func (ecli *EthInteractiveBotClient)InitUrl(host,scheme,path,fragment string,Query map[string][]string)  {
	termUrl := &url.URL{}
	termUrl.Host =  host
	termUrl.Scheme = scheme
	termUrl.Path = path
	termUrl.Fragment = fragment
	ecli.Request.URL =termUrl
	ecli.Query = Query
}

// AddUrlQuery 用于添加URL中的查询参数
func (ecli *EthInteractiveBotClient)AddUrlQuery(a map[string][]string){
	for key,val := range a{
		if v,ok :=ecli.Query[key];ok{
			ecli.Query[key] = append(v,val...)
		}else {
			ecli.Query[key] = val
		}
	}
}

// DeleteUrlQuery 用于删除url中的查询
func (ecli *EthInteractiveBotClient)DeleteUrlQuery(keys []string)  {
	for _,key:=range keys{
		if _,ok := ecli.Query[key];ok{
			delete(ecli.Query,key)
		}
	}
}

// Get 使用get方法，返回信息储存在ecli的responsebody中
func (ECli *EthInteractiveBotClient)Get() error{
	if ECli.Query!= nil{
		ECli.Request.URL.RawQuery = ECli.Query.Encode()
	}
	fmt.Println("url is",ECli.Request.URL.String())
	ECli.Request.Method = "GET"
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

func (ECli *EthInteractiveBotClient)Post() error{
	if ECli.Query!= nil{
		ECli.Request.URL.RawQuery = ECli.Query.Encode()
	}
	fmt.Println("url is",ECli.Request.URL.String())
	b := ECli.data.Encode()
	var b2 IReader
	b2.Reader = strings.NewReader(b)
	ECli.Request.Body = b2
	ECli.Request.Method = http.MethodPost
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

func (r IReader)Close() error{
	return nil
}
