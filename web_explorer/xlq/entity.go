package xlq

import "time"

type Summary struct {
	Data []struct {
		Difficulty       float64 `json:"difficulty"`
		DifficultyHybrid string  `json:"difficultyHybrid"`
		Supply           float64 `json:"supply"`
		Hashrate         string  `json:"hashrate"`
		LastPrice        int     `json:"lastPrice"`
		Connections      int     `json:"connections"`
		Blockcount       int     `json:"blockcount"`
		AlqoUSD          string  `json:"alqoUSD"`
		AlqoBTC          string  `json:"alqoBTC"`
		MarketCap        float64 `json:"marketCap"`
	} `json:"data"`
}

type Connections struct {
	Data []struct {
		ID        string    `json:"_id"`
		V         int       `json:"__v"`
		Country   string    `json:"country"`
		Version   string    `json:"version"`
		Protocol  string    `json:"protocol"`
		Address   string    `json:"address"`
		CreatedAt time.Time `json:"createdAt"`
	} `json:"data"`
}

type Block struct {
	Hash              string   `json:"hash"`
	Confirmations     int      `json:"confirmations"`
	Size              int      `json:"size"`
	Height            int      `json:"height"`
	Version           int      `json:"version"`
	Merkleroot        string   `json:"merkleroot"`
	AccCheckpoint     string   `json:"acc_checkpoint"`
	Tx                []string `json:"tx"`
	Time              int      `json:"time"`
	Nonce             int      `json:"nonce"`
	Bits              string   `json:"bits"`
	Difficulty        float64  `json:"difficulty"`
	Chainwork         string   `json:"chainwork"`
	Previousblockhash string   `json:"previousblockhash"`
	Nextblockhash     string   `json:"nextblockhash"`
	Moneysupply       float64  `json:"moneysupply"`
	ZXLQsupply        struct {
		Num1    int `json:"1"`
		Num5    int `json:"5"`
		Num10   int `json:"10"`
		Num50   int `json:"50"`
		Num100  int `json:"100"`
		Num500  int `json:"500"`
		Num1000 int `json:"1000"`
		Num5000 int `json:"5000"`
		Total   int `json:"total"`
	} `json:"zXLQsupply"`
}

type Transaction struct {
	Hex      string `json:"hex"`
	Txid     string `json:"txid"`
	Version  int    `json:"version"`
	Locktime int    `json:"locktime"`
	Vin      []struct {
		Coinbase string `json:"coinbase"`
		Sequence int    `json:"sequence"`
	} `json:"vin"`
	Vout []struct {
		Value        int `json:"value"`
		N            int `json:"n"`
		ScriptPubKey struct {
			Asm       string   `json:"asm"`
			Hex       string   `json:"hex"`
			ReqSigs   int      `json:"reqSigs"`
			Type      string   `json:"type"`
			Addresses []string `json:"addresses"`
		} `json:"scriptPubKey"`
	} `json:"vout"`
	Blockhash     string `json:"blockhash"`
	Confirmations int    `json:"confirmations"`
	Time          int    `json:"time"`
	Blocktime     int    `json:"blocktime"`
}

type TransactionInfo struct {
	Tx struct {
		ID         string `json:"_id"`
		Txid       string `json:"txid"`
		Blockhash  string `json:"blockhash"`
		V          int    `json:"__v"`
		Blockindex int    `json:"blockindex"`
		Timestamp  int    `json:"timestamp"`
		Total      int64  `json:"total"`
		Vout       []struct {
			Addresses string `json:"addresses"`
			Amount    int64  `json:"amount"`
		} `json:"vout"`
		Vin []struct {
			Addresses string `json:"addresses"`
			Amount    int64  `json:"amount"`
		} `json:"vin"`
	} `json:"tx"`
	Confirmations int `json:"confirmations"`
	Blockcount    int `json:"blockcount"`
}

type Distribution struct {
	Supply float64 `json:"supply"`
	T125   struct {
		Percent string `json:"percent"`
		Total   string `json:"total"`
	} `json:"t_1_25"`
	T2650 struct {
		Percent string `json:"percent"`
		Total   string `json:"total"`
	} `json:"t_26_50"`
	T5175 struct {
		Percent string `json:"percent"`
		Total   string `json:"total"`
	} `json:"t_51_75"`
	T76100 struct {
		Percent string `json:"percent"`
		Total   string `json:"total"`
	} `json:"t_76_100"`
	T101Plus struct {
		Percent string `json:"percent"`
		Total   string `json:"total"`
	} `json:"t_101plus"`
}

type Wallet struct {
	Address  string `json:"address"`
	Sent     int64  `json:"sent"`
	Received int64  `json:"received"`
	Balance  string `json:"balance"`
	LastTxs  []struct {
		Addresses string `json:"addresses"`
		Type      string `json:"type"`
	} `json:"last_txs"`
}

type WalletTransactions struct {
	Txs []struct {
		ID         string `json:"_id"`
		Txid       string `json:"txid"`
		Blockhash  string `json:"blockhash"`
		V          int    `json:"__v"`
		Blockindex int    `json:"blockindex"`
		Timestamp  int    `json:"timestamp"`
		Total      int64  `json:"total"`
		Vout       []struct {
			Addresses string `json:"addresses"`
			Amount    int64  `json:"amount"`
		} `json:"vout"`
		Vin []struct {
			Addresses string `json:"addresses"`
			Amount    int64  `json:"amount"`
		} `json:"vin"`
	} `json:"txs"`
}

type LastTransaction struct {
	Data []struct {
		ID         string `json:"_id"`
		Txid       string `json:"txid"`
		Blockhash  string `json:"blockhash"`
		V          int    `json:"__v"`
		Blockindex int    `json:"blockindex"`
		Timestamp  int    `json:"timestamp"`
		Total      int64  `json:"total"`
		Vout       []struct {
			Addresses string `json:"addresses"`
			Amount    int64  `json:"amount"`
		} `json:"vout"`
		Vin []struct {
			Addresses string `json:"addresses"`
			Amount    int64  `json:"amount"`
		} `json:"vin"`
	} `json:"data"`
}
