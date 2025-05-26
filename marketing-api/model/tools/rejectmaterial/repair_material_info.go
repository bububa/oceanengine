package rejectmaterial

  // RepairMaterialInfo 原素材及修复素材列表
type RepairMaterialInfo struct {
  // AIRepairID 素材修复id，唯一标识，可用于采纳接口
// id在修复时间+7天 是有效的，过期会失效
  AIRepairID uint64 `json:"ai_repair_id,omitempty"`
  // PromotionID 被拒审素材所属的广告id
  PromotionID uint64 `json:"promotion_id,omitempty"`
  // AdvertiserID 被拒审素材所属的广告账户id
  AdvertiserID uint64 `json:"advertiser_id,omitempty"`
  // MaterialID 被拒审原素材id
  MaterialID uint64 `json:"material_id,omitempty"`
  // AIRepairTime 素材修复时间，格式：yyyy-mm-dd hh:mm:ss
  AIRepairTime string `json:"ai_repair_time,omitempty"`
  // VideoID 被拒审原视频id
  VideoID string `json:"video_id,omitempty"`
  // PreviewURL 被拒审原视频素材url，有效期1小时
// 注意：存在同主体校验，要求开发者主体与广告主主体一致才可获取，否则广告主授权时需勾选敏感物料授权
  PreviewURL string `json:"preview_url,omitempty"`
  // AIRepairVideoID AI修复后视频id
  AIRepairVideoID string `json:"ai_repair_video_id,omitempty"`
  // AIRepairVideoCoverID AI修复后视频封面id
  AIRepairVideoCoverID string `json:"ai_repair_video_cover_id,omitempty"`
  // AIRepairVideoCoverURL AI修复后视频封面预览地址url，有效期1小时
// 注意：存在同主体校验，要求开发者主体与广告主主体一致才可获取，否则广告主授权时需勾选敏感物料授权
  AIRepairVideoCoverURL string `json:"ai_repair_video_cover_url,omitempty"`
  // AIRepairPreviewURL AI修复后视频预览url，有效期1小时
// 注意：存在同主体校验，要求开发者主体与广告主主体一致才可获取，否则广告主授权时需勾选敏感物料授权
  AIRepairPreviewURL string `json:"ai_repair_preview_url,omitempty"`
  // AIRepairContent 修复内容详情
  AIRepairContent []string `json:"ai_repair_content,omitempty"`
}
