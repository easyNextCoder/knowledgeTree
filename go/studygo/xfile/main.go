package main

import (
	"fmt"
	"studygo/xfile/search_xls"
)

func main() {

	//err := ta_hand.ReadTa()
	//if err != nil {
	//	fmt.Printf("err occur %s\n", err)
	//	return
	//}

	err := search_xls.SearchXls("/Users/zp/Documents/knowledgeTree/go/studygo/xfile/search_xls/9a4009092417x.xlsx")
	if err != nil {
		fmt.Printf("searchXls err(%s)", err)
		return
	}

}
