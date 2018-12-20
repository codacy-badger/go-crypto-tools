package xrp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	APIUrl = "http://data.ripple.com/v2"

	GetAccountTransactions            = "/accounts/%s/transactions"
	GetAccountTransactionsBySequence  = "/accounts/%s/transactions/%d"
	GetAccountPayments                = "/accounts/%s/payments"
	GetAccountExchanges               = "/accounts/%s/exchanges"
	GetAccountBalanceChanges          = "/accounts/%s/balance_changes"
	GetAccountReports                 = "/accounts/%s/reports/%s"
	GetAccountBalances                = "/accounts/%s/balances"
	GetAccountOrders                  = "/accounts/%s/orders"
	GetAccountTransactionStats        = "/accounts/%s/stats/transactions"
	GetAccountValueStats              = "/accounts/%s/stats/value"
	GetAccount                        = "/accounts/%s"
	GetAccounts                       = "/accounts"
	GetLedgers                        = "/ledgers/%s"
	GetTransactions                   = "/transactions"
	GetTransaction                    = "/transactions/%s"
	GetPayments                       = "/payments/%s"
	GetExchanges                      = "/exchanges/%s/%s"
	GetActiveAccounts                 = "/active_accounts/%s/%d"
	GetExchangeVolume                 = "/network/exchange_volume"
	GetPaymentVolume                  = "/network/payment_volume"
	GetExternalMarketVolume           = "/network/external_markets"
	GetXRPDistribution                = "/network/xrp_distribution"
	GetTopCurrencies                  = "/network/top_currencies/%s"
	GetTopMarkets                     = "/network/top_markets/%s"
	GetNetworkTopology                = "/network/topology"
	GetNetworkTopologyNodes           = "/network/topology/nodes"
	GetNetworkTopologyNodeByPublicKey = "/network/topology/nodes/%s"
	GetNetworkTopologyLinks           = "/network/topology/links"
	GetValidators                     = "/network/validators"
	GetValidator                      = "/network/validators/%s"
	GetValidatorValidations           = "/network/validators/%s/validations"
	GetValidatorManifests             = "/network/validators/%s/manifests"
	GetValidatorReportsByPublicKey    = "/network/validators/%s/reports"
	GetValidatorReports               = "/network/validator_reports"
	GetValidations                    = "/network/validations"
	GetRippledVersions                = "/network/rippled_versions"
	GetExchangeRate                   = "/exchange_rates/%s/%d"
	NormalizeAmount                   = "/normalize"
	GetDailySummary                   = "/reports/%s"
	GetTransactionStatistics          = "/stats/%s/%s"
	CheckHealth                       = "/health/%s"
)

//API definition https://data.ripple.com
//API documentation https://developers.ripple.com/data-api.html
type API struct {
	baseUrl string
	c       *http.Client
}

//InitApi initializes the client with the given base-url
func InitApi(url string) *API {
	client := &http.Client{}
	return &API{baseUrl: url, c: client}
}

func (x *API) GetAccountTransactions(address string) (at AccountTransaction, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, fmt.Sprintf(GetAccountTransactions, address)))
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &at)
	return
}

