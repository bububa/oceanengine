package model

import "io"

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
	// Key field key
	Key string
	// Value field value
	Value string
	// Reader upload file reader
	Reader io.Reader
}

// UploadRequest multipart/form-data reqeust interface
type UploadRequest interface {
	// Encode encode request to UploadFields
	Encode() []UploadField
}
