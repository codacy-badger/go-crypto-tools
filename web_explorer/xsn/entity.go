package xsn

type Transaction struct {
	ID            string `json:"id"`
	Size          int    `json:"size"`
	Blockhash     string `json:"blockhash"`
	Time          int    `json:"time"`
	Blocktime     int    `json:"blocktime"`
	Confirmations int    `json:"confirmations"`
	Input         []struct {
		Address string  `json:"address"`
		Value   float64 `json:"value"`
	} `json:"input"`
	Output []struct {
		Address string  `json:"address"`
		Value   float64 `json:"value"`
	} `json:"output"`
}

type RawTransaction struct {
	Txid     string `json:"txid"`
	Hash     string `json:"hash"`
	Version  int    `json:"version"`
	Size     int    `json:"size"`
	Vsize    int    `json:"vsize"`
	Weight   int    `json:"weight"`
	Locktime int    `json:"locktime"`
	Vin      []struct {
		Txid      string `json:"txid"`
		Vout      int    `json:"vout"`
		ScriptSig struct {
			Asm string `json:"asm"`
			Hex string `json:"hex"`
		} `json:"scriptSig"`
		Txinwitness []string `json:"txinwitness"`
		Sequence    int64    `json:"sequence"`
	} `json:"vin"`
	Vout []struct {
		Value        float64 `json:"value"`
		N            int     `json:"n"`
		ScriptPubKey struct {
			Asm       string   `json:"asm"`
			Hex       string   `json:"hex"`
			ReqSigs   int      `json:"reqSigs"`
			Type      string   `json:"type"`
			Addresses []string `json:"addresses"`
		} `json:"scriptPubKey"`
	} `json:"vout"`
	Hex           string `json:"hex"`
	Blockhash     string `json:"blockhash"`
	Confirmations int    `json:"confirmations"`
	Time          int    `json:"time"`
	Blocktime     int    `json:"blocktime"`
}

type Address struct {
	Address   string  `json:"address"`
	Received  float64 `json:"received"`
	Spent     float64 `json:"spent"`
	Available int     `json:"available"`
}

type AddressTransactions struct {
	Offset int                       `json:"offset"`
	Limit  int                       `json:"limit"`
	Total  int                       `json:"total"`
	Data   []AddressTransactionsItem `json:"data"`
}

type AddressTransactionsV2 struct {
	Data []struct {
		ID        string        `json:"id"`
		Blockhash string        `json:"blockhash"`
		Size      int           `json:"size"`
		Time      int           `json:"time"`
		Inputs    []interface{} `json:"inputs"`
		Outputs   []struct {
			Index int     `json:"index"`
			Value float64 `json:"value"`
		} `json:"outputs"`
	} `json:"data"`
}

type AddressTransactionsItem struct {
	ID        string  `json:"id"`
	Blockhash string  `json:"blockhash"`
	Time      int     `json:"time"`
	Size      int     `json:"size"`
	Sent      float64 `json:"sent"`
	Received  float64 `json:"received"`
}

type AddressUTXOs []AddressUTXOItem

type AddressUTXOItem struct {
	Satoshis    int64  `json:"satoshis"`
	Txid        string `json:"txid"`
	OutputIndex int    `json:"outputIndex"`
	Address     string `json:"address"`
	Script      string `json:"script"`
}

type Blocks []Block

type Block struct {
	Hash              string   `json:"hash"`
	PreviousBlockhash string   `json:"previousBlockhash"`
	MerkleRoot        string   `json:"merkleRoot"`
	Transactions      []string `json:"transactions"`
	Confirmations     int      `json:"confirmations"`
	Size              int      `json:"size"`
	Height            int      `json:"height"`
	Version           int      `json:"version"`
	Time              int      `json:"time"`
	MedianTime        int      `json:"medianTime"`
	Nonce             int      `json:"nonce"`
	Bits              string   `json:"bits"`
	Chainwork         string   `json:"chainwork"`
	Difficulty        float64  `json:"difficulty"`
	NextBlockhash     string   `json:"nextBlockhash,omitempty"`
	TposContract      string   `json:"tposContract,omitempty"`
}

type BlockInfo struct {
	Block struct {
		Hash              string   `json:"hash"`
		PreviousBlockhash string   `json:"previousBlockhash"`
		NextBlockhash     string   `json:"nextBlockhash"`
		MerkleRoot        string   `json:"merkleRoot"`
		Transactions      []string `json:"transactions"`
		Confirmations     int      `json:"confirmations"`
		Size              int      `json:"size"`
		Height            int      `json:"height"`
		Version           int      `json:"version"`
		Time              int      `json:"time"`
		MedianTime        int      `json:"medianTime"`
		Nonce             int      `json:"nonce"`
		Bits              string   `json:"bits"`
		Chainwork         string   `json:"chainwork"`
		Difficulty        float64  `json:"difficulty"`
	} `json:"block"`
	Rewards struct {
		Coinstake struct {
			Address string `json:"address"`
			Value   int    `json:"value"`
		} `json:"coinstake"`
		Masternode struct {
			Address string `json:"address"`
			Value   int    `json:"value"`
		} `json:"masternode"`
	} `json:"rewards"`
}

type RawBlock struct {
	Hash              string   `json:"hash"`
	Confirmations     int      `json:"confirmations"`
	Strippedsize      int      `json:"strippedsize"`
	Size              int      `json:"size"`
	Weight            int      `json:"weight"`
	Height            int      `json:"height"`
	Version           int      `json:"version"`
	VersionHex        string   `json:"versionHex"`
	Merkleroot        string   `json:"merkleroot"`
	Tx                []string `json:"tx"` //transaction json attribute different naming from other block entities
	Time              int      `json:"time"`
	Mediantime        int      `json:"mediantime"`
	Nonce             int      `json:"nonce"`
	Bits              string   `json:"bits"`
	Difficulty        float64  `json:"difficulty"`
	Chainwork         string   `json:"chainwork"`
	Previousblockhash string   `json:"previousblockhash"`
	Nextblockhash     string   `json:"nextblockhash"`
}

type BlockTransactions struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
	Total  int `json:"total"`
	Data   []struct {
		ID        string  `json:"id"`
		Blockhash string  `json:"blockhash"`
		Time      int     `json:"time"`
		Size      int     `json:"size"`
		Sent      float64 `json:"sent"`
		Received  float64 `json:"received"`
	} `json:"data"`
}

type Stats struct {
	Blocks            int     `json:"blocks"`
	Masternodes       int     `json:"masternodes"`
	CirculatingSupply float64 `json:"circulatingSupply"`
	TotalSupply       float64 `json:"totalSupply"`
	Transactions      int     `json:"transactions"`
}

type Balance struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
	Total  int `json:"total"`
	Data   []struct {
		Address   string  `json:"address"`
		Received  float64 `json:"received"`
		Spent     float64 `json:"spent"`
		Available float64 `json:"available"`
	} `json:"data"`
}

type Masternodes struct {
	Offset int          `json:"offset"`
	Limit  int          `json:"limit"`
	Total  int          `json:"total"`
	Data   []Masternode `json:"data"`
}

type Masternode struct {
	Txid          string `json:"txid"`
	IP            string `json:"ip"`
	Protocol      string `json:"protocol"`
	Status        string `json:"status"`
	ActiveSeconds int    `json:"activeSeconds"`
	LastSeen      int    `json:"lastSeen"`
	Payee         string `json:"payee"`
}
