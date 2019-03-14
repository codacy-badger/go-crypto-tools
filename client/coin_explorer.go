package client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	CoinExplorerWebApiExplorerMNList  = "https://www.coinexplorer.net/api/v1/XZC/masternode/list"
)

type CoinExplorerMNCount struct {
	Success bool `json:"success"`
	Result  struct {
		Total    int    `json:"total"`
		Active   string `json:"active"`
		Inactive string `json:"inactive"`
	} `json:"result"`
	Error interface{} `json:"error"`
}

type CoinExplorerMNList struct {
	Success bool `json:"success"`
	Result  []struct {
		Txid            string `json:"txid"`
		N               string `json:"n"`
		Protocolversion string `json:"protocolversion"`
		Status          string `json:"status"`
		Active          string `json:"active"`
		Pubkey          string `json:"pubkey"`
		Addr            string `json:"addr"`
		IP              string `json:"ip"`
		Country         string `json:"country"`
		Organization    string `json:"organization"`
	} `json:"result"`
	Error interface{} `json:"error"`
}

type CoinExplorerWebApiClient struct {
	Client *http.Client
}

func NewCoinExplorerWebApiClient() *CoinExplorerWebApiClient {
	m := CoinExplorerWebApiClient{}
	m.Client = &http.Client{}
	return &m
}

func (i *CoinExplorerWebApiClient) GetLink() string {
	return CoinExplorerWebApiExplorerMNList
}

func (i *CoinExplorerWebApiClient) GetMNList() (CoinExplorerMNList, error) {
	u, _ := url.Parse(CoinExplorerWebApiExplorerMNList)
	res, err := i.Client.Get(u.String())
	if err != nil {
		return CoinExplorerMNList{}, err
	}
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return CoinExplorerMNList{}, err
	}
	var resp CoinExplorerMNList
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return CoinExplorerMNList{}, err
	}
	return resp, nil
}