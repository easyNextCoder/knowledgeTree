package files

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"
	"time"
)

/*
close函数是一个内建函数， 用来关闭channel，这个channel要么是双向的， 要么是只写的（chan<- Type）。
这个方法应该只由发送者调用， 而不是接收者。
当最后一个发送的值都被接收者从关闭的channel(下简称为c)中接收时,
接下来所有接收的值都会非阻塞直接成功，返回channel元素的零值。
如下的代码：
如果c已经关闭（c中所有值都被接收）， x, ok := <- c， 读取ok将会得到false。

作者：夏海社长
链接：https://www.jianshu.com/p/eb1a9b316f05
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

var done = make(chan struct{}) //用一个done控制整个流程，不再启用新的go程

func cancelled() bool {
	select {
	case val, ok := <-done:
		fmt.Println("cancelled", val, ok)
		return true
	default:
		return false
	}
}

type XFile struct {
	file os.FileInfo
	dir  string
}

var sema = make(chan interface{}, 1)
var files = make(chan XFile)
var nowDirs = make(chan string)
var sn = sync.WaitGroup{}

func dirents(dir string) []os.FileInfo {
	//fmt.Println("dirents start")
	select {
	case sema <- struct{}{}: // acquire token
	case <-done:
		return nil // cancelled
	}
	defer func() { <-sema }() // release token
	// ...read directory...
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil
	}

	return fileInfos
}

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- XFile) {
	n.Add(1)
	if cancelled() {
		n.Done()
		return
	}
	for _, entry := range dirents(dir) {
		// ...
		fileSizes <- XFile{
			file: entry,
			dir:  dir,
		}
	}
	n.Done()
}

func selfWorkDir(dirs chan string, n *sync.WaitGroup, xfiles chan<- XFile) {

	n.Add(1)
	if cancelled() {
		fmt.Println("canceled!!!!")
		n.Done()
	}

	for dir := range dirs {

		for _, entry := range dirents(dir) {

			xfiles <- XFile{
				file: entry,
				dir:  dir,
			}
		}
	}

}

func sizeReceiver() {

	for {
		select {
		case <-done:
			// Drain files to allow existing goroutines to finish.
			fmt.Println("close sizeReceiver")
			for range files {
				// Do nothing.
			}
			return
		case xfile, ok := <-files:
			// ...
			if ok {
				//fmt.Println(xfile.dir)
				if xfile.file.IsDir() {
					newDir := xfile.dir + xfile.file.Name() + "/"
					go func() { //如果不开启go，则要为nowDirs增加很大的缓存，否则会卡死
						nowDirs <- newDir
					}()

				}
			}

		}

	}
}

func ControlWork() {
	// Cancel traversal when input is detected.
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		close(done)
	}()

	sn := sync.WaitGroup{}

	go selfWorkDir(nowDirs, &sn, files)
	go selfWorkDir(nowDirs, &sn, files)
	go sizeReceiver()
	nowDirs <- "/"

	//time.Sleep(time.Second * 1)
	sn.Wait()
	//sn.Add(1)
	//walkDir("/", &sn, files)
	//sn.Wait()
}

func sizeReceiver2() {

	for {
		select {
		case <-done:
			// Drain files to allow existing goroutines to finish.
			fmt.Println("close sizeReceiver")
			for range files {
				// Do nothing.
			}
			return
		case xfile, ok := <-files:
			// ...
			if ok {
				//fmt.Println(xfile.dir)
				if xfile.file.IsDir() {
					newDir := xfile.dir + xfile.file.Name() + "/"
					go walkDir(newDir, &sn, files)

				}
			}

		}

	}
}
func ControlWork2() {
	// Cancel traversal when input is detected.
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		fmt.Println("input char close done")
		close(done)
	}()

	go walkDir("/", &sn, files)
	go sizeReceiver2()

	time.Sleep(time.Second * 1)
	sn.Wait()
	//sn.Add(1)
	//walkDir("/", &sn, files)
	//sn.Wait()
}
