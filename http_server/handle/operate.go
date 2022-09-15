package handle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"super-view-server/http_server/api_code"
	"super-view-server/utils"
	"time"
)

func (h *HttpHandle) Operate(ctx *gin.Context) {
	var (
		req       api_code.JsonRequest
		resp      api_code.JsonResponse
		apiResp   api_code.ApiResp
		clientIp  = GetClientIp(ctx)
		startTime = time.Now()
	)
	resp.Result = &apiResp

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Error("ShouldBindJSON err:", zap.Error(err))
		apiResp.ApiRespErr(api_code.ApiCodeParamsInvalid, "params invalid")
		ctx.JSON(http.StatusOK, resp)
		return
	}

	resp.ID, resp.JsonRpc = req.ID, req.JsonRpc
	log.Info("Operate:", zap.String("req.Method", req.Method), zap.String("clientIp", clientIp), zap.String("req", utils.JsonString(req)))

	switch req.Method {
	case api_code.MethodOrderPayHash:
		//h.RpcOrderPayHash(req.Params, &apiResp)
	default:
		log.Error("method not exist:", zap.String("req.Method", req.Method))
		apiResp.ApiRespErr(api_code.ApiCodeMethodNotExist, fmt.Sprintf("method [%s] not exits", req.Method))
	}

	api_code.DoMonitorLogRpc(&apiResp, req.Method, clientIp, startTime)

	ctx.JSON(http.StatusOK, resp)
	return
}
