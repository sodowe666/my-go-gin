package main

import (
	"fmt"
	"net/http"
	"runtime"
	"strconv"
	"time"
)

func index(w http.ResponseWriter, r *http.Request) {
	i++
	fmt.Println("aaaaa")
	//fmt.Fprintf(w,"dddd")
	//fmt.Fprintf(w,strconv.Itoa(i))
	go func() {
		time.Sleep(5 * time.Second)
	}()
	num := runtime.NumGoroutine()
	fmt.Fprintf(w, strconv.Itoa(num))
}

var i int

type MyStruct struct {
	name string
}

func (this MyStruct) GetName(str string) string {
	this.name = str
	return this.name
}

type aaa interface {
}

func main() {
	//runtime.NumGoroutine();
	//scanner := bufio.NewScanner(
	//	strings.NewReader("ABCDEFG\nHIJKELM"),
	//)
	//scanner.Split(bufio.ScanWords) /*四种方式之一，你也可以自定义, 实现SplitFunc方法*/
	//for scanner.Scan() {
	//	fmt.Println(scanner.Text()) // scanner.Bytes()
	//}
	//http.ListenAndServe("0.0.0.0:8080", nil)
	//data := []byte{10, 0, 11, 0, 10, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10, 3, 97, 97, 97, 18, 3, 98, 98, 98}
	//pkg := bytes.Buffer{}
	//l, _ := pkg.Write(data)
	//var length int32
	//k := new(int8)
	//err := binary.Read(&pkg, binary.LittleEndian, k)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(pkg.Bytes())
	//fmt.Println(length)
	//fmt&.Println(l)
	a := 6
	fmt.Printf("%p", &a)
	fmt.Println()
	b := 7
	fmt.Printf("%p", &b)

}
