package prtx

import (
	"testing"
)

func TestAPI_Difficulty(t *testing.T) {
	client := InitApi(APIUrl)
	d, err := client.GetDifficulty()
	if err != nil {
		t.Fail()
		return
	}
	if d == 0 {
		t.Fail()
		return
	}
}

func TestAPI_BlockCount(t *testing.T) {
	client := InitApi(APIUrl)
	d, err := client.GetBlockCount()
	if err != nil {
		t.Fail()
		return
	}
	if int(d) < 80000 {
		t.Fail()
		return
	}
}

func TestAPI_MasternodeCount(t *testing.T) {
	client := InitApi(APIUrl)
	mns, err := client.GetMasternodeCount()
	if err != nil {
		t.Fail()
		return
	}
	if mns.Total < 10 {
		t.Fail()
		return
	}
	if mns.Enabled < 10 {
		t.Fail()
		return
	}
}

func TestAPI_GetConnectionCount(t *testing.T) {
	client := InitApi(APIUrl)
	d, err := client.GetConnectionCount()
	if err != nil {
		t.Fail()
		return
	}
	if d == 0 {
		t.Fail()
		return
	}
}

func TestAPI_BlockHashByIndex(t *testing.T) {
	client := InitApi(APIUrl)
	hash, err := client.GetBlockHashByIndex(100)
	if err != nil {
		t.Fail()
		return
	}
	if hash != "00000388e2cde0f5704bf1e139cb0ee17b006261fae01b81cdc8161b887b2b37" {
		t.Fail()
		return
	}
	b, err := client.GetBlockByHash(hash)
	if err != nil {
		t.Fail()
		return
	}
	if b.Hash != "00000388e2cde0f5704bf1e139cb0ee17b006261fae01b81cdc8161b887b2b37" {
		t.Fail()
		return
	}
	if b.Merkleroot != "fa7e7855a1df35f65af7c797630a8db3aa47d8904ae60398d9aac49fc1f79427" {
		t.Fail()
		return
	}
	if b.Previousblockhash != "0000001fde3992d5c25b4b3dc31614ee22c324dfb7791dbd3e375e8669dd3755" {
		t.Fail()
		return
	}
}

func TestAPI_TransactionById(t *testing.T) {
	client := InitApi(APIUrl)
	b, err := client.GetTransactionById("d72741d5a2a607fd8eb1a1cf03fce77974298f0d552616cf0bbb4ab2c1261249", true)
	if err != nil {
		t.Fail()
		return
	}

	if b.Blockhash != "b16f2dab08b43e7bb95f0c92f8b1d7975243cd6e67cabc5afc22ac08a9acb84e" {
		t.Fail()
		return
	}

	if b.Vout[2].Value != 58.5 {
		t.Fail()
		return
	}
	if b.Vout[2].ScriptPubKey.Hex != "76a91498eebc828d02b4a765d6f6612a934e2a104f8f7688ac" {
		t.Fail()
		return
	}
}

func TestAPI_GetLastTransaction(t *testing.T) {
	client := InitApi(APIUrl)
	lt, err := client.GetLastTransaction(3, 2)
	if err != nil {
		t.Fail()
		return
	}
	if len(lt.Data) == 0 {
		t.Fail()
		return
	}
}

func TestAPI_GetGetNetworkHashRate(t *testing.T) {
	client := InitApi(APIUrl)
	n, err := client.GetGetNetworkHashRate()
	if err != nil {
		t.Fail()
		return
	}

	if n == 0 {
		t.Fail()
		return
	}
}

func TestAPI_GetAddress(t *testing.T) {
	client := InitApi(APIUrl)
	w, err := client.GetAddress("pKVCuZHC3ShTcg7uVcrvgJLtpzAmLC3ZnZ")
	if err != nil {
		t.Fail()
		return
	}
	if w.Address != "pKVCuZHC3ShTcg7uVcrvgJLtpzAmLC3ZnZ" {
		t.Fail()
		return
	}
	if len(w.LastTxs) == 0 {
		t.Fail()
		return
	}
}

func TestAPI_GetBalance(t *testing.T) {
	client := InitApi(APIUrl)
	b, err := client.GetBalance("pKVCuZHC3ShTcg7uVcrvgJLtpzAmLC3ZnZ")
	if err != nil {
		t.Fail()
		return
	}
	if b < 0 {
		t.Fail()
		return
	}
}

func TestAPI_GetDistribution(t *testing.T) {
	client := InitApi(APIUrl)
	d, err := client.GetDistribution()
	if err != nil {
		t.Fail()
		return
	}
	if d.Supply <= 0 {
		t.Fail()
		return
	}
}
