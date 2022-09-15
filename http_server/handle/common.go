package handle

import (
	"crypto/md5"
	"fmt"
	"github.com/dotbitHQ/das-lib/common"
	"github.com/dotbitHQ/das-lib/txbuilder"
	"github.com/scorpiotzh/toolib"
	"time"
)

type Pagination struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

func (p Pagination) GetLimit() int {
	if p.Size < 1 || p.Size > 100 {
		return 100
	}
	return p.Size
}

func (p Pagination) GetOffset() int {
	page := p.Page
	if p.Page < 1 {
		page = 1
	}
	size := p.GetLimit()
	return (page - 1) * size
}

// ======

type SignInfo struct {
	SignKey  string               `json:"sign_key"`  // sign tx key
	SignList []txbuilder.SignData `json:"sign_list"` // sign list
	MMJson   *common.MMJsonObj    `json:"mm_json"`   // 712 mmjson
}

func (s *SignInfo) SignListString() string {
	return toolib.JsonString(s.SignList)
}

type SignInfoCache struct {
	ChainType common.ChainType                   `json:"chain_type"`
	Address   string                             `json:"address"`
	Action    string                             `json:"action"`
	Account   string                             `json:"account"`
	Capacity  uint64                             `json:"capacity"`
	BuilderTx *txbuilder.DasTxBuilderTransaction `json:"builder_tx"`
}

func (s *SignInfoCache) SignKey() string {
	key := fmt.Sprintf("%d%s%s%d", s.ChainType, s.Address, s.Action, time.Now().UnixNano())
	return fmt.Sprintf("%x", md5.Sum([]byte(key)))
}
