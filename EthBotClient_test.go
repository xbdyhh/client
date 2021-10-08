package client

import (
	"fmt"
	"testing"
)

func TestNewEthINteractiveBotClient(t *testing.T) {
	cli := NewEthINteractiveBotClient()
	fmt.Println("new client success!!")
	cli.InitUrl("23.88.111.115:9090","http","/api/v1/query","", map[string][]string{
		"query": {"tendermint_consensus_validator_power"},
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
