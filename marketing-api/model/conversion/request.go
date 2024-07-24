package conversion

import (
	"crypto/rsa"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
)

// Request 转化回传参数
type Request struct {
	// EventType 回传的事件，例如”激活“、”付费“。详细列表见附录
	EventType enum.ConversionEventType `json:"event_type,omitempty"`
	// EventWeight 广告主可以针对每一次转化，回传一个这次转化的价值，单位是 RMB。
	EventWeight float64 `json:"event_weight,omitempty"`
	// AttributeLabel 当前场景只支持convert，勿传其他值
	AttributeLabel string `json:"attribute_label,omitempty"`
	// BizType 当前场景只支持4，，勿传其他值
	BizType int `json:"biz_type,omitempty"`
	// Context 包含一些关键的上下文信息
	Context *Context `json:"context,omitempty"`
	// Properties 对于上报事件的附加属性，详情请见自定义属性的介绍。
	Properties *Properties `json:"properties,omitempty"`
	// Timestamp 事件发生的毫秒级时间戳
	Timestamp int64 `json:"timestamp,omitempty"`
	// Source 广告主标识，自定义值，用于标识数据来源，例如：jd
	Source string `json:"source,omitempty"`
	// PrivateKey
	PrivateKey *rsa.PrivateKey `json:"-"`
	// Credential
	Credential enum.Credential `json:"-"`
	// AppAccessToken
	AppAccessToken string `json:"-"`
}

// Context 包含一些关键的上下文信
type Context struct {
	Ad     *ContextAd     `json:"ad,omitempty"`     // 包含一些关键的广告相关信息
	Device *ContextDevice `json:"device,omitempty"` // 传递一些归因的设备信息
	App    *ContextApp    `json:"app,omitempty"`
}

// ContextAd 广告相关信息
type ContextAd struct {
	// Callback callback 字段有两个获取途径，对于监测链接归因的方式，需要从监测链接的__CALLBACK_PARAM__这个宏获取这个字段值；对于落地页或小程序归因的方式，需要从 url 中的 clickid 参数获取值。
	Callback string `json:"callback,omitempty"`
	// SkuID skuid
	SkuID string `json:"skuid,omitempty"`
	// MatchType 这个参数可以帮助广告主标识是通过哪种渠道进行的归因，如果不传，就会默认为点击归因。
	MatchType enum.ConversionAdMatchType `json:"match_type,omitempty"`
	// CustomerID 数据上报白名单customerid；每次上报必须需要带上已配置白名单的customerid，否则会上报失败；
	// 注意，此字段必填，且仅用于上报权限验证。
	CustomerID uint64 `json:"customer_id,omitempty"`
	// Attributed true: 客户已归因，上报的事件为转化事件
	// false: 客户未归因，上报为需要巨量引擎归因的事件
	Attributed bool `json:"attributed"`
	// AdID 广告计划id；若您可以确定当前订单来自对应广告计划id，则可以上报此字段获得更加精准的归因结果
	AdID uint64 `json:"ad_id,omitempty"`
	// CampaignID 广告组id；同上，若您可以确定当前订单来自对应广告组id，则可以上报此字段获得更加精准的归因结果
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// AdvertiserIDs 广告账户id；同上，若您可以确定当前订单来自对应广告账户id，则可以上报此字段获得更加精准的归因结果
	AdvertiserIDs []uint64 `json:"advertiser_ids,omitempty"`
	// CustomerIDs 客户id；同上，若您可以确定当前订单来自对应客户id，则可以上报此字段获得更加精准的归因结果
	// 若您不确定订单来自哪个账户，请务必在此字段中填充 所有投放下单的 customer_id，字节侧将为您进行精准归因
	CustomerIDs []uint64 `json:"customer_ids,omitempty"`
	// ClickTime 点击时间，指用户点击页面跳转时间，上报秒级时间戳
	ClickTime int64 `json:"click_time,omitempty"`
}

