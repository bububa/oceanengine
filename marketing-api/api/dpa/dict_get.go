package dpa

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/dpa"
)

// DictGet 获取DPA词包
func DictGet(clt *core.SDKClient, accessToken string, req *dpa.DictGetRequest) ([]dpa.Dict, error) {
	var resp dpa.DictGetResponse
	if err := clt.Get("2/dpa/dict/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
