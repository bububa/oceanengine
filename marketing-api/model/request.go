package model

import (
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// PostRequest post request interface
type PostRequest interface {
	// Encode encode request to bytes
	Encode() []byte
}

// GetRequest get request interface
type GetRequest interface {
	// Encode encode request to string
	Encode() string
}

// UploadField multipart/form-data post request field struct
type UploadField struct {
	// Reader upload file reader
	Reader io.Reader
	// Key field key
	Key string
	// Value field value
	Value string
}

// UploadRequest multipart/form-data reqeust interface
type UploadRequest interface {
	// Encode encode request to UploadFields
	Encode() []UploadField
}

// ConversionRequest
type ConversionRequest interface {
	PostRequest
	Sign(req *http.Request, content []byte) (string, error)
	GetAppAccessToken() string
}

// TrackRequest
type TrackRequest interface {
	ConversionRequest
	RequestURI() string
}

// CredentialSign implement ConvertionRequest interface
func CredentialSign(req *http.Request, content []byte, privateKey *rsa.PrivateKey, credential enum.Credential) (string, error) {
	// get stringToSign
	method := req.Method
	pathAndQuery := req.URL.Path
	// 在path末尾加上'/'
	if pathAndQuery == "" || pathAndQuery[len(pathAndQuery)-1] != '/' {
		pathAndQuery = pathAndQuery + "/"
	}
	if req.URL.RawQuery != "" {
		pathAndQuery = pathAndQuery + "?" + req.URL.RawQuery
	}
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	contentHash := getContentHashBase64(content)
	buf := util.GetBufferPool()
	buf.WriteString(strings.ToUpper(method))
	buf.WriteByte('\n')
	buf.WriteString(pathAndQuery)
	buf.WriteByte('\n')
	buf.WriteString(timestamp)
	buf.WriteByte('\n')
	buf.WriteString(contentHash)
	signBytes, err := util.SignWithPrivateKey(buf.Bytes(), privateKey)
	util.PutBufferPool(buf)
	if err != nil {
		return "", err
	}
	signature := base64.StdEncoding.EncodeToString(signBytes)
	builder := util.GetStringsBuilder()
	builder.WriteString("credential=")
	builder.WriteString(string(credential))
	builder.WriteString("&timestamp=")
	builder.WriteString(timestamp)
	builder.WriteString("&signature=")
	builder.WriteString(signature)
	token := builder.String()
	util.PutStringsBuilder(builder)
	return token, nil
}

func getContentHashBase64(content []byte) string {
	hasher := sha256.New()
	hasher.Write(content)
	return base64.StdEncoding.EncodeToString(hasher.Sum(nil))
}
