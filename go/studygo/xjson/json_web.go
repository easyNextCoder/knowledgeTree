package xjson

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	MAX_RETRY_TIMES = 2
)

type Wrap struct {
	Data    interface{} `json:"data"`
	Result  int         `json:"result"`
	Message string      `json:"message"`
}

type Project struct {
	// 共用部分
	Type int   `json:"type"` // Project类型
	Raw  []int `json:"raw"`  // 当前牌的原始牌
	// 顺子牌所需要的额外数据
	Preset []int `json:"preset"` // 当前牌所代表的牌 没有Okey牌的时候为空
}

func (s Project) String() string {
	return fmt.Sprintf("type(%v)-raw(%v)-preset(%v)", s.Type, s.Raw, s.Preset)
}

type AskMaxProj struct {
	Cards []int `json:"cards"`
}

type RspMaxProj struct {
	Projs []*Project `json:"projs"`
}

func post(c *AskMaxProj) []byte {
	marshal, err := json.Marshal(&Wrap{
		Data:    c,
		Result:  0,
		Message: "web",
	})
	if err != nil {
		return nil
	}
	return marshal
}

func jsonMarshalWork() {
	req := &AskMaxProj{Cards: []int{1, 2, 3, 4, 5}}
	//rsp := &RspMaxProj{
	//	Projs: []*Project{
	//		&Project{
	//			Type:   0,
	//			Raw:    []int{5, 4, 3, 2, 1},
	//			Preset: []int{9, 8, 7, 6, 5},
	//		},
	//	},
	//}

	marshalBuf, err := json.Marshal(req)
	if err != nil {
		return
	}

	fmt.Printf("Marshal []byte %v, string %v\n", marshalBuf, string(marshalBuf))

	rspBuf := post(req)

	receiveRsp := &RspMaxProj{}
	wrap := new(Wrap)
	wrap.Result = 0
	wrap.Data = receiveRsp

	err = json.Unmarshal(rspBuf, wrap)
	if err != nil {
		return
	}

	fmt.Printf("unmarmhal  %+v\n", wrap)

	//mp := map[string]interface{}{
	//	"result": 0,
	//	"data":   []int{1, 2, 3, 4, 5, 6, 7},
	//}
	//
	//buf, err := json.Marshal(mp)
	//if err != nil {
	//	return
	//}
	//
	//fmt.Println(string(buf))
	//
	//x := wrap.Data.([]interface{})
	//fmt.Println(reflect.TypeOf(x[1]).String())

}

func GetWebPostReq(url string, data interface{}, rsp interface{}) error {
	var resp *http.Response
	var err error
	buf, err := json.Marshal(data)
	if err != nil {
		return err
	}
	for i := 1; i <= MAX_RETRY_TIMES; i++ {
		resp, err = http.Post(url, "application/json", bytes.NewReader(buf))
		if err != nil {
			fmt.Printf("GetWebPostReq request url : %s retry times : %d, err : %s", url, i, err)
			if i == MAX_RETRY_TIMES {
				return err
			}
			<-time.After(time.Second * 2)
			continue
		}
		break
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("GetWebPostReq read resp.Body err : %s", err.Error())
		return err
	}

	wrapRsp := new(Wrap)
	wrapRsp.Data = rsp

	if err = json.Unmarshal([]byte(body), wrapRsp); err != nil {
		fmt.Printf("GetWebPostReq url:%s req:%s", url, string(buf))
		fmt.Printf("GetWebPostReq unmarshal resp status '%s' resp.Body '%s' err : %s", resp.Status, string(body), err.Error())
		return err
	}
	if wrapRsp.Result != 0 {
		if wrapRsp.Result != -1 {
			return errors.New("err")
		}
		return errors.New("err2")
	}

	return nil
}
