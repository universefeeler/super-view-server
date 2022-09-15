package handle

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"super-view-server/http_server/api_code"
	"super-view-server/utils"
	"time"
)

type ReqVersion struct {
}

type RespVersion struct {
	Version string `json:"version"`
}

func (h *HttpHandle) RpcVersion(p json.RawMessage, apiResp *api_code.ApiResp) {
	var req []ReqVersion
	err := json.Unmarshal(p, &req)
	if err != nil {
		log.Error("json.Unmarshal err:", zap.Error(err))
		apiResp.ApiRespErr(api_code.ApiCodeParamsInvalid, "params invalid")
		return
	} else if len(req) == 0 {
		log.Error("len(req) is 0")
		apiResp.ApiRespErr(api_code.ApiCodeParamsInvalid, "params invalid")
		return
	}

	if err = h.doVersion(&req[0], apiResp); err != nil {
		log.Error("doVersion err:", zap.Error(err))
	}
}

func (h *HttpHandle) Version(ctx *gin.Context) {
	var (
		funcName = "Version"
		clientIp = GetClientIp(ctx)
		req      ReqVersion
		apiResp  api_code.ApiResp
		err      error
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Error("ShouldBindJSON err: ", zap.String("err", err.Error()), zap.String("funcName", funcName), zap.String("clientIp", clientIp))
		apiResp.ApiRespErr(api_code.ApiCodeParamsInvalid, "params invalid")
		ctx.JSON(http.StatusOK, apiResp)
		return
	}
	log.Info("ApiReq:", zap.String("funcName", funcName), zap.String("clientIp", clientIp), zap.String("req", utils.JsonString(req)))

	if err = h.doVersion(&req, &apiResp); err != nil {
		log.Error("doVersion err:", zap.String("err", err.Error()), zap.String("funcName", funcName), zap.String("clientIp", clientIp))
	}

	ctx.JSON(http.StatusOK, apiResp)
}

func (h *HttpHandle) doVersion(req *ReqVersion, apiResp *api_code.ApiResp) error {
	var resp RespVersion

	resp.Version = time.Now().String()

	apiResp.ApiRespOK(resp)
	return nil
}
