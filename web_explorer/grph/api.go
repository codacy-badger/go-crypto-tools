package grph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"time"
)

const (
	APIUrl = "https://explorer.graphcoin.net"
	Name   = "GraphCoin"

	//API
	GetDifficulty       = "/api/getdifficulty"
	GetConnectionCount  = "/api/getconnectioncount"
	GetBlockCount       = "/api/getblockcount"
	GetBlockHashByIndex = "/api/getblockhash"      //query param [index]
	GetBlockByHash      = "/api/getblock"          //query param [hash]
	GetTransactionById  = "/api/getrawtransaction" //query param [txId] [decrypt]

	//EXT
	GetMoneySupply = "/ext/getmoneysupply"
	GetAddress     = "/ext/getaddress/%s"
	GetBalance     = "/ext/getbalance/%s"
)

//Github repo: https://github.com/Graphcoin
//API doc: https://explorer.graphcoin.net/api
//Explorer: https://explorer.graphcoin.net
type API struct {
	BaseUrl       string
	Client        *http.Client
	Coin          string
	WaitingPeriod time.Duration
}

//InitApi initializes the client with the given base-url
func InitApi(url string) *API {
	m := API{}
	m.WaitingPeriod = time.Second * 0
	m.BaseUrl = url
	m.Client = &http.Client{}
	return &m
}

//Returns the current difficulty.
func (g *API) GetDifficulty() (d float64, err error) {
	time.Sleep(g.WaitingPeriod)
	body, err := g.get(fmt.Sprintf("%s%s", g.BaseUrl, GetDifficulty))
	if err != nil {
		return
	}
	d, err = strconv.ParseFloat(string(body), 16)
	return
}

//Returns the number of connections the block explorer has to other nodes.
func (g *API) GetConnectionCount() (d int64, err error) {
	time.Sleep(g.WaitingPeriod)
	body, err := g.get(fmt.Sprintf("%s%s", g.BaseUrl, GetConnectionCount))
	if err != nil {
		return
	}
	d, err = strconv.ParseInt(string(body), 10, 64)
	return
}

//Returns the current block count.
func (g *API) GetBlockCount() (d int64, err error) {
	time.Sleep(g.WaitingPeriod)
	body, err := g.get(fmt.Sprintf("%s%s", g.BaseUrl, GetBlockCount))
	if err != nil {
		return
	}
	d, err = strconv.ParseInt(string(body), 10, 64)
	return
}

//Returns the hash of the block at ; index 0 is the genesis block.
func (g *API) GetBlockHashByIndex(index int) (s string, err error) {
	time.Sleep(g.WaitingPeriod)
	v := url.Values{}
	v.Set("index", strconv.Itoa(index))
	body, err := g.get(fmt.Sprintf("%s%s?%s", g.BaseUrl, GetBlockHashByIndex, v.Encode()))
	if err != nil {
		return
	}
	return string(body), nil
}

//Returns information about the block with the given hash.
func (g *API) GetBlockByHash(hash string) (b Block, err error) {
	time.Sleep(g.WaitingPeriod)
	v := url.Values{}
	v.Set("hash", hash)
	body, err := g.get(fmt.Sprintf("%s%s?%s", g.BaseUrl, GetBlockByHash, v.Encode()))
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &b)
	return
}

//Returns raw transaction representation for given transaction id. decrypt can be set to 0(false) or 1(true).
func (g *API) GetTransactionById(txId string, decrypt bool) (t Transaction, err error) {
	time.Sleep(g.WaitingPeriod)
	v := url.Values{}
	v.Set("txId", txId)
	d := 0
	if decrypt {
		d = 1
	}
	v.Set("decrypt", fmt.Sprintf("%d", d))
	body, err := g.get(fmt.Sprintf("%s%s?%s", g.BaseUrl, GetTransactionById, v.Encode()))
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &t)
	return
}

//Returns current money supply
func (g *API) GetMoneySupply() (d float64, err error) {
	time.Sleep(g.WaitingPeriod)
	body, err := g.get(fmt.Sprintf("%s%s", g.BaseUrl, GetMoneySupply))
	if err != nil {
		return
	}
	d, err = strconv.ParseFloat(string(body), 16)
	return
}

//Returns information for given address
func (g *API) GetAddress(addr string) (a Address, err error) {
	time.Sleep(g.WaitingPeriod)
	body, err := g.get(fmt.Sprintf("%s%s", g.BaseUrl, fmt.Sprintf(GetAddress, addr)))
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &a)
	return
}

//Returns current balance of given address
func (g *API) GetBalance(addr string) (d float64, err error) {
	time.Sleep(g.WaitingPeriod)
	body, err := g.get(fmt.Sprintf("%s%s", g.BaseUrl, fmt.Sprintf(GetBalance, addr)))
	if err != nil {
		return
	}
	d, err = strconv.ParseFloat(string(body), 16)
	return
}

func (g *API) GetBaseUrl() string {
	return g.BaseUrl
}
func (g *API) GetClient() *http.Client {
	return g.Client
}
func (g *API) GetCoin() string {
	return Name
}

func (g *API) get(url string) ([]byte, error) {
	res, err := g.Client.Get(url)
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
