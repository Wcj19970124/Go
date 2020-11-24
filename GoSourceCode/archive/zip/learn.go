package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

func main() {

	unZipFile := "abc.txt"
	zipFile := "dst.zip"
	param := compressZip(zipFile, unZipFile)
	switch param.(type) {
	case error:
		fmt.Println("压缩失败", param)
	case bool:
		fmt.Println("压缩成功!")
	}

}

//读取文件内容，并将其压缩为zip文件
//基本步骤和tar文件类似
func compressZip(zipFile string, unZipFile string) interface{} {

	//打开原始文件，获取原始文件句柄
	file, err := os.Open(unZipFile)
	if err != nil {
		return err
	}

	defer file.Close()

	//创建目标压缩文件
	dstFile, err := os.Create(zipFile)
	if err != nil {
		return err
	}

	//获取zip写入器
	zw := zip.NewWriter(dstFile)
	defer zw.Close()

	//获取文件信息，并且写入
	fileInfo, err := os.Stat(unZipFile)
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(fileInfo)
	if err != nil {
		return err
	}

	iw, err := zw.CreateHeader(header)
	if err != nil {
		return err
	}

	//写入文件数据
	_, err = io.Copy(iw, file)
	if err != nil {
		return err
	}

	return true
}
