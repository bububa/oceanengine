package appmanagement

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AndroidBasicPackageGetRequest 查询安卓应用母包 API Request
type AndroidBasicPackageGetRequest struct {
	// AccountID 账户id指可以接的账号体系如广告主id、巨量纵横组织id等
	AccountID uint64 `json:"account_id,omitempty"`
	// AccountType 账户类型，允许值:
	// AD 广告主账户、BP 巨量纵横组织账号
	AccountType enum.AccountType `json:"account_type,omitempty"`
	// PackageID 修改的安卓母包id
	PackageID string `json:"package_id,omitempty"`
}

// Encode implement GetRequest interface
func (r AndroidBasicPackageGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("account_id", strconv.FormatUint(r.AccountID, 10))
	values.Set("account_type", string(r.AccountType))
	values.Set("package_id", r.PackageID)
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// AndroidBasicPackageGetResponse 查询安卓应用母包  API Response
type AndroidBasicPackageGetResponse struct {
	model.BaseResponse
	Data *AndroidBasicPackageGetResult `json:"data,omitempty"`
}

type AndroidBasicPackageGetResult struct {
	// CurrentVersion 当前线上版本信息
	CurrentVersion *AndroidBasicPackage `json:"current_version,omitempty"`
	// NextVersion 待发布版本信息
	NextVersion *AndroidAppListRequest `json:"next_version,omitempty"`
}
