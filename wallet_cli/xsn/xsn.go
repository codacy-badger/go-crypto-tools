package wallet_cli_xsn

import (
	"regexp"
	"strings"

	"github.com/cwntr/go-crypto-tools/wallet_cli"
)

const (
	//Masternode CLI commands
	ArgumentMasternodeList           = "masternodelist"
	ArgumentMasternodeListOptionFull = "full"

	MNStatusEnabled          = "ENABLED"
	MNStatusNewStartRequired = "NEW_START_REQUIRED"

	//Governance CLI commands
	ArgumentGovernanceObject         = "gobject"
	ArgumentGovernanceObjectCount    = "count"
	ArgumentGovernanceObjectList     = "list"
	ArgumentGovernanceObjectAll      = "all"
	ArgumentGovernanceObjectGet      = "get"
	ArgumentGovernanceObjectGetVotes = "getvotes"
	ArgumentGovernanceInfo           = "getgovernanceinfo"

	//Merchant CLI commands
	ArgumentMerchantNodeList     = "merchantnodelist"
	ArgumentMerchantNodeListFull = "full"

	DateTimeFmt = "2006-01-02 15:04:05"
	DateFmt     = "2006-01-02"
)

type XSNMNItem struct {
	IP       string `json:"ip"`
	Protocol string `json:"protocol"`
	Output   string `json:"output"`
	Address  string `json:"address"`
}

func GetXSNMNList(cli string, isTestRequest bool) (xsnItems []XSNMNItem, err error) {
	var data string
	if isTestRequest {
		data, err = GetTestMNLIST(), nil
	} else {
		data, err = wallet_cli.ExecCLI(cli, ArgumentMasternodeList, ArgumentMasternodeListOptionFull)
	}

	if err != nil {
		return []XSNMNItem{}, err
	}

	var rIp = regexp.MustCompile(`(?m)\d+.\d+.\d+.\d+:\d+`)
	var rProtocol = regexp.MustCompile(`7020\d`)
	var rOutput = regexp.MustCompile(`(?m)COutPoint\((\w+)`)
	var rAddress = regexp.MustCompile(`(?m)\s+(X\w+)\s+`)

	rows := strings.Split(data, `",`)

	for _, r := range rows {
		if strings.Contains(r, MNStatusEnabled) || strings.Contains(r, MNStatusNewStartRequired) {
			x := XSNMNItem{}
			//@TODO: use multi regex match instead of each field individually
			for _, match := range rIp.FindAllString(r, -1) {
				x.IP = match
			}
			for _, match := range rProtocol.FindAllString(r, -1) {
				x.Protocol = match
			}
			for _, match := range rOutput.FindAllStringSubmatch(r, -1) {
				x.Output = match[1]
			}
			for _, match := range rAddress.FindAllString(r, -1) {
				x.Address = strings.TrimSpace(match)
			}
			xsnItems = append(xsnItems, x)
		}
	}
	return xsnItems, nil
}

func GetXSNGovernanceTotals(cli string, isTestRequest bool) (gov XSNGovTotal, err error) {
	var data string
	if isTestRequest {
		data, err = TestXSNGovernanceTotals(cli, ArgumentGovernanceObject, ArgumentGovernanceObjectCount)
	} else {
		data, err = wallet_cli.ExecCLI(cli, ArgumentGovernanceObject, ArgumentGovernanceObjectCount)
	}
	return GovTotalToObject(data), err
}

func GetXSNGovInfo(cli string, isTestRequest bool) (gov XSNGovernanceInfo, err error) {
	var data string
	if isTestRequest {
		data, err = TestXSNGovInfo(cli, ArgumentGovernanceInfo)
	} else {
		data, err = wallet_cli.ExecCLI(cli, ArgumentGovernanceInfo)
	}
	return GovInfoToObject(data), err
}

func GetXSNGovObjects(cli string, isTestRequest bool) (XSNGovernanceList, error) {
	var data string
	var err error

	if isTestRequest {
		data, err = TestXSNGovObjects(cli, ArgumentGovernanceObject, ArgumentGovernanceObjectList, ArgumentGovernanceObjectAll)
	} else {
		data, err = wallet_cli.ExecCLI(cli, ArgumentGovernanceObject, ArgumentGovernanceObjectList, ArgumentGovernanceObjectAll)
	}
	if err != nil {
		return XSNGovernanceList{}, err
	}
	return GovObjectsToList(data)
}

func GetXSNGovObject(cli string, hash string, isTestRequest bool) (XSNGovObjectWithVotes, error) {
	var objStr string
	var err error

	//get full governance object
	if isTestRequest {
		objStr, err = TestXSNFullGovObject(cli, ArgumentGovernanceObject, ArgumentGovernanceObjectGet, hash)
	} else {
		objStr, err = wallet_cli.ExecCLI(cli, ArgumentGovernanceObject, ArgumentGovernanceObjectGet, hash)
	}
	if err != nil {
		return XSNGovObjectWithVotes{}, err
	}
	obj, err := ParseXSNFullGovObject(objStr)
	if err != nil {
		return XSNGovObjectWithVotes{}, err
	}

	//fetch current MN list
	mn, err := GetXSNMNList(cli, isTestRequest)
	if err != nil {
		return XSNGovObjectWithVotes{}, err
	}

	//fetch current votes for full governance object hash
	var votesStr string
	if isTestRequest {
		votesStr, err = TestXSNGovObjectVotes(cli, ArgumentGovernanceObject, ArgumentGovernanceObjectGetVotes, hash)
	} else {
		votesStr, err = wallet_cli.ExecCLI(cli, ArgumentGovernanceObject, ArgumentGovernanceObjectGetVotes, hash)
	}
	if err != nil {
		return XSNGovObjectWithVotes{}, err
	}

	//match current votes with current masternode list
	votes, err := ParseVotes(votesStr, mn)
	if err != nil {
		return XSNGovObjectWithVotes{}, err
	}
	return XSNGovObjectWithVotes{Votes: votes, Object: obj}, nil
}

func GetXSNMerchants(cli string, isTestRequest bool) ([]XSNMerchantItem, error) {
	var data string
	var err error
	if isTestRequest {
		data, err = TestXSNMerchantListCLI(cli, ArgumentMerchantNodeList, ArgumentMerchantNodeListFull)
	} else {
		data, err = wallet_cli.ExecCLI(cli, ArgumentMerchantNodeList, ArgumentMerchantNodeListFull)
	}
	if err != nil {
		return []XSNMerchantItem{}, err
	}
	return ParseXSNMerchants(data), nil
}
