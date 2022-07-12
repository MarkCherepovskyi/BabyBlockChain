package Blockchain

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"log"
)

const (
	KEY_SIZE = 512
)

type Keys struct {
	privateKey *rsa.PrivateKey
}

func (k *Keys) GetPrivate() *rsa.PrivateKey {
	return k.privateKey
}

func (k *Keys) GenKeys() error {
	key, err := rsa.GenerateKey(rand.Reader, int(KEY_SIZE))
	if err != nil {
		return err
	}
	k.privateKey = key

	return nil
}

func (k *Keys) GetPublicKey() *rsa.PublicKey {
	return &k.privateKey.PublicKey
}

func (k Keys) Sign(data string, key *rsa.PrivateKey) ([]byte, error) {
	hash := crypto.SHA256.New()
	hash.Write([]byte(data))
	mesHash := hash.Sum(nil)

	sign, err := rsa.SignPSS(rand.Reader, key, crypto.SHA256, mesHash, nil)
	if err != nil {
		log.Println("error")

		return nil, err
	}
	return sign, nil
}

func Verify(pub *rsa.PublicKey, data string, sign []byte) (bool, error) {
	hash := crypto.SHA256.New()
	hash.Write([]byte(data))
	mesHash := hash.Sum(nil)

	err := rsa.VerifyPSS(pub, crypto.SHA256, mesHash, sign, nil)
	if err != nil {

		return false, err
	}
	return true, err
}

func (k Keys) ToString() (string, string) {
	privStr := base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PrivateKey(k.GetPrivate()))
	pubStr := base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PublicKey(k.GetPublicKey()))

	return privStr, pubStr
}
