package grph

import (
	"testing"
	"time"
)

const (
	WaitingDurationDueToIPRequestLimit = time.Second * 0 //moved to client
)

func TestAPI_Difficulty(t *testing.T) {
	time.Sleep(WaitingDurationDueToIPRequestLimit)
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
	time.Sleep(WaitingDurationDueToIPRequestLimit)
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
	time.Sleep(WaitingDurationDueToIPRequestLimit)
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
	time.Sleep(WaitingDurationDueToIPRequestLimit)
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
	time.Sleep(WaitingDurationDueToIPRequestLimit)
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
	time.Sleep(WaitingDurationDueToIPRequestLimit)
	client := InitApi(APIUrl)
	b, err := client.GetTransactionById("f26936cebf7e1bb251c532890039021d556d843e038609c0b321ae2369058c23", false)
	if err != nil {
		t.Fail()
		return
	}

	if b.Hex != "01000000017da776c1396a2036dcc0283358d8728b0cc0fb1e7042c1da6b26a9db8d3ac1d60100000049483045022100d677fab506883523027afc69b316752fbd7c4d7165e2a8b0bdeebafe57dbf01002200fb9a8b431a385230c6dbe731c595bc0060ce2cbb6c3f86a9466996bfad2a57b01ffffffff0300000000000000000060f6f81c250000002321032df8a266910bec6ff4148834ed24171dcda38607c02e827065d4bcd6c4a0d20dac008c8647000000001976a914abde1ab4d271d420255662ad95363c92ca1f45bf88ac00000000" {
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
	time.Sleep(WaitingDurationDueToIPRequestLimit)
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
	time.Sleep(WaitingDurationDueToIPRequestLimit)
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
