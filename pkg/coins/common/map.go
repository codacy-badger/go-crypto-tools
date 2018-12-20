package common

import (
	"strconv"

	"github.com/cwntr/crypto-sdk/pkg/coins/grph"
	"github.com/cwntr/crypto-sdk/pkg/coins/mc"
)

func ToAddrGrph(adr grph.Address) (ca Address) {
	ca.Address = adr.Address
	ca.Sent = adr.Sent
	ca.Received = adr.Received
	ca.Balance = adr.Balance
	lastTxs := make([]LastTx, 0)
	for _, tx := range adr.LastTxs {
		lastTxs = append(lastTxs, LastTx{Type: tx.Type, Address: tx.Addresses})
	}
	ca.LastTxs = lastTxs
	return ca
}

func ToAddrMc(adr mc.Address) (ca Address) {
	ca.Address = adr.Address
	ca.Sent = adr.Sent
	ca.Received = adr.Received
	f, _ := strconv.ParseFloat(adr.Balance, 64)
	ca.Balance = f
	lastTxs := make([]LastTx, 0)
	for _, tx := range adr.LastTxs {
		lastTxs = append(lastTxs, LastTx{Type: tx.Type, Address: tx.Addresses})
	}
	ca.LastTxs = lastTxs
	return ca
}

func ToTxMc(tx mc.Transaction) (ctxi Transaction) {
	ctxi.ID = tx.Txid
	ctxi.Hex = tx.Hex
	ctxi.Version = tx.Version
	ctxi.LockTime = tx.Locktime
	vins := make([]Vin, 0)
	for _, vin := range tx.Vin {
		v := Vin{}
		v.Vout = strconv.Itoa(vin.Vout)
		v.TxID = vin.Txid
		v.Sequence = strconv.Itoa(int(vin.Sequence))
		v.ScriptSig = ScriptSig{Asm: vin.ScriptSig.Asm, Hex: vin.ScriptSig.Hex}
		vins = append(vins, v)
	}
	ctxi.Vin = vins
	vouts := make([]Vout, 0)
	for _, vout := range tx.Vout {
		v := Vout{}
		v.Value = vout.Value
		v.N = vout.N
		v.ScriptPubKey = ScriptPubKey{
			Asm:       vout.ScriptPubKey.Asm,
			Hex:       vout.ScriptPubKey.Hex,
			ReqSigs:   vout.ScriptPubKey.ReqSigs,
			Type:      vout.ScriptPubKey.Type,
			Addresses: vout.ScriptPubKey.Addresses,
		}
		vouts = append(vouts, v)
	}

	ctxi.Vout = vouts
	ctxi.BlockHash = tx.Blockhash
	ctxi.Confirmations = tx.Confirmations
	ctxi.Time = tx.Time
	ctxi.BlockTime = tx.Blocktime
	ctxi.TxID = tx.Txid
	return ctxi
}

func ToTxGrph(tx grph.Transaction) (ctxi Transaction) {
	ctxi.ID = tx.TxID
	ctxi.Hex = tx.Hex
	ctxi.Version = tx.Version
	ctxi.LockTime = tx.LockTime
	vins := make([]VinItem, 0)
	for _, vin := range tx.VinList {
		v := VinItem{}
		v.Address = vin.Address
		v.TxId = vin.TxID
		v.Amount= vin.Amount
		//v.ScriptSig = ScriptSig{Asm: vin.ScriptSig, Hex: vin.ScriptSig.Hex}
		vins = append(vins, v)
	}
	ctxi.VinList = vins
	vouts := make([]Vout, 0)
	for _, vout := range tx.Vout {
		v := Vout{}
		v.Value = vout.Value
		v.N = vout.N
		v.ScriptPubKey = ScriptPubKey{
			Asm:       vout.ScriptPubKey.Asm,
			Hex:       vout.ScriptPubKey.Hex,
			ReqSigs:   vout.ScriptPubKey.ReqSigs,
			Type:      vout.ScriptPubKey.Type,
			Addresses: vout.ScriptPubKey.Addresses,
		}
		vouts = append(vouts, v)
	}

	ctxi.Vout = vouts
	ctxi.BlockHash = tx.BlockHash
	ctxi.Confirmations = tx.Confirmations
	ctxi.Time = tx.Time
	ctxi.BlockTime = tx.BlockTime
	ctxi.TxID = tx.TxID
	return ctxi
}