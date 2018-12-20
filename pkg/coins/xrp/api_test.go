package xrp

import "testing"

func TestAPI_GetAccountTransactions(t *testing.T){
	client := InitApi(APIUrl)
	accTx, err := client.GetAccountTransactions("r3kmLJN5D28dHuH8vZNUZpMC43pEHpaocV")
	if err != nil {
		t.Fail()
		return
	}
	if accTx.Count == 0 {
		t.Fail()
		return
	}
}
