package main

import (
	"fmt"
	"encoding/json"
)

type Point struct {
	x float64
	y float64
}

type Shape struct {
	leftup, rightdown Point
	Name byte
	Number int8
	Age int8
}

type Person struct {
	Name string
	Age int
}

type Cat struct {
	Name string
	Age int
	Color string
	Scores []int
	Features map[string]string
	Nextcat *Cat
}

type Monster struct {
	Name string `json:"name"` // `json:"name"` => struct tag
	Age int `json:"age"`
	Skill string `json:"skill"`
}

func main() {
	
	arr := make([]int,10)
	fmt.Println(arr)

	var cat1 Cat 
	cat1.Name = "baby"
	cat1.Age = 7
	cat1.Color = "White"
	cat1.Scores = make([]int,10)
	cat1.Features = make(map[string]string)

	cat1.Features["no1"] = "cute"
	testStruct(&cat1)

	fmt.Println(cat1)

	testStuct2(&(cat1.Name))

	fmt.Println(cat1)

	var cat2 Cat

	cat2.Name = "baby"
	cat2.Age = 7
	cat2.Color = "black"

	cat1.Nextcat = &cat2 // link list

	fmt.Println(cat1.Nextcat.Name)

	// 創建方法2 => 推薦使用

	
	var p1 Person = Person{"harry", 4}
	// p1 := Person{"harry", 4}
	fmt.Println(p1)
	

	//創建方法3

	p2 := new(Person) // 要一個 Person 的空間 返回 ptr => review : new() => 要值類型的空間 ， make() => 要引用類型的空間 

	fmt.Printf("\n%T\n",p2)
	// 這邊有一個點是說 照理來說p2得到的是一個pointer 所以輸出Name時 應該要(*p2).Name，但golang底層被優化過
	// 會進行再處理 所以可以簡化成 p2.Name
	p2.Name = "Jack"
	p2.Age = 12

	fmt.Println(p2.Name)

	//創建方法4

	var p3 *Person = &Person{}

	p3.Name = "george"

	fmt.Println(p3.Name)

	
	rectangle := Shape{Point{2,3},Point{5,6},'e',5,6}
	fmt.Println()
	fmt.Printf("%v , %v , %v , %v , %v , %v, %v\n",&rectangle.leftup.x,&rectangle.leftup.y,&rectangle.rightdown.x,&rectangle.rightdown.y,&rectangle.Name,&rectangle.Number,&rectangle.Age)

	
	// struct 序列化 轉成string json 格式

	monster := Monster{"Hacker",5,"coding"}

	jsonStr, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("json error", err)
	}
	fmt.Println("jsonStr",string(jsonStr))
}

func testStruct(ptr *Cat) {
	(*ptr).Name = "kitty"
	(*ptr).Age = 5
}

func testStuct2(ptr *string) {
	*ptr = "jacky"
}