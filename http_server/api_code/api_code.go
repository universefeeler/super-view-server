package api_code

import "time"

type ApiCode = int

const (
	ApiCodeSuccess        ApiCode = 200
	ApiServerError        ApiCode = 500
	ApiCodeParamsInvalid  ApiCode = 10000
	ApiCodeMethodNotExist ApiCode = 10001
	ApiCodeDbError        ApiCode = 10002
	ApiGetConfigError     ApiCode = 10003

	ApiCodeTransactionNotExist ApiCode = 11001
	ApiCodeNotSupportAddress   ApiCode = 11005
	ApiCodeInsufficientBalance ApiCode = 11007
	ApiCodeTxExpired           ApiCode = 11008
	ApiCodeAmountInvalid       ApiCode = 11010
	ApiCodeRejectedOutPoint    ApiCode = 11011
	ApiCodeNotEnoughChange     ApiCode = 11014
	ApiCodeSyncBlockNumber     ApiCode = 11012
	ApiCodeOperationFrequent   ApiCode = 11013

	ApiCodeReverseAlreadyExist ApiCode = 12001
	ApiCodeReverseNotExist     ApiCode = 12002

	ApiAwsSendEmailError ApiCode = 20006

	ApiCodeNotOpenForRegistration       ApiCode = 30001
	ApiCodeQueryAccountFailed           ApiCode = 30002
	ApiCodeAccountNotExist              ApiCode = 30003
	ApiCodeAccountAlreadyRegister       ApiCode = 30004
	ApiCodeAccountNotLogin              ApiCode = 30005
	ApiCodeAccountLoginDeviceNotMatch   ApiCode = 30006
	ApiCodeAccountIsExpired             ApiCode = 30007
	ApiCodeAccountLoginFailed           ApiCode = 30008
	ApiCodeAccountSessionInvalid        ApiCode = 30009
	ApiCodePermissionDenied             ApiCode = 30011
	ApiCodeAccountHavenExist            ApiCode = 30012
	ApiCodeAccountLenInvalid            ApiCode = 30014
	ApiCodeReservedAccount              ApiCode = 30017
	ApiCodeGenerateAccountFailed        ApiCode = 30018
	ApiCodeSystemUpgrade                ApiCode = 30019
	ApiCodeAddressHavenExist            ApiCode = 30020
	ApiCodeSameLock                     ApiCode = 30023
	ApiCodeCheckAccountExistError       ApiCode = 30026
	ApiCodeOrderPaid                    ApiCode = 30027
	ApiCodeUnAvailableAccount           ApiCode = 30029
	ApiCodeAccountStatusOnSaleOrAuction ApiCode = 30031
	ApiCodePayTypeInvalid               ApiCode = 30032
	ApiCodeSameOrderInfo                ApiCode = 30033
	ApiCodeAccountPasswordNotRight      ApiCode = 30034
	ApiCodeCallEmailCodeNotExist        ApiCode = 30035
	ApiCodeCallEmailCodeNotMatch        ApiCode = 30036
	ApiCodeCallScasInvalid              ApiCode = 30037
	ApiCodeCallScasEmailBindFailed      ApiCode = 30038
)

const (
	MethodSuperAccountRegister     = "super_accountRegister"
	MethodSuperAccountChange       = "super_accountChange"
	MethodSuperAccountDelete       = "super_accountDelete"
	MethodSuperAccountDetail       = "super_accountDetail"
	MethodSuperAccountSendCode     = "super_userSendCode"
	MethodSuperAccountSignin       = "super_userSignin"
	MethodSuperAccountTransactions = "super_userTransactions"
	MethodSuperAddressRiskControl  = "super_userRiskControl"

	MethodSuperAddressAccountBind = "super_addressAccountBind"

	MethodSuperAddressDeclare = "super_addressDeclare"
	MethodSuperAddressChange  = "super_addressChange"
	MethodSuperAddressList    = "super_addressList"

	MethodObserverAddressCud  = "super_observerAddressCud"
	MethodObserverAddressList = "super_observerAddressList"

	MethodTokenList         = "super_tokenList"
	MethodConfigInfo        = "super_config"
	MethodAccountList       = "super_accountList"
	MethodAccountMine       = "super_myAccounts"
	MethodAccountDetail     = "super_accountDetail"
	MethodAccountRecords    = "super_accountParsingRecords"
	MethodReverseLatest     = "super_reverseLatest"
	MethodReverseList       = "super_reverseList"
	MethodTransactionStatus = "super_transactionStatus"
	MethodBalanceInfo       = "super_balanceInfo"
	MethodTransactionList   = "super_transactionList"
	MethodRewardsMine       = "super_myRewards"
	MethodWithdrawList      = "super_withdrawList"
	MethodAccountSearch     = "super_accountSearch"
	MethodRegisteringList   = "super_registeringAccounts"
	MethodOrderDetail       = "super_orderDetail"

	MethodReverseDeclare   = "super_reverseDeclare"
	MethodReverseRedeclare = "super_reverseRedeclare"
	MethodReverseRetract   = "super_reverseRetract"
	MethodTransactionSend  = "super_transactionSend"
	MethodBalanceWithdraw  = "super_balanceWithdraw"
	MethodBalanceTransfer  = "super_balanceTransfer"
	MethodEditManager      = "super_editManager"
	MethodEditOwner        = "super_transferAccount"
	MethodEditRecords      = "super_editRecords"
	MethodOrderRenew       = "super_submitRenewOrder"
	MethodBalancePay       = "super_balancePay"
	MethodOrderRegister    = "super_submitRegisterOrder"
	MethodOrderChange      = "super_changeOrder"
	MethodOrderPayHash     = "super_doOrderPayHash"

	MethodCkbRpc = "super_ckbRpc"
)

type ApiResp struct {
	ResultCode ApiCode     `json:"result_code"`
	TimeStamp  int64       `json:"timestamp"`
	Msg        string      `json:"msg"`
	Data       interface{} `json:"data"`
}

func ApiRespOK(data interface{}) ApiResp {
	return ApiResp{
		ResultCode: ApiCodeSuccess,
		Msg:        "success",
		TimeStamp:  time.Now().UnixNano() / 1e6,
		Data:       data,
	}
}

func ApiRespErr(errNo ApiCode, errMsg string) ApiResp {
	return ApiResp{
		ResultCode: errNo,
		Msg:        errMsg,
		TimeStamp:  time.Now().UnixNano() / 1e6,
		Data:       nil,
	}
}

type ApiRsp[T any] struct {
	ResultCode ApiCode `json:"result_code"`
	TimeStamp  int64   `json:"timestamp"`
	Msg        string  `json:"msg"`
	Data       T       `json:"data"`
}

func ApiRspFail[T any](errNo ApiCode, errMsg string) ApiRsp[T] {
	return ApiRsp[T]{
		ResultCode: errNo,
		Msg:        errMsg,
		TimeStamp:  time.Now().UnixMilli(),
	}
}

func ApiRspSuccess[T any](data T) ApiRsp[T] {
	return ApiRsp[T]{
		ResultCode: ApiCodeSuccess,
		Msg:        "success",
		TimeStamp:  time.Now().UnixMilli(),
		Data:       data,
	}
}

func (a *ApiResp) ApiRespErr(errNo ApiCode, errMsg string) {
	a.ResultCode = errNo
	a.Msg = errMsg
	a.TimeStamp = time.Now().UnixNano() / 1e6
}

func (a *ApiResp) ApiRespOK(data interface{}) {
	a.ResultCode = ApiCodeSuccess
	a.Data = data
	a.Msg = "success"
	a.TimeStamp = time.Now().UnixNano() / 1e6
}
