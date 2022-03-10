package gin_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestServer(t *testing.T) {
	// 创建默认的引擎
	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})

		return
	})
	// 运行 默认为80端口
	err := r.Run(":6666")
	if err != nil {
		return
	}
}

func TestServer1(t *testing.T) {
	// 创建默认的引擎
	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "server1",
		})

		return
	})
	// 运行 默认为80端口
	err := r.Run(":6677")
	if err != nil {
		return
	}
}
func TestServerFile(t *testing.T) {
	// 创建默认的引擎
	r := gin.Default()

	// //创建请求 当访问地址为/uploadfile时执行后面的函数
	// r.POST("/uploadfile", func(c *gin.Context) {
	//	//获取表单数据 参数为name值
	//	f, err := c.FormFile("f1")
	//	//错误处理
	//	if err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{
	//			"error": err,
	//		})
	//		return
	//	} else {
	//		//将文件保存至本项目根目录中
	//		err = c.SaveUploadedFile(f, f.Filename)
	//		//保存成功返回正确的Json数据
	//		c.JSON(http.StatusOK, gin.H{
	//			"message": "OK",
	//		})
	//	}
	//
	// })

	r.GET("/download/test", func(c *gin.Context) {
		data := make(map[string]interface{})
		err := c.ShouldBind(&data)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.Writer.WriteHeader(http.StatusOK)
		c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "test.txt")) // fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名

		file, err := ioutil.ReadFile("./data/test.txt")
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		c.Data(200, "application/octet-stream", file)

		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})

		return
	})
	// 运行 默认为80端口
	r.Run(":6666")
}
