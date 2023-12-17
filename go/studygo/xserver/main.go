package main

import (
	bytes2 "bytes"
	"fmt"
	"github.com/flosch/pongo2"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type rootHandler struct {
}

func (self *rootHandler) ServeHTTP(http.ResponseWriter, *http.Request) {

}

func server() {

	http.ListenAndServe("8081", &rootHandler{})
}

type xFileHandler struct {
	h http.Handler
	s string
}

func (self *xFileHandler) ServeHTTP(rsp http.ResponseWriter, req *http.Request) {
	fmt.Println("xFileHandler", req.URL.Path, "work final")
	self.h.ServeHTTP(rsp, req)
	return

	if !strings.Contains(req.URL.Path, "html") {
		self.h.ServeHTTP(rsp, req)
		return
	}

	tp, _ := pongo2.FromFile("/Users/zp/Documents/knowledgeTree/go/studygo/xserver/statics/editor.md/show.html")

	f, err := os.Open("/Users/zp/Documents/knowledgeTree/go/studygo/xserver/statics/editor.md/README.md")
	if err != nil {

		//ex, _ := os.Executable()

		fmt.Println("open file err", err)
		return
	}

	bytes := make([]byte, 100)
	tx := make([]byte, 0)
	for {
		n, errf := f.Read(bytes)
		if errf != nil {
			fmt.Println("read file err")
			break
		}
		tx = append(tx, bytes[:n]...)
	}

	v := string(tx)
	strings.Replace(v, "/`", "/`", 1)
	bytes2.Replace(tx, []byte("`"), []byte("\\`"), -1)

	newBytes := make([]byte, 0)

	for _, b := range tx {
		if b == '`' {
			newBytes = append(newBytes, '\\')
		}
		newBytes = append(newBytes, b)
	}
	//fmt.Println(string(tx))

	//strings.Replace(v, "\\\\`", "cvb", -1)
	//v = strconv.Quote(v)
	//fmt.Println("before", v, "final")
	//fmt.Println("read file is:", ",", string(tx))
	ctx := pongo2.Context{"contentmd": string(newBytes)}

	afbytes, err := tp.ExecuteBytes(ctx)
	if err != nil {
		return
	}
	//fmt.Println("headers", rsp.Header())
	//fmt.Println("tp execute is:", string(afbytes))
	rsp.Header().Set("Content-Type", "text/html; charset=utf-8")

	rsp.WriteHeader(200)
	//fmt.Println(string(afbytes))
	io.CopyN(rsp, bytes2.NewReader(afbytes), int64(len(afbytes)))

}

var Entrys = map[string]http.Handler{
	"edit":   &EditHandler{h: http.FileServer(http.Dir("./xserver")), s: "first"},
	"pages":  &PagesHandler{},
	"static": &xFileHandler{h: http.FileServer(http.Dir("./xserver")), s: "first"},
}

type EditHandler struct {
	h http.Handler
	s string
}

func (self *EditHandler) ServeHTTP(rsp http.ResponseWriter, req *http.Request) {
	fmt.Println("editHandler", req.URL.Path)

	tp, _ := pongo2.FromFile("/Users/zp/Documents/knowledgeTree/go/studygo/xserver/edit/edi.html")

	f, err := os.Open("/Users/zp/Documents/knowledgeTree/go/studygo/xserver/md/xuyongkang/README.md")
	if err != nil {

		//ex, _ := os.Executable()

		fmt.Println("open file err", err)
		return
	}

	bytes := make([]byte, 100)
	tx := make([]byte, 0)
	for {
		n, errf := f.Read(bytes)
		if errf != nil {
			fmt.Println("read file err")
			break
		}
		tx = append(tx, bytes[:n]...)
	}

	v := string(tx)
	strings.Replace(v, "/`", "/`", 1)
	bytes2.Replace(tx, []byte("`"), []byte("\\`"), -1)

	newBytes := make([]byte, 0)

	for _, b := range tx {
		if b == '`' {
			newBytes = append(newBytes, '\\')
		}
		newBytes = append(newBytes, b)
	}
	//fmt.Println(string(tx))

	//strings.Replace(v, "\\\\`", "cvb", -1)
	//v = strconv.Quote(v)
	//fmt.Println("before", v, "final")
	//fmt.Println("read file is:", ",", string(tx))
	ctx := pongo2.Context{"contentmd": string(newBytes)}

	afbytes, err := tp.ExecuteBytes(ctx)
	if err != nil {
		return
	}
	//fmt.Println("headers", rsp.Header())
	//fmt.Println("tp execute is:", string(afbytes))
	rsp.Header().Set("Content-Type", "text/html; charset=utf-8")

	rsp.WriteHeader(200)
	//fmt.Println(string(afbytes))
	io.CopyN(rsp, bytes2.NewReader(afbytes), int64(len(afbytes)))

	//self.h.ServeHTTP(w, req)
	return
}

type PagesHandler struct {
	h http.Handler
	s string
}

func (self *PagesHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

}

type DispatchHandler struct {
}

func (self *DispatchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("req is %+v", r.URL.Path)

	url := r.URL.Path

	if len(url) == 0 {
		w.WriteHeader(404)
	}
	nurl := url[1:]
	segs := strings.Split(nurl, "/")

	if len(segs) == 0 {
		w.WriteHeader(404)
		return
	}

	fmt.Println("segs", segs, segs[0])

	if h, ok := Entrys[segs[0]]; ok {
		h.ServeHTTP(w, r)
		return
	} else {
		fmt.Println("not found handler")
	}

	w.WriteHeader(404)
}

func fileServer() {
	mux := http.NewServeMux()

	//mux.Handle("/", &xFileHandler{h: http.FileServer(http.Dir("./xserver/statics/editor.md")), s: "first"})
	mux.Handle("/", &DispatchHandler{})

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	fileServer()
}
