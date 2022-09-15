package handle

import (
	"github.com/gin-gonic/gin"
	"github.com/scorpiotzh/mylog"
	"go.uber.org/zap"
	"net/http"
	"super-view-server/http_server/api_code"
	"super-view-server/utils"
)

var (
	log = mylog.NewLogger("handle", mylog.LevelDebug)
)

type ReqCheckHealth struct {
}

type RespCheckHealth struct {
}

func (h *HttpHandle) CheckHealth(ctx *gin.Context) {
	var (
		funcName = "CheckHealth"
		clientIp = GetClientIp(ctx)
		req      ReqCheckHealth
		apiResp  api_code.ApiResp
		err      error
	)

	log.Info("ApiReq:", zap.String("funcName", funcName), zap.String("clientIp", clientIp), zap.String("req", utils.JsonString(req)))

	if err = h.doCheckHealth(&req, &apiResp); err != nil {
		log.Error("doCheckHealth err:", zap.String("err", err.Error()), zap.String("funcName", funcName), zap.String("clientIp", clientIp))
	}

	ctx.JSON(http.StatusOK, apiResp)
}

func (h *HttpHandle) doCheckHealth(req *ReqCheckHealth, apiResp *api_code.ApiResp) error {
	var resp RespCheckHealth
	apiResp.ApiRespOK(resp)
	return nil
}
