package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"

	"github.com/sirupsen/logrus"
)

func main() {
	const publicKeyPem = `
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxgLDSzaeSOgAwZRTuq0v
+VMlOP09LZ8wy3NGJhgBjtVo2LZxHoWPfUNezntK3KOqZRk62VJdwcdaBkPlAOYa
d63Tew+uQLDEJnmRVeXLb/VPgg3cMpd7IFhdSBEEPxpj8dm6iZ1yDay136MorVe0
WwKcVcW4FR1O2Jf34+J2zAKAQ1LpUDQ8o9c4vVDga59BnZFdehso2HEazKph6dT/
keVoVTF6k2nY/xJVnxih5CgXGa4WD1iEBtPWi2WY+ClnDANNrG8TzVpVtiXIJLeb
bpHRA0QVuXr2Bmdoz722udX17wviFmUzsCB0nRm8r522VT4hPj5E4ScYY3u0fMQt
EQIDAQAB
-----END PUBLIC KEY-----
`
	const privateKeyPem = `
-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEAxgLDSzaeSOgAwZRTuq0v+VMlOP09LZ8wy3NGJhgBjtVo2LZx
HoWPfUNezntK3KOqZRk62VJdwcdaBkPlAOYad63Tew+uQLDEJnmRVeXLb/VPgg3c
Mpd7IFhdSBEEPxpj8dm6iZ1yDay136MorVe0WwKcVcW4FR1O2Jf34+J2zAKAQ1Lp
UDQ8o9c4vVDga59BnZFdehso2HEazKph6dT/keVoVTF6k2nY/xJVnxih5CgXGa4W
D1iEBtPWi2WY+ClnDANNrG8TzVpVtiXIJLebbpHRA0QVuXr2Bmdoz722udX17wvi
FmUzsCB0nRm8r522VT4hPj5E4ScYY3u0fMQtEQIDAQABAoIBAQC3FQ7lmFa6Kmmo
kzDnxgI5kbs0+lF6wVoogyk1R7UJECGlXuswwWVu+L/MQwtw/pMqby2Oiy3NYjnG
2TgAoMGQLtT/T/xaLBbxTC4pMhPsLPGa2p3i2VUrDwkQ/Urj3c7Ta/GiFEXEhxCS
+xU9UX6dN8i9NdeDDlo8wiaeOmKFNZog1nUqmhIUJplCoE2yLsQaKo/5QBm8Kryu
bEDbMue3R9bD+uNySdDUxqEaABrPbGNzA0zQYR8FB32kZOOiK0DLlE4QCbWw0xQG
lgRpTIThiVkUEiYJGxX3ldtsrWBD6DSP60sZebJtl35Kitv6hIDMO3HK4LoTygyS
ZusEUhBRAoGBAOoNuXJuUSRnCE3oZj3BC3uW4kCfKMg/2BEr3k4fhbn/7Pn5251k
XHte7eG+/MinfevTdvVmLxsl/JCHZsx+7AxfJCkFBFU5g18HTbZhsSb5GhSW6rzr
kKFdb39QTWEaGMay+xwFe/J+RfzbCwe1iFaSCQ/tMjxYz8cOBqSbapJnAoGBANiT
28iiYnVF4UwYlvN/7DZU5bh6Jm4QNJ8PMXB7h0drehHwH1V/kdbUaaUzS+fFnbPW
B4qRI+/frmJOAFpxxmy+uqKDy5WNqxuMlN34SUwhLxgK5ZkbgorK4pK2NFb4cFf6
JNYzf+htpm/XPmJ87czKXxHL/fgwcypYq5JTS0nHAoGBAOfxZd+it0Kt4Crkrb1w
+yIk6l5D27RmoIaTxKAC2RD0MDGhjCOze3q/HlpiXDu00DLNwst50iDDSkEoDZWG
CgJQnpMP57VVY3zezgJ5WYdXwKK61DEX+KIQ6Bdwyt35cKgoYnTocmZjAJSNtaaU
HWI9IkHoO6pWEMIggjDJ25jtAoGAWh8U2uUvgUtF2At3bPXzThx/xwcVTMFzlCiX
Uw4kGcUQc7xu3X4LKtES8QpqcJSb7gSG+NMymBOFFiUNIpWrkv6+3Pdo1N+xrvIN
1FwH7CT22tNV/SMkbCrGo3QsTXBuFz3g7D1b4VUXZ/yyWftQXOJtMxOmqAtJMPra
0jXU42sCgYEAt8WBLiN/eMWNHjekCfxmZ6lfxxdQu9oOy2rNYPcRr/5LTRai08ps
dEWbvz7IIfL8Dmlxy5iNEWdi7cBHsmlaE6JX/z6Qwx0vbd/6lRTNvwNHRSnYOGkF
g5vfYJJUUJt8HVcntSHVfNewmMa1JFw2ZThjgTDhEWR415nOcdClRWM=
-----END RSA PRIVATE KEY-----`

	const plaintext = "1234567890abcdef"

	publicKey := BytesToPublicKey([]byte(publicKeyPem))
	if publicKey == nil {

		logrus.Error("public key is nil")
		return
	}
	logrus.WithField("size", publicKey.Size()).Info("got key")

	ciphertext := EncryptWithPublicKey([]byte(plaintext), publicKey)
	if ciphertext == nil {

		return
	}

	ciphertextBase64 := base64.StdEncoding.EncodeToString(ciphertext)

	logrus.WithField("ciphertext", ciphertextBase64).Info("encrypted")

	key := BytesToPrivateKey([]byte(privateKeyPem))

	ciphertext, err := base64.StdEncoding.DecodeString(ciphertextBase64)
	if err != nil {
		logrus.WithError(err).Error("decode base64 error")
		return
	}

	logrus.WithField("plaintext", string(DecryptWithPrivateKey(ciphertext, key))).Info("decrypted")
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

// BytesToPrivateKey bytes to private key
func BytesToPrivateKey(privateKey []byte) *rsa.PrivateKey {
	block, _ := pem.Decode(privateKey)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	var err error
	if enc {
		logrus.Println("is encrypted pem block")
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			logrus.Error(err)
		}
	}
	key, err := x509.ParsePKCS1PrivateKey(b)
	if err != nil {
		logrus.Error(err)
	}
	return key
}

// EncryptWithPublicKey encrypts data with public key
func EncryptWithPublicKey(msg []byte, pub *rsa.PublicKey) []byte {
	hash := sha512.New()
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, pub, msg, nil)
	if err != nil {
		logrus.Error(err)
	}
	return ciphertext
}

// DecryptWithPrivateKey decrypts data with private key
func DecryptWithPrivateKey(ciphertext []byte, privateKey *rsa.PrivateKey) []byte {
	hash := sha512.New()
	plaintext, err := rsa.DecryptOAEP(hash, rand.Reader, privateKey, ciphertext, nil)
	if err != nil {
		logrus.Error(err)
	}
	return plaintext
}
