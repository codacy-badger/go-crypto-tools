package common

const (
	NameGRPH = "GraphCoin"
	NameMC   = "MasterCoin"
	NamePRTX = "Printex"
	NameXLQ  = "ALQO"
	NameXRP  = "Ripple"
	NameXSN  = "Stakenet"

	CurrencyGRPH = "GRPH"
	CurrencyMC   = "MC"
	CurrencyPRTX = "PRTX"
	CurrencyXLQ  = "XLQ"
	CurrencyXRP  = "XRP"
	CurrencyXSN  = "XSN"
)

type CoinInfo struct {
	Currency      string
	Name          string
	ConsensusType ConsensusType //consider slice e.g. XSN has PoS and TPoS
}

func GetAvailableCoins() []CoinInfo {
	return []CoinInfo{
		{CurrencyGRPH, NameGRPH, PoS},
		{CurrencyMC, NameMC, PoS},
		{CurrencyPRTX, NamePRTX, PoS},
		{CurrencyXLQ, NameXLQ, PoS},
		{CurrencyXRP, NameXRP, PoW},
		{CurrencyXSN, NameXSN, PoS},
	}
}

func GetCoinByName(name string) CoinInfo {
	for _, c := range GetAvailableCoins() {
		if c.Name == name {
			return c
		}
	}
	return CoinInfo{}
}
