package adraise

import (
	"encoding/json"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model/tools/adraise"
)

// Status 获取当前起量状态: 获取当前起量状态
func Status(clt *core.SDKClient, accessToken string, req *adraise.StatusRequest) (map[uint64]enum.AdRaiseStatus, error) {
	var resp adraise.StatusResponse
	if err := clt.Get("2/tools/ad_raise_status/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	ret := make(map[uint64]enum.AdRaiseStatus, len(req.AdIDs))
	if err := json.Unmarshal([]byte(resp.Data.Status), &ret); err != nil {
		return nil, err
	}
	return ret, nil
}
