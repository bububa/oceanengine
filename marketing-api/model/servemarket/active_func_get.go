package servemarket

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ActiveFuncGetRequest 获取用户已购功能点列表
type ActiveFuncGetRequest struct {
	// AppID 应用 id
	AppID uint64 `json:"app_id,omitempty"`
	// UseUid 使用用户id
	UseUid uint64 `json:"use_uid,omitempty"`
	// FuncKeys 功能Key列表，不传则返回全部
	FuncKeys []string `json:"func_keys,omitempty"`
}

// Encode implement GetRequest interface
func (r ActiveFuncGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("app_id", strconv.FormatUint(r.AppID, 10))
	values.Set("use_uid", strconv.FormatUint(r.UseUid, 10))
	if len(r.FuncKeys) > 0 {
		bs, _ := json.Marshal(r.FuncKeys)
		values.Set("func_keys", string(bs))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// ActiveFuncGetResponse 获取用户已购功能点列表
type ActiveFuncGetResponse struct {
	model.BaseResponse
	Data struct {
		// FuncList 已购功能key列表，仅返回用户已购且在购买有效期内的功能列表
		FuncList []OrderFunction `json:"func_list,omitempty"`
	} `json:"data,omitempty"`
}
