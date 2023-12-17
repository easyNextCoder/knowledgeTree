package xbit

import "fmt"

//16进制的位移动操作，只用了后两位
func CalSource(oldStatus, sourceVal int) int {
	last := oldStatus & 15
	res := (last << 4) | sourceVal
	return res
}

func check(x int64, index int) {
	x = 17
	x = x & (1 << 10)
	a := int32(-1)
	for i := 0; i < 32; i++ {
		fmt.Println(a & (1 << i))
	}

	fmt.Printf("%064b, %b, %064b", x, int32(2147483647), (int64(-1)))
	y := (2 | 1<<0)
	fmt.Println("Y is:", y)
}

func runCheck() {
	check(0, 0)
}

func borderCheck(border [2]int) error {
	left, right := border[0], border[1]

	if left < 0 {
		return fmt.Errorf("borderCheck err border(%+v)", border)
	}

	if !(right > left) {
		return fmt.Errorf("borderCheck err border(%+v)", border)
	}

	if right-left > 31 { //最长允许设置31位
		return fmt.Errorf("borderCheck err border(%+v)", border)
	}

	return nil
}

func SetInt64(in uint64, v int, border [2]int) (out uint64, err error) {

	out = in
	err = nil

	if err = borderCheck(border); err != nil {
		return
	}

	l, r := border[0], border[1]

	length := r - l

	uint64v := uint64(v)

	mask := (uint64(1) << length) - 1 //0000111

	afterMove := mask << l

	in = in | afterMove //对应位置先设为1

	afterMoveOr := ^afterMove //border对应位置为0，其余位置为1

	uint64v = uint64(uint64v) << l

	uint64v = afterMoveOr | uint64v //空余位置全部设置为1

	out = in & uint64v

	return
}

func GetInt64(in uint64, l, r uint) int {

	length := r - l

	mask := (uint64(1) << length) - 1

	afterMove := in >> l

	return int(afterMove & mask)
}
