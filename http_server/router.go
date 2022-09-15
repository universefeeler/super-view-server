package http_server

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/scorpiotzh/toolib"
	"go.uber.org/zap"
	"net/http"
	"super-view-server/config"
	"super-view-server/http_server/api_code"
	"time"
)

func (h *HttpServer) initRouter() {
	originList := []string{
		`https:\/\/[^.]*\.superview\.xyz`,
		`https:\/\/superview\.xyz`,
	}
	if len(config.Cfg.Origins) > 0 {
		toolib.AllowOriginList = append(toolib.AllowOriginList, config.Cfg.Origins...)
	} else {
		toolib.AllowOriginList = append(toolib.AllowOriginList, originList...)
	}
	h.engine.Use(toolib.MiddlewareCors())
	v1 := h.engine.Group("v1")
	{
		// cache
		shortExpireTime, _, lockTime := time.Second*5, time.Second*15, time.Minute
		shortDataTime, _ := time.Minute*3, time.Minute*10
		cacheHandleShort := toolib.MiddlewareCacheByRedis(h.rc.GetRedisClient(), false, shortDataTime, lockTime, shortExpireTime, respHandle)
		//cacheHandleShortCookies := tool_lib.MiddlewareCacheByRedis(h.rc.GetRedisClient(), true, shortDataTime, lockTime, shortExpireTime, respHandle)

		v1.GET("/version", api_code.DoMonitorLog("Version"), cacheHandleShort, h.h.Version)
		v1.GET("/system/check", api_code.DoMonitorLog("CheckHealth"), h.h.CheckHealth)
	}

	//internalV1 := h.internalEngine.Group("v1"){}
}

func respHandle(c *gin.Context, res string, err error) {
	if err != nil {
		log.Error("respHandle err:", zap.Error(err))
		c.AbortWithStatusJSON(http.StatusOK, api_code.ApiRespErr(http.StatusInternalServerError, err.Error()))
	} else if res != "" {
		var respMap map[string]interface{}
		_ = json.Unmarshal([]byte(res), &respMap)
		c.AbortWithStatusJSON(http.StatusOK, respMap)
	}
}
