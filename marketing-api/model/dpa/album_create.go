package dpa

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AlbumCreateRequest 上传短剧剧目 API Request
type AlbumCreateRequest struct {
	// AlbumTitle 短剧标题，请注意命名符合 https://www.volcengine.com/docs/4/1200126 规则
	AlbumTitle string `json:"album_title,omitempty"`
	// SeqNum 总集数
	SeqNum int `json:"seq_num,omitempty"`
	// VideoMetaMico 短剧信息
	VideoMetaMico []VideoMetaMico `json:"video_meta_mico,omitempty"`
	// AppID 请传入app_access_token对应的app_id，可在开放平台官网-开发者后台查询
	AppID uint64 `json:"app_id,omitempty"`
}

// VideoMetaMico 短剧信息
type VideoMetaMico struct {
	// URL 视频url，需要带有文件扩展名， 示例：https://xxxx.com/
	// xxxxx.mp4，目前视频只支持mp4 avi mov m3u8格式
	URL string `json:"url,omitempty"`
	// Title 集名称
	Title string `json:"title,omitempty"`
	// Seq 集序号
	Seq int `json:"seq,omitempty"`
}

// Encode implements PostRequest interface
func (r AlbumCreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// AlbumCreateResponse 上传短剧剧目 API Response
type AlbumCreateResponse struct {
	model.BaseResponse
	Data struct {
		// AlbumID 短剧ID
		AlbumID string `json:"album_id,omitempty"`
	} `json:"data,omitempty"`
}
