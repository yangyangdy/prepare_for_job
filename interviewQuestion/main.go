package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"
	"unsafe"
)

type Handler interface {
	Do(k, v interface{})
}

func Each(m map[interface{}]interface{}, h Handler) {
	if m != nil && len(m) > 0 {
		for k, v := range m {
			h.Do(k, v)
		}
	}
}

type welcome string

// 必须实现handler接口，Do方法名不能改，不能定义一个更有意义的名字
// 必须定义一个新类型，才能实现handler接口，才能使用each函数
func (w welcome) Do(k, v interface{}) {
	fmt.Println("%s,my name is %s, age is %d\n", w, k, v)
}

func (w welcome) selfInfo(k, v interface{}) {
	fmt.Println("hello world!")
}

type HandlerFunc func(k, v interface{})

func (f HandlerFunc) Do(k, v interface{}) {
	f(k, v)
}

func EachFunc(m map[interface{}]interface{}, f func(k, v interface{})) {
	Each(m, HandlerFunc(f))
}

// 强转
func String2Bytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

type A func(int, int)

func (f A) Serve() {
	fmt.Println("serve2")
}

func serve(int, int) {
	fmt.Println("serve1")
}

var Error *log.Logger

func init() {
	errFile, err := os.OpenFile("./crash.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("打开日志文件失败：", err)
	}
	Error = log.New(errFile, "Error:", log.Ldate|log.Ltime|log.Lshortfile)
}

func panicFunc() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("程序crash了：", err)
		}
	}()
	log.Println("panicFunc 开始运行")
	panic("------crash------")
}

type Printer interface {
	Print()
}

type CanonPrinter struct {
	printerName string
}

func (printer CanonPrinter) Print() {
	fmt.Println(printer.printerName, "print.")
}

type PrintWorker struct {
	Printer
	name string
	age  int
}

type mytype int

type node struct {
	val int
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)

}

func A1(a int) {
	fmt.Println(a)
}

func AA() {
	a, b := 1, 2
	defer A1(a)
	a = a + b
	fmt.Println(a, b)
}

var (
	slice  = make([]*int, 5)
	slice2 = slice[2:]
	tmp    = 1
)

func printSlice() {
	fmt.Printf("len=%d cap=%d %v\n", len(slice), cap(slice), slice)
	fmt.Printf("len=%d cap=%d %v\n", len(slice2), cap(slice2), slice2)
}

type my1 struct {
	a int
	b string
	c bool
}

type my2 struct {
	d int
	e string
	f my1
}

const (
	aa = -1
	bb = iota - 1
	cc
	dd
	ee
)

var wg sync.WaitGroup

func func1(chan_one chan bool, chan_two chan bool) {
	for {
		select {
		case <-chan_one:
			fmt.Println("one")
			chan_two <- true
			break
		default:
			break
		}
	}
}

func f3() {
	panic("panic err")
}
func f2() {
	f3()
}
func f1() {
	f2()
}

type key int

type mykey struct {
}

func getParalleInputUnionDiscount(inputUnionDiscount map[string][]float64, parallelDegree int) []map[string][]float64 {
	result := make([]map[string][]float64, parallelDegree)
	i := 0
	sliceLen := len(inputUnionDiscount) / parallelDegree // 每个并行的分组包含的折扣组合数
	if len(inputUnionDiscount)%parallelDegree != 0 {
		sliceLen = sliceLen + 1
	}
	for k, v := range inputUnionDiscount {
		result[i/sliceLen][k] = v
		i++
	}
	return result
}

type my struct {
	val int
}

type WordNode struct {
	count int
	word  string
	next  *WordNode
}
func SortList(head *WordNode) *WordNode {
	if head == nil || head.next == nil {
		return head
	}
	mid := midList(head)
	LA := head
	LB := mid.next
	mid.next = nil
	// LA = sortList(LA)
	// LB = sortList(LB)
	return merge(SortList(LA), SortList(LB))
}

func midList(head *WordNode) *WordNode {
	slow, fast := head, head
	for fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}

