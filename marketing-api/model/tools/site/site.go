package site

import (
	"encoding/json"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
)

// SiteDetail 站点信息
type SiteDetail struct {
	// Bricks 具体见返回示例业务数据（新建或更新时传递的数据）
	Bricks []Brick `json:"bricks,omitempty"`
	// ID 站点ID
	ID model.Uint64 `json:"id,omitempty"`
	// Status 站点状态
	Status enum.SiteStatus `json:"status,omitempty"`
	// SiteType 建站类型
	SiteType enum.SiteType `json:"site_type,omitempty"`
	// Thumbnail 缩略图
	Thumbnail string `json:"thumbnail,omitempty"`
}

type tmpSiteDetail struct {
	// Bricks 具体见返回示例业务数据（新建或更新时传递的数据）
	Bricks []json.RawMessage `json:"bricks,omitempty"`
	// ID 站点ID
	ID model.Uint64 `json:"id,omitempty"`
	// Status 站点状态
	Status enum.SiteStatus `json:"status,omitempty"`
	// SiteType 建站类型
	SiteType enum.SiteType `json:"site_type,omitempty"`
	// Thumbnail 缩略图
	Thumbnail string `json:"thumbnail,omitempty"`
}

// Site 站点信息
type Site struct {
	// SiteID 站点ID
	SiteID model.Uint64 `json:"siteId,omitempty"`
	// Name 建站名称
	Name string `json:"name,omitempty"`
	// Status 站点状态
	Status enum.SiteStatus `json:"status,omitempty"`
	// SiteType 建站类型
	SiteType enum.SiteType `json:"siteType,omitempty"`
	// FunctionType 建站类别，SITE_FUNC_TYPE_NEW_MODULAR（智能建站）SITE_FUNC_TYPE_NORMAL （普通建站）
	FunctionType enum.SiteFunctionType `json:"function_type,omitempty"`
	// Thumbnail 缩略图
	Thumbnail string `json:"thumbnail,omitempty"`
}

// UnmarshalJSON implement json Unmarshal interface
func (s *SiteDetail) UnmarshalJSON(b []byte) (err error) {
	var tmp tmpSiteDetail
	if err = json.Unmarshal(b, &tmp); err != nil {
		return
	}
	detail := SiteDetail{
		ID:        tmp.ID,
		Status:    tmp.Status,
		SiteType:  tmp.SiteType,
		Thumbnail: tmp.Thumbnail,
	}
	for _, b := range tmp.Bricks {
		var base BaseBrick
		if err := json.Unmarshal(b, &base); err != nil {
			return err
		}
		switch base.Type() {
		case XrVideo:
			var data VideoBrick
			if err := json.Unmarshal(b, &data); err != nil {
				return err
			}
			detail.Bricks = append(detail.Bricks, data)
		case XrPicture:
			var data ImageBrick
			if err := json.Unmarshal(b, &data); err != nil {
				return err
			}
			detail.Bricks = append(detail.Bricks, data)
		case XrPictureGroup:
			var data ImagesBrick
			if err := json.Unmarshal(b, &data); err != nil {
				return err
			}
			detail.Bricks = append(detail.Bricks, data)
		case XrSimpleText:
			var data TextBrick
			if err := json.Unmarshal(b, &data); err != nil {
				return err
			}
			detail.Bricks = append(detail.Bricks, data)
		case XrRichText:
			var data RichTextBrick
			if err := json.Unmarshal(b, &data); err != nil {
				return err
			}
			detail.Bricks = append(detail.Bricks, data)
		case XrForm:
			var data FormBrick
			if err := json.Unmarshal(b, &data); err != nil {
				return err
			}
			detail.Bricks = append(detail.Bricks, data)
		case XrButton:
			var data ButtonBrick
			if err := json.Unmarshal(b, &data); err != nil {
				return err
			}
			detail.Bricks = append(detail.Bricks, data)
		case XrWechataApplet, XrWechatGame:
			var data WechatBrick
			if err := json.Unmarshal(b, &data); err != nil {
				return err
			}
			detail.Bricks = append(detail.Bricks, data)
		}
	}
	*s = detail
	return nil
}