// ContextDevice 归因设备信息
type ContextDevice struct {
	// Imei 归因上的设备 imei 的 MD5 值
	Imei string `json:"imei,omitempty"`
	// Idfa 归因上的设备的 idfa 的原值
	Idfa string `json:"idfa,omitempty"`
	// Oaid 归因上的设备的 oaid 的原值
	Oaid string `json:"oaid,omitempty"`
	// Gaid
	Gaid string `json:"gaid,omitempty"`
	// AndroidId
	AndroidId string `json:"android_id,omitempty"`
	// AndroidIdMd5
	AndroidIdMd5 string `json:"android_id_md5,omitempty"`
	// PhoneNumBlurred 下单用户的模糊手机号，目前支持以下3种类型：
	// 1. （新）仅后四位：例如*******1234，前七位需要用星号表示；当上传此手机号格式时，receiver_province、receiver_city必填，否则无法上报和正确归因
	// 2. 省略中间四位：例如130****1234，中间四位需用星号表示
	// 3. 原始手机号sha256后的结果，64位字符串
	// 【注意】手机号的加密步骤仅在能获取明文手机号情况下，使用sha256加密，其他两种手机号形式切勿加密！否则会导致归因为0
	PhoneNumBlurred string `json:"phone_num_blurred,omitempty"`
}

// ContextApp 归因应用信息
type ContextApp struct {
	// PackageName 应用包名
	PackageName string `json:"package_name,omitempty"`
}