/*
func (x *API) GetAccountTransactionsBySequence(address string, sequence int) (d float64, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, fmt.Sprintf(GetAccountTransactionsBySequence, address, sequence)))
	if err != nil {
		return
	}
	return
}
func (x *API) GetAccountPayments(address string) (d float64, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, fmt.Sprintf(GetAccountPayments, address)))
	if err != nil {
		return
	}
	return
}
func (x *API) GetAccountExchanges(address string) (d float64, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, fmt.Sprintf(GetAccountExchanges, address)))
	if err != nil {
		return
	}
	return
}
func (x *API) GetAccountBalanceChanges(address string) (d float64, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, fmt.Sprintf(GetAccountBalanceChanges, address)))
	if err != nil {
		return
	}
	return
}
func (x *API) GetAccountReports(address string) (d float64, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, fmt.Sprintf(GetAccountReports, address)))
	if err != nil {
		return
	}
	return
}
func (x *API) GetAccountBalances(address string) (d float64, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, fmt.Sprintf(GetAccountBalances, address)))
	if err != nil {
		return
	}
	return
}
func (x *API) GetAccountOrders(address string) (d float64, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, fmt.Sprintf(GetAccountOrders, address)))
	if err != nil {
		return
	}
	return
}
func (x *API) GetAccountTransactionStats(address string) (d float64, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, fmt.Sprintf(GetAccountTransactionStats, address)))
	if err != nil {
		return
	}
	return
}
func (x *API) GetAccountValueStats(address string) (d float64, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, fmt.Sprintf(GetAccountValueStats, address)))
	if err != nil {
		return
	}
	return
}
func (x *API) GetAccount(address string) (d float64, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, fmt.Sprintf(GetAccount, address)))
	if err != nil {
		return
	}
	return
}
func (x *API) GetAccounts(d float64, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, GetAccounts))
	if err != nil {
		return
	}
	return
}
func (x *API) GetLedgers(d float64, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, GetLedgers))
	if err != nil {
		return
	}
	return
}
func (x *API) GetTransactions(d float64, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, GetTransactions))
	if err != nil {
		return
	}
	return
}
func (x *API) GetTransaction(d float64, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, GetTransaction))
	if err != nil {
		return
	}
	return
}
func (x *API) GetPayments(d float64, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, GetPayments))
	if err != nil {
		return
	}
	return
}
func (x *API) GetExchanges(d float64, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, GetExchanges))
	if err != nil {
		return
	}
	return
}
func (x *API) GetActiveAccounts(d float64, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, GetActiveAccounts))
	if err != nil {
		return
	}
	return
}
func (x *API) GetExchangeVolume(d float64, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, GetExchangeVolume))
	if err != nil {
		return
	}
	return
}
func (x *API) GetPaymentVolume(d float64, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, GetPaymentVolume))
	if err != nil {
		return
	}
	return
}
func (x *API) GetExternalMarketVolume(d float64, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, GetExternalMarketVolume))
	if err != nil {
		return
	}
	return
}
func (x *API) GetXRPDistribution(d float64, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, GetXRPDistribution))
	if err != nil {
		return
	}
	return
}
func (x *API) GetTopCurrencies(d float64, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, GetTopCurrencies))
	if err != nil {
		return
	}
	return
}
func (x *API) GetTopMarkets(d float64, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, GetTopMarkets))
	if err != nil {
		return
	}
	return
}
func (x *API) GetNetworkTopology(d float64, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, GetNetworkTopology))
	if err != nil {
		return
	}
	return
}
func (x *API) GetNetworkTopologyNodes(d float64, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, GetNetworkTopologyNodes))
	if err != nil {
		return
	}
	return
}
func (x *API) GetNetworkTopologyNodeByPublicKey(d float64, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, GetNetworkTopologyNodeByPublicKey))
	if err != nil {
		return
	}
	return
}
func (x *API) GetNetworkTopologyLinks(d float64, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, GetNetworkTopologyLinks))
	if err != nil {
		return
	}
	return
}
func (x *API) GetValidators(d float64, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, GetValidators))
	if err != nil {
		return
	}
	return
}
func (x *API) GetValidator(d float64, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, GetValidator))
	if err != nil {
		return
	}
	return
}
func (x *API) GetValidatorValidations(d float64, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, GetValidatorValidations))
	if err != nil {
		return
	}
	return
}
func (x *API) GetValidatorManifests(d float64, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, GetValidatorManifests))
	if err != nil {
		return
	}
	return
}
func (x *API) GetValidatorReportsByPublicKey(d float64, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, GetValidatorReportsByPublicKey))
	if err != nil {
		return
	}
	return
}
func (x *API) GetValidatorReports(d float64, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, GetValidatorReports))
	if err != nil {
		return
	}
	return
}
func (x *API) GetValidations(d float64, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, GetValidations))
	if err != nil {
		return
	}
	return
}
func (x *API) GetRippledVersions(d float64, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, GetRippledVersions))
	if err != nil {
		return
	}
	return
}
func (x *API) GetExchangeRate(d float64, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, GetExchangeRate))
	if err != nil {
		return
	}
	return
}
func (x *API) NormalizeAmount(d float64, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, NormalizeAmount))
	if err != nil {
		return
	}
	return
}
func (x *API) GetDailySummary(d float64, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, GetDailySummary))
	if err != nil {
		return
	}
	return
}
func (x *API) GetTransactionStatistics(d float64, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, GetTransactionStatistics))
	if err != nil {
		return
	}
	return
}
func (x *API) CheckHealth(d float64, err error) {
	body, err := x.get(fmt.Sprintf("%s%s", x.baseUrl, CheckHealth))
	if err != nil {
		return
	}
	return
}
*/
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
