package creative

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// MaterialReadRequest 创意素材信息
type MaterialReadRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CreativeIDs 创意ID集合，支持最大长度为100。创意ID需属于当前广告主，否则会报错
	CreativeIDs []uint64 `json:"creative_ids,omitempty"`
	// Fields 查询字段集合, 默认查询所有字段。详见下方response字段定义; 允许值: "id", "ad_id", "advertiser_id", "title", "image_info","image_mode", "opt_status"
	Fields []string `json:"fields,omitempty"`
}

// Encode implement GetRequest interface
func (r MaterialReadRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if len(r.CreativeIDs) > 0 {
		ids, _ := json.Marshal(r.CreativeIDs)
		values.Set("creative_ids", string(ids))
	}
	if len(r.Fields) > 0 {
		fields, _ := json.Marshal(r.Fields)
		values.Set("fields", string(fields))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// MaterialReadResponse 创意素材信息 API Response
type MaterialReadResponse struct {
	model.BaseResponse
	// Data json返回值
	Data []Material `json:"data,omitempty"`
}
