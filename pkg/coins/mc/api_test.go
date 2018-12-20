package mc

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
	if hash != "0000049111a2a7ae7f14e049c98fc10c97d20d707d37d9744a6f899913e0b3b1" {
		t.Fail()
		return
	}
	b, err := client.GetBlockByHash(hash)
	if err != nil {
		t.Fail()
		return
	}
	if b.Hash != "0000049111a2a7ae7f14e049c98fc10c97d20d707d37d9744a6f899913e0b3b1" {
		t.Fail()
		return
	}
	if b.Merkleroot != "cca7dd31410bfa23ee7eb000051f50c2c9a345bbf08a2ae0e1bbe9637eeed683" {
		t.Fail()
		return
	}
	if b.Previousblockhash != "000004f8a38e565b2d0eaa3028410fa9b604074cdb1937321f42b47f54fc5871" {
		t.Fail()
		return
	}
}

func TestAPI_TransactionById(t *testing.T) {
	client := InitApi(APIUrl)
	b, err := client.GetTransactionById("407d3c8c0c1cecadc9dc50626ef0aef70159278283a46590287599e9d8b9b0bd", true)
	if err != nil {
		t.Fail()
		return
	}

	if b.Blockhash != "2eac86e77fa9d100cd01f586e31196b1f3e295c4db2417a550a09a647abc331c" {
		t.Fail()
		return
	}

	if b.Vout[2].ScriptPubKey.Hex != "76a9142d9c3cae5e11d8b0442cee04480c8d54eaa6c94d88ac" {
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
	w, err := client.GetAddress("ma9Aii73QibsTCXbwcMg7LCGK7rcfZt2M8")
	if err != nil {
		t.Fail()
		return
	}
	if w.Address != "ma9Aii73QibsTCXbwcMg7LCGK7rcfZt2M8" {
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
	b, err := client.GetBalance("ma9Aii73QibsTCXbwcMg7LCGK7rcfZt2M8")
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
