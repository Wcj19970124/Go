package main

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
)

func main() {

	//压缩文件
	// unTarFile := "abc.txt"
	tarFile := "dst.tar"
	// param := Tar(tarFile, unTarFile)
	// switch param.(type) {
	// case error:
	// 	fmt.Println("发生了错误:", param)
	// case bool:
	// 	fmt.Println("压缩文件成功!")
	// }

	//解压文件
	unTarFilePath := "Go/GoSourceCode/archive/tar"
	param := unTar(tarFile, unTarFilePath)
	switch param.(type) {
	case error:
		fmt.Println("解压失败", param)
	case bool:
		fmt.Println("解压成功")
	}
}

//Tar 压缩
//因为tar文件分为两个部分：文件信息和文件数据;所以需要将文件分成两部分进行写入
//分析：
//第一步：打开原始文件，获取一个原始文件的句柄，用于之后文件的压缩
//第二步：因为写入tar文件，需要一个tar.Writer对象，所以我们创建一个tar文件，然后用于创建生成一个tar.Writer对象
//第三步：需要注意的是，tar包下文件的写入和读取是利用的缓存，所以我们必须关闭tar.Writer对象
//第四步：tar文件都包括两部分，分别为文件信息和文件数据，我们需要分别将两部分写入到tar文件中
//第五步：根据原始文件获取文件信息，然后写入tar文件
//第六步：将原始文件的数据写入到tar文件
func Tar(tarFile string, unTarFile string) interface{} {

	//打开原始文件,获取一个操作原始文件的文件句柄
	file, err := os.Open(unTarFile)
	if err != nil {
		return err
	}

	defer file.Close() //最终必须关闭打开的文件句柄

	//创建一个tar文件，作为os.Writer
	destFile, err := os.Create(tarFile)
	if err != nil {
		return err
	}

	//获取tar.Writer对象
	tw := tar.NewWriter(destFile)
	defer tw.Close() //最终需要关闭

	//获取原始文件文件信息，写入tar文件
	fileInfo, err := os.Stat(unTarFile)
	if err != nil {
		return err
	}

	//将原始文件信息写入到tar文件
	fileInfoHeader, err := tar.FileInfoHeader(fileInfo, "")
	if err != nil {
		return err
	}
	err = tw.WriteHeader(fileInfoHeader)
	if err != nil {
		return err
	}

	//将原始文件的数据写入tar文件
	_, err = io.Copy(tw, file)
	if err != nil {
		return err
	}

	return true
}

//解压缩tar文件到普通文件中
//分析：
//第一步：打开tar文件，获取操作文件的句柄
//第二步：因为读取tar文件，需要一个tar.Reader对象，进而需要一个os.Reader对象，可以使用打开的tar文件句柄
//第三步：创建普通文件目录，用于存放解压的普通文件
//第四步：利用tar.Reader的next方法遍历tar文件，读取其中的信息，然后写入到普通文件
func unTar(tarFile string, unTarFilePath string) interface{} {

	//第一步
	file, err := os.Open(tarFile)
	if err != nil {
		return err
	}

	defer file.Close()

	//第二步
	tr := tar.NewReader(file)

	//第三步
	fileInfo, err := os.Stat(unTarFilePath)
	if fileInfo != nil {
		os.RemoveAll(unTarFilePath)
	}
	os.MkdirAll(unTarFilePath, os.ModePerm)

	//第四步
	for {

		header, err := tr.Next() //获取压缩文件的文件信息
		if err == io.EOF {
			break //读取到文件末尾截至
		}
		if err != nil {
			return err
		}

		unTarFile := unTarFilePath + header.Name //将文件路径和原始文件名拼接作为目标文件名

		dstFile, err := os.Create(unTarFile) //创建目标文件
		if err != nil {
			return err
		}

		_, err = io.Copy(dstFile, tr)
		if err != nil {
			return err
		}
	}

	return true
}
