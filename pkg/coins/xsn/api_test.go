package xsn

import (
	"net/url"
	"strings"
	"testing"
	"fmt"
)

func TestAPI_TransactionById(t *testing.T) {
	testTxId := "1f39356a0db916a50dc303c20622edf107fae09d37c949c26ee790853e74e983"
	client := InitApi(APIUrl)
	tx, err := client.GetTransactionById(testTxId)
	if err != nil {
		t.Fail()
		return
	}
	if tx.ID != testTxId {
		t.Fail()
		return
	}
}

func TestAPI_RawTransactionById(t *testing.T) {
	testTxId := "1f39356a0db916a50dc303c20622edf107fae09d37c949c26ee790853e74e983"
	client := InitApi(APIUrl)
	tx, err := client.GetRawTransactionById(testTxId)
	if err != nil {
		t.Fail()
		return
	}
	if tx.Txid != testTxId {
		t.Fail()
		return
	}
	if tx.Blockhash != "c9a89cddc18c67969c0dfa26efeb78d5350435c30bb4d27c9aee18fa02cf7a90" {
		t.Fail()
		return
	}
}

func TestAPI_GetAddress(t *testing.T) {
	testAddress := "7iacT67sNWkm1JKW1dBwpEN4BA7ZsdpPMW"
	client := InitApi(APIUrl)
	addr, err := client.GetAddress(testAddress)
	if err != nil {
		t.Fail()
		return
	}
	if addr.Address != testAddress {
		t.Fail()
		return
	}
}

func TestAPI_GetAddressTransactions(t *testing.T) {
	testAddress := "Xi2xEzvFkPtuK459PYBaangiwbyGHgMHYh"
	client := InitApi(APIUrl)
	addr, err := client.GetAddressTransactions(testAddress, url.Values{})
	if err != nil {
		t.Fail()
		fmt.Printf("err :%v", err)
		return
	}
	if len(addr.Data) == 0 {
		t.Fail()
		return
	}
}

func TestAPI_GetAddressTransactionsWithQuery(t *testing.T) {
	testAddress := "Xi2xEzvFkPtuK459PYBaangiwbyGHgMHYh"
	client := InitApi(APIUrl)
	q := url.Values{}
	q.Set("limit", "50")
	addr, err := client.GetAddressTransactions(testAddress, q)
	if err != nil {
		fmt.Printf("err :%v", err)
		t.Fail()
		return
	}
	if addr.Limit != 50 {
		t.Fail()
		return
	}
}

func TestAPI_GetAddressUTXOs(t *testing.T) {
	testAddress := "Xi2xEzvFkPtuK459PYBaangiwbyGHgMHYh"
	client := InitApi(APIUrl)
	utxos, err := client.GetAddressUTXOs(testAddress)
	if err != nil {
		t.Fail()
		return
	}
	if len(utxos) > 0 {
		if utxos[0].Address != testAddress {
			t.Fail()
			return
		}
	}
}

func TestAPI_GetLatestBlocks(t *testing.T) {
	client := InitApi(APIUrl)
	blocks, err := client.GetLatestBlocks()
	if err != nil {
		t.Fail()
		return
	}

	if len(blocks) < 5 {
		t.Fail()
		return
	}
	for _, b := range blocks {
		if b.Confirmations == 0 {
			t.Fail()
			return
		}
		if b.Hash == "" {
			t.Fail()
			return
		}
		if b.PreviousBlockhash == "" {
			t.Fail()
			return
		}
	}
}

func TestAPI_BlockByHash(t *testing.T) {
	testBlockHash := "b438e6f08812c9591c6f886ef0ba2ce873659a3797b32ba78e2a160a9b43aacf"
	client := InitApi(APIUrl)
	blockInfo, err := client.GetBlockByQuery(testBlockHash)
	if err != nil {
		t.Fail()
		return
	}

	if blockInfo.Block.Hash != testBlockHash {
		t.Fail()
		return
	}

	if blockInfo.Rewards.Coinstake.Value != 9 {
		t.Fail()
		return
	}

	if blockInfo.Rewards.Masternode.Value != 9 {
		t.Fail()
		return
	}
	if len(blockInfo.Block.Transactions) == 0 {
		t.Fail()
		return
	}
}

