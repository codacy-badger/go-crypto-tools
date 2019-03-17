package wallet_cli_grph

import (
	"encoding/json"

	"github.com/cwntr/go-crypto-tools/wallet_cli"
)

const (
	ArgumentMasternodeList       = "masternode"
	ArgumentMasternodeListOption = "list"
)

type GRAPHResponse []struct {
	Rank       int    `json:"rank"`
	Network    string `json:"network"`
	Txhash     string `json:"txhash"`
	Outidx     int    `json:"outidx"`
	Status     string `json:"status"`
	IP         string `json:"ip"`
	Addr       string `json:"addr"`
	Version    int    `json:"version"`
	Lastseen   int    `json:"lastseen"`
	Activetime int    `json:"activetime"`
	Lastpaid   int    `json:"lastpaid"`
}

func GetGRAPHList(cli string) (GRAPHResponse, error) {
	data, err := wallet_cli.ExecCLI(cli, ArgumentMasternodeList, ArgumentMasternodeListOption)
	if err != nil {
		return GRAPHResponse{}, err
	}
	var resp GRAPHResponse
	err = json.Unmarshal([]byte(data), &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
