package Account

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"log"
)

const (
	KEY_SIZE = 512
)

type Keys struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

func (k Keys) GenKeys() (*rsa.PrivateKey, *rsa.PublicKey) {
	key, err := rsa.GenerateKey(rand.Reader, int(KEY_SIZE))
	if err != nil {
		return nil, nil
	}

	//log.Println(k.PrivateKey, k.PublicKey)
	return key, &key.PublicKey

}

func (k Keys) ToString() (string, string) {
	privStr := base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PrivateKey(k.PrivateKey))
	pubStr := base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PublicKey(k.PublicKey))
	fmt.Printf("Private key - %s \nPublic Key - %s\n", privStr, pubStr)
	return privStr, pubStr
}

func (k Keys) Sign(data string, key *rsa.PrivateKey) []byte {
	hash := crypto.SHA256.New()
	hash.Write([]byte(data))
	mesHash := hash.Sum(nil)

	sign, err := rsa.SignPSS(rand.Reader, key, crypto.SHA256, mesHash, nil)
	if err != nil {
		log.Println("error")
		log.Println(err)
		return nil
	}
	return sign
}

func Verify(pub *rsa.PublicKey, data string, sign []byte) bool {
	hash := crypto.SHA256.New()
	hash.Write([]byte(data))
	mesHash := hash.Sum(nil)

	err := rsa.VerifyPSS(pub, crypto.SHA256, mesHash, sign, nil)
	if err != nil {

		log.Println(err)
		return false
	}
	return true
}
