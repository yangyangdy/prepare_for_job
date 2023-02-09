package main

import (
	"fmt"
)

type test66 struct {
	a int
}

func(p *test66) f1(){
	fmt.Println(p.a)
}

func genNumberAndString() (int, string){
	a1:= 5
	b1:="hello"
	return a1, b1
}
func test66_()(a int, b string ){
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(&b)
	a, astr := genNumberAndString()
	_ = astr
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(&b)
	return
}
func main() {
	mp := map[string]int{
		"1":1,
		"2":2,
		"3":3,
	}
	
	var val int
	val = mp["4"]
	fmt.Println(val)
	//test66_()
	//t := time.Now().Unix()
	//formatTimeStr := time.Unix(t,0).Format("2006-01-02 15:04:05")
	//
	//fmt.Println(formatTimeStr)
	//defer func(){
	//	if err := recover(); err != nil {
	//		debug.PrintStack()
	//		fmt.Println(debug.Stack())
	//	}
	//}()
	//var t *test66
	//t.f1()
	//var fixedPriceTotalFee, fixedPriceAfterPrice int64 = -1, -1
	//fmt.Println(fixedPriceTotalFee)
	//fmt.Println(fixedPriceAfterPrice)
	//str := "4.5057948277295247E-4"
	//res, _ := strconv.ParseFloat(str, 64)
	//fmt.Println(res)
	//var mp map[int]int
	//v,ok := mp[4]
	//if ok {
	//	fmt.Println("1")
	//} else {
	//	fmt.Println(v)
	//}
	// var a float64 = 0.12
	// var b *float64 = &a
	//
	// c := int32(*b * 100)
	// fmt.Println(c)
	//fmt.Println("date1Start: ", date1Start)

	//type ZeroTrCityConfData struct {
	//	zeroTrCityList []int `json:"zero_tr_city_list"`
	//	zeroTrCityMap  map[int]bool
	//}
	//a := &ZeroTrCityConfData{}
	//a.zeroTrCityList = []int{1,2,3,4,5,6}
	//a.zeroTrCityMap = make(map[int]bool)
	//for _, v := range a.zeroTrCityList{
	//	a.zeroTrCityMap[v] = true
	//}
	//fmt.Println(a.zeroTrCityList)
	//fmt.Println("--------------")
	//fmt.Println(a.zeroTrCityMap)
}
