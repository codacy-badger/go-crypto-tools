package client

import (
	"net/http"
	"strings"

	"io/ioutil"
	"net/url"

	"encoding/json"
	"fmt"

	"github.com/dnaeon/go-vcr/recorder"
)

const (
	//Endpoint
	ApiEndpointScheme = "http"
	ApiEndpointHost   = "api.ipstack.com"

	//Query options
	ApiQueryAccessKey = "access_key"
	ApiQueryFormat    = "format"
	ApiQueryHostname  = "hostname"
)

type AccessKey struct {
	Coin string `json:"coin"`
	Key  string `json:"key"`
}

type CLIPath struct {
	Coin string `json:"coin"`
	Path string `json:"path"`
}
type Config struct {
	AccessKeys       []AccessKey `json:"access_keys"`
	CLIPaths         []CLIPath   `json:"cli_paths"`
	SupportedCoins   []string    `json:"coins"`
	MysqlDSL         string      `json:"mysql_dsl"`
	Env              string      `json:"env"`
	WhitelistDomains []string    `json:"whitelist_domains"`
	MyDomain         string      `json:"my_domain"`
	CMCAccessKey     string      `json:"cmc_access_key"`
}

func (c *Config) GetAccessKey(coin string) string {
	for _, c := range c.AccessKeys {
		if strings.ToUpper(c.Coin) == strings.ToUpper(coin) {
			return c.Key
		}
	}
	fmt.Printf("coin: '%s' has no access key in config \n", coin)
	return ""
}

func (c *Config) GetCLIPath(coin string) string {
	for _, c := range c.CLIPaths {
		if strings.ToUpper(c.Coin) == strings.ToUpper(coin) {
			return c.Path
		}
	}
	fmt.Printf("coin: '%s' has no cli path config \n", coin)
	return ""
}

func (c *Config) GetWhitelistDomains() []string {
	return c.WhitelistDomains
}

type IPStackResponse struct {
	IP            string  `json:"ip"`
	Hostname      string  `json:"hostname"`
	Type          string  `json:"type"`
	ContinentCode string  `json:"continent_code"`
	ContinentName string  `json:"continent_name"`
	CountryCode   string  `json:"country_code"`
	CountryName   string  `json:"country_name"`
	RegionCode    string  `json:"region_code"`
	RegionName    string  `json:"region_name"`
	City          string  `json:"city"`
	Zip           string  `json:"zip"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	Location      struct {
		GeonameID int    `json:"geoname_id"`
		Capital   string `json:"capital"`
		Languages []struct {
			Code   string `json:"code"`
			Name   string `json:"name"`
			Native string `json:"native"`
		} `json:"languages"`
		CountryFlag             string `json:"country_flag"`
		CountryFlagEmoji        string `json:"country_flag_emoji"`
		CountryFlagEmojiUnicode string `json:"country_flag_emoji_unicode"`
		CallingCode             string `json:"calling_code"`
		IsEu                    bool   `json:"is_eu"`
	} `json:"location"`
}

type IPStackClient struct {
	Client    *http.Client
	AccessKey string
}

//InitApi initializes the client with the given base-url. Injects recorder if it was passed.
func NewIpStackClient(rec *recorder.Recorder, key string) *IPStackClient {
	m := IPStackClient{}
	m.AccessKey = key
	if rec != nil {
		m.Client = &http.Client{Transport: rec}
	} else {
		m.Client = &http.Client{}
	}
	return &m
}

func (i *IPStackClient) Request(ip string) (IPStackResponse, error) {
	u := url.URL{}
	u.Scheme = ApiEndpointScheme
	u.Host = ApiEndpointHost
	q := u.Query()
	q.Set(ApiQueryAccessKey, i.AccessKey)
	q.Set(ApiQueryFormat, "2")
	q.Set(ApiQueryHostname, "1")
	u.RawQuery = q.Encode()
	u.Path = ip
	res, err := i.Client.Get(u.String())
	if err != nil {
		return IPStackResponse{}, err
	}
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return IPStackResponse{}, err
	}
	var resp IPStackResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return IPStackResponse{}, err
	}
	return resp, nil
}
