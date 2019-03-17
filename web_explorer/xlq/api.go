package xlq

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	APIUrl = "https://explorer.alqo.org/api"

	//API
	GetSummary         = "/summary"
	GetDifficulty      = "/difficulty"
	GetBlockCount      = "/blockcount"
	GetNetworkHash     = "/networkhash"
	GetConnectionCount = "/connectioncount"
	GetConnections     = "/connections"

	GetBlockHashByIndex = "/blockhash/%d"
	GetBlockByHash      = "/block/%s"
	GetBlockById        = "/blockdata/%d"

	GetTransactionById     = "/transaction/%s"
	GetTransactionInfoById = "/transactioninfo/%s"

	GetAddNodes = "/addNodes"

	GetDistribution                = "/distribution"
	GetWalletByAddress             = "/wallet/%s"
	GetWalletTransactionsByAddress = "/wallettransactions/%s"
	GetWalletBalance               = "/balance/%s"
	GetLastTransactions            = "/lasttxs/%d/%d"
	GetLastTransactionsBetween     = "/lasttxsbetween/%d/%d"
	GetRichlist                    = "/richlist" //seems to be broken ATM
)

//Github repo: https://github.com/ALQOCRYPTO/ALQO/
//API doc:  https://explorer.alqo.org/apiusage
//Explorer: https://explorer.alqo.org/
type API struct {
	baseUrl string
	c       *http.Client
}

//InitApi initializes the client with the given base-url
func InitApi(url string) *API {
	client := &http.Client{}
	return &API{baseUrl: url, c: client}
}

//Returns a basic summary.
func (x *API) GetSummary() (s Summary, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, GetSummary))
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &s)
	return
}

//Returns the current difficulty.
func (x *API) GetDifficulty() (d float64, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, GetDifficulty))
	if err != nil {
		return
	}
	d, err = strconv.ParseFloat(string(body), 16)
	return
}

//Returns the current block index.
func (x *API) GetBlockCount() (d int64, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, GetBlockCount))
	if err != nil {
		return
	}
	d, err = strconv.ParseInt(string(body), 10, 64)
	return
}

//Returns the current network hashrate. (hash/s)
func (x *API) GetNetworkHash() (d int64, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, GetNetworkHash))
	if err != nil {
		return
	}
	d, err = strconv.ParseInt(string(body), 10, 64)
	return
}

//Returns the number of connections the block explorer has to other nodes.
func (x *API) GetConnectionCount() (d int64, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, GetConnectionCount))
	if err != nil {
		return
	}
	d, err = strconv.ParseInt(string(body), 10, 64)
	return
}

//Returns information about nodes the block explorer is connected to.
func (x *API) GetConnections() (c Connections, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, GetConnections))
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &c)
	return
}

//Returns the hash of the block at [index]
func (x *API) GetBlockHashByIndex(blockId int) (blockHash string, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, fmt.Sprintf(GetBlockHashByIndex, blockId)))
	if err != nil {
		return
	}
	return string(body), nil
}

//Returns information about the block with the given hash.
func (x *API) GetBlockByHash(hash string) (b Block, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, fmt.Sprintf(GetBlockByHash, hash)))
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &b)
	return
}

//Returns information about the block with the given index.
func (x *API) GetBlockById(blockId int) (b Block, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, fmt.Sprintf(GetBlockById, blockId)))
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &b)
	return
}

//Returns raw transaction representation for given transaction id.
func (x *API) GetTransactionById(txId string) (t Transaction, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, fmt.Sprintf(GetTransactionById, txId)))
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &t)
	return
}

//Returns additional transaction information for given transaction id.
func (x *API) GetTransactionInfoById(txId string) (t TransactionInfo, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, fmt.Sprintf(GetTransactionInfoById, txId)))
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &t)
	return
}

//Returns some nodes for the alqo.conf wallet
func (x *API) GetAddNodes() (nodes string, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, GetAddNodes))
	if err != nil {
		return
	}
	return string(body), err
}

//Returns wealth distribution stats
func (x *API) GetDistribution() (d Distribution, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, GetDistribution))
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &d)
	return
}

//Returns information for given wallet-address
func (x *API) GetWalletByAddress(addr string) (w Wallet, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, fmt.Sprintf(GetWalletByAddress, addr)))
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &w)
	return
}

//Returns additional transaction information for the given wallet-address
func (x *API) GetWalletTransactionsByAddress(addr string) (w WalletTransactions, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, fmt.Sprintf(GetWalletTransactionsByAddress, addr)))
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &w)
	return
}

//Returns current balance of given address
func (x *API) GetWalletBalance(addr string) (b float64, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, fmt.Sprintf(GetWalletBalance, addr)))
	if err != nil {
		return
	}
	b, err = strconv.ParseFloat(string(body), 16)
	return
}

//Returns last [count] transactions greater than [min]
func (x *API) GetLastTransaction(count int, min int) (lt LastTransaction, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, fmt.Sprintf(GetLastTransactions, count, min)))
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &lt)
	return
}

//Returns transactions between two timestamps
func (x *API) GetLastTransactionsBetween(tsStart int64, tsEnd int64) (lt LastTransaction, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, fmt.Sprintf(GetLastTransactionsBetween, tsStart, tsEnd)))
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &lt)
	return
}

func (x *API) get(url string) ([]byte, error) {
	res, err := x.c.Get(url)
	if err != nil {
		return []byte{}, err
	}
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}
