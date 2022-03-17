package datasource

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// UpdateRequest 数据源更新API Request
type UpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// DataSourceID 数据源ID
	DataSourceID string `json:"data_source_id,omitempty"`
	// OperationType 更新操作类型, 枚举值："1" ：添加,"2"：删除, "3"：重置
	OperationType enum.DmpDatasourceOperationType `json:"operation_type,omitempty"`
	// DataFormat 数据格式, 枚举值:"0"：ProtocolBuffer
	DataFormat int `json:"data_format"`
	// FileStorageType 数据存储类型, 枚举值:"0"：API
	FileStorageType int `json:"file_storage_type"`
	// FilePaths 通过【数据源文件上传】接口得到的文件路径，注意：一次上传最多1000个
	FilePaths []string `json:"file_paths,omitempty"`
}

// Encode implement PostRequest interface
func (r UpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
