package mc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/dnaeon/go-vcr/recorder"
)

const (
	APIUrl = "http://explorer.mastercoin.one"
	Name   = "MasterCoin"

	//API
	GetDifficulty         = "/api/getdifficulty"
	GetConnectionCount    = "/api/getconnectioncount"
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

//Github repo: https://github.com/MasterCoinOne/MasterCoinV2
//API doc:  http://explorer.mastercoin.one/info
//Explorer: http://explorer.mastercoin.one
type API struct {
	BaseUrl      string
	Client       *http.Client
	RecordClient *http.Client
	Coin         string
}

//InitApi initializes the client with the given base-url
func InitApi(url string, rec *recorder.Recorder) *API {
	m := API{}
	m.BaseUrl = url

	if rec != nil {
		m.RecordClient = &http.Client{Transport: rec}
	} else {
		m.RecordClient = &http.Client{}
	}
	m.Client = &http.Client{}
	return &m
}

//Returns the current difficulty.
func (m *API) GetDifficulty() (d float64, err error) {
	body, err := m.get(fmt.Sprintf("%s%s", m.BaseUrl, GetDifficulty))
	if err != nil {
		return
	}
	d, err = strconv.ParseFloat(string(body), 16)
	return
}

//Returns the number of connections the block explorer has to other nodes.
func (m *API) GetConnectionCount() (d int64, err error) {
	body, err := m.get(fmt.Sprintf("%s%s", m.BaseUrl, GetConnectionCount))
	if err != nil {
		return
	}
	d, err = strconv.ParseInt(string(body), 10, 64)
	return
}

//Returns the current block count.
func (m *API) GetBlockCount() (d int64, err error) {
	body, err := m.get(fmt.Sprintf("%s%s", m.BaseUrl, GetBlockCount))
	if err != nil {
		return
	}
	d, err = strconv.ParseInt(string(body), 10, 64)
	return
}

//Returns the hash of the block at ; index 0 is the genesis block.
func (m *API) GetBlockHashByIndex(index int) (s string, err error) {
	v := url.Values{}
	v.Set("index", strconv.Itoa(index))
	body, err := m.getRecorded(fmt.Sprintf("%s%s?%s", m.BaseUrl, GetBlockHashByIndex, v.Encode()))
	if err != nil {
		return
	}
	return string(body), nil
}

//Returns information about the block with the given hash.
func (m *API) GetBlockByHash(hash string) (b Block, err error) {
	v := url.Values{}
	v.Set("hash", hash)
	body, err := m.getRecorded(fmt.Sprintf("%s%s?%s", m.BaseUrl, GetBlockByHash, v.Encode()))
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &b)
	return
}

//Returns raw transaction representation for given transaction id. decrypt can be set to 0(false) or 1(true).
func (m *API) GetTransactionById(txId string, decrypt bool) (t Transaction, err error) {
	v := url.Values{}
	v.Set("txId", txId)
	body, err := m.getRecorded(fmt.Sprintf("%s%s?%s&decrypt=1", m.BaseUrl, GetTransactionById, v.Encode()))
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &t)
	return
}

//Returns current network hash rate
func (m *API) GetGetNetworkHashRate() (d float64, err error) {
	body, err := m.get(fmt.Sprintf("%s%s", m.BaseUrl, GetGetNetworkHashRate))
	if err != nil {
		return
	}
	d, err = strconv.ParseFloat(string(body), 16)
	return
}

//Returns current money supply
func (m *API) GetDistribution() (d Distribution, err error) {
	body, err := m.get(fmt.Sprintf("%s%s", m.BaseUrl, GetDistribution))
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &d)
	return
}

//Returns current money supply
func (m *API) GetMoneySupply() (d float64, err error) {
	body, err := m.get(fmt.Sprintf("%s%s", m.BaseUrl, GetMoneySupply))
	if err != nil {
		return
	}
	d, err = strconv.ParseFloat(string(body), 16)
	return
}

//Returns information for given address
func (m *API) GetAddress(addr string) (a Address, err error) {
	body, err := m.get(fmt.Sprintf("%s%s", m.BaseUrl, fmt.Sprintf(GetAddress, addr)))
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &a)
	return
}

//Returns current balance of given address
func (m *API) GetBalance(addr string) (d float64, err error) {
	body, err := m.get(fmt.Sprintf("%s%s", m.BaseUrl, fmt.Sprintf(GetBalance, addr)))
	if err != nil {
		return
	}
	d, err = strconv.ParseFloat(string(body), 16)
	return
}

//Returns last [count] transactions greater than [min]
func (m *API) GetLastTransaction(count int, min int) (lt LastTransaction, err error) {
	body, err := m.get(fmt.Sprintf("%s%s", m.BaseUrl, fmt.Sprintf(GetLastTransactions, count, min)))
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &lt)
	return
}

func (m *API) GetBaseUrl() string {
	return m.BaseUrl
}
func (m *API) GetClient() *http.Client {
	return m.Client
}
func (m *API) GetCoin() string {
	return Name
}

func (m *API) get(url string) ([]byte, error) {
	res, err := m.Client.Get(url)
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

// getRecorded will perform a GET request and persist response in file storage if not present
func (g *API) getRecorded(url string) ([]byte, error) {
	res, err := g.RecordClient.Get(url)
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
