package xml_test

import (
	"encoding/xml"
	"fmt"
	"os"
	"testing"

	// "encoding/json"
	"io/ioutil"

	"github.com/feng-crazy/go-utils/file"
)

type Location struct {
	CountryRegion []CountryRegion
}

type CountryRegion struct {
	Name  string `xml:",attr"`
	Code  string `xml:",attr"`
	State []State
}

type State struct {
	Name string `xml:",attr"`
	Code string `xml:",attr"`
	City []City
}

type City struct {
	Name   string `xml:",attr"`
	Code   string `xml:",attr"`
	Region []Region
}

type Region struct {
	Name string `xml:",attr"`
	Code string `xml:",attr"`
}

func TestXMlParse(t *testing.T) {
	f, err := os.Open("LocList.xml")
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	// v := make(map[string]interface{})
	var v Location
	err = xml.Unmarshal(data, &v)
	if err != nil {
		panic(err)
	}

	// fmt.Printf("%#v\n", v)

	// table
	for _, countryRegion := range v.CountryRegion {
		// fmt.Printf("%s,%s\n", countryRegion.Code, countryRegion.Name)
		if len(countryRegion.State) == 0 {
			continue
		}
		for _, state := range countryRegion.State {
			// fmt.Printf("%s,%s,%s\n", countryRegion.Code, state.Code, state.Name)
			if len(state.City) == 0 {
				continue
			}
			for _, city := range state.City {
				// fmt.Printf("%s,%s,%s,%s\n", countryRegion.Code, state.Code, city.Code, city.Name)
				if len(city.Region) == 0 {
					continue
				}
				for _, region := range city.Region {
					fmt.Printf("%s,%s,%s,%s,%s\n", countryRegion.Code, state.Code, city.Code, region.Code, region.Name)
				}
			}
		}
	}

	// // json
	// js, err := json.Marshal(&v.CountryRegion[0])
	// if err != nil {
	//  panic(err)
	// }
	// fmt.Printf("%s\n", js)
}

type usr_connet_cloud struct {
	XMLName xml.Name `xml:"root"`
	Ord     string   `xml:"ord,attr"`
	Ret     string   `xml:"ret,attr"`
	UsrIP   string   `xml:"usr_ip,attr"`
	UsrPort string   ` xml:"usr_port,attr"`
}

func TestParseXMl2(t *testing.T) {
	xmldata, _ := file.ToBytes("test.xml")
	var v usr_connet_cloud
	err := xml.Unmarshal([]byte(xmldata), &v)
	if err != nil {
		panic(err)
	}
	fmt.Println(v)
}
