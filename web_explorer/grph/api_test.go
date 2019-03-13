package grph

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
	if int(d) < 176700 {
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
	if hash != "0000009db21c27d26e76c8a7cb3b1f3459ea9c434cbf434e6b18d926a1821f78" {
		t.Fail()
		return
	}
	b, err := client.GetBlockByHash(hash)
	if err != nil {
		t.Fail()
		return
	}
	if b.Hash != "0000009db21c27d26e76c8a7cb3b1f3459ea9c434cbf434e6b18d926a1821f78" {
		t.Fail()
		return
	}
	if b.MerkleRoot != "be31d803d502a5096871bc489fc56f45c7896217a86510696d373e35359c1282" {
		t.Fail()
		return
	}
	if b.PreviousBlockHash != "0000009631c3ad018ef7a594a2d72ddacef098828f9e69b07a41d6da32c21c37" {
		t.Fail()
		return
	}
}

func TestAPI_TransactionById(t *testing.T) {
	client := InitApi(APIUrl)
	b, err := client.GetTransactionById("f26936cebf7e1bb251c532890039021d556d843e038609c0b321ae2369058c23", false)
	if err != nil {
		t.Fail()
		return
	}

	if b.BlockHash != "4ef7742cf0a8288899a83965c505f119bb2822059fcc203c705cebf553dbdef7" {
		t.Fail()
		return
	}
	if b.Vout[1].ScriptPubKey.Addresses[0] != "g3LKcvuusEkHE31LSdL32XFcC6eUJyQRuy" {
		t.Fail()
		return
	}
}

func TestAPI_GetAddress(t *testing.T) {
	client := InitApi(APIUrl)
	w, err := client.GetAddress("gFRUAyksVe8Wd3fSXXjUWrQ4EPU3AW7W8j")
	if err != nil {
		t.Fail()
		return
	}
	if w.Address != "gFRUAyksVe8Wd3fSXXjUWrQ4EPU3AW7W8j" {
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
	b, err := client.GetBalance("AH5nHnkL2Was7u23g7idrz7iJSgtRWFJSD")
	if err != nil {
		t.Fail()
		return
	}
	if b < 0 {
		t.Fail()
		return
	}
}
