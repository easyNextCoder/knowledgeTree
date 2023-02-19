package xstring

import (
	"fmt"
	"math/rand"
)

func tSprintf() {
	head := "AssignCard"
	s := fmt.Sprintf(head+"worhname %s", 5)
	fmt.Println(s)
}

func RandString(n int) string {
	b := make([]byte, n)
	for i, _ := range b {
		b[i] = byte(i%26 + 97)
	}
	rand.Shuffle(n, func(i, j int) {
		v := b[i]
		b[i] = b[j]
		b[j] = v
	})
	return string(b)
}
