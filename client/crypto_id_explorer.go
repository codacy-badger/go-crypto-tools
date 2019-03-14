package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	CryptoIDWebApiExplorer = "https://chainz.cryptoid.info/%s/api.dws?q=nodes"
	CryptoIDBLOCK          = "block"
	CryptoIDDASH           = "dash"
	CryptoIDPIVX           = "pivx"
	CryptoIDPHR            = "phr"
	CryptoIDNIX            = "nix"
	CryptoIDSYS            = "sys"
)

type CryptoIDResponse []struct {
	Subver   string   `json:"subver"`
	Protocol int      `json:"protocol"`
	Nodes    []string `json:"nodes"`
}

type CryptoIDWebApiClient struct {
	Client *http.Client
	Coin   string
}

func NewCryptoIDWebApiClient(coin string) *CryptoIDWebApiClient {
	m := CryptoIDWebApiClient{}
	m.Client = &http.Client{}
	m.Coin = coin
	return &m
}

func (i *CryptoIDWebApiClient) GetLink() string {
	return fmt.Sprintf(CryptoIDWebApiExplorer, i.Coin)
}

func (i *CryptoIDWebApiClient) GetMNList() (CryptoIDResponse, error) {
	u, _ := url.Parse(i.GetLink())
	res, err := i.Client.Get(u.String())
	if err != nil {
		return CryptoIDResponse{}, err
	}
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return CryptoIDResponse{}, err
	}

	var resp CryptoIDResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return CryptoIDResponse{}, err
	}
	return resp, nil
}
