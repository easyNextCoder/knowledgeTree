package ta_hand

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"os"
	"runtime"
)

func gbk2utf8(s string) []byte {
	trans := transform.NewReader(bytes.NewReader([]byte(s)), simplifiedchinese.GBK.NewDecoder())
	all, err := ioutil.ReadAll(trans)
	if err != nil {
		return nil
	}
	return all
}

func ReadTa() error {

	_, f, _, _ := runtime.Caller(1)
	fmt.Println("this is f", f)

	file, err := os.Open("/Users/zp/Documents/knowledgeTree/go/studygo/xfile/ta_hand/new_hand_user_setOnece.csv")
	if err != nil {
		return err
	} //

	reader := csv.NewReader(file)
	//reader.Comma = ','
	reader.Comment = '#'
	reader.LazyQuotes = true
	reader.FieldsPerRecord = -1

	for {
		row, err := reader.Read()
		if err != nil {
			return err
		}
		for i, r := range row {
			out := r
			fmt.Println("res:", i, string(out))
		}
	}
}
