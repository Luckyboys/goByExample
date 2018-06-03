package main

import (
	"encoding/base64"
	"fmt"
)

const OneRuneEqualsFourByte = 4

func main() {
	cipherText := encode("Hello World , 私のいとし人")
	fmt.Println("cipher text: ", cipherText)
	plainText, err := decode(cipherText)
	if err != nil {
		fmt.Println(nil)
	} else {
		fmt.Println("plaintext: ", plainText)
	}
}

/**
 * 传入明文，计算出密文返回
 * 对每个字都进行数学上的求补，用来混淆一定量的视觉识别
 * @param string plaintext 明文
 * @return string 密文
 */
func encode(plaintext string) (base64String string) {

	bCipherText := make([]byte, len([]byte(plaintext)))
	bPlainText := make([]byte, len([]byte(plaintext)))

	copy( bPlainText , []byte(plaintext) )
	for index, b := range []byte(plaintext) {

		bCipherText[index] = ^b
	}

	fmt.Println("encoded plain text(byte): ", bPlainText)
	fmt.Println("encoded cipher text(complement): ", bCipherText)

	base64String = base64.StdEncoding.EncodeToString(bCipherText)
	return
}

/**
 * 对密文重新再求一次补，从而获取回来明文
 * @param string cipherText 密文
 * @return string 明文
 * @return error 是否解密出错
 */
func decode(cipherText string) (string, error) {

	bCipherText, err := base64.StdEncoding.DecodeString(cipherText)

	fmt.Println("decoded bCipher text(bytes): ", bCipherText)
	if err != nil {
		return "", err
	}

	fmt.Println("decoded cipher text(complement): ", bCipherText)

	plaintext := make([]byte, len(bCipherText))

	for index, value := range bCipherText {
		plaintext[index] = ^value
	}

	fmt.Println("decoded plain text(byte): ", plaintext)

	return string(plaintext), nil
}
