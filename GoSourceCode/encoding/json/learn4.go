package main

import (
	"crypto/aes"
	"fmt"
)

//sign的值为：6f56569d12874015e80a6df7ae0c84ad25eaa6bb7ab2fa9a1fd3203ac45bf595b88758cdc71ce28ca21cf56aea10e99ca75387075e21571ee2cf79be9c82824a8272b85978c2ab2067279b91af1d4199a3607dc49248d39d8e568c28d82e0469a2a674ce4e4ec
//17eae755daae7f44369a1f47571aa466ee8cfcfca2850cf37caa779a75f57021defb9dca64c2cdfc78884e0b7d7508f42729c9fef1e04c04dbf2d476d4934604d5bf58cbca6dfe727fc526d38fbcb1345e08eb55f0baa81f97227b1ae223cc63b9921a7aad79e60573e0fb15776b962d9dab3b978cbf20eef726f8d194ea82603d6667f5fa4a1a5a493c16a3a411e1aab2da7f0f6327330acf7
//rsa_dec的结果，bin2hex：b07e9fcefd6210143d9975a0dcc87130e05dbc76ced18850edd8c0def82d8bf70b21b7c8c84826b1c4d1ad0dc80049ddb9f77ab76fcd04c65f91cf89a9406246dda53df34ad387850d3270392e02613869fb587c3032270deae21ff8d6930bbf
//aes_dec的结果：1130227577+d6c44f9c1f6897e0dc15549ddb98c5f13a1aa37ea5efe7f2e3b202e524870b91+1614664104205+118
//aes解密
func main() {

	// cipherText := "b07e9fcefd6210143d9975a0dcc87130e05dbc76ced18850edd8c0def82d8bf70b21b7c8c84826b1c4d1ad0dc80049ddb9f77ab76fcd04c65f91cf89a9406246dda53df34ad387850d3270392e02613869fb587c3032270deae21ff8d6930bbf"
	// cipher, err := hex.DecodeString(cipherText)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// key := "84e6c6dc0f9p4a56"
	// plainText, _ := AESDecrypt([]byte(cipher), []byte(key))
	// fmt.Println(string(plainText))

	text := "hello world!"
	key := "84e6c6dc0f9p4a56"
	cilpherText, err := AESEncrypt([]byte(text), []byte(key))
	if err != nil {
		panic(err)
	} else {
		fmt.Println(string(cilpherText))
	}

	plainText, err := AESDecrypt([]byte(cilpherText), []byte(key))
	if err != nil {
		panic(err)
	} else {
		fmt.Println(string(plainText))
	}
}

//AESDecrypt AES解密
func AESDecrypt(src, key []byte) ([]byte, error) {
	//TODD：创建一个block
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	//创建缓存，用于存储分组解码的数据
	buf := make([]byte, block.BlockSize())
	decrypted := make([]byte, 0)
	for i := 0; i < len(src); i += block.BlockSize() {
		block.Decrypt(buf, src[i:i+block.BlockSize()])
		decrypted = append(decrypted, buf...)
	}

	return decrypted, nil
}

//AESEncrypt AES加密
func AESEncrypt(src, key []byte) ([]byte, error) {
	//TODD：创建一个block
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	//创建缓存，用于存储加密的每个分组数据
	buf := make([]byte, block.BlockSize())
	encrypted := make([]byte, 0)
	for i := 0; i < len(src); i += block.BlockSize() {
		block.Encrypt(buf, src[i:i+block.BlockSize()])
		encrypted = append(encrypted, buf...)
	}

	return encrypted, nil
}
