package xjson

import (
	"encoding/json"
	"fmt"
)

type (
	S struct {
		Val int64 `json:"val"`
	}
)

type (
	SV struct {
		Val interface{} `json:"val"`
	}
)

func jsonWork() {
	jsonv, _ := json.Marshal(&S{Val: 10000000000000})
	fmt.Println(string(jsonv), jsonv)

	sv := new(SV)

	fmt.Printf("svp(%p) svpp(%p) sv(%p) sv(%+v)\n", sv, &sv, *(&sv), *sv)
	//fmt.Printf("(%d)\n", reflect.ValueOf(&sv).Pointer())
	err := json.Unmarshal(jsonv, sv)
	if err != nil {
		return
	}

	fmt.Printf("svp(%+v) svpp(%p) sv(%v) sv(%+v)\n", sv, &sv, sv, *sv)

	switch sv.Val.(type) {
	case float64:
		fmt.Println("sv.Val type is float64")
	default:
		fmt.Println("use default type")
	}

	ival := int(sv.Val.(float64))
	fmt.Println(ival)
	//ret := reflect.ValueOf(svp.Val)
	//fmt.Println(ret.Kind(), ret.Interface())

	//vali := svp.Val.(int)
	//fmt.Println(vali)

	//v.Set(reflect.ValueOf(string(s)))
}

type Info struct {
	Name string `json:"name"`
}

type UserMaxProjRsp struct {
	Projects []*Info `json:"projects"`
	Cards    []int   `json:"cards"`
	Score    int     `json:"score"`
}

type UserUsePrevCardCrushRsp struct {
	MaxProj *UserMaxProjRsp
	Used    bool `json:"used"`
}

type CrushAssistInfo struct {
	*UserUsePrevCardCrushRsp
}

func nestStructMarshalWork() {
	x := &UserUsePrevCardCrushRsp{
		MaxProj: &UserMaxProjRsp{
			Projects: []*Info{&Info{Name: "xuyongkang"}, &Info{Name: "xuyong"}},
			Cards:    []int{1, 2, 3},
			Score:    100,
		},
		Used: false,
	}
	y := &CrushAssistInfo{UserUsePrevCardCrushRsp: x}

	bytes, ok := json.Marshal(y)
	fmt.Println(string(bytes), ok)

	unmarshalObj := new(CrushAssistInfo)
	json.Unmarshal(bytes, unmarshalObj)
	fmt.Printf("unmarshalObj(%+v)", unmarshalObj.MaxProj)

}

type (
	StructA struct {
		Used bool `json:"used"`
	}
	StructB struct {
		UnUsed bool `json:"un_used"`
	}
)

func unmarshalToDifferentType() {
	var a StructA = StructA{Used: true}
	var b StructB

	bytes, err := json.Marshal(&a)
	if err != nil {
		return
	}

	err = json.Unmarshal(bytes, &b)
	if err != nil {
		fmt.Println("this is err")
		return
	}

	fmt.Printf("%+v", b)
}
