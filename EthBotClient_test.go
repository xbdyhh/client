package go_http_client

import (
	"fmt"
	"testing"
)

func TestNewEthINteractiveBotClient(t *testing.T) {
	cli := NewEthINteractiveBotClient()
	cli.Url = "http://www.baidu.com/"
	cli.Get()
	if cli.Response == nil{
		t.Errorf("eth bot get err")
	}else {
		fmt.Printf("response is :%v",string(cli.Response))
	}
}
