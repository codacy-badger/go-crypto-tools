package main

import (
	"fmt"

	"time"

	"github.com/cwntr/crypto-sdk/pkg/coins/common"
)

const (
	TestAddrGrphMaster = "gMo9tMCfij7rPPsPz2m59xcrVJMQPQW7gY"
	TestAddrGrphStake  = "g8HvrJEmA8eiRhJpqbzSHynkW4xXzpwnYL"
	TestAddrMcMaster   = "mZzSaKnj4qVpBEY4C6gyJo6kiGVgE9S2j6"
	TestAddrMcStake    = "mdu8n8Dp7cPPJKGyg6Vh1GqbjFKLoCxmTL"
)

func main() {
	common.InitClients()
	sa := StakingAnalyzer{}
	wallets := make([]Wallet, 0)

	wallets = append(wallets, Wallet{
		Coin:        common.GetCoinByName(common.NameGRPH),
		AddressList: []string{TestAddrGrphStake, TestAddrGrphMaster},
	})

	wallets = append(wallets, Wallet{
		Coin:        common.GetCoinByName(common.NameMC),
		AddressList: []string{TestAddrMcStake, TestAddrMcMaster},
	})

	sa.wallets = wallets
	addresses := sa.GetAddresses()
	fmt.Printf("main.addreses %+v \n", addresses)

}

type Wallet struct {
	Coin        common.CoinInfo
	AddressList []string
	Addresses   []common.Address
}

type StakingAnalyzer struct {
	wallets []Wallet
	clients []interface{}
}

func (sa StakingAnalyzer) Set(wallets []Wallet) {
	sa.wallets = wallets
}

func (sa StakingAnalyzer) GetAddresses() []common.AddressTransactions {
	addresses := make([]common.AddressTransactions, 0)
	for _, w := range sa.wallets {
		switch w.Coin.Name {

		case common.NameGRPH:
			client := common.GetGrphClient()
			for _, adr := range w.AddressList {
				clientAdr, err := client.GetAddress(adr)
				if err != nil {
					fmt.Printf("err: %v \n", err)
					continue
				}
				addr := common.ToAddrGrph(clientAdr)
				transactions := make([]common.Transaction, 0)
				stakes := make([]common.StakeReward, 0)
				for _, tx := range addr.LastTxs {
					transaction, err := client.GetTransactionById(tx.Address, true)
					if err != nil {
						fmt.Printf("error on MC get transaction, err: %v \n", err)
						continue
					}

					cTx := common.ToTxGrph(transaction)
					//calc stake rewards
					var baseAmount float64
					for _, in := range cTx.VinList {
						if in.Address == adr {
							baseAmount = in.Amount
						}
					}

					for _, out := range cTx.Vout {
						for _, oAddr := range out.ScriptPubKey.Addresses {
							if oAddr == adr {
								diff := out.Value - baseAmount
								if diff > 0 && diff < 20 {
									stakes = append(stakes, common.StakeReward{
										Address: adr,
										Amount:  diff,
										Time:    time.Unix(int64(cTx.BlockTime), 0),
									})
								}
							}
						}
					}
					transactions = append(transactions, cTx)
				}
				addresses = append(addresses, common.AddressTransactions{
					Address:      addr,
					Transactions: transactions,
					Stakes:       stakes,
				})
			}
		case common.NameMC:
			client := common.GetMcClient()
			for _, adr := range w.AddressList {
				clientAdr, err := client.GetAddress(adr)
				if err != nil {
					fmt.Printf("err: %v \n", err)
					continue
				}

				addr := common.ToAddrMc(clientAdr)
				transactions := make([]common.Transaction, 0)
				for _, tx := range addr.LastTxs {
					transaction, err := client.GetTransactionById(tx.Address, true)
					if err != nil {
						fmt.Printf("error on MC get transaction, err: %v \n", err)
						continue
					}
					transactions = append(transactions, common.ToTxMc(transaction))
				}
				addresses = append(addresses, common.AddressTransactions{
					Address:      addr,
					Transactions: transactions,
				})
			}
		}
	}
	return addresses
}
