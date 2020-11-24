package main

import (
	"compress/gzip"
	"fmt"
	"os"
)

func main() {

	gzipFileName := "demo.gzip"
	// unGzipFileName := "demo.txt"
	// if compressGzip(gzipFileName, unGzipFileName) == nil {
	// 	fmt.Println("压缩成功")
	// } else {
	// 	fmt.Println("压缩失败")
	// }

	if unCompressGzip(gzipFileName) == nil {
		fmt.Println("解压成功")
	} else {
		fmt.Println("解压失败")
	}
}

//compressGzip 压缩普通文件为gzip格式
func compressGzip(gzipFileName string, unGzipFileName string) error {

	gfile, err := os.Create(gzipFileName)
	if err != nil {
		return err
	}

	defer gfile.Close()

	gw := gzip.NewWriter(gfile)
	defer gw.Close()

	file, err := os.Open(unGzipFileName)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	gw.Header.Name = fileInfo.Name()

	buf := make([]byte, fileInfo.Size())
	_, err = file.Read(buf)
	if err != nil {
		return err
	}

	_, err = gw.Write(buf)
	if err != nil {
		return err
	}

	return nil
}

//unCompressGzip 解压缩gzip
func unCompressGzip(gzipFileName string) error {

	file, err := os.Open(gzipFileName)
	if err != nil {
		return err
	}
	defer file.Close()

	gr, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	defer gr.Close()

	dstFile, err := os.Create(gr.Header.Name)
	if err != nil {
		return err
	}

	buf := make([]byte, 1024*1024*10)
	n, err := gr.Read(buf)

	_, err = dstFile.Write(buf[:n])
	if err != nil {
		return err
	}

	return nil
}
