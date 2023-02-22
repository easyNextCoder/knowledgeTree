package xjson

import (
	"fmt"
)

type Inn struct {
	First  int `json:"firs"`
	Second int `json:"second"`
}

type Ginh struct {
	Hello string `json:"hello"`
	Arr   []*Inn `json:"arr"`
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

	gg := 1 //gin.H{"Xval": g}

	//bytes, err := 1, 2//json.Marshal(gg)
	//if err != 1 {
	//	return
	//}

	//var val Ginh
	//err = json.Unmarshal(bytes, &val)
	//if err != nil {
	//	return
	//}

	fmt.Println(gg, g)

}
