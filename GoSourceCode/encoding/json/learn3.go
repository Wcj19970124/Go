package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

//golang的rsa加解密
func learn3() {
	plainText := "hello world!"
	priFile := "./private.pem"
	pubFile := "./public.pem"

	cipherText := RSAEncrypt([]byte(plainText), pubFile) //加密
	fmt.Println(string(cipherText))

	text := RSADecrypt(cipherText, priFile)
	fmt.Println(string(text))
}

//RSAEncrypt 加密
func RSAEncrypt(plainText []byte, path string) []byte {

	//读取文件内容即公钥
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("读取公钥数据失败:" + err.Error())
	}

	//pem解码
	block, _ := pem.Decode(data)

	//x509解码
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Println("x509解码失败:" + err.Error())
	}

	publicKey := publicKeyInterface.(*rsa.PublicKey) //类型断言

	//加密
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText)
	if err != nil {
		fmt.Println("加密失败:" + err.Error())
	}

	return cipherText
}

//RSADecrypt 解密
func RSADecrypt(cipherText []byte, path string) []byte {

	//读取私钥文件内容
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("读取私钥数据失败：" + err.Error())
	}

	//pem解码
	block, _ := pem.Decode(data)

	//x509解码
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("x509解码失败：" + err.Error())
	}

	//对密文进行解密
	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)
	if err != nil {
		fmt.Println("解密失败：" + err.Error())
	}

	return plainText
}
