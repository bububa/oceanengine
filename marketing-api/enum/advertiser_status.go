package enum

// 广告主状态
type AdvertiserStatus string

const (
	STATUS_DISABLE                AdvertiserStatus = "STATUS_DISABLE"                // 已禁用
	STATUS_PENDING_CONFIRM        AdvertiserStatus = "STATUS_PENDING_CONFIRM"        // 申请待审核
	STATUS_PENDING_VERIFIED       AdvertiserStatus = "STATUS_PENDING_VERIFIED"       // 待验证
	STATUS_CONFIRM_FAIL           AdvertiserStatus = "STATUS_CONFIRM_FAIL"           // 审核不通过
	STATUS_ENABLE                 AdvertiserStatus = "STATUS_ENABLE"                 // 已审核
	STATUS_CONFIRM_FAIL_END       AdvertiserStatus = "STATUS_CONFIRM_FAIL_END"       // CRM审核不通过
	STATUS_PENDING_CONFIRM_MODIFY AdvertiserStatus = "STATUS_PENDING_CONFIRM_MODIFY" // 修改待审核
	STATUS_CONFIRM_MODIFY_FAIL    AdvertiserStatus = "STATUS_CONFIRM_MODIFY_FAIL"    // 修改审核不通过
	STATUS_LIMIT                  AdvertiserStatus = "STATUS_LIMIT"                  // 限制
	STATUS_WAIT_FOR_BPM_AUDIT     AdvertiserStatus = "STATUS_WAIT_FOR_BPM_AUDIT"     // 等待CRM审核
	STATUS_WAIT_FOR_PUBLIC_AUTH   AdvertiserStatus = "STATUS_WAIT_FOR_PUBLIC_AUTH"   // 待对公验证
	STATUS_SELF_SERVICE_UNAUDITED AdvertiserStatus = "STATUS_SELF_SERVICE_UNAUDITED" // 自助开户待验证资质
)
