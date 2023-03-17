package enum

// SendType 数据发送方式
type SendType string

const (
	// SendType_SERVER_SEND 服务器端上传
	SendType_SERVER_SEND SendType = "SERVER_SEND"
	// SendType_CLIENT_SEND 客户端上传
	SendType_CLIENT_SEND SendType = "CLIENT_SEND"
)
