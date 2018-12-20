package grph

import "time"

type Block struct {
	BlockTime           time.Time   `json:"BlockTime"`
	ConfirmationsClient int         `json:"ConfirmationsClient"`
	TimeFromNowUtc      string      `json:"TimeFromNowUtc"`
	DiffBig             int         `json:"DiffBig"`
	DiffSmall           int         `json:"DiffSmall"`
	Tx                  []string    `json:"Tx"`
	Hash                string      `json:"Hash"`
	Confirmations       int         `json:"Confirmations"`
	Size                int         `json:"Size"`
	Height              int         `json:"Height"`
	Version             int         `json:"Version"`
	MerkleRoot          string      `json:"MerkleRoot"`
	Difficulty          float64     `json:"Difficulty"`
	ChainWork           string      `json:"ChainWork"`
	PreviousBlockHash   string      `json:"PreviousBlockHash"`
	NextBlockHash       string      `json:"NextBlockHash"`
	Bits                string      `json:"Bits"`
	Flags               interface{} `json:"Flags"`
	Time                int         `json:"Time"`
	Nonce               string      `json:"Nonce"`
	Additional          interface{} `json:"Additional"`
}

type Transaction struct {
	ID      string `json:"Id"`
	VinList []struct {
		TxID    string  `json:"TxId"`
		Address string  `json:"Address"`
		Amount  float64 `json:"Amount"`
	} `json:"VinList"`
	TxIDLower string `json:"TxIdLower"`
	Height    int    `json:"Height"`
	Hex       string `json:"Hex"`
	Version   int    `json:"Version"`
	LockTime  int    `json:"LockTime"`
	Vin       []struct {
		TxID      string      `json:"TxId"`
		Vout      string      `json:"Vout"`
		ScriptSig interface{} `json:"ScriptSig"`
		CoinBase  string      `json:"CoinBase"`
		Sequence  string      `json:"Sequence"`
	} `json:"Vin"`
	Vout []struct {
		Value        float64 `json:"Value"`
		N            int     `json:"N"`
		ScriptPubKey struct {
			Asm       string   `json:"Asm"`
			Hex       string   `json:"Hex"`
			ReqSigs   int      `json:"ReqSigs"`
			Type      string   `json:"Type"`
			Addresses []string `json:"Addresses"`
		} `json:"ScriptPubKey"`
	} `json:"Vout"`
	BlockHash     string `json:"BlockHash"`
	Confirmations int    `json:"Confirmations"`
	Time          int    `json:"Time"`
	BlockTime     int    `json:"BlockTime"`
	TxID          string `json:"TxId"`
}

type Address struct {
	Address  string  `json:"address"`
	Sent     float64 `json:"sent"`
	Received float64 `json:"received"`
	Balance  float64 `json:"balance"`
	LastTxs  []struct {
		Type      string `json:"type"`
		Addresses string `json:"addresses"`
	} `json:"last_txs"`
}
