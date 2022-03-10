package compress_test

import (
	"encoding/json"
	"log"
	"os"
	"testing"

	"github.com/feng-crazy/go-utils/file"
	jsoniter "github.com/json-iterator/go"
)

func BenchmarkJson(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		fr, err := os.Open("data.json") // 创建gzip包文件，返回*io.Writer
		if err != nil {
			log.Fatalln(err)
		}
		data := make([]byte, 374)
		_, err = fr.Read(data)
		if err != nil {
			log.Fatal(err)
		}
		_ = fr.Close()
		m := make(map[string]interface{})
		err = json.Unmarshal(data, &m)
		if err != nil {
			log.Fatal(err)
		}

		fw, err := os.Create("data.tmp")
		_, err = fw.Write(data)
		if err != nil {
			log.Fatal(err)
		}

		data, err = json.Marshal(m)
		if err != nil {
			log.Fatal(err)
		}

		_ = fw.Close()
	}
}

func BenchmarkJsoniter(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		fr, err := os.Open("data.json") // 创建gzip包文件，返回*io.Writer
		if err != nil {
			log.Fatalln(err)
		}
		data := make([]byte, 374)
		_, err = fr.Read(data)
		if err != nil {
			log.Fatal(err)
		}
		m := make(map[string]interface{})
		err = jsoniter.Unmarshal(data, &m)
		if err != nil {
			log.Fatal(err)
		}

		fw, err := os.Create("data.tmp")
		_, err = fw.Write(data)
		if err != nil {
			log.Fatal(err)
		}

		data, err = jsoniter.Marshal(m)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func BenchmarkGzip(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		GzipFile()
		UnGzipFile()
	}
}

func BenchmarkZlib(b *testing.B) {
	b.ReportAllocs()
	data, _ := file.ToBytes("data.json")
	log.Println("data len", len(data))
	for i := 0; i < b.N; i++ {
		LibzFile()
		UnLibzFile()
	}
}
