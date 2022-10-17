package xml

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/feng-crazy/go-utils/file"
	"github.com/sirupsen/logrus"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

type Root struct {
	XMLName xml.Name `xml:"root"`
	Ord     string   `xml:"ord,attr"`
	Ret     string   `xml:"ret,attr"`
}

func ParseRoot(data []byte) (v Root, err error) {
	err = xml.Unmarshal([]byte(data), &v)
	if err != nil {
		return
	}
	// logrus.Printf("Root : %+v\n", v)
	return
}

type UsrLoginRet struct {
	XMLName     xml.Name   `xml:"root"`
	Ord         string     `xml:"ord,attr"`
	Ret         string     `xml:"ret,attr"`
	UsrIP       string     `xml:"usr_guid,attr"`
	UsrNickName string     `xml:"usr_nickname,attr"`
	UsrPriority string     `xml:"usr_priority,attr"`
	ServerTime  string     `xml:"time,attr"`
	Areas       []Area     `xml:"areas>area"`
	Usrs        []UsrLogin `xml:"usrs>usr"`
}

type Area struct {
	// XMLName   xml.Name `xml:"area"`

	Id   string `xml:"id,attr"`   // 分区 ID
	Name string `xml:"name,attr"` // 分区名称
}

type UsrLogin struct {
	// XMLName   xml.Name `xml:"usr"`

	Name        string `xml:"name,attr"`         // 设备用户名
	NickName    string `xml:"nick_name,attr"`    // 设备昵称
	Id          string `xml:"id,attr"`           // 设备 ID
	Pid         string `xml:"pid,attr"`          // 所属 area id” 所属的区域 ID
	UsrIP       string `xml:"usr_ip,attr"`       // 设备 IP
	UsrPriority string `xml:"usr_priority,attr"` //  本用户等级 1-16
	Status      string `xml:"status,attr"`       // 状态，0[离线]/1[在线]/2[喊话中]/3[对讲中]/4[呼叫 中]/5[被呼中]/6[广播中]/7[监控中]/8[会议中]/9[会 议呼叫中]。 非 0 代表在线
	IoIn2Tp     string `xml:"io_in2_tp,attr"`    //  设备类型，0[按钮]/1[门禁系统]/其它系统
}

func TestXml3(t *testing.T) {
	data, err := file.ToBytes("./test")
	if err != nil {
		fmt.Println(err)
	}
	data, _ = GbkToUtf8(data)
	fmt.Println(ParseRoot(data))

	var v UsrLoginRet
	err = xml.Unmarshal(data, &v)
	if err != nil {
		logrus.Error(err)
	}
	logrus.Printf("UsrLoginRet : %+v", v)
}

func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		logrus.Error(e)
		return nil, e
	}
	return d, nil
}
