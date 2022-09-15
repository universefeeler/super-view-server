package config

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

var GethMap = &gethMap{innerMap: make(map[string]*ethclient.Client)}

func init() {
	for _, chain := range Cfg.SupportedChains {
		client, err := ethclient.Dial(chain.GethUrl)
		if err != nil {
			log.Error("init geth client fail", zap.String("chainId", chain.ChainId), zap.Error(err))
		}
		GethMap.setClient(chain.ChainId, client)
	}
}

type gethMap struct {
	innerMap map[string]*ethclient.Client
}

func (ass *gethMap) GetClientWithChainId(chainId string) *ethclient.Client {
	return ass.innerMap[chainId]
}

func (ass *gethMap) setClient(chainId string, client *ethclient.Client) {
	ass.innerMap[chainId] = client
}
