package map_test

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/sirupsen/logrus"
	"testing"
)

type HDF struct {
	Name      string  `hdf:"name"`
	AgeHaha   int     `hdf:"ageHaha"`
	ScoreRate float64 `hdf:"score_rate"`
}

func TestMapStructure(t *testing.T) {
	var result map[string]interface{}
	config := &mapstructure.DecoderConfig{
		Metadata:   nil,
		ZeroFields: false,
		TagName:    "hdf",
		Result:     &result,
	}
	dc, err := mapstructure.NewDecoder(config)
	if err != nil {
		logrus.Error(err)
	}
	err = dc.Decode(&HDF{
		Name:      "test",
		AgeHaha:   12,
		ScoreRate: 111,
	})
	if err != nil {
		logrus.Error(err)
	}

	fmt.Println(result)

}
