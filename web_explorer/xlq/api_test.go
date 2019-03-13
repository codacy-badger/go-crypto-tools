package xlq

import (
	"strings"
	"testing"
)

func TestAPI_GetSummary(t *testing.T) {
	client := InitApi(APIUrl)
	tx, err := client.GetSummary()
	if err != nil {
		t.Fail()
		return
	}
	if len(tx.Data) == 0 {
		t.Fail()
		return
	}
	if tx.Data[0].Blockcount < 51600 {
		t.Fail()
		return
	}
}

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
	if d < int64(516600) {
		t.Fail()
		return
	}
}

func TestAPI_NetworkHash(t *testing.T) {
	client := InitApi(APIUrl)
	d, err := client.GetNetworkHash()
	if err != nil {
		t.Fail()
		return
	}
	if d == 0 {
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
	conns, cErr := client.GetConnections()
	if cErr != nil {
		t.Fail()
		return
	}
	if len(conns.Data) == 0 {
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
	if hash != "000000038251a866c4a3612a8251c3799865a264dda2125308386d981c4ed1e9" {
		t.Fail()
		return
	}
}

func TestAPI_BlockByHash(t *testing.T) {
	client := InitApi(APIUrl)
	b, err := client.GetBlockByHash("000000038251a866c4a3612a8251c3799865a264dda2125308386d981c4ed1e9")
	if err != nil {
		t.Fail()
		return
	}
	if b.Hash != "000000038251a866c4a3612a8251c3799865a264dda2125308386d981c4ed1e9" {
		t.Fail()
		return
	}
	if b.Merkleroot != "f1ef6faf093dce21f0667fb8bc9acef714c46ca64e68d4d43e5c534a67281dbe" {
		t.Fail()
		return
	}
	if b.Previousblockhash != "00000006d1bb54343d12cbe1135d4f3eba75e513ca1374dccc6205513cc93e85" {
		t.Fail()
		return
	}
}

func TestAPI_BlockById(t *testing.T) {
	client := InitApi(APIUrl)
	b, err := client.GetBlockById(100)
	if err != nil {
		t.Fail()
		return
	}
	if b.Hash != "000000038251a866c4a3612a8251c3799865a264dda2125308386d981c4ed1e9" {
		t.Fail()
		return
	}
	if b.Merkleroot != "f1ef6faf093dce21f0667fb8bc9acef714c46ca64e68d4d43e5c534a67281dbe" {
		t.Fail()
		return
	}
	if b.Previousblockhash != "00000006d1bb54343d12cbe1135d4f3eba75e513ca1374dccc6205513cc93e85" {
		t.Fail()
		return
	}
}

func TestAPI_TransactionById(t *testing.T) {
	client := InitApi(APIUrl)
	b, err := client.GetTransactionById("40cfee57cb5ea3c154cc82a02e7a9e02fd18f8512d3d4cd26a9d518ee0070981")
	if err != nil {
		t.Fail()
		return
	}
	if b.Hex != "01000000010000000000000000000000000000000000000000000000000000000000000000ffffffff1f02e600045a7cf95908300000141a2f00000d2f6e6f64655374726174756d2f0000000003001cd5a1030000001976a91423cb7766b1d8ed31b65f017a129f5f305f93796388ac00286bee000000001976a9140e5b2429777dca0b484c229c839e87377450c2f688ac0084d717000000001976a914ec3b5679b0320bf3b5b43ba8de6065ea3e928cf788ac00000000" {
		t.Fail()
		return
	}
	if b.Blockhash != "0000000000163fb38294457e53beb9055a28dfe3185147b892e01188c772c790" {
		t.Fail()
		return
	}
	if b.Vout[0].ScriptPubKey.Addresses[0] != "AK393Jj9tCzyQfMJaXLCFEvG2PWWZp94Aj" {
		t.Fail()
		return
	}
	if b.Vout[1].ScriptPubKey.Addresses[0] != "AH5nHnkL2Was7u23g7idrz7iJSgtRWFJSD" {
		t.Fail()
		return
	}
}

func TestAPI_TransactionInfoById(t *testing.T) {
	client := InitApi(APIUrl)
	txInfo, err := client.GetTransactionInfoById("40cfee57cb5ea3c154cc82a02e7a9e02fd18f8512d3d4cd26a9d518ee0070981")
	if err != nil {
		t.Fail()
		return
	}
	if txInfo.Tx.Blockhash != "0000000000163fb38294457e53beb9055a28dfe3185147b892e01188c772c790" {
		t.Fail()
		return
	}

	if txInfo.Tx.Vout[0].Addresses != "AK393Jj9tCzyQfMJaXLCFEvG2PWWZp94Aj" {
		t.Fail()
		return
	}
	if txInfo.Tx.Vout[0].Amount != int64(15600000000) {
		t.Fail()
		return
	}
}

func TestAPI_GetAddNodes(t *testing.T) {
	client := InitApi(APIUrl)
	addNodesContent, err := client.GetAddNodes()
	if err != nil {
		t.Fail()
		return
	}
	if !strings.Contains(addNodesContent, "addnode=") {
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

func TestAPI_WalletByAddress(t *testing.T) {
	client := InitApi(APIUrl)
	w, err := client.GetWalletByAddress("AH5nHnkL2Was7u23g7idrz7iJSgtRWFJSD")
	if err != nil {
		t.Fail()
		return
	}
	if w.Address != "AH5nHnkL2Was7u23g7idrz7iJSgtRWFJSD" {
		t.Fail()
		return
	}
	if len(w.LastTxs) == 0 {
		t.Fail()
		return
	}
}

func TestAPI_WalletTransactionsByAddress(t *testing.T) {
	client := InitApi(APIUrl)
	w, err := client.GetWalletTransactionsByAddress("AH5nHnkL2Was7u23g7idrz7iJSgtRWFJSD")
	if err != nil {
		t.Fail()
		return
	}
	if len(w.Txs) == 0 {
		t.Fail()
		return
	}
}

func TestAPI_WalletBalance(t *testing.T) {
	client := InitApi(APIUrl)
	_, err := client.GetWalletBalance("AH5nHnkL2Was7u23g7idrz7iJSgtRWFJSD")
	if err != nil {
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
	if len(lt.Data) != 2 {
		t.Fail()
		return
	}
}

func TestAPI_GetLastTransactionBetween(t *testing.T) {
	client := InitApi(APIUrl)
	lt, err := client.GetLastTransactionsBetween(1511049600, 1511050800)
	if err != nil {
		t.Fail()
		return
	}
	if len(lt.Data) == 0 {
		t.Fail()
		return
	}
}
