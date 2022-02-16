package coupon

import (
	"io"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
)

// CodeUploadRequest 上传券码 API Request
type CodeUploadRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// File 券码文件
	// 建议txt格式，券码数量不超过100万，否则可能解析失败。
	// 每行一个券码，券码格式4-30位，字母+数字
	// 查看券码文件示例
	File io.Reader `json:"file,omitempty"`
	// Filename 文件名
	Filename string `json:"filename,omitempty"`
	// Remark 本次上传的标识; 任意字符串，建议每次不一样，会在response中透传回来
	Remark string `json:"remark,omitempty"`
	// ResourceID 优惠券ID，调用卡券详情接口返回
	ResourceID uint64 `json:"resource_id,omitempty"`
}

// Encode implement UploadReqeust interface
func (r CodeUploadRequest) Encode() []model.UploadField {
	ret := make([]model.UploadField, 0, 3)
	ret = append(ret, model.UploadField{
		Key:   "advertiser_id",
		Value: strconv.FormatUint(r.AdvertiserID, 10),
	})
	ret = append(ret, model.UploadField{
		Key:    "file",
		Value:  r.Filename,
		Reader: r.File,
	})
	ret = append(ret, model.UploadField{
		Key:   "remark",
		Value: r.Remark,
	})
	ret = append(ret, model.UploadField{
		Key:   "resource_id",
		Value: strconv.FormatUint(r.ResourceID, 10),
	})
	return ret
}

// CodeUploadResponse 上传券码 API Response
type CodeUploadResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *CodeUploadResponseData `json:"data,omitempty"`
}

// CodeUploadResponseData json返回值
type CodeUploadResponseData struct {
	// Count 本次上传的有效券码总数
	// 卡券未开始前，支持重复上传和补量
	// 卡券开始后，不支持补量
	Count int `json:"count,omitempty"`
	// Key 生成文件的唯一标识（保留字段，暂时没有地方使用）
	Key string `json:"key,omitempty"`
	// Lines 文件解析得到的行数
	Lines int `json:"lines,omitempty"`
	// Mark 本次上传的标识
	Mark string `json:"mark,omitempty"`
	// URL 本次上传有效的券码文件地址
	URL string `json:"url,omitempty"`
}
