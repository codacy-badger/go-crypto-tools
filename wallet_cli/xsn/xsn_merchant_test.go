package wallet_cli_xsn

import "testing"

func TestXSNMerchantParse(t *testing.T) {
	merchants := ParseXSNMerchants(TestXSNMerchantList())
	x := merchants[0]
	if x.ID != "33633537633664623263633231316665383835656539396231306136316530343936363037336635" {
		t.Fail()
		return
	}
	if x.Status != MNStatusEnabled {
		t.Fail()
		return
	}
	if x.Protocol != "70209" {
		t.Fail()
		return
	}
	if x.MerchantAddress != "Xy4fbojK3LiwtfERNnQbDrQEwd3mpSLYbM" {
		t.Fail()
		return
	}
	if x.HashTPoSContractTx != "7f4521bd2249d95c5a88cde71c7f4506bef0446dc5ea2228e33bb73ed65075b1" {
		t.Fail()
		return
	}
	if x.LastSeen != 1550048420 {
		t.Fail()
		return
	}
	if x.ActiveSeconds != 16682298 {
		t.Fail()
		return
	}
	if x.IP != "95.179.146.41:62583" {
		t.Fail()
		return
	}
}
