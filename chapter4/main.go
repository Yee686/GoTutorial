package main

import (
	"fmt"
)

func arrayOption() {
	fmt.Println("arrayOption 数组操作")

	array := [5]int{10, 20, 30, 40, 50}

	for i, v := range array {
		fmt.Println(i, v)
	}

	array[1] = 25

	for i, v := range array {
		fmt.Println(i, v)
	}

	var array1 [5]int

	for i, v := range array1 {
		fmt.Println(i, v)
	}

	array3 := [2]*int{0: new(int), 1: new(int)}
	for i, v := range array3 {
		fmt.Println(i, v, *v)
	}

	*array3[1] = 10
	for i, v := range array3 {
		fmt.Println(i, v, *v)
	}

	arr := [5][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}, {50, 51}}
	for _, v := range arr {
		fmt.Println(v)
	}

	arr2 := [2][2]int{1: {1: 1}}
	for _, v := range arr2 {
		fmt.Println(v)
	}
}

func sliceOption() {
	fmt.Println("sliceOption 切片操作")
	var slice []int
	slice2 := make([]string, 3, 5)
	fmt.Println(slice, len(slice), cap(slice))
	fmt.Println(slice2, len(slice2), cap(slice2))
	fmt.Println(slice == nil)

	var array = [5]int{10, 20, 30, 40, 50}
	slice3 := array[1:3]
	fmt.Println(slice3, len(slice3), cap(slice3))
	slice4 := append(slice3, 60)
	fmt.Println(slice4, len(slice4), cap(slice4))
	slice5 := append([]byte("abc"), []byte("def")...)
	fmt.Println(slice5, len(slice5), cap(slice5))

	slice6 := []int{10, 20, 30, 40}
	slice7 := append(slice6, 50)
	fmt.Println(slice6, len(slice6), cap(slice6))
	fmt.Println(slice7, len(slice7), cap(slice7))

	source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	fmt.Println(source, len(source), cap(source))
	fmt.Println(source[2:3:4], len(source[2:3:4]), cap(source[2:3:4]))

	// 不使用第三个索引，会改变底层数组的值
	sr := source[2:3:3]
	sr = append(sr, "Kiwi")
	fmt.Println(sr, len(sr), cap(sr))
	fmt.Println(source, len(source), cap(source))

	// 使用第三个索引，由于容量限制会创建新的底层数组
	sr2 := source[2:3]
	sr2 = append(sr2, "Kiwi")
	fmt.Println(sr2, len(sr2), cap(sr2))
	fmt.Println(source, len(source), cap(source))

	// 迭代切片,在迭代切片中返回的value时对该元素的引用,且地址总是相同的
	for i, v := range source {
		fmt.Printf("Value: %s value-addr: %X element-addr: %X\n", v, &v, &source[i])
	}

	// for循环
	for index := 2; index < len(source); index++ {
		fmt.Printf("source[%d]: %s\n", index, source[index])
	}

	// 多维切片
	sli := [][]int{{10}, {100, 200}}
	fmt.Println(sli)
	sli[0] = append(sli[0], 20)
	fmt.Println(sli)
	sli = append(sli, []int{30})
	fmt.Println(sli)

}

func mapOption() {
	fmt.Println("mapOption 映射操作")

	dict := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}

	for key, value := range dict {
		fmt.Printf("Key: %s Value: %s\n", key, value)
	}

	colors := map[string]string{}
	colors["Red"] = "#da1337"

	colors2 := map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "#ff7F50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
	}

	for key, value := range colors2 {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

	fmt.Println("删除元素")

	delete(colors2, "Coral")
	for key, value := range colors2 {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

}

func main() {
	fmt.Println("数组、切片与映射")

	arrayOption()

	sliceOption()

	mapOption()
}
