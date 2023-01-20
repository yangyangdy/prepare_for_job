package main

import (
	"fmt"
	"os"
)

//union-find
//并查集的基本思想就是构建一颗颗的树
//并查集中的树并不是平衡树，在极端情况下树会退化成的链表，因此极端情况下寻找根节点的时间复杂度是O(n)
//因此，为了优化性能，我们在合并两个union的时候需要保证合并后树的平衡性
//小一些的树接到大一些的树下面，这样就能避免头重脚轻，更平衡一些
func main() {
	fmt.Println(os.Stdout)
	fmt.Println(len(os.Args))
	fmt.Println(os.Args)
	//if len(os.Args) >= 2 {
	//	switch os.Args[1] {
	//	case "indexer":
	//		cmd.Indexer()
	//		return
	//	case "crawler":
	//		cmd.Crawler()
	//		return
	//	}
	//}
}

// [0, 1, 2, 3, 4, 5, 6, 7,...]
func sift(arr[]int, i int, arrLen int) []int {
	done := false
	tmp := 0
	maxChild := 0

	for (i*2+1 < arrLen) && (!done) { //i的左子树
		if i*2+1 == arrLen-1 {
			maxChild = i*2+1
		} else if arr[i*2+1] > arr[i*2+2] { //左子树大于右子树
			maxChild = i*2+1
		} else {
			maxChild = i*2+2
		}

		if arr[i] < arr[maxChild] {
			tmp = arr[i]
			arr[i] = arr[maxChild]
			arr[maxChild] = tmp
			i = maxChild
		} else {
			done = true
		}
	}
	return arr
}

func HeapSort(arr []int) {
	i := 0
	tmp := 0
	for i = len(arr)/2-1; i>=0;i--{
		//从倒数第二层的做后一个节点开始调整
		arr = sift(arr, i, len(arr))
	}

	for i = len(arr)-1;i>=1;i--{
		tmp = arr[0]
		arr[0] = arr[i]
		arr[i] = tmp
		arr = sift(arr, 0, i)
	}
}