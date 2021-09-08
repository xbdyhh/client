package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type  EthInteractiveBotClient struct {
	Url string//域名网址
	ContentType string
	UrlData map[string]string//网址中需要的搜索字段
	Data map[interface{}]interface{}//使用POST方法时的传输字段
	Response []byte//返回值
}

func NewEthINteractiveBotClient() *EthInteractiveBotClient {
	return &EthInteractiveBotClient{
		Url:      "",
		ContentType: "application/json",
		UrlData: nil,
		Data:     nil,
		Response: nil,
	}
}


func (ECli *EthInteractiveBotClient)Get() error{
	var data  url.Values
	if ECli.UrlData != nil{
		for k,v := range ECli.UrlData{
			data.Set(k,v)
		}
	}
	rawQuery := data.Encode()
	u, err := url.ParseRequestURI(ECli.Url)
	if err != nil{
		return err
	}
	u.RawQuery = rawQuery
	resp, err := http.Get(u.String())
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return err
	}
	ECli.Response = body
	if resp.Body == nil{
		fmt.Print("body is nil")
	}else {
		fmt.Printf("body is %v",string(body))
	}
	return nil
}

func (ECli *EthInteractiveBotClient)Post()error{
	var data  url.Values
	if ECli.UrlData != nil{
		for k,v := range ECli.UrlData{
			data.Set(k,v)
		}
	}
	rawQuery := data.Encode()
	u, err := url.ParseRequestURI(ECli.Url)
	if err != nil{
		return err
	}
	u.RawQuery = rawQuery
	s,err :=json.Marshal(ECli.Data)
	if err != nil{
		return err
	}
	resp,err :=http.Post(u.String(),ECli.ContentType,strings.NewReader(string(s)))
	if err != nil{
		return err
	}
	if resp.StatusCode!= 200{
		return errors.New("connect err happened！！！！")
	}
	ECli.Response,err = ioutil.ReadAll(resp.Body)
	if err != nil{
		return err
	}
	return nil
}