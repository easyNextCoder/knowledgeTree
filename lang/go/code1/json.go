package main

import (
	"encoding/json"
	"fmt"
)

func jsonMarshal() {

	tmp := struct {
		UserInfo string
		Id       int
	}{
		UserInfo: "xuyongkang",
		Id:       1,
	}
	v, _ := json.Marshal(tmp)

	//json.Unmarshal(v, UserCommand.rsp)
	//fmt.Println(string(v), UserCommand.rsp)

	oy := new(CommandWrapper)
	oy.Data = UserCommand.rsp
	json.Unmarshal(v, oy.Data)
	fmt.Println(UserCommand.rsp, oy.Data, &(UserCommand.rsp), &(oy.Data))

}
