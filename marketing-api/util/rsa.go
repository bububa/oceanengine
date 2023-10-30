package util

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io"
)

// func init() {
// 	mrand.Seed(time.Now().UnixNano())
// }

// ReadPublicKeyFromPem read public key from pem file
func ReadPublicKeyFromPem(r io.Reader) (*rsa.PublicKey, error) {
	bs, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return ParsePublicKeyFromPem(bs)
}

// ParsePublicKeyFromPem read public key from pem file
func ParsePublicKeyFromPem(bs []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(bs)
	pubKeyI, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pubKey, ok := pubKeyI.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("not rsa public key")
	}
	return pubKey, nil
}

// ReadPrivateKeyFromPem read private key from pem file
func ReadPrivateKeyFromPem(r io.Reader) (*rsa.PrivateKey, error) {
	bs, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return ParsePrivateKeyFromPem(bs)
}

func ParsePrivateKeyFromPem(bs []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(bs)
	privKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return privKey, nil
}

// WritePrivateKeyPem write private key to pem file
func WritePrivateKeyPem(w io.Writer, privKey *rsa.PrivateKey) error {
	X509PrivateKey := x509.MarshalPKCS1PrivateKey(privKey)
	privateBlock := pem.Block{Type: "RSA Private Key", Bytes: X509PrivateKey}
	return pem.Encode(w, &privateBlock)
}

// WritePublicKeyPem write public key to pem file
func WritePublicKeyPem(w io.Writer, pubKey *rsa.PublicKey) error {
	X509PublicKey, err := x509.MarshalPKIXPublicKey(pubKey)
	if err != nil {
		return err
	}
	publicBlock := pem.Block{Type: "RSA Public Key", Bytes: X509PublicKey}
	return pem.Encode(w, &publicBlock)

}

// SignWithPrivateKey sign data with private key
func SignWithPrivateKey(msg []byte, privKey *rsa.PrivateKey) ([]byte, error) {
	hash := sha256.Sum256(msg)
	return rsa.SignPKCS1v15(rand.Reader, privKey, crypto.SHA256, hash[:])
}

// VerifyWithPublicKey sign data with private key
func VerifyWithPublicKey(msg []byte, signature []byte, pubKey *rsa.PublicKey) error {
	hash := sha256.Sum256(msg)
	return rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, hash[:], signature)
}

// EncryptWithPublicKey encrypts data with public key
func EncryptWithPublicKey(msg []byte, pub *rsa.PublicKey) ([]byte, error) {
	hash := sha256.New()
	return rsa.EncryptOAEP(hash, rand.Reader, pub, msg, nil)
}

// DecryptWithPrivateKey decrypts data with private key
func DecryptWithPrivateKey(ciphertext []byte, priv *rsa.PrivateKey) ([]byte, error) {
	hash := sha256.New()
	return rsa.DecryptOAEP(hash, rand.Reader, priv, ciphertext, nil)
}
