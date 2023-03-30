package xjson

import (
	"encoding/json"
	"fmt"
)

type Inn struct {
	First  int `json:"firs"`
	Second int `json:"second"`
}

type InnOtherName struct {
	First  int `json:"firs"`
	Second int `json:"second"`
}

type Ginh struct {
	Hello string `json:"hello"`
	Arr   []*Inn `json:"arr"`
}

type GinhOtherName struct {
	Hello string          `json:"hello"`
	Arr   []*InnOtherName `json:"arr"`
}

func ginWork() {
	g := Ginh{
		Hello: "xuyongkang",
		Arr: []*Inn{
			&Inn{
				First:  1,
				Second: 2,
			},
			&Inn{
				First:  4,
				Second: 8,
			},
		},
	}

	bytes, err := json.Marshal(&g)
	if err != nil {
		return
	}

	var pp GinhOtherName
	err = json.Unmarshal(bytes, &pp)
	if err != nil {
		return
	}

	//v := pp["Xval"].(Ginh)

	fmt.Println(g, pp, pp.Hello, pp.Arr, pp.Arr[0], pp.Arr[1])

}
