package datasource

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CreateRequest 数据源创建API Request
type CreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// DataSourceName 数据源名称, 限30个字符内
	DataSourceName string `json:"data_source_name,omitempty"`
	// Description 数据源描述, 限256个字符内
	Description string `json:"description,omitempty"`
	// DataFormat 数据格式, 枚举值:"0"：ProtocolBuffer
	DataFormat int `json:"data_format"`
	// FileStorageType 数据存储类型, 枚举值:"0"：API
	FileStorageType int `json:"file_storage_type"`
	// FilePaths 通过【数据源文件上传】接口得到的文件路径，注意：一次上传最多1000个
	FilePaths []string `json:"file_paths,omitempty"`
	// DataSourceType 投放数据源类型，枚举值如下: "UID"：用户ID, "DID"：设备ID ,默认值： "UID"
	DataSourceType string `json:"data_source_type,omitempty"`
}

// Encode implement PostRequest interface
func (r CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CreateResponse 数据源创建API Response
type CreateResponse struct {
	model.BaseResponse
	// Data json返回值, 其中data_source_id参数包含唯一标识数据源的字符串
	Data *CreateResponseData `json:"data,omitempty"`
}

// CreateResponseData json返回值, 其中data_source_id参数包含唯一标识数据源的字符串
type CreateResponseData struct {
	// DataSourceID 数据源id, 唯一标识数据源的字符串
	DataSourceID string `json:"data_source_id,omitempty"`
}
