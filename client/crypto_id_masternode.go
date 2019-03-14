package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type CryptoIDResponseMN []struct {
	Net  string `json:"net"`
	Nb   int    `json:"nb"`
	IP   string `json:"ip"`
	Cc   string `json:"cc"`
	City string `json:"city"`
	Org  string `json:"org"`
}

const (
	CryptoIDMasternodeEndpoint = "https://chainz.cryptoid.info/explorer/masternodes.data.dws?coin=%s"
)

type CryptoIDWebApiMNClient struct {
	Client *http.Client
	Coin   string
}

func NewCryptoIDWebApiMNClient(coin string) *CryptoIDWebApiMNClient {
	m := CryptoIDWebApiMNClient{}
	m.Client = &http.Client{}
	m.Coin = coin
	return &m
}

func (i *CryptoIDWebApiMNClient) GetLink() string {
	return fmt.Sprintf(CryptoIDMasternodeEndpoint, i.Coin)
}

func (i *CryptoIDWebApiMNClient) GetMNList() (CryptoIDResponseMN, error) {
	u, _ := url.Parse(i.GetLink())
	res, err := i.Client.Get(u.String())
	if err != nil {
		return CryptoIDResponseMN{}, err
	}
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return CryptoIDResponseMN{}, err
	}

	var resp CryptoIDResponseMN
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return CryptoIDResponseMN{}, err
	}
	return resp, nil
}
