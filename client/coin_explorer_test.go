package client

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	webApiXZC := NewCoinExplorerWebApiClient()
	mnList, err := webApiXZC.GetMNList()
	if err != nil {
		fmt.Printf("web client err: %v \n", err)
		t.Fail()
		return
	}
	if len(mnList.Result) == 0 {
		t.Fail()
		return
	}
}
