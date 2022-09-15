package timer

import (
	"context"
	"go.uber.org/zap"
	"super-view-server/dao"
	"sync"
	"time"
)

type TxTimer struct {
	ctx   context.Context
	wg    *sync.WaitGroup
	dbDao *dao.DbDao
}

type TxTimerParam struct {
	DbDao *dao.DbDao
	Ctx   context.Context
	Wg    *sync.WaitGroup
}

func NewTxTimer(p TxTimerParam) *TxTimer {
	var t TxTimer
	t.ctx = p.Ctx
	t.wg = p.Wg
	t.dbDao = p.DbDao
	return &t
}

func (t *TxTimer) Run() error {
	//go t.doViewFromFile()
	//go t.doRandomPrivateKeyView()
	go t.doRandomMnemonicView()

	//tickerViewFileDataTask := time.NewTicker(time.Hour * 2400)
	//tickerRandomPrivateKeyViewTask := time.NewTicker(time.Hour * 2400)
	tickerRandommnemonicViewTask := time.NewTicker(time.Hour * 2400)
	//improveNftMetaTask := time.NewTicker(time.Minute * 1)
	//pendingLimit := 10
	//tickerRejected := time.NewTicker(time.Minute * 3)
	//tickerExpired := time.NewTicker(time.Minute * 30)
	t.wg.Add(1)

	go func() {
		for {
			select {
			/*
				case <-tickerViewFileDataTask.C:
					log.Info("doViewFromFile start")
					if err := t.doViewFromFile(); err != nil {
						log.Error("doViewFromFile ", zap.ByteString("err", []byte(err.Error())))
					}
					log.Info("doViewFromFile end")

				case <-tickerRandomPrivateKeyViewTask.C:
					log.Info("tickerRandomPrivateKeyViewTask start")
					if err := t.doRandomPrivateKeyView(); err != nil {
						log.Error("doRandomPrivateKeyView ", zap.ByteString("err", []byte(err.Error())))
					}
					log.Info("tickerRandomPrivateKeyViewTask end")
			*/
			case <-tickerRandommnemonicViewTask.C:
				log.Info("tickerRandommnemonicViewTask start")
				if err := t.doRandomMnemonicView(); err != nil {
					log.Error("doRandomMnemonicView ", zap.ByteString("err", []byte(err.Error())))
				}
				log.Info("tickerRandommnemonicViewTask end")

			/*
				case <-improveNftMetaTask.C:
					log.Info("improveNftMetaTask start ...")
					if err := t.doImproveNftMetaTask(); err != nil {
						log.Error("improveNftMetaTask err: ", zap.ByteString("err", []byte(err.Error())))
					}
					log.Info("improveNftMetaTask end ...")
					/*
						case <-tickerVerifiedAddress.C:
							log.Info("doVerifiedAddress start ...")
							if err := t.doVerifiedAddress(); err != nil {
								log.Error("doVerifiedAddress err: ", zap.ByteString("err", []byte(err.Error())))
							}
							log.Info("doVerifiedAddress end ...")
			*/
			case <-t.ctx.Done():
				log.Info("timer done")
				t.wg.Done()
				return
			}
		}
	}()
	return nil
}
