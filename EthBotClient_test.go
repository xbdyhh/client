package client

import (
	"fmt"
	"testing"
)

func TestNewEthINteractiveBotClient(t *testing.T) {
	cli := NewEthINteractiveBotClient()
	fmt.Println("new client success!!")
	cli.InitUrl("baidu.com","http","","", map[string][]string{
		"wd":{"goland"},
	})
	fmt.Println("init url success!!")
	cli.Get()
	fmt.Println("get method success")
	if cli.ResponseBody == nil{
		t.Errorf("eth bot get err")
	}else {
		fmt.Printf("response is :%v",string(cli.ResponseBody))
	}
}
