package prtx

type Block struct {
	Hash              string   `json:"hash"`
	Confirmations     int      `json:"confirmations"`
	Size              int      `json:"size"`
	Height            int      `json:"height"`
	Version           int      `json:"version"`
	Merkleroot        string   `json:"merkleroot"`
	Tx                []string `json:"tx"`
	Time              int      `json:"time"`
	Nonce             int64    `json:"nonce"`
	Bits              string   `json:"bits"`
	Difficulty        float64  `json:"difficulty"`
	Chainwork         string   `json:"chainwork"`
	Previousblockhash string   `json:"previousblockhash"`
	Nextblockhash     string   `json:"nextblockhash"`
}

type Transaction struct {
	Hex      string `json:"hex"`
	Txid     string `json:"txid"`
	Version  int    `json:"version"`
	Locktime int    `json:"locktime"`
	Vin      []struct {
		Txid      string `json:"txid"`
		Vout      int    `json:"vout"`
		ScriptSig struct {
			Asm string `json:"asm"`
			Hex string `json:"hex"`
		} `json:"scriptSig"`
		Sequence int64 `json:"sequence"`
	} `json:"vin"`
	Vout []struct {
		Value        float64 `json:"value"`
		N            int     `json:"n"`
		ScriptPubKey struct {
			Asm  string `json:"asm"`
			Hex  string `json:"hex"`
			Type string `json:"type"`
		} `json:"scriptPubKey"`
	} `json:"vout"`
	Blockhash     string `json:"blockhash"`
	Confirmations int    `json:"confirmations"`
	Time          int    `json:"time"`
	Blocktime     int    `json:"blocktime"`
}

type Address struct {
	Address  string  `json:"address"`
	Sent     float64 `json:"sent"`
	Received float64 `json:"received"`
	Balance  string  `json:"balance"`
	LastTxs  []struct {
		Addresses string `json:"addresses"`
		Type      string `json:"type"`
	} `json:"last_txs"`
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
			Amount    int    `json:"amount"`
		} `json:"vout"`
		Vin []interface{} `json:"vin"`
	} `json:"data"`
}

type Masternodes struct {
	Total     int `json:"total"`
	Stable    int `json:"stable"`
	Obfcompat int `json:"obfcompat"`
	Enabled   int `json:"enabled"`
	Inqueue   int `json:"inqueue"`
	Ipv4      int `json:"ipv4"`
	Ipv6      int `json:"ipv6"`
	Onion     int `json:"onion"`
}
