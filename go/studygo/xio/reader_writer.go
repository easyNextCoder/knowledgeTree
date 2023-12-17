package xio

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"
)

func countLetters(r io.Reader) (map[string]int, error) {
	buf := make([]byte, 2048)
	out := map[string]int{}
	for {
		n, err := r.Read(buf)
		for _, b := range buf[:n] {
			if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
				out[string(b)]++
			}
		}

		if err == io.EOF {
			return out, nil
		}

		if err != nil {
			return nil, err
		}
	}
}

func readerWork() error {
	s := "The quick brown fox jumped over the lazy dog"
	sr := strings.NewReader(s)
	counts, err := countLetters(sr)
	if err != nil {
		return err
	}
	fmt.Println(counts)
	return nil
}

func buildGZipReader(fileName string) (*gzip.Reader, func(), error) {
	r, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("buildGZipReader err(%s)", err)
		return nil, nil, err
	}

	gr, err := gzip.NewReader(r)
	if err != nil {
		return nil, nil, err
	}

	return gr, func() {
		gr.Close()
		r.Close()
	}, nil
}

func buildGZipReaderWork() error {
	r, closer, err := buildGZipReader("/Users/zp/Documents/knowledgeTree/go/studygo/xinterface/tmp/tmp.gzip")
	if err != nil {
		fmt.Printf("buildGZipReaderWork err(%s)\n", err)
		return err
	}
	defer closer()

	counts, err := countLetters(r)
	if err != nil {
		return err
	}

	fmt.Println(counts)
	return nil
}

type SelfReader struct {
	selfStr string
}

func (self *SelfReader) Read(p []byte) (int, error) {
	copy(p, []byte(self.selfStr))
	return len(self.selfStr), io.EOF
}

func selfReaderWork() {
	r := &SelfReader{selfStr: "abcdefghigklmnop"}
	letters, err := countLetters(r)
	fmt.Println("selfReaderWork reader", letters)
	if err != nil {
		return
	}
}
