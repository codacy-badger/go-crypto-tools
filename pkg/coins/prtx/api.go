package prtx

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

const (
	APIUrl = "https://blocks.printex.tech"

	//API
	GetDifficulty         = "/api/getdifficulty"
	GetConnectionCount    = "/api/getconnectioncount"
	GetMasternodeCount    = "/api/getmasternodecount"
	GetBlockCount         = "/api/getblockcount"
	GetBlockHashByIndex   = "/api/getblockhash"      //query param [index]
	GetBlockByHash        = "/api/getblock"          //query param [hash]
	GetTransactionById    = "/api/getrawtransaction" //query param [txId] [decrypt]
	GetGetNetworkHashRate = "/api/getnetworkhashps"

	//EXT
	GetMoneySupply      = "/ext/getmoneysupply"
	GetDistribution     = "/ext/getdistribution"
	GetAddress          = "/ext/getaddress/%s"
	GetBalance          = "/ext/getbalance/%s"
	GetLastTransactions = "/ext/getlasttxs/%d/%d" //Returns last [count] transactions greater than [min] Note: returned values are in satoshis
)

//Github repo: https://github.com/Printex-official/printex-core
//API doc: https://blocks.printex.tech/info
//Explorer: https://blocks.printex.tech
type API struct {
	baseUrl string
	c       *http.Client
}

func InitApi(url string) *API {
	client := &http.Client{}
	return &API{baseUrl: url, c: client}
}

//Returns the current difficulty.
func (p *API) GetDifficulty() (d float64, err error) {
	body, err := p.get(fmt.Sprintf("%s%s", p.baseUrl, GetDifficulty))
	if err != nil {
		return
	}
	d, err = strconv.ParseFloat(string(body), 16)
	return
}

//Returns the number of connections the block explorer has to other nodes.
func (p *API) GetConnectionCount() (d int64, err error) {
	body, err := p.get(fmt.Sprintf("%s%s", p.baseUrl, GetConnectionCount))
	if err != nil {
		return
	}
	d, err = strconv.ParseInt(string(body), 10, 64)
	return
}

//Returns the total number of master nodes on the network.
func (p *API) GetMasternodeCount() (mn Masternodes, err error) {
	body, err := p.get(fmt.Sprintf("%s%s", p.baseUrl, GetMasternodeCount))
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &mn)
	return
}

//Returns the current block count.
func (p *API) GetBlockCount() (d int64, err error) {
	body, err := p.get(fmt.Sprintf("%s%s", p.baseUrl, GetBlockCount))
	if err != nil {
		return
	}
	d, err = strconv.ParseInt(string(body), 10, 64)
	return
}

//Returns the hash of the block at ; index 0 is the genesis block.
func (p *API) GetBlockHashByIndex(index int) (s string, err error) {
	v := url.Values{}
	v.Set("index", strconv.Itoa(index))
	body, err := p.get(fmt.Sprintf("%s%s?%s", p.baseUrl, GetBlockHashByIndex, v.Encode()))
	if err != nil {
		return
	}
	return string(body), nil
}

//Returns information about the block with the given hash.
func (p *API) GetBlockByHash(hash string) (b Block, err error) {
	v := url.Values{}
	v.Set("hash", hash)
	body, err := p.get(fmt.Sprintf("%s%s?%s", p.baseUrl, GetBlockByHash, v.Encode()))
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &b)
	return
}

//Returns raw transaction representation for given transaction id. decrypt can be set to 0(false) or 1(true).
func (p *API) GetTransactionById(txId string, decrypt bool) (t Transaction, err error) {
	v := url.Values{}
	v.Set("txId", txId)
	body, err := p.get(fmt.Sprintf("%s%s?%s&decrypt=1", p.baseUrl, GetTransactionById, v.Encode()))
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &t)
	return
}

//Returns current network hash rate
func (p *API) GetGetNetworkHashRate() (d float64, err error) {
	body, err := p.get(fmt.Sprintf("%s%s", p.baseUrl, GetGetNetworkHashRate))
	if err != nil {
		return
	}
	d, err = strconv.ParseFloat(string(body), 16)
	return
}

//Returns current money supply
func (p *API) GetDistribution() (d Distribution, err error) {
	body, err := p.get(fmt.Sprintf("%s%s", p.baseUrl, GetDistribution))
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &d)
	return
}

//Returns current money supply
func (p *API) GetMoneySupply() (d float64, err error) {
	body, err := p.get(fmt.Sprintf("%s%s", p.baseUrl, GetMoneySupply))
	if err != nil {
		return
	}
	d, err = strconv.ParseFloat(string(body), 16)
	return
}

//Returns information for given address
func (p *API) GetAddress(addr string) (a Address, err error) {
	body, err := p.get(fmt.Sprintf("%s%s", p.baseUrl, fmt.Sprintf(GetAddress, addr)))
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &a)
	return
}

//Returns current balance of given address
func (p *API) GetBalance(addr string) (d float64, err error) {
	body, err := p.get(fmt.Sprintf("%s%s", p.baseUrl, fmt.Sprintf(GetBalance, addr)))
	if err != nil {
		return
	}
	d, err = strconv.ParseFloat(string(body), 16)
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

func (p *API) get(url string) ([]byte, error) {
	res, err := p.c.Get(url)
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