func merge(L1 *WordNode, L2 *WordNode) *WordNode {
	var res = &WordNode{}
	next := res
	currA := L1
	currB := L2
	for currA != nil && currB != nil {
		if currA.count > currB.count {
			next.next = &WordNode{
				count: currA.count,
				word:  currA.word,
			}
			currA = currA.next
		} else {
			next.next = &WordNode{
				count: currB.count,
				word:  currB.word,
			}
			currB = currB.next
		}
		next = next.next
	}
	if currA != nil {
		next.next = currA
		currA = currA.next
	}
	if currB != nil {
		next.next = currB
		currB = currB.next
	}
	return res.next
}

func test1() *int{
	a := 45

	return &a
}

func trace(a []int, k int) int{
	var ret int
	for i := k; i < len(a); i++ {
		if a[i] == 1 {
			return a[i]
		}
		if a[i] == 2{
			return a[i]
		}
		 ret = trace(a, k+1)
	}
	return ret
}

func test_sync_once() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("only once")
	}
	done := make(chan bool)
	for i :=0; i<10; i++ {
		go func(){
			once.Do(onceBody)
			done <- true
		}()
	}
	for i := 0; i<10; i++ {
		<-done
	}
	close(done)
}

type Mystring string
func (m Mystring) String() string { //这样或造成无限递归
	return fmt.Sprintf("MyString=%s", m)
}

type testdefer struct {
	a int
	errNo int
	errMsg string
}

func testDefer() (result *testdefer) {
	var errStr string
	result = &testdefer{}
	defer func() {
		if errStr != "" {
			result.errNo=-1
			result.errMsg = "helloworld"
		}
	}()
	errStr = "1"
	return
}

