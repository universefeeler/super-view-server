package chain

type EvmChainInfo struct {
	Name           string         `json:"name"`
	ShortName      string         `json:"shortName"`
	Chain          string         `json:"chain"`
	Network        string         `json:"network"`
	ChainId        uint64         `json:"chainId"`
	NetworkId      uint64         `json:"networkId"`
	Rpc            []string       `json:"rpc"`
	Faucets        []string       `json:"faucets"`
	Explorers      []Explorer     `json:"explorers"`
	InfoURL        string         `json:"infoURL"`
	Title          string         `json:"title"`
	NativeCurrency NativeCurrency `json:"nativeCurrency"`
}

type NativeCurrency struct {
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Decimals string `json:"decimals"`
}

type Explorer struct {
	Name     string `json:"name"`
	Url      string `json:"url"`
	Standard string `json:"standard"`
}
