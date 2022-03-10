package compress_test

import (
	"compress/zlib"
	"log"
	"os"
	"testing"
)

func TestLibzFile(t *testing.T) {
	LibzFile()
}

func LibzFile() {
	fw, err := os.Create("data.zlib") // 创建zlib包文件，返回*io.Writer
	if err != nil {
		log.Fatalln(err)
	}
	defer fw.Close()

	// 实例化心得zlib.Writer
	zw := zlib.NewWriter(fw)
	defer zw.Close()

	// 获取要打包的文件信息
	fr, err := os.Open("data.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer fr.Close()

	// 获取文件头信息
	fi, err := fr.Stat()
	if err != nil {
		log.Fatalln(err)
	}

	// 读取文件数据
	buf := make([]byte, fi.Size())
	_, err = fr.Read(buf)
	if err != nil {
		log.Fatalln(err)
	}
	// 写入数据到zip包
	_, err = zw.Write(buf)
	if err != nil {
		log.Fatalln(err)
	}

}

func TestUnLibzFile(t *testing.T) {
	UnLibzFile()
}
func UnLibzFile() {
	// 打开zlib文件
	fr, err := os.Open("data.zlib")
	if err != nil {
		log.Fatalln(err)
	}
	defer fr.Close()

	// 创建zlib.Reader
	gr, err := zlib.NewReader(fr)
	if err != nil {
		log.Fatalln(err)
	}
	defer gr.Close()

	// 读取文件内容
	// 如果单独使用，需自己决定要读多少内容，根据官方文档的说法，你读出的内容可能超出你的所需（当你压缩zlib文件中有多个文件时，强烈建议直接和tar组合使用）
	buf := make([]byte, 374)
	n, err := gr.Read(buf)

	// 将包中的文件数据写入
	fw, err := os.Create("data.json")
	if err != nil {
		log.Fatalln(err)
	}
	_, err = fw.Write(buf[:n])
	if err != nil {
		log.Fatalln(err)
	}
}
