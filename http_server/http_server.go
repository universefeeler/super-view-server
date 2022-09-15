package http_server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/scorpiotzh/mylog"
	"go.uber.org/zap"
	"net/http"
	"super-view-server/cache"
	"super-view-server/dao"
	"super-view-server/http_server/handle"
)

var (
	log = mylog.NewLogger("config", mylog.LevelDebug)
)

type HttpServer struct {
	ctx             context.Context
	address         string
	engine          *gin.Engine
	srv             *http.Server
	internalAddress string
	internalEngine  *gin.Engine
	internalSrv     *http.Server
	h               *handle.HttpHandle
	rc              *cache.RedisCache
}

type HttpServerParams struct {
	Ctx                    context.Context
	Address                string
	InternalAddress        string
	DbDao                  *dao.DbDao
	Rc                     *cache.RedisCache
	MapReservedAccounts    map[string]struct{}
	MapUnAvailableAccounts map[string]struct{}
}

func Initialize(p HttpServerParams) (*HttpServer, error) {
	hs := HttpServer{
		ctx:             p.Ctx,
		address:         p.Address,
		internalAddress: p.InternalAddress,
		engine:          gin.New(),
		internalEngine:  gin.New(),
		h: handle.Initialize(handle.HttpHandleParams{
			DbDao:                  p.DbDao,
			Rc:                     p.Rc,
			Ctx:                    p.Ctx,
			MapReservedAccounts:    p.MapReservedAccounts,
			MapUnAvailableAccounts: p.MapUnAvailableAccounts,
		}),
		rc: p.Rc,
	}
	return &hs, nil
}

func (h *HttpServer) Run() {
	h.initRouter()
	h.srv = &http.Server{
		Addr:    h.address,
		Handler: h.engine,
	}
	h.internalSrv = &http.Server{
		Addr:    h.internalAddress,
		Handler: h.internalEngine,
	}
	go func() {
		if err := h.srv.ListenAndServe(); err != nil {
			log.Error("http_server run err:", zap.Error(err))
		}
	}()

	go func() {
		if err := h.internalSrv.ListenAndServe(); err != nil {
			log.Error("http_server internal run err:", zap.Error(err))
		}
	}()
}

func (h *HttpServer) Shutdown() {
	if h.srv != nil {
		log.Warn("Http server shutdown ... ")
		if err := h.srv.Shutdown(h.ctx); err != nil {
			log.Error("Http server Shutdown err:", zap.Error(err))
		}
	}
	if h.internalSrv != nil {
		log.Warn("Http server internal shutdown ... ")
		if err := h.internalSrv.Shutdown(h.ctx); err != nil {
			log.Error("Http server internal Shutdown err:", zap.Error(err))
		}
	}
}
