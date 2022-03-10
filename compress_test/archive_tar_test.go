package compress_test

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"log"
	"os"
	"path"
	"strings"
)

// srcFilePath源文件路径  destDirPath目标文件路径
func Uncompress(srcFilePath string, destDirPath string) (string, error) {
	var temp string
	// 创建目标文件路径
	os.Mkdir(destDirPath, os.ModePerm)
	fr, err := os.Open(srcFilePath)
	if err != nil {
		return temp, err
	}
	defer fr.Close()
	// Gzip reader
	gr, err := gzip.NewReader(fr)
	if err != nil {
		return temp, err
	}
	defer gr.Close()
	// Tar reader
	tr := tar.NewReader(gr)

	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			// End of tar archive
			break
		}
		// Check if it is diretory or file
		if hdr.Typeflag != tar.TypeDir {
			// Get files from archive
			// Create diretory before create file
			os.MkdirAll(destDirPath+"/"+path.Dir(hdr.Name), os.ModePerm)
			if temp == "" {
				temp = path.Dir(hdr.Name)
			}
			// Write data to file
			fw, _ := os.Create(destDirPath + "/" + hdr.Name)
			if err != nil {
				return temp, err
			}
			_, err = io.Copy(fw, tr)
			if err != nil {
				return temp, err
			}
		}
	}
	return temp, nil
}

// 单个文件
func UnGz(srcFilePath, destDirPath string) error {
	fr, err := os.Open(srcFilePath)
	if err != nil {
		log.Printf("open secGZ failed. err:%v\n", err)
		return err
	}
	defer fr.Close()
	gr, err := gzip.NewReader(fr)
	if err != nil {
		return err
	}
	defer gr.Close()
	index := strings.LastIndex(srcFilePath, ".")
	if index == -1 {
		log.Printf("find . failed. err:%v\n", err)
		return err
	}
	// 因为fr.Name()和gr.Name为空，只能切割出文件名
	fw, err := os.Create(destDirPath + "/" + strings.Split(fr.Name(), ".")[len(strings.Split(fr.Name(), "."))-2])
	if err != nil {
		return err
	}
	defer fw.Close()
	// 写文件
	_, err = io.Copy(fw, gr)
	if err != nil {
		return err
	}
	return nil
}
