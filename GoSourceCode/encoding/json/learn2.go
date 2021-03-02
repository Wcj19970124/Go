package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

//golang的rsa密钥生成和保存
func learn2() {

	generateKey(2048)
}

func generateKey(bits int) {

	//利用rsa包的函数生成rsa的一对密钥
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		fmt.Println("生成rsa密钥对失败:" + err.Error())
	}

	//利用x509标准将得到的rsa私钥序列化未ANS.1的DER编码字符串
	x509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)

	//保存密钥
	priFile, err := os.Create("private.pem")
	if err != nil {
		fmt.Println("创建私钥pem文件失败:" + err.Error())
	}

	defer priFile.Close() //关闭文件句柄

	//构建pem.Block结构体对象
	privateBlock := pem.Block{Type: "RSA Private Key", Bytes: x509PrivateKey}
	if err := pem.Encode(priFile, &privateBlock); err != nil {
		fmt.Println("保存私钥失败！")
	}

	//获取公钥
	publicKey := privateKey.PublicKey
	//利用x509标准将得到的rsa公钥序列化未ANS.1的DER编码字符串
	x509PublicKey, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		fmt.Println("序列化公钥失败:" + err.Error())
	}

	pubFile, err := os.Create("public.pem")
	if err != nil {
		fmt.Println("创建公钥pem文件失败:" + err.Error())
	}

	defer pubFile.Close()

	publicBlock := pem.Block{Type: "RSA Public Key", Bytes: x509PublicKey}
	if err := pem.Encode(pubFile, &publicBlock); err != nil {
		fmt.Println("保存公钥失败")
	}

}
