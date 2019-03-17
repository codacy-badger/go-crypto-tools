package wallet_cli_xsn

import (
	"testing"
)

func TestXSNMNParse(t *testing.T) {
	data, err := GetXSNMNList("test", true)
	if err != nil {
		t.Fail()
		return
	}
	testMN := data[0]
	if testMN.IP != "149.28.141.200:62583" {
		t.Fail()
		return
	}
	if testMN.Protocol != "70209" {
		t.Fail()
		return
	}
	if testMN.Output != "6067a102d131ea27fd4a3fb8c6b189e9f809a42398457df8a5c5f4ba9e650200" {
		t.Fail()
		return
	}
	if testMN.Address != "XypnJqfangXpSFVe3wYcKjQb4UzkNVjE63" {
		t.Fail()
		return
	}
}
