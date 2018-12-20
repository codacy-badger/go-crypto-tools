package common

import "time"

type Address struct {
	Address  string
	Sent     float64
	Received float64
	Balance  float64
	LastTxs  []LastTx
}

type LastTx struct {
	Type    string
	Address string
}

type StakeReward struct {
	Address string
	Amount  float64
	Time    time.Time
}

type MasternodeReward struct {
	Address string
	Amount  float64
	Time    time.Time
}

type AddressTransactions struct {
	Address           Address
	Transactions      []Transaction
	Stakes            []StakeReward
	MasternodeRewards []MasternodeReward
}

type Transaction struct {
	ID            string
	Height        int
	Hex           string
	Version       int
	LockTime      int
	VinList       []VinItem
	Vin           []Vin
	Vout          []Vout
	BlockHash     string
	Confirmations int
	Time          int
	BlockTime     int
	TxID          string
}

type VinItem struct {
	TxId    string
	Address string
	Amount  float64
}

type Vin struct {
	TxID      string
	Vout      string
	ScriptSig ScriptSig
	CoinBase  string
	Sequence  string
}

type Vout struct {
	Value        float64
	N            int
	ScriptPubKey ScriptPubKey
}

type ScriptPubKey struct {
	Asm       string
	Hex       string
	ReqSigs   int
	Type      string
	Addresses []string
}

type ScriptSig struct {
	Asm string
	Hex string
}
