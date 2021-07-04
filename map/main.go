package main

import (
	"fmt"
	"sort"
)

func main() {
	var a map[string]string
	a = make(map[string]string)

	// key不能重複
	// data為無序的

	a["no1"] = "Steven"
	a["no2"] = "Gary"
	a["no1"] = "Harry"
	a["no3"] = "Lion"

	fmt.Println(a)

	b := make(map[string]string)
	b["no1"] = "kenny"
	fmt.Println(b)

	//另一種方式使用map

	heroes := map[string]string {
		"hero1" : "superman",
		"hero2" : "batman",
		"hero3" : "the flash",
	}

	fmt.Println(heroes)

	// Example 鍵入學生資料 有 name , sex


	student := make(map[string]map[string]string)

	student["stud01"] = make(map[string]string,2)
	student["stud01"]["name"] = "Steven"
	student["stud01"]["sex"] = "M"

	student["stud02"] = make(map[string]string,2)
	student["stud02"]["name"] = "Mary"
	student["stud02"]["sex"] = "W"

	student["stud03"] = make(map[string]string,2)
	student["stud03"]["name"] = "Kevin"
	student["stud03"]["sex"] = "M"

	fmt.Println(student)


	// map 增刪改查

	// 增加上述已經提過了
	student["stud03"]["name"] = "Ken" //這樣就修改了
	fmt.Println(student)
	
	// 刪除功能

	delete(student["stud02"],"name")
	fmt.Println(student)

	delete(student,"stud02")
	fmt.Println(student)

	// 若要全部刪除有兩種辦法
	// 1. 遍歷刪除
	// 2. 直接make新的

	// student = make(map[string]map[string]string)
	// fmt.Println(student)
 
	// 查詢功能

	val, ok := student["stud01"]
	// 會返回兩個值 有找到 ok = ture ,沒找到 ok = false  
	if ok {
		fmt.Println(val)
		fmt.Println(ok)
		
	} else {
		fmt.Println(val)
		fmt.Println(ok)
	}

	// map 遍歷

	for key, _ := range student {
		fmt.Println(key)
		
		for key2, val2 := range student[key] {
			fmt.Println(key2,val2)

		}
	}
	
	// map 長度
	fmt.Println(len(student))


	// map slice (切片)
	// 很像再map外面包一層slice(array)的概念 加了index的感覺 且可動態配置
	
	monsters := make([]map[string]string, 2)

	if monsters[0] == nil {
		monsters[0] = make(map[string]string,2)
		monsters[0]["name"] = "m1"
		monsters[0]["age"] = "1500"
	}

	if monsters[1] == nil {
		monsters[1] = make(map[string]string,2)
		monsters[1]["name"] = "m2"
		monsters[1]["age"] = "400"
	}

	fmt.Println(monsters)

	// 動態增加 map slice => 一樣用 append
	
	newMonster := map[string]string {
		"name" : "m3",
		"age" : "600",
	}

	monsters = append(monsters, newMonster)
	fmt.Println(monsters)


	// map 本身是沒有排序的
	// 若硬要排序的方法

	map1 := make(map[int]int)

	map1[1] = 12
	map1[4] = 36
	map1[14] = 56
	map1[18] = 123

	var keys []int

	fmt.Printf("%T", keys)

	for k,_ := range map1 {
		keys = append(keys,k)
	}

	sort.Ints(keys)

	fmt.Println(keys)

	for _,val := range keys {
		fmt.Printf("map1[%d] = %d\n", val ,map1[val])
	}

	// map 搭配 struct 使用

	type Stu struct{
		name string
		age int
		address string
	}

	studs := make(map[string]Stu)
	stud1 := Stu{"steven",23,"taipei"}
	stud2 := Stu{"tom",53,"taipei"}
	stud3 := Stu{"JERRY",63,"NY"}

	studs["stud1"] = stud1
	studs["stud2"] = stud2
	studs["stud3"] = stud3

	fmt.Println(studs)

	for key, val := range studs {
		fmt.Printf("%s : \n",key)
		fmt.Printf("name : %s\n",val.name)
		fmt.Printf("age : %d\n",val.age)
		fmt.Printf("address : %s \n",val.address)
		fmt.Println("-----------")
	}

	// Example

	members := make(map[string]map[string]string)

	modifyMember(members,"steven")
	modifyMember(members,"marry")
	modifyMember(members,"tom")

	modifyMember(members,"tom")

	fmt.Println(members)
	
}

func modifyMember(members map[string]map[string]string, name string) {
	
	if _, ok := members[name]; ok {
		members[name]["pwd"] = "888888" 
		fmt.Println("Modfication of Account already complete")

	}else {

		members[name] = make(map[string]string)
		members[name]["nickname"] = name
		members[name]["pwd"] = "12345"
		fmt.Println("Applyment of Account already complete")
	}
}