func TestAPI_RawBlockByHash(t *testing.T) {
	testBlockHash := "b438e6f08812c9591c6f886ef0ba2ce873659a3797b32ba78e2a160a9b43aacf"
	client := InitApi(APIUrl)
	rawBlock, err := client.GetRawBlocksByQuery(testBlockHash)
	if err != nil {
		t.Fail()
		return
	}

	if rawBlock.Hash != testBlockHash {
		t.Fail()
		return
	}

	if rawBlock.Confirmations == 0 {
		t.Fail()
		return
	}
	if len(rawBlock.Tx) == 0 {
		t.Fail()
		return
	}
	if rawBlock.Previousblockhash == "" {
		t.Fail()
		return
	}
}

func TestAPI_BlockTransactionsByHash(t *testing.T) {
	testBlockHash := "b438e6f08812c9591c6f886ef0ba2ce873659a3797b32ba78e2a160a9b43aacf"
	client := InitApi(APIUrl)
	tx, err := client.GetBlocksTransactionsByHash(testBlockHash, url.Values{})
	if err != nil {
		t.Fail()
		return
	}
	if len(tx.Data) == 0 {
		t.Fail()
		return
	}
	for _, txi := range tx.Data {
		if txi.ID == "" || txi.Blockhash == "" {
			t.Fail()
			return
		}
	}
}

func TestAPI_BlockTransactionsByHashWithQuery(t *testing.T) {
	testBlockHash := "b438e6f08812c9591c6f886ef0ba2ce873659a3797b32ba78e2a160a9b43aacf"
	client := InitApi(APIUrl)
	q := url.Values{}
	q.Set("limit", "2")
	q.Set("offset", "1")
	q.Set("orderBy", "time")
	tx, err := client.GetBlocksTransactionsByHash(testBlockHash, q)
	if err != nil {
		t.Fail()
		return
	}
	if len(tx.Data) != 2 {
		t.Fail()
		return
	}
	if tx.Data[0].ID != "491706404c7eadbbff3c68fc21c1be7bca134ad3b9502b626320549f001267db" ||
		tx.Data[0].Blockhash != testBlockHash {
		t.Fail()
		return
	}
	if tx.Data[1].ID != "916819d7fc141b4ee44c2453e5e5a7c40b0dd7c6bd3e3f66413e2f4c741fed79" ||
		tx.Data[1].Blockhash != testBlockHash {
		t.Fail()
		return
	}
}

func TestAPI_Stats(t *testing.T) {
	client := InitApi(APIUrl)
	s, err := client.GetStats()
	if err != nil {
		t.Fail()
		return
	}
	if s.Blocks < 353000 {
		t.Fail()
		return
	}
	if s.Masternodes == 0 {
		t.Fail()
		return
	}
}

func TestAPI_Balance(t *testing.T) {
	client := InitApi(APIUrl)
	q := url.Values{}
	q.Set("limit", "2")
	q.Set("offset", "1")
	tx, err := client.GetBalances(q)
	if err != nil {
		t.Fail()
		return
	}
	if tx.Offset != 1 || tx.Limit != 2 || len(tx.Data) != 2 {
		t.Fail()
		return
	}
}

func TestAPI_GetMasternodes(t *testing.T) {
	client := InitApi(APIUrl)
	q := url.Values{}
	q.Set("limit", "2")
	q.Set("offset", "1")
	q.Set("orderBy", "activeSeconds:desc")
	mns, err := client.GetMasternodes(q)
	if err != nil {
		t.Fail()
		return
	}
	if mns.Offset != 1 || mns.Limit != 2 || len(mns.Data) != 2 {
		t.Fail()
		return
	}

	ip := strings.Split(mns.Data[0].IP, ":")[0] //take first IP found
	m, mErr := client.GetMasternodeByIp(ip)
	if mErr != nil {
		t.Fail()
		return
	}
	if m.Status == "" {
		t.Fail()
		return
	}
}
