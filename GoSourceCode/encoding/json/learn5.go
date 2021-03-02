package main

import (
	"bytes"
	"crypto/aes"
	"encoding/hex"
	"errors"
	"fmt"
)

func learn5() {
	aes := AesECB{key: []byte("84e6c6dc0f9p4a56"), blockSize: 16}
	src := "b07e9fcefd6210143d9975a0dcc87130e05dbc76ced18850edd8c0def82d8bf70b21b7c8c84826b1c4d1ad0dc80049ddb9f77ab76fcd04c65f91cf89a9406246dda53df34ad387850d3270392e02613869fb587c3032270deae21ff8d6930bbf"
	cipher, err := hex.DecodeString(src)
	if err != nil {
		fmt.Println("十六进制转换为ASCII字符串失败:" + err.Error())
	}
	plainText, err := aes.Decrypt([]byte(cipher))
	if err != nil {
		fmt.Println("解密失败:" + err.Error())
	} else {
		fmt.Println("解密的值为:" + string(plainText))
	}
}

type AesECB struct {
	key       []byte
	blockSize int
}

func (a *AesECB) Encrypt(src []byte, isPad ...bool) ([]byte, error) {
	block, err := aes.NewCipher(a.key)
	if err != nil {
		return nil, err
	}

	if len(src) == 0 {
		return nil, errors.New("content is empty")
	}

	if len(isPad) > 0 && isPad[0] == false {
		src = a.noPadding(src)
	} else {
		src = a.padding(src)
	}

	buf := make([]byte, a.blockSize)
	encrypted := make([]byte, 0)
	for i := 0; i < len(src); i += a.blockSize {
		block.Encrypt(buf, src[i:i+a.blockSize])
		encrypted = append(encrypted, buf...)
	}
	return encrypted, nil
}

func (a *AesECB) Decrypt(src []byte, isPad ...bool) ([]byte, error) {
	block, err := aes.NewCipher(a.key)
	if err != nil {
		return nil, err
	}
	if len(src) == 0 {
		return nil, errors.New("content is empty")
	}
	buf := make([]byte, a.blockSize)
	decrypted := make([]byte, 0)
	for i := 0; i < len(src); i += a.blockSize {
		block.Decrypt(buf, src[i:i+a.blockSize])
		decrypted = append(decrypted, buf...)
	}

	if len(isPad) > 0 && isPad[0] == false {
		decrypted = a.unNoPadding(decrypted)
	} else {
		decrypted = a.unPadding(decrypted)
	}

	return decrypted, nil
}

//nopadding模式
func (a *AesECB) noPadding(src []byte) []byte {
	count := a.blockSize - len(src)%a.blockSize
	if len(src)%a.blockSize == 0 {
		return src
	} else {
		return append(src, bytes.Repeat([]byte{byte(0)}, count)...)
	}
}

//nopadding模式
func (a *AesECB) unNoPadding(src []byte) []byte {
	for i := len(src) - 1; ; i-- {
		if src[i] != 0 {
			return src[:i+1]
		}
	}
	return nil
}

//padding模式
func (a *AesECB) padding(src []byte) []byte {
	count := a.blockSize - len(src)%a.blockSize
	padding := bytes.Repeat([]byte{byte(0)}, count)
	padding[count-1] = byte(count)
	return append(src, padding...)
}

//padding模式
func (a *AesECB) unPadding(src []byte) []byte {
	l := len(src)
	p := int(src[l-1])
	return src[:l-p]
}
