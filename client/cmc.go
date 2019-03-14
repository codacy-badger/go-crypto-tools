package client

import (
	"net/http"

	"encoding/json"
	"io/ioutil"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	CMCGetIDs          = "https://pro-api.coinmarketcap.com/v1/cryptocurrency/map"
	CMCWebLink         = "https://pro-api.coinmarketcap.com"
	CMCCmcDynamicsList = "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest?limit=5000"
	CMCCmcInfo         = "https://pro-api.coinmarketcap.com/v1/cryptocurrency/info"
)

type CMCListingResponse struct {
	Status struct {
		Timestamp    time.Time   `json:"timestamp"`
		ErrorCode    int         `json:"error_code"`
		ErrorMessage interface{} `json:"error_message"`
		Elapsed      int         `json:"elapsed"`
		CreditCount  int         `json:"credit_count"`
	} `json:"status"`
	Data []struct {
		ID                int         `json:"id"`
		Name              string      `json:"name"`
		Symbol            string      `json:"symbol"`
		Slug              string      `json:"slug"`
		CirculatingSupply float64     `json:"circulating_supply"`
		TotalSupply       float64     `json:"total_supply"`
		MaxSupply         float64     `json:"max_supply"`
		DateAdded         time.Time   `json:"date_added"`
		NumMarketPairs    int         `json:"num_market_pairs"`
		Tags              []string    `json:"tags"`
		Platform          interface{} `json:"platform"`
		CmcRank           int         `json:"cmc_rank"`
		LastUpdated       time.Time   `json:"last_updated"`
		Quote             struct {
			USD struct {
				Price            float64   `json:"price"`
				Volume24H        float64   `json:"volume_24h"`
				PercentChange1H  float64   `json:"percent_change_1h"`
				PercentChange24H float64   `json:"percent_change_24h"`
				PercentChange7D  float64   `json:"percent_change_7d"`
				MarketCap        float64   `json:"market_cap"`
				LastUpdated      time.Time `json:"last_updated"`
			} `json:"USD"`
		} `json:"quote"`
	} `json:"data"`
}

type CMCAllInfo struct {
	Status struct {
		Timestamp    time.Time   `json:"timestamp"`
		ErrorCode    int         `json:"error_code"`
		ErrorMessage interface{} `json:"error_message"`
		Elapsed      int         `json:"elapsed"`
		CreditCount  int         `json:"credit_count"`
	} `json:"status"`
	Data struct {
		Map map[string]interface{} `json:"-"`
	} `json:"data"`
}

type CMCInfoMap map[string]CMCInfoItem

type CMCInfoItem struct {
	Urls struct {
		Website      []string `json:"website"`
		Twitter      []string `json:"twitter"`
		Reddit       []string `json:"reddit"`
		MessageBoard []string `json:"message_board"`
		Announcement []string `json:"announcement"`
		Chat         []string `json:"chat"`
		Explorer     []string `json:"explorer"`
		SourceCode   []string `json:"source_code"`
	} `json:"urls"`
	Logo      string      `json:"logo"`
	ID        int         `json:"id"`
	Name      string      `json:"name"`
	Symbol    string      `json:"symbol"`
	Slug      string      `json:"slug"`
	DateAdded time.Time   `json:"date_added"`
	Tags      interface{} `json:"tags"`
	Platform  interface{} `json:"platform"`
	Category  string      `json:"category"`
}

type CMCLatestList struct {
	Data []struct {
		ID                int         `json:"id"`
		Name              string      `json:"name"`
		Symbol            string      `json:"symbol"`
		Slug              string      `json:"slug"`
		CmcRank           int         `json:"cmc_rank"`
		NumMarketPairs    int         `json:"num_market_pairs"`
		CirculatingSupply int         `json:"circulating_supply"`
		TotalSupply       int         `json:"total_supply"`
		MaxSupply         int         `json:"max_supply"`
		LastUpdated       time.Time   `json:"last_updated"`
		DateAdded         time.Time   `json:"date_added"`
		Tags              []string    `json:"tags"`
		Platform          interface{} `json:"platform"`
		Quote             struct {
			USD struct {
				Price            float64   `json:"price"`
				Volume24H        int64     `json:"volume_24h"`
				PercentChange1H  float64   `json:"percent_change_1h"`
				PercentChange24H float64   `json:"percent_change_24h"`
				PercentChange7D  float64   `json:"percent_change_7d"`
				MarketCap        int64     `json:"market_cap"`
				LastUpdated      time.Time `json:"last_updated"`
			} `json:"USD"`
			BTC struct {
				Price            int       `json:"price"`
				Volume24H        int       `json:"volume_24h"`
				PercentChange1H  int       `json:"percent_change_1h"`
				PercentChange24H int       `json:"percent_change_24h"`
				PercentChange7D  int       `json:"percent_change_7d"`
				MarketCap        int       `json:"market_cap"`
				LastUpdated      time.Time `json:"last_updated"`
			} `json:"BTC"`
		} `json:"quote"`
	} `json:"data"`
	Status struct {
		Timestamp    time.Time `json:"timestamp"`
		ErrorCode    int       `json:"error_code"`
		ErrorMessage string    `json:"error_message"`
		Elapsed      int       `json:"elapsed"`
		CreditCount  int       `json:"credit_count"`
	} `json:"status"`
}

type CMCClient struct {
	Client    *http.Client
	AccessKey string
}

func NewCMCClient(accessKey string) *CMCClient {
	m := CMCClient{}
	m.AccessKey = accessKey
	m.Client = &http.Client{}
	return &m
}

func (i *CMCClient) GetLink() string {
	return CMCWebLink
}

func (i *CMCClient) GetInfo(ids []int) (CMCInfoMap, error) {
	req, err := http.NewRequest("GET", CMCCmcInfo, nil)
	if err != nil {
		return CMCInfoMap{}, err
	}

	q := url.Values{}
	var strIds []string
	for _, id := range ids {
		strIds = append(strIds, strconv.Itoa(id))
	}
	q.Add("id", strings.Join(strIds, ","))

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", i.AccessKey)
	req.URL.RawQuery = q.Encode()

	res, err := i.Client.Do(req)
	if err != nil {
		return CMCInfoMap{}, err
	}
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return CMCInfoMap{}, err
	}
	return parseNestedJSONMap(body)
}

func parseNestedJSONMap(body []byte) (CMCInfoMap, error) {
	var p CMCAllInfo
	var infoMap CMCInfoMap
	if err := json.Unmarshal(body, &p.Data.Map); err == nil {
		for key, m := range p.Data.Map {
			if key == "data" {
				str, err := json.Marshal(m)
				if err != nil {
					return CMCInfoMap{}, err
				}
				err = json.Unmarshal(str, &infoMap)
				if err != nil {
					return CMCInfoMap{}, err
				}
				return infoMap, nil
			}
		}
	} else {
		return CMCInfoMap{}, err
	}
	return CMCInfoMap{}, nil
}

func (i *CMCClient) GetLatestListing() (response CMCListingResponse, err error) {
	req, err := http.NewRequest("GET", CMCCmcDynamicsList, nil)
	if err != nil {
		return
	}
	q := url.Values{}
	q.Add("limit", "5000")
	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", i.AccessKey)
	req.URL.RawQuery = q.Encode()
	res, err := i.Client.Do(req)
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return
	}
	return response, nil
}