// Properties 附加属性
type Properties struct {
	// PayAmount 支付金额, 单位分
	PayAmount float64 `json:"pay_amount,omitempty"`
	// BillTrackType 支付方式
	BillTrackType enum.ConversionBillTrackType `json:"bill_track_type,omitempty"`
	// AggCost 聚合页消耗金额总和
	AggCost int64 `json:"agg_cost,omitempty"`
	// ClassPurchaseAmount 正价课购买次数;背景：教育行业深度转化，优化用户的正价课购买行为；其中有一些用户会一次性产生多科的正价课购买，希望该行为能够表达给深度转化模型优化；释义：用户产生一科正价课购买行为，则属性为1；用户产生3科正价课购买行为，则属性为3；
	ClassPurchaseAmount int `json:"class_purchase_amount,omitempty"`
	// PeopleQuality 差异价值人群质量;金融行业授信事件差异价值获取属性1:高质量;2:中质量;3:低质量
	PeopleQuality int `json:"people_quality,omitempty"`
	// ValueType 用户价值类型;premium_payment(保险支付)事件下的属性，表示付款用户的用户价值1：低净值2：高净值
	ValueType int `json:"value_type,omitempty"`
	// OrderAmount 用户在电商平台APP内，提交订单的金额
	OrderAmount float64 `json:"order_amount,omitempty"`
	// ImServiceType 智能客服服务方; 腾讯微信智能客服：wechatim，巨量引擎小六客服：xiaoliuim，抖音私信douyinim，其他待补充腾讯微信智能客服：wechatim，巨量引擎小六客服：xiaoliuim，抖音私信douyinim，其他待补充
	ImServiceType string `json:"im_service_type,omitempty"`
	// ReplyMode 智能客服回复方式; auto是系统自动回复，staff是人工客服回复
	ReplyMode string `json:"reply_mode,omitempty"`
	// IfHello 智能客服是否发送欢迎语; 1已发送，0已调起客服页面但未触发欢迎语
	IfHello *int `json:"if_hello,omitempty"`
	// DialogMode 智能客服对话输入方式; input是用户主动输入,pick是用户点击菜单输入
	DialogMode string `json:"dialog_mode,omitempty"`
	// ConsumptionDetail 记录LU场景下广告主消耗的ctr/cvr/bid等
	ConsumptionDetail string `json:"consumption_detail,omitempty"`
	// Consumption2 LU搜索侧广告主总消耗金额; LU广告搜索侧广告主的总消耗金额
	Consumption2 int64 `json:"consumption2,omitempty"`
	// OriginalEvent 客户自定义的原始事件
	OriginalEvent string `json:"original_event,omitempty"`
	// OriginalClassPrice 教育行业客户后链路转化正价课在客户侧的对应课程价格
	OriginalClassPrice float64 `json:"original_class_price,omitempty"`
	// OriginalClassID 用户转化的正价课id; 教育行业客户后链路转化正价课在客户侧的对应课程id
	OriginalClassID string `json:"original_class_id,omitempty"`
	// M2Surrender M2退保率预估;广告主对保险用户的M2退保率的预估
	M2Surrender float64 `json:"m2_surrender,omitempty"`
	// SkuCount 商品数量; 商品数量
	SkuCount int `json:"sku_count,omitempty"`
	// SkuAmount 商品单价,单位分
	SkuAmount int64 `json:"sku_amount,omitempty"`
	// SkuID 团购商品id
	SkuID string `json:"sku_id,omitempty"`
	// AnchorID 主播id
	AnchorID string `json:"anchor_id,omitempty"`
	// EnterMethod 直播间来源二级目录
	EnterMethod string `json:"enter_method,omitempty"`
	// EnterFromMerge 直播间来源页面
	EnterFromMerge string `json:"enter_from_merge,omitempty"`
	// ProductID 商品id，按照淘宝的sku_id进行回传。若订单中包含多个sku，则可在此字段中填写全部sku_id，英文逗号分隔
	ProductID string `json:"product_id,omitempty"`
	// ProductName 商品名
	ProductName string `json:"product_name,omitempty"`
	// ProductImag 商品图片
	ProductImag string `json:"product_imag,omitempty"`
	// ProductNumber 商品数量
	ProductNumber int `json:"product_number,omitempty"`
	// ProductTitle 商品标题
	ProductTitle string `json:"product_title,omitempty"`
	// ProductPrice 商品单价，单位为“分”, 商品加个范围
	ProductPrice string `json:"product_price,omitempty"`
	// ProductCategory 商品类目
	ProductCategory string `json:"product_category,omitempty"`
	// ProductPicUrl 商品大图URL
	ProductPicUrl string `json:"product_pic_url,omitempty"`
	// ShopName 店铺名称
	ShopName string `json:"shop_name,omitempty"`
	// ProductType 商品类型
	ProductType string `json:"product_type,omitempty"`
	// M2Score M2质量分; 自有建模预估能力的客户对保险用户M2后是否续保的打分
	M2Score float64 `json:"m2_score,omitempty"`
	// Consumption LU搜索侧广告主消耗金额; LU广告搜索侧广告主的消耗金额
	Consumption int64 `json:"consumption,omitempty"`
	// LengthOfStudy 听课时长; 标示这个转化事件发生时用户传了; 分钟
	LengthOfStudy int `json:"length_of_study,omitempty"`
	// Homework 标示这个转化事件发生时用户是否交当堂课作业，0-未交作业，1-交了作业
	Homework *int `json:"homework,omitempty"`
	// ClassNum 当前课程节数;标志这个转化事件对应的是教育客户低价课的哪一节
	ClassNum int `json:"class_num,omitempty"`
	// FinalClass 是否为末课; 标示这个转化事件对应的是否为教育客户低价课的最后一节，0-不为末课，1-为末课。
	FinalClass *int `json:"final_class,omitempty"`
	// OriginClassGrade 用户转化的年级; 教育行业客户转化课程对应的年级
	OriginClassGrade int `json:"origin_class_grade,omitempty"`
	// NovalAcLastActiveTs 用户上次在该频道活跃时间戳; 用户上次在该频道活跃时间戳
	NovalAcLastActiveTs int64 `json:"novel_ac_last_active_ts,omitempty"`
	// NovalAcCategory 小说频道类型; 用户活跃频道的具体细分
	NovalAcCategory string `json:"noval_ac_category,omitempty"`
	// NovalAcGender 小说频道性别类型; 该频道为男频还是女频
	NovalAcGender *int `json:"noval_ac_gender,omitempty"`
	// NovalBsAddTs 加入书架时间戳; 加入书架时间戳
	NovalBsAddTs int64 `json:"noval_bs_add_ts,omitempty"`
	// NovalBsAuthor 加入书架小说的作者
	NovalBsAuthor string `json:"noval_bs_author,omitempty"`
	// NovalBsCategory 加入书架的小说类别
	NovalBsCategory string `json:"noval_bs_category,omitempty"`
	// NovalBsTitle 加入书架的小说名
	NovalBsTitle string `json:"noval_bs_title,omitempty"`
	// NovalRcLastReadTs 上次阅读时间戳
	NovalRcLastReadTs int64 `json:"novel_rc_last_read_ts,omitempty"`
	// NovalRcPaychapter 小说起始付费章节
	NovalRcPaychapter int `json:"novel_rc_paychapter,omitempty"`
	// NovalRcDuration 用户在该小说总阅读时长
	NovalRcDuration int64 `json:"novel_rc_duration,omitempty"`
	// NovalRcChapter 用户阅读到的小说章节
	NovalRcChapter int `json:"noval_rc_chapter,omitempty"`
	// NovalRcAuthor 用户阅读的小说作者
	NovalRcAuthor string `json:"noval_rc_author,omitempty"`
	// NovalRcCategory 用户阅读的小说的类别
	NovalRcCategory string `json:"noval_rc_category,omitempty"`
	// NovalRcTitle 用户阅读的小说名
	NovalRcTitle string `json:"noval_rc_title,omitempty"`
	// NovalType 小说数据行为类型; 回传的是该用户的哪种行为信息
	NovalType int `json:"noval_type,omitempty"`
	// LoanValue 贷款价值，互金行业
	LoanValue float64 `json:"loan_value,omitempty"`
	// OrderType 订单类型，旅游等行业
	OrderType int `json:"order_type,omitempty"`
	// OrderCnt 订单数量，旅游等行业
	OrderCnt int `json:"order_cnt,omitempty"`
	// ActiveLevel 激活质量，各行业通用
	ActiveLevel float64 `json:"active_level,omitempty"`
	// ReadingTime 阅读时长，小说行业
	ReadingTime int64 `json:"reading_time,omitempty"`
	// FromChannel 贷款场景; 互金行业
	FromChannel int `json:"from_channel,omitempty"`
	// KeyLevel 关键行为水平
	KeyLevel float64 `json:"key_level,omitempty"`
	// PayType 支付方式
	PayType int `json:"pay_type,omitempty"`
	// UserValue 用户价值，1-10整数值
	UserValue int `json:"user_value,omitempty"`
	// Gender 用户性别
	Gender int `json:"gender,omitempty"`
	// OrderID 订单id；建议按照淘宝订单中的父订单（tid）进行回传；系统会根据订单id辅助去重提升归因准确性
	OrderID string `json:"order_id,omitempty"`
	// OrderState 订单状态
	OrderState string `json:"order_state,omitempty"`
	// ReceiverProvince 收货人所在的省份；当phone_num_blurred值为*******1234类型时必填
	ReceiverProvince string `json:"receiver_province,omitempty"`
	// ReceiverCity 收货人所在的城市（若城市为直辖市仍然填写市，如北京市）；当phone_num_blurred值为*******1234类型时必填
	ReceiverCity string `json:"receiver_city,omitempty"`
	// AdzoneID 推广位id，即通过新淘客创建追踪链接时，选择的推广位管理下的推广位名称对应的ID，同时也是pid=mm_1_2_3中的“3”这段数字
	AdzoneID uint64 `json:"adzone_id,omitempty"`
	// TecAgent 数据回传身份, 回传公司主体名称
	TecAgent string `json:"tec_agent,omitempty"`
	// EcomPlatform 来源平台（PDD TB JD 微信小程序）
	EcomPlatform string `json:"ecom_platform,omitempty"`
	// UgEventVal 穿山甲快应用LTV回传，实时且连续地回传每个新用户调起后24小时内的全渠道变现ltv累计值 单位：分；整数
	UgEventVal uint64 `json:"ug_event_val,omitempty"`
}

// Encode implement GetRequest interface
func (r Request) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}

// Sign implement ConvertionRequest interface
func (r Request) Sign(req *http.Request, content []byte) (string, error) {
	if r.PrivateKey == nil || r.Credential == "" {
		return "", errors.New("no private_key/credential")
	}
	return model.CredentialSign(req, content, r.PrivateKey, r.Credential)
}

// GetAccessToken implement ConvertionRequest interface
func (r Request) GetAppAccessToken() string {
	return r.AppAccessToken
}
