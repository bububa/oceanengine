package model

type PostRequest interface {
	Encode() []byte
}

type GetRequest interface {
	Encode() string
}
