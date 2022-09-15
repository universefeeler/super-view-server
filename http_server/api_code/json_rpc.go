package api_code

import "encoding/json"

type JsonRequest struct {
	ID      int64           `json:"id"`
	JsonRpc string          `json:"jsonrpc"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params"`
}

type JsonResponse struct {
	ID      int64       `json:"id"`
	JsonRpc string      `json:"jsonrpc"`
	Result  interface{} `json:"result"`
}

func (j *JsonResponse) ResultData(data interface{}) {
	j.Result = data
}