func main() {
	f1()
	var a_val int32 = 1230
	b_val := a_val/int32(100)
	fmt.Println("________b_val______", b_val)
	a := false
	var ab *WordNode
	if (a && ab.count == 0 ) || true {
		fmt.Println("________get_in______")
	}
	strr := "[12345]"
	strr = strr[1:len(strr)-1]
	fmt.Println(strr)
	slic := make([]int, 3)
	slic[0] = 1
	slic[1] = 2
	slic[2] = 3
	fmt.Println(slic)
	var mmp []map[int]int
	item1 := make(map[int]int, 0)
	mmp = append(mmp, item1)
	fmt.Println(mmp)
	mapA := make(map[float64]float64,0)
	mapA[1.0] = 202
	v_, ok_ := mapA[1]
	if ok_ {
		fmt.Println(v_)
	}
	resDefer := testDefer()
	fmt.Println(resDefer)
	fmt.Println("///")
	estimateDataMap := make(map[string]string, 0)
	fmt.Println("len:", len(estimateDataMap))
	for _, value := range estimateDataMap {
		fmt.Println("get_in:%v", value)
		break
	}
	ListStr := " [0.8499187966436534, 6777.371331348362, 0.8649394747932534, 726.7189977185724, 2097.639360426898]"
	ListStr = strings.TrimSpace(ListStr)
	ListStr = ListStr[1:len(ListStr)-1]
	fmt.Println(ListStr)
	countSplit := strings.Split(ListStr, ",")
	var coefList []float64
	for i := range countSplit {
		str := strings.TrimSpace(countSplit[i])
		v,e := strconv.ParseFloat(str, 64)
		if e != nil {
			return
		}
		coefList = append(coefList, v)
	}
	fmt.Println(coefList)
	vv, er := strconv.ParseFloat("-0.19",64)
	if er != nil {
		fmt.Println("parse fail")
	} else {
		fmt.Println(vv)
	}
	fmt.Println(countSplit)
	var ff float64
	fmt.Println(ff)
	fmt.Println("///")
	//a := []byte("hello")
	//b := []byte("world")


	c := make(chan int) //分配一个信道

	go func(){
	c <- 1 //发送信号
	}()
	<-c //丢弃发来的值

	var mstr Mystring
	mstr = "test"
	mstr.String()
	aaa := make([]int, 0)
	if len(aaa) == 0 {
		fmt.Println("yesyeysye")
	}
	dayOfWeek := time.Time.Weekday(time.Now())
	dayOfWeek = 7
	fmt.Println(int(dayOfWeek))
	test_sync_once()
	res := make([]map[string][]float64, 3)
	for i:=0; i<3; i++ {
		res[i] = make(map[string][]float64, 0)
	}
	res[0]["0"] = []float64{1,2,3}
	res_tmp := res[0]["0"]
	res_tmp[0] = 5
	fmt.Println(res[0]["0"][0])
	fmt.Println("/n")
	aa := test1()
	fmt.Println(aa, *aa)
	predicts := make(map[string]*my, 0)
	_,ok := predicts["1"]

	if !ok {
		predicts["1"] = &my{
			val:1,
		}
	}

	result := make([]map[string][]float64, 5)
	result[0] = make(map[string][]float64)
	result[0]["0"] = []float64{1,2,3}
	sli := make([]string, 100)
	sli = append(sli, "hihi")
	fmt.Println(sli, "可以访问")

	str := "fp_ite_100_95"
	strList := strings.Split(str, "_")
	fmt.Println(len(strList), strList)
	var fff float64 = 0.8500
	fff_str := strconv.FormatFloat(fff,'f', -1, 64)
	fmt.Println(fff_str)
	mp1 := make(map[string]int)
	m1 := mp1["a"]
	if m1 == 0 {
		fmt.Println("value: %v", m1)
	}

	f1()
	fmt.Println("aa:", aa)
	fmt.Println("bb:", bb)
	fmt.Println("cc:", cc)
	fmt.Println("dd:", dd)
	fmt.Println("ee:", ee)
	var j = false
	if j {
		goto doit
	}
	fmt.Println()
doit:
	fmt.Println("gggggg")
	fmt.Println("ccccccc")
	var index float64 = 1006

	indexStr := strconv.FormatFloat(index, 'f', 2, 64)
	fmt.Println(indexStr)

	ta := 2
	if ta == 2 {
		fmt.Println("------yes-----")
	} else if ta == 2 {
		fmt.Println("********yes******")
	}
	var mp map[int]int
	fmt.Println("mp len :", len(mp))
	printSlice()
	//修改slice2第0个元素

	slice2[0] = &tmp
	printSlice()
	//扩容slice2
	tmp = 2
	slice2 = append(slice2, &tmp)
	printSlice()
	//修改slice2第0个元素
	tmp = 3
	slice2[0] = &tmp
	printSlice()
	//修改slice第0个元素
	tmp = 4
	slice[0] = &tmp
	printSlice()

	fmt.Println("---------")
	AA()
	start := time.Now().UnixNano()
	time.Sleep(1 * time.Millisecond)
	end := time.Now().UnixNano()
	diff := (end - start) / int64(time.Millisecond)
	fmt.Println("time diff: ", diff)

	//a := int64(4)
	//b := int64(5)
	//c := a / b
	fmt.Println(c)
	arr := make([]*node, 6)
	for i := range arr {
		arr[i] = &node{val: i}
	}
	tt := make([]*node, 0)
	tt = arr
	tt[0].val = 34
	fmt.Println("-------*******--------")
	for i := range tt {
		fmt.Println(tt[i].val)
	}
	fmt.Println("-------*******--------")
	arr1 := make([]*node, 0)
	arr2 := make([]*node, 0)
	for i := 0; i < 3; i++ {
		arr1 = append(arr1, arr[i])
	}
	for i := 3; i < 6; i++ {
		arr2 = append(arr2, arr[i])
	}
	arr1[0].val = 9
	arr2[0].val = 8

	for i := range arr {
		fmt.Println(arr[i].val)
	}
	// fmt.Println("arr: ", arr)

	m := make(map[string]string)
	fmt.Println("m: ", len(m))

	v, ok := m["r"]
	if v == "" {
		fmt.Println("__nil__")
	}
	fmt.Println(v, ok)
	//canon := CanonPrinter{printerName: "canon_1"}
	//printWorker := PrintWorker{Printer: canon, name: "Zhang", age: 21}
	//printWorker.Printer.Print()
	// go panicFunc()
	// time.Sleep(2*time.Second)
	/*
		a := A(serve)
		a(1,2)
		a.Serve()


		person := make(map[interface{}]interface{})
		person["张三"] = 20
		person["李四"] = 23
		person["王五"] = 26

		var w welcome = "hello"

		EachFunc(person, w.selfInfo)
		// a := "hello"
		// b := String2Bytes(a)
		// b[0] = 'H'
		// http.Handle()

	*/

}
