package appmanagement

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AndroidBasicPackagePublishRequest 发布安卓应用母包 API Request
type AndroidBasicPackagePublishRequest struct {
	// AccountID 账户id指可以接的账号体系如广告主id、巨量纵横组织id等
	AccountID uint64 `json:"account_id,omitempty"`
	// AccountType 账号类型，允许值：
	// AD 广告主类型、BP 巨量纵横账号类型
	AccountType enum.AccountType `json:"account_type,omitempty"`
	// PackageID 应用包id
	PackageID string `json:"package_id,omitempty"`
}

// Encode implement PostRequest interface
func (r AndroidBasicPackagePublishRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
