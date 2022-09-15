package handle

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"super-view-server/cache"
	"super-view-server/config"
	"super-view-server/dao"
	"super-view-server/http_server/api_code"
)

type HttpHandle struct {
	ctx                    context.Context
	dbDao                  *dao.DbDao
	rc                     *cache.RedisCache
	mapReservedAccounts    map[string]struct{}
	mapUnAvailableAccounts map[string]struct{}
}

type HttpHandleParams struct {
	DbDao                  *dao.DbDao
	Rc                     *cache.RedisCache
	Ctx                    context.Context
	MapReservedAccounts    map[string]struct{}
	MapUnAvailableAccounts map[string]struct{}
}

func Initialize(p HttpHandleParams) *HttpHandle {
	hh := HttpHandle{
		dbDao:                  p.DbDao,
		rc:                     p.Rc,
		ctx:                    p.Ctx,
		mapReservedAccounts:    p.MapReservedAccounts,
		mapUnAvailableAccounts: p.MapUnAvailableAccounts,
	}
	return &hh
}

// 获取IP
func GetClientIp(ctx *gin.Context) string {
	clientIP := fmt.Sprintf("%v", ctx.Request.Header.Get("X-Real-IP"))
	return fmt.Sprintf("(%s)(%s)", clientIP, ctx.Request.RemoteAddr)
}

// post 请求绑定参数
func shouldBindJSON(ctx *gin.Context, req interface{}) (*api_code.ApiResp, error) {
	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp := api_code.ApiRespErr(api_code.ApiCodeParamsInvalid, "params invalid")
		return &resp, err
	} else {
		return &api_code.ApiResp{}, nil
	}
}

func (h *HttpHandle) checkSystemUpgrade(apiResp *api_code.ApiResp) error {
	if config.Cfg.Server.IsUpdate {
		apiResp.ApiRespErr(api_code.ApiCodeSystemUpgrade, "The service is under maintenance, please try again later.")
		return fmt.Errorf("backend system upgrade")
	}
	return nil
}
