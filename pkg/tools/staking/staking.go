package main

import (
	"fmt"

	"time"

	"github.com/cwntr/crypto-sdk/pkg/coins/common"
	"github.com/dnaeon/go-vcr/recorder"
)

const (
	TestAddrGrphMaster = "gMo9tMCfij7rPPsPz2m59xcrVJMQPQW7gY"
	TestAddrGrphStake  = "g8HvrJEmA8eiRhJpqbzSHynkW4xXzpwnYL"
	TestAddrMcMaster   = "mZzSaKnj4qVpBEY4C6gyJo6kiGVgE9S2j6"
	TestAddrMcStake    = "mdu8n8Dp7cPPJKGyg6Vh1GqbjFKLoCxmTL"
)

func main() {
	// Start our recorder
	r, err := recorder.New("fixtures/staking")
	if err != nil {
		fmt.Printf("err: %v \n", err)
		panic(0)
	}
	defer r.Stop() // Make sure recorder is stopped once done with it

	common.InitClients(r)
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
	for _, adr := range addresses {
		sum := float64(0)
		var dStart time.Time
		var dEnd time.Time
		for i, r := range adr.Rewards {
			if i == 0 {
				dStart = r.Time
			}
			if len(adr.Rewards)-1 == i {
				dEnd = r.Time
			}
			sum += r.Amount
		}
		fmt.Printf("Address: %s | Rewards: %v | Range [%s - %s] \n", adr.Address.Address, sum, dStart.Format(time.RFC3339), dEnd.Format(time.RFC3339))
	}
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
				stakes := make([]common.Reward, 0)
				for _, tx := range addr.LastTxs {
					transaction, err := client.GetTransactionById(tx.Address, true)
					if err != nil {
						fmt.Printf("error on GRPH get transaction, err: %v \n", err)
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
									stakes = append(stakes, common.Reward{
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
					Rewards:      stakes,
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
				stakes := make([]common.Reward, 0)
				for _, tx := range addr.LastTxs {
					transaction, err := client.GetTransactionById(tx.Address, true)
					if err != nil {
						fmt.Printf("error on MC get transaction, err: %v \n", err)
						continue
					}

					cTx := common.ToTxMc(transaction)
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
									stakes = append(stakes, common.Reward{
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
					Rewards:      stakes,
				})
			}
		}
	}
	return addresses
}
