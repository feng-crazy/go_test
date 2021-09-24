package xml

import (
	"encoding/xml"
	"fmt"
	"testing"
)

/*
XMLName的标签为`xml:"Person"`,则给的XML编码中XML元素必须有给定的名Person，否则会报错error: expected element type <person> but have <Person>
Name的标签为`xml:"FullName"`，则给的XML编码中解析XML元素为FullName的值为字段Name对应的值
Age的标签为`xml:"-"`，则不会被反序列填写，可以观察出Age的值为0
Phone无标签，则解析时将XML元素为Phone的值为该字段对应的值
Groups的标签为`xml:"Group>Value"`，则解析时将XML元素为Group的子元素Value的值为该字段对应的值
Email结构的Where字段的标签为`xml:"where,attr"`，则解析时将XML元素为Email的属性Where的值赋给该字段
Description标签为",innerxml",通过打印Description的值，可以看出直接将对应原始XML文本写入该字段
补充说明：
%#v-值的Go语法表示
%v-值的默认格式表示。当输出结构体时，扩展标志（%+v）会添加字段名
%q-该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
*/
func Test_XMLUnMarshal(t *testing.T) {
	type Email struct {
		Where string `xml:"where,attr"`
		Addr  string
	}
	type Address struct {
		City, State string
	}
	type Result struct {
		XMLName xml.Name `xml:"Person"`
		Name    string   `xml:"FullName"`
		Age     int      `xml:"-"`
		Phone   string
		Email   []Email
		Groups  []string `xml:"Group>Value"`
		Address
		Description string `xml:",innerxml"`
	}
	v := Result{Name: "none", Phone: "none"}
	data := `
		<Person>
			<FullName>Grace R. Emlin</FullName>
			<Company>Example Inc.</Company>
           <Age>20</Age>
			<Email where="home">
				<Addr>gre@example.com</Addr>
			</Email>
			<Email where='work'>
				<Addr>gre@work.com</Addr>
			</Email>
			<Group>
				<Value>Friends</Value>
				<Value>Squash</Value>
			</Group>
			<City>Hanga Roa</City>
			<State>Easter Island</State>
		</Person>
	`
	err := xml.Unmarshal([]byte(data), &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Printf("XMLName: %#v\n", v.XMLName)
	fmt.Printf("Name: %q\n", v.Name)
	fmt.Printf("Age: %v\n", v.Age)
	fmt.Printf("Phone: %q\n", v.Phone)
	fmt.Printf("Email: %v\n", v.Email)
	fmt.Printf("Email[0].Where: %v\n", v.Email[0].Where)
	fmt.Printf("Groups: %v\n", v.Groups)
	fmt.Printf("Address: %v\n", v.Address)
	fmt.Printf("Description: %v\n", v.Description)
}
