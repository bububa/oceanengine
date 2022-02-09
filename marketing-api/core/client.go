package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/bububa/oceanengine/marketing-api/core/internal/debug"
	"github.com/bububa/oceanengine/marketing-api/model"
)

// SDKClient sdk client
type SDKClient struct {
	AppID   string
	Secret  string
	debug   bool
	sandbox bool
}

// NewSDKClient 创建SDKClient
func NewSDKClient(appID string, secret string) *SDKClient {
	return &SDKClient{
		AppID:  appID,
		Secret: secret,
	}
}

// SetDebug 设置debug模式
func (c *SDKClient) SetDebug(debug bool) {
	c.debug = debug
}

// UseSandbox 启用sandbox
func (c *SDKClient) UseSandbox() {
	c.sandbox = true
}

// DisableSandbox 禁用sandbox
func (c *SDKClient) DisableSandbox() {
	c.sandbox = false
}

// Post post api
func (c *SDKClient) Post(gw string, req model.PostRequest, resp model.Response, accessToken string) error {
	var reqBytes []byte
	if req != nil {
		reqBytes = req.Encode()
	}
	var builder strings.Builder
	builder.WriteString(BASE_URL)
	builder.WriteString(gw)
	reqUrl := builder.String()
	debug.PrintPostJSONRequest(reqUrl, reqBytes, c.debug)
	httpReq, err := http.NewRequest("POST", reqUrl, bytes.NewReader(reqBytes))
	if err != nil {
		return err
	}
	httpReq.Header.Add("Content-Type", "application/json")
	if accessToken != "" {
		httpReq.Header.Add("Access-Token", accessToken)
	}
	if c.sandbox {
		httpReq.Header.Add("X-Debug-Mode", "1")
	}
	return c.fetch(httpReq, resp)
}

// Get get api
func (c *SDKClient) Get(gw string, req model.GetRequest, resp model.Response, accessToken string) error {
	var builder strings.Builder
	builder.WriteString(BASE_URL)
	builder.WriteString(gw)
	if req != nil {
		builder.WriteString("?")
		builder.WriteString(req.Encode())
	}
	reqUrl := builder.String()
	debug.PrintGetRequest(reqUrl, c.debug)
	httpReq, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return err
	}
	if accessToken != "" {
		httpReq.Header.Add("Access-Token", accessToken)
	}
	if c.sandbox {
		httpReq.Header.Add("X-Debug-Mode", "1")
	}
	return c.fetch(httpReq, resp)
}

// GetWithData get api with json body
func (c *SDKClient) GetWithData(gw string, req model.GetWithDataRequest, resp model.Response, accessToken string) error {
	var data []byte
	if req != nil {
		data = req.Encode()
	}
	var builder strings.Builder
	builder.WriteString(BASE_URL)
	builder.WriteString(gw)
	reqUrl := builder.String()
	httpReq, err := http.NewRequest("GET", reqUrl, bytes.NewReader(data))
	if err != nil {
		return err
	}
	httpReq.Header.Add("Content-Type", "application/json")

	if accessToken != "" {
		httpReq.Header.Add("Access-Token", accessToken)
	}
	if c.sandbox {
		httpReq.Header.Add("X-Debug-Mode", "1")
	}

	debug.PrintJSONRequest("GET", reqUrl, httpReq.Header, data, c.debug)
	return c.fetch(httpReq, resp)
}

// Upload multipart/form-data post
func (c *SDKClient) Upload(gw string, req model.UploadRequest, resp model.Response, accessToken string) error {
	var buf bytes.Buffer
	mw := multipart.NewWriter(&buf)
	params := req.Encode()
	mp := make(map[string]string, len(params))
	var builder strings.Builder
	for _, v := range params {
		builder.Reset()
		var (
			fw  io.Writer
			r   io.Reader
			err error
		)
		if v.Reader != nil {
			if fw, err = mw.CreateFormFile(v.Key, v.Value); err != nil {
				return err
			}
			r = v.Reader
			builder.WriteString("@")
			builder.WriteString(v.Value)
			mp[v.Key] = builder.String()
		} else {
			if fw, err = mw.CreateFormField(v.Key); err != nil {
				return err
			}
			r = strings.NewReader(v.Value)
			mp[v.Key] = v.Value
		}
		if _, err = io.Copy(fw, r); err != nil {
			return err

		}
	}
	mw.Close()
	builder.Reset()
	builder.WriteString(BASE_URL)
	builder.WriteString(gw)
	reqUrl := builder.String()
	debug.PrintPostMultipartRequest(reqUrl, mp, c.debug)
	httpReq, err := http.NewRequest("POST", reqUrl, &buf)
	if err != nil {
		return err
	}
	httpReq.Header.Add("Content-Type", mw.FormDataContentType())
	if accessToken != "" {
		httpReq.Header.Add("Access-Token", accessToken)
	}
	if c.sandbox {
		httpReq.Header.Add("X-Debug-Mode", "1")
	}
	return c.fetch(httpReq, resp)
}

// AnalyticsPost 转化回传API专用
func (c *SDKClient) AnalyticsPost(gw string, req model.PostRequest, resp model.Response) error {
	reqBytes := req.Encode()
	reqBuf := bytes.NewBuffer(reqBytes)
	reqBuf.WriteString(c.Secret)
	sha256Sum := sha256.Sum256(reqBuf.Bytes())
	sign := hex.EncodeToString(sha256Sum[:])

	var builder strings.Builder
	builder.WriteString(ANALYTICS_URL)
	builder.WriteString(gw)
	reqUrl := builder.String()
	debug.PrintPostJSONRequest(reqUrl, reqBuf.Bytes(), c.debug)
	httpReq, err := http.NewRequest("POST", reqUrl, bytes.NewReader(reqBytes))
	if err != nil {
		return err
	}
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("x-signature", sign)
	if c.sandbox {
		httpReq.Header.Add("X-Debug-Mode", "1")
	}
	return c.fetch(httpReq, resp)
}

// fetch execute http request
func (c *SDKClient) fetch(httpReq *http.Request, resp model.Response) error {
	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()
	if resp == nil {
		resp = &model.BaseResponse{}
	}
	err = debug.DecodeJSONHttpResponse(httpResp.Body, resp, c.debug)
	if err != nil {
		debug.PrintError(err, c.debug)
		return err
	}
	if resp.IsError() {
		return resp
	}
	return nil
}
