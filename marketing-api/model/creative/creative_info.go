package creative

import "github.com/bububa/oceanengine/marketing-api/enum"

// CreativeInfo 创意信息
type CreativeInfo struct {
	// ImageMode 素材类型
	ImageMode enum.ImageMode `json:"image_mode,omitempty"`
	// CreativeID 创意id，填写id为修改创意，不填为新增创意，需要注意该接口为全量接口，不填写id的已有创意会被覆盖
	CreativeID uint64 `json:"creative_id,omitempty"`
	// TitleMaterial 标题素材
	TitleMaterial *TitleMaterial `json:"title_material,omitempty"`
	// TitleMaterials 创意标题素材，最多支持10个标题，不使用程序化创意包时必填
	TitleMaterials []TitleMaterial `json:"title_materials,omitempty"`
	// ImageMaterials 创意图片素材，组图、橱窗类型传3张，其他图片类型传1张，使用DPA模板时template_image只能传入1个，image_mode为图片素材时使用
	ImageMaterials []ImageMaterial `json:"image_materials,omitempty"`
	// VideoMaterials 视频素材信息，image_mode为视频素材时使用
	VideoMaterials []VideoMaterial `json:"video_materials,omitempty"`
	// VideoMaterial 视频素材信息，image_mode为视频素材时使用
	VideoMaterial *VideoMaterial `json:"video_material,omitempty"`
	// SubTitleMaterial 副标题素材
	SubTitleMaterial *TitleMaterial `json:"sub_title_material,omitempty"`
	// PlayableMaterial 基础试玩素材，image_mode为基础试玩素材时使用，不支持和image_materials同时传入
	PlayableMaterial *PlayableMaterial `json:"playable_material,omitempty"`
	// InteractiveMaterial 直出互动素材信息
	InteractiveMaterial *InteractiveMaterial `json:"interactive_material,omitempty"`
	// CompnentMaterials 组件信息，数组传入最大长度2，每种类型（基础/附加组件）最多支持1个; 附加组件：选择磁贴、投票磁贴、图片磁贴、电商磁贴、优惠券磁贴、游戏礼包码; 基础组件：推广卡片
	ComponentMaterials []ComponentMaterial `json:"component_materials,omitempty"`
	// AbstractMaterials 摘要素材，使用标签摘要时必须要传入3个;使用搜索广告时必传
	AbstractMaterials []AbstractMaterial `json:"abstract_materials,omitempty"`
	// DecorationMaterial 家装卡券素材
	DecorationMaterial *DecorationMaterial `json:"decoration_material,omitempty"`
	// DerivePosterCid 是否将视频的封面和标题同步到图片创意。视频素材可填写。在创建完视频创意后，在计划下同步新建图片创意，允许值：0（不开启），1（开启）
	DerivePosterCid *int `json:"derive_poster_cid,omitempty"`
	// ThirdPartyID 创意自定义参数，例如开发者可设定此参数为创意打标签，用于区分使用的素材
	ThirdPartyID string `json:"third_party_id,omitempty"`
}
