package main

import (
	"context"
	"fmt"
	"github.com/scorpiotzh/mylog"
	"github.com/scorpiotzh/toolib"
	"super-view-server/cache"
	"super-view-server/config"
	"super-view-server/dao"
	"super-view-server/http_server"
	"super-view-server/timer"
	"sync"
)

var (
	log               = mylog.NewLogger("main", mylog.LevelDebug)
	ctxServer, cancel = context.WithCancel(context.Background())
	wgServer          = sync.WaitGroup{}
)

func main() {
	log.Debug("main start：")

	red, err := toolib.NewRedisClient(config.Cfg.Cache.Redis.Addr, config.Cfg.Cache.Redis.Password, config.Cfg.Cache.Redis.DbNum)
	if err != nil {
		log.Info("NewRedisClient err: %s", err.Error())
		//return fmt.Errorf("NewRedisClient err:%s", err.Error())
	} else {
		log.Info("redis ok")
	}

	// http service
	hs, err := http_server.Initialize(http_server.HttpServerParams{
		Address:                config.Cfg.Server.HttpServerAddr,
		InternalAddress:        config.Cfg.Server.HttpServerInternalAddr,
		DbDao:                  dao.AllDbDao,
		Rc:                     cache.Initialize(red),
		Ctx:                    ctxServer,
		MapReservedAccounts:    config.MapReservedAccounts,
		MapUnAvailableAccounts: config.MapUnAvailableAccounts,
	})
	if err != nil {
		panic(fmt.Errorf("http server Initialize err:%s", err.Error()))
	}
	hs.Run()
	log.Info("httpserver ok")

	// service timer
	txTimer := timer.NewTxTimer(timer.TxTimerParam{
		Ctx: ctxServer,
		Wg:  &wgServer,
	})
	if err = txTimer.Run(); err != nil {
		log.Errorf("txTimer.Run() err:%s", err.Error())
		return
	}
	log.Info("timer ok")

	<-ctxServer.Done()
}
