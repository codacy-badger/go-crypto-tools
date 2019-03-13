package xrp

import "time"

type AccountTransaction struct {
	Result       string `json:"result"`
	Count        int    `json:"count"`
	Marker       string `json:"marker"`
	Transactions []struct {
		Hash        string    `json:"hash"`
		LedgerIndex int       `json:"ledger_index"`
		Date        time.Time `json:"date"`
		Tx          struct {
			TransactionType string `json:"TransactionType"`
			Flags           int    `json:"Flags"`
			Sequence        int    `json:"Sequence"`
			Amount          interface{} `json:"Amount"`
			Fee             string `json:"Fee"`
			SigningPubKey   string `json:"SigningPubKey"`
			TxnSignature    string `json:"TxnSignature"`
			Account         string `json:"Account"`
			Destination     string `json:"Destination"`
		} `json:"tx"`
		Meta struct {
			TransactionIndex int `json:"TransactionIndex"`
			AffectedNodes    []struct {

				CreatedNode struct {
					LedgerEntryType string `json:"LedgerEntryType"`
					LedgerIndex     string `json:"LedgerIndex"`
					NewFields       struct {
						Sequence int    `json:"Sequence"`
						Balance  interface{} `json:"Balance"`
						Account  string `json:"Account"`
					} `json:"NewFields"`
				} `json:"CreatedNode,omitempty"`
				ModifiedNode struct {
					LedgerEntryType   string `json:"LedgerEntryType"`
					PreviousTxnLgrSeq int    `json:"PreviousTxnLgrSeq"`
					PreviousTxnID     string `json:"PreviousTxnID"`
					LedgerIndex       string `json:"LedgerIndex"`
					PreviousFields    struct {
						Sequence int    `json:"Sequence"`
						Balance  interface{} `json:"Balance"`
					} `json:"PreviousFields"`
					FinalFields struct {
						Flags      int    `json:"Flags"`
						Sequence   int    `json:"Sequence"`
						OwnerCount int    `json:"OwnerCount"`
						Balance    interface{} `json:"Balance"`
						Account    string `json:"Account"`
					} `json:"FinalFields"`
				} `json:"ModifiedNode,omitempty"`
			} `json:"AffectedNodes"`
			TransactionResult string `json:"TransactionResult"`
			DeliveredAmount   interface{} `json:"delivered_amount"`
		} `json:"meta"`
	} `json:"transactions"`
}
