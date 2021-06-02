package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"github.com/sirupsen/logrus"
)

func main() {
	const publicKey = `
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCvjwlzaVLXQEn+U51rPRVTp7Eh
TnYRcv0Ju9zoRjzG2l0Gj7Es4AuVxPWuIAmL7PClfJ8wzqjpkIYabCBa+b5JSKq/
YDKO1vxI9991uX5SM4CjXDzcseoNJ9KeGRKkgqIDtHI1pjp27GwIm56uFXFjQE7E
fQJH5KWIQ1UzGXPGNQIDAQAB
-----END PUBLIC KEY-----
`
	key := BytesToPublicKey([]byte(publicKey))
	if key !=nil{

		logrus.WithField("size", key.Size()).Info("got key")
	}
}

// BytesToPublicKey bytes to public key
func BytesToPublicKey(pub []byte) *rsa.PublicKey {
	block, _ := pem.Decode(pub)
	var err error
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	if enc {
		logrus.Println("is encrypted pem block")
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			logrus.Error(err)
		}
	}
	ifc, err := x509.ParsePKIXPublicKey(b)
	if err != nil {
		logrus.Error(err)
	}
	key, ok := ifc.(*rsa.PublicKey)
	if !ok {
		logrus.Error("not ok")
	}
	return key
}
