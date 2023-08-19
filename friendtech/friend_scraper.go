package friendtech

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type User struct {
	ID             int     `json:"id"`
	Address        string  `json:"address"`
	TwitterUsername string `json:"twitterUsername"`
	TwitterName    string  `json:"twitterName"`
	TwitterPfpUrl  string  `json:"twitterPfpUrl"`
	TwitterUserId  string  `json:"twitterUserId"`
	LastOnline     int     `json:"lastOnline"`
	DisplayPrice   string  `json:"displayPrice"`
	HolderCount    int     `json:"holderCount"`
	ShareSupply    int     `json:"shareSupply"`
}

type Response struct {
	Users []User `json:"users"`
}


// pulls recently joined Users from v1.
// keep note this data is hr and a half delayed so not the most useful
func GetRecentlyJoinedv1() (*Response, error) {
	url := "https://prod-api.kosetto.com/lists/recently-joined"

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result Response
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}