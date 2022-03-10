package compress_test

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"testing"
	"time"
)

func TestGzipFile(t *testing.T) {
	GzipFile()
}

func GzipFile() {
	fw, err := os.Create("data.gzip") // 创建gzip包文件，返回*io.Writer
	if err != nil {
		log.Fatalln(err)
	}

	// 实例化心得gzip.Writer
	gw := gzip.NewWriter(fw)

	// 获取要打包的文件信息
	fr, err := os.Open("data.json")
	if err != nil {
		log.Fatalln(err)
	}

	// 获取文件头信息
	fi, err := fr.Stat()
	if err != nil {
		log.Fatalln(err)
	}

	// 创建gzip.Header
	gw.Header.Name = fi.Name()

	// fmt.Println("file size:", fi.Size())
	// 读取文件数据
	buf := make([]byte, fi.Size())
	_, err = fr.Read(buf)
	if err != nil {
		log.Fatalln(err)
	}
	// fmt.Println("readLen:", readLen)
	// 写入数据到zip包
	_, err = gw.Write(buf)
	if err != nil {
		log.Fatalln(err)
	}
	// fmt.Println("Write file size:",
	// size)

	_ = fr.Close()
	_ = gw.Close()
	_ = fw.Close()
}
func TestUnGzipFile(t *testing.T) {
	UnGzipFile()
}
func UnGzipFile() {
	// 打开gzip文件
	fr, err := os.Open("data.gzip")
	if err != nil {
		log.Fatalln(err)
	}

	// fi, err := fr.Stat()
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	//
	// fmt.Println("file size:", fi.Size())
	// 创建gzip.Reader
	gr, err := gzip.NewReader(fr)
	if err != nil {
		log.Fatalln(err)
	}

	// 读取文件内容
	// 如果单独使用，需自己决定要读多少内容，根据官方文档的说法，你读出的内容可能超出你的所需（当你压缩gzip文件中有多个文件时，强烈建议直接和tar组合使用）
	buf := make([]byte, 374)
	n, err := gr.Read(buf)
	// fmt.Println("gr read n ", n, err)
	// 将包中的文件数据写入
	fw, err := os.Create(gr.Header.Name)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = fw.Write(buf[:n])
	if err != nil {
		log.Fatalln(err)
	}

	fw.Close()
	gr.Close()
	fr.Close()

}

func TestUnGzip(t *testing.T) {
	UnGzip()
}
func UnGzip() {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)

	// Setting the Header fields is optional.
	zw.Name = "a-new-hope.txt"
	zw.Comment = "an epic space opera by George Lucas"
	zw.ModTime = time.Date(1977, time.May, 25, 0, 0, 0, 0, time.UTC)

	_, err := zw.Write([]byte("A long time ago in a galaxy far, far away..."))
	if err != nil {
		log.Fatal(err)
	}

	if err := zw.Close(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(buf.Len())
	fmt.Println(string(buf.Bytes()))

	zr, err := gzip.NewReader(&buf)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Name: %s\nComment: %s\nModTime: %s\n\n", zr.Name, zr.Comment, zr.ModTime.UTC())

	if _, err := io.Copy(os.Stdout, zr); err != nil {
		log.Fatal(err)
	}

	if err := zr.Close(); err != nil {
		log.Fatal(err)
	}
}
