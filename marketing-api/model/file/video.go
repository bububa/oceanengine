package file

import "github.com/bububa/oceanengine/marketing-api/enum"

// Video 视频
type Video struct {
	// ID 视频ID
	ID string `json:"id,omitempty"`
	// VideoID 视频ID
	VideoID string `json:"video_id,omitempty"`
	// VideoCoverURL 视频封面图片url
	VideoCoverURL string `json:"video_cover_url,omitempty"`
	// Size 视频大小
	Size uint64 `json:"size,omitempty"`
	// Width 视频宽度
	Width int `json:"width,omitempty"`
	// Height 视频高度
	Height int `json:"height,omitempty"`
	// URL 视频地址，仅限同主体进行素材预览查看，若非同主体会返回“素材所属主体与开发者主体不一致无法获取URL”
	// 链接仅做预览使用，预览链接有效期为1小时
	URL string `json:"url,omitempty"`
	// VideoURL 视频地址
	VideoURL string `json:"video_url,omitempty"`
	// Format 视频格式
	Format string `json:"format,omitempty"`
	// Signature 视频md5
	Signature string `json:"signature,omitempty"`
	// VideoSignature 视频md5
	VideoSignature string `json:"video_signature,omitempty"`
	// PosterURL 视频首帧截图，仅限同主体进行素材预览查看，若非同主体会返回“素材所属主体与开发者主体不一致无法获取URL”
	PosterURL string `json:"poster_url,omitempty"`
	// BitRate 码率，单位bps
	BitRate int64 `json:"bit_rate,omitempty"`
	// Duration 视频时长
	Duration float64 `json:"duration,omitempty"`
	// MaterialID 素材id，即多合一报表中的素材id，一个素材唯一对应一个素材id
	MaterialID uint64 `json:"material_id,omitempty"`
	// ImageMode 素材类型，枚举值:
	// LARGE 大图
	// SMALL 小图
	// LARGE_VERTICAL 大图竖图
	// UNION_SPLASH 穿山甲开屏图片
	// VIDEO_LARGE 横版视频
	// VIDEO_VERTICAL 竖版视频
	ImageMode enum.MaterialMode `json:"image_mode,omitempty"`
	// Source 素材来源，详见【附录-素材来源】
	Source enum.MaterialSource `json:"source,omitempty"`
	// CreateTime 素材的上传时间，格式："yyyy-mm-dd HH:MM:SS"
	CreateTime string `json:"create_time,omitempty"`
	// Filename 素材的文件名
	Filename string `json:"filename,omitempty"`
	// Tags 素材标签
	Tags []string `json:"tags,omitempty"`
	// Labels 视频标签
	Labels []string `json:"labels,omitempty"`
	// AwemeItemID 抖音短视频 ID
	AwemeItemID uint64 `json:"aweme_item_id,omitempty"`
	// Title 抖音中的视频标题
	Title string `json:"title,omitempty"`
	// IsRecommend 是否推荐; 0 不推荐 1 推荐
	IsRecommend int `json:"is_recommend,omitempty"`
	// ViewCnt 观看数
	ViewCnt int64 `json:"view_cnt,omitempty"`
	// LikeCnt 点赞数
	LikeCnt int64 `json:"like_cnt,omitempty"`
	// ShareCnt 分享数
	ShareCnt int64 `json:"share_cnt,omitempty"`
	// CommentCnt 评论数
	CommentCnt int64 `json:"comment_cnt,omitempty"`
	// StarAuthorID 星图达人ID
	StarAuthorID string `json:"star_author_id,omitempty"`
}
