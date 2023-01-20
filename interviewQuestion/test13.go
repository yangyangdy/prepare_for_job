package main

import "fmt"

type BcrModelType int
const (
	BcrModelTypeV1 BcrModelType = iota+6
	BcrModelTypeV2
	BcrModelTypeV3 = iota +1
	BcrModelTypeV4
)


func main(){
	fmt.Println(BcrModelTypeV1)
	fmt.Println(BcrModelTypeV2)
	fmt.Println(BcrModelTypeV3)
	fmt.Println(BcrModelTypeV4)

	return
}