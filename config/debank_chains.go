package config

import (
	"encoding/json"
)

type ChainItem struct {
	Id             string `json:"id"`
	CommunityId    int64  `json:"community_id"`
	Name           string `json:"name"`
	NativeTokenId  string `json:"native_token_id"`
	LogoUrl        string `json:"logo_url"`
	WrappedTokenId string `json:"wrapped_token_id"`
}

var (
	DebankChainMap             = make(map[string]ChainItem) // id-> ChainItem
	DebankChainId2SuperChainId = map[string]string{
		"eth":   "1",
		"bsc":   "56",
		"matic": "137",
	}
)

func init() {
	jsonData, err := StaticDir.ReadFile("static/json/debank_chains.json")
	if err != nil {
		log.Fatal("read debank_chains.json fail, err=" + err.Error())
		panic(err)
	}
	var chainList []ChainItem
	if err = json.Unmarshal(jsonData, &chainList); err != nil {
		log.Fatal("parse debank_chains.json fail, err=" + err.Error())
		panic(err)
	}

	for _, item := range chainList {
		DebankChainMap[item.Id] = item
	}
}
