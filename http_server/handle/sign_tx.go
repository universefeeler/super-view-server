package handle

import (
	"encoding/json"
	"fmt"
	"github.com/dotbitHQ/das-lib/common"
	"github.com/dotbitHQ/das-lib/sign"
	"github.com/dotbitHQ/das-lib/txbuilder"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strings"
	"super-view-server/http_server/api_code"
	"super-view-server/utils"
)

type ReqSignTx struct {
	ChainId int    `json:"chain_id"`
	Private string `json:"private"`
	SignInfo
}

type RespSignTx struct {
	SignInfo
}

func (h *HttpHandle) RpcSignTx(p json.RawMessage, apiResp *api_code.ApiResp) {
	var req []ReqSignTx
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

	if err = h.doSignTx(&req[0], apiResp); err != nil {
		log.Error("doSignTx err:", zap.Error(err))
	}
}

func (h *HttpHandle) SignTx(ctx *gin.Context) {
	var (
		funcName = "SignTx"
		clientIp = GetClientIp(ctx)
		req      ReqSignTx
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

	if err = h.doSignTx(&req, &apiResp); err != nil {
		log.Error("doSignTx err:", zap.String("err", err.Error()), zap.String("funcName", funcName), zap.String("clientIp", clientIp))
	}

	ctx.JSON(http.StatusOK, apiResp)
}

func (h *HttpHandle) doSignTx(req *ReqSignTx, apiResp *api_code.ApiResp) error {
	var resp RespSignTx
	resp.SignKey = req.SignKey
	var signData []byte
	var err error
	for _, v := range req.SignList {
		switch v.SignType {
		case common.DasAlgorithmIdTron:
			signData, err = sign.TronSignature(true, common.Hex2Bytes(v.SignMsg), req.Private)
			if err != nil {
				err = fmt.Errorf("sign.TronSignature err: %s", err.Error())
				apiResp.ApiRespErr(api_code.ApiServerError, err.Error())
				return err
			}
		case common.DasAlgorithmIdEth:
			signData, err = sign.PersonalSignature(common.Hex2Bytes(v.SignMsg), req.Private)
			if err != nil {
				err = fmt.Errorf("sign.PersonalSignature err: %s", err.Error())
				apiResp.ApiRespErr(api_code.ApiServerError, err.Error())
				return err
			}
		case common.DasAlgorithmIdEd25519:
			signData = sign.Ed25519Signature(common.Hex2Bytes(req.Private), common.Hex2Bytes(v.SignMsg))
			signData = append(signData, []byte{1}...)
		case common.DasAlgorithmIdEth712:
			var obj3 apitypes.TypedData
			mmJson := req.MMJson.String()

			log.Info("old mmJson:", zap.String("mmJson", mmJson))
			oldChainId := fmt.Sprintf("chainId\":%d", req.ChainId)
			newChainId := fmt.Sprintf("chainId\":\"%d\"", req.ChainId)
			mmJson = strings.ReplaceAll(mmJson, oldChainId, newChainId)
			oldDigest := "\"digest\":\"\""
			newDigest := fmt.Sprintf("\"digest\":\"%s\"", v.SignMsg)
			mmJson = strings.ReplaceAll(mmJson, oldDigest, newDigest)
			log.Info("new mmJson:", zap.String("mmJson", mmJson))
			_ = json.Unmarshal([]byte(mmJson), &obj3)
			var mmHash, signature []byte
			mmHash, signature, err = sign.EIP712Signature(obj3, req.Private)
			if err != nil {
				err = fmt.Errorf("sign.EIP712Signature err: %s", err.Error())
				apiResp.ApiRespErr(api_code.ApiServerError, err.Error())
				return err
			}
			log.Info("EIP712Signature mmHash:", zap.String("mmHash", common.Bytes2Hex(mmHash)))
			log.Info("EIP712Signature signature:", zap.String("signature", common.Bytes2Hex(signature)))
			signData = append(signature, mmHash...)

			hexChainId := fmt.Sprintf("%x", req.ChainId)
			chainIdData := common.Hex2Bytes(fmt.Sprintf("%016s", hexChainId))
			signData = append(signData, chainIdData...)
		default:
			apiResp.ApiRespErr(api_code.ApiServerError, fmt.Sprintf("not support sign type [%d]", v.SignType))
			return nil
		}
		resp.SignList = append(resp.SignList, txbuilder.SignData{
			SignType: v.SignType,
			SignMsg:  common.Bytes2Hex(signData),
		})
	}
	apiResp.ApiRespOK(resp)
	return nil
}
