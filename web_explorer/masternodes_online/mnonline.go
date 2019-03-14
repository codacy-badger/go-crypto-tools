package masternodes_online

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

const (
	ExplorerMasternodesOnline = "https://explorer.masternodes.online/currencies/%s/"
)

type MNOnlineResponse struct {
	IP       string
	Port     string
	Protocol string
}
type MasternodesOnlineClient struct {
	Client *http.Client
	coin   string
}

func NewMasternodesOnlineClient(coin string) *MasternodesOnlineClient {
	m := MasternodesOnlineClient{}
	m.Client = &http.Client{}
	m.coin = coin
	return &m
}

func (i *MasternodesOnlineClient) GetLink() string {
	return fmt.Sprintf(ExplorerMasternodesOnline, i.coin)
}

func (i *MasternodesOnlineClient) GetIPsForPage(page int) (list []MNOnlineResponse, pageNr int, err error) {
	u, _ := url.Parse(i.GetLink())
	q := u.Query()
	q.Set("page", strconv.Itoa(page))
	u.RawQuery = q.Encode()
	res, err := i.Client.Get(u.String())
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return
	}
	return parseIPTable(body)
}

func parseIPTable(body []byte) (list []MNOnlineResponse, pageNr int, err error) {
	var re = regexp.MustCompile(`(?m)<tr><td>(.*?)<\/td><td>(\d+)(.*)ENABLED(.*)<span(.*?)>(\d+)`)
	for _, match := range re.FindAllStringSubmatch(string(body), -1) {
		item := MNOnlineResponse{}
		item.IP = match[1]
		item.Protocol = match[6]
		item.Port = match[2]
		list = append(list, item)
	}

	pageNr = 1
	var maxPage = regexp.MustCompile(`(?m)\?page=(\d+)`)
	for _, match := range maxPage.FindAllStringSubmatch(string(body), -1) {
		lastPage := match[1]
		tPageNr, err := strconv.Atoi(lastPage)
		if err != nil {
			return list, pageNr, err
		}
		if tPageNr > pageNr {
			pageNr = tPageNr
		}
	}
	return
}